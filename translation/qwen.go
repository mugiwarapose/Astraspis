package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

var narrative string = "这是一份来自临床试验 CFFS-098（一种抗 JKBD 抗体偶联药物）在晚期实体瘤患者中的首次人体Ⅰ CFFS-098（一种抗-01）06中心研究者的病例报告。" +
	"受试者是一位46岁的白人男性。某日，受试者被诊断为转移性胆管细胞癌（骨骼转移（2024年），肝转移（2023年）和恶性胸腔积液）。" +
	"受试者既往用药包括：用于治疗胆管细胞癌肝转移的顺铂，吉西他滨和度伐利尤单抗。" +
	"具体时间未知，受试者在外院另一临床试验发生疾病进展。" +
	"2024年6月20日，受试者签署知情同意书，受试者编号为06001。" +
	"2024年7月8日（C1D1（第1周期第1天）），受试者开始接受研究药物AMT-253治疗（492.2mg（用于剂量计算的体重为102.55kg），一次，静脉注射，批号：未知）。当日，受试者发生SAE输液相关反应（详见病例#AMT2024000025）。"

func Translate() {
	// 记录开始时间
	start := time.Now()
	m := map[string]string{
		"name":      "转移性胆管细胞癌",
		"narrative": narrative,
	}

	// 将map转换为JSON字符串
	mBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	prompt := "【临床医药领域翻译任务指令】" +
		"请将输入的JSON对象中的所有value值进行专业医学翻译，翻译成英文，保留原有键名(key)不变，返回结构相同的JSON文本。具体要求如下：" +

		"翻译领域：临床医学/药学/生命科学（根据具体需求选择），" +
		"保留JSON结构：{\"原键名\": \"翻译后的医学专业表述\", ...}，" +
		"术语要求：使用最新版医学主题词表(MeSH)标准译法，" +
		"格式要求：严格保持双引号，确保JSON语法有效性。" +
		"返回json文本，不要返回markdown格式" +
		"示例输入：" +
		"{\"symptom\": \"intermittent claudication\", \"diagnosis\": \"peripheral arterial disease\"}" +
		"示例输出：" +
		"{\"symptom\": \"间歇性跛行\",\"diagnosis\": \"外周动脉疾病 (Peripheral Arterial Disease, PAD)\"}" +
		"请对以下JSON进行翻译："
	//narrative := `这是一份来自临床试验 “AMT-253（一种抗 MUC18 抗体偶联药物）在晚期实体瘤患者中的首次人体Ⅰ 期研究”（方案编号：AMT-253-01）06中心研究者的病例报告。受试者是一位46岁的白人男性。某日，受试者被诊断为转移性胆管细胞癌（骨骼转移（2024年），肝转移（2023年）和恶性胸腔积液）。受试者既往用药包括：用于治疗胆管细胞癌肝转移的顺铂，吉西他滨和度伐利尤单抗。具体时间未知，受试者在外院另一临床试验发生疾病进展。2024年6月20日，受试者签署知情同意书，受试者编号为06001。2024年7月8日（C1D1（第1周期第1天）），受试者开始接受研究药物AMT-253治疗（492.2mg（用于剂量计算的体重为102.55kg），一次，静脉注射，批号：未知）。当日，受试者发生SAE输液相关反应（详见病例#AMT2024000025）。C1期间，受试者伴有恶性腹水和需要引流的胸腔积液，受试者有引流导管，每日引流两升。2024年7月12日，受试者发生SAE呼吸短促（详见病例#AMT2024000038）。2024年7月13日，血小板计数：131×10^9/L。2024年7月15日，血小板计数：154×10^9/L。2024年7月31日（C2D1），受试者接受研究药物AMT-253治疗（492.2mg（用于剂量计算的体重为92.6kg），一次，静脉注射，批号：未知）。SAE发生前，最近一次给药日期为2024年7月31日。2024年8月6日，受试者发生腹泻和恶心（报告为全身性治疗7天后）。受试者因呕吐和脱水（与疾病进展相关，非SAE）在外院住院并有高钾血症（口服摄入量减少引起，非SAE），腹泻和恶心（可能有关），于肿瘤科接受治疗。入院时合并中性粒细胞减少（WCC（白细胞计数）为1.6×10^9/L，中性粒细胞为0.9×10^9/L）和血小板减少症（1级）。当日，受试者发生SAE全血细胞减少症。无受试者SAE发生前的服用了非甾体类抗炎药及其他引起血小板减少的药物记录。受试者使用胰岛素/葡萄糖和聚苯乙烯磺酸钠治疗高钾血症，强的松和止吐药治疗腹泻和恶心后，受试者腹泻停止。受试者脓毒症筛查结果：目前呈阴性，受试者静脉注射特治星预防潜在感染。2024年8月，受试者住院期间，受试者接受补液（IV（静脉注射））和塞克利嗪（SC（皮下注射）），并使用一剂非格司亭（300 mcg）以增加中性粒细胞数。偶有直肠出血，便常规：阴性排除结直肠炎。K（钾）：6.9。受试者发生AKI（急性肾损伤）（持续）。呼吸道病毒拭子：PJP（耶氏肺孢子菌肺炎）生长（不显著）；尿液MCS（抗微生物药敏感性试验）：ve 。受试者住院期间未出现药物过敏。2024年8月7日，CXR（胸部X线）: R（右）多发性积液+L（左）胸腔积液。2024年8月8日，受试者达到血小板减少症G3 (3级)，胸部CT（电子计算机断层扫描）：1.右肺淋巴管癌性病变的影像学表现值得高度关注；2. 没有令人信服的PJP特征。右侧包裹性胸腔积液；下方成分变小，与l引流一致。左侧单纯性胸腔积液增大；3. 肺和肝转移灶大小轻度增大；4. 溶骨性转移灶增大；5. 总体上与疾病进展一致。当日，因口服摄入减少导致的高钾血症结束。Mx 后K：5.1。2024年8月11日，受试者达到血小板减少症G4级 (4级)。当日，受试者发生低钾血症。2024年8月12日，受试者血小板为10×10^9/L，血红蛋白降低，ANC（绝对中性粒细胞计数）为0.3×10^9/L（无发热），葡萄糖：6.7mmol/L，蛋白：9g/L，LDH（乳酸脱氢酶）：40U/L，MCS：待定。受试者输一单位血小板。受试者疾病恶化，经过与受试者家属，肿瘤科和姑息治疗团队共同讨论，决定回家临终护理。当日，受试者出院。2024年8月13日，受试者回家进行姑息治疗，不能到现场进行进展扫描或复查。受试者的ECOG（东部肿瘤协作组体能状态）评分为3分。由于临床进展，研究者决定停止受试者的治疗。2024年8月18日，受试者因疾病进展去世（已做病例报道详见病例#AMT2024000065）。未进行尸检。受试者治疗用药包括：塞克利嗪（0.5mL，皮下注射，每天3次，2024年08月），甲氧氯普胺（10mg，口服，每天3次），泮托拉唑（40mg，口服，每天2次），羟考酮（5mg，口服，每4小时1次，24小时后最大剂量：20mg）。受试者非药物治疗包括：人血白蛋白（白蛋白20%）、输注血小板。受试者因疾病进展退出研究，非SAE全血细胞减少症原因。SAE全血细胞减少症（CTCAE：4级）的严重程度标准为导致住院。因SAE全血细胞减少症对研究药物AMT-253采取的措施为停止用药（2024年08月13日)。截至报告时间，SAE全血细胞减少症转归为未恢复。研究者评价：SAE全血细胞减少症与研究药物AMT-253有关，和研究程序无关。申办方评价: 受试者为46岁男性，转移性胆管细胞癌患者。开始接受研究药物AMT-253治疗29天（末次用药6天）后，受试者因呕吐和脱水住院。当日，受试者发生SAE呕吐和脱水（详见case#AMT2024000037）和SAE血小板减少症。1周后，受试者出院回家进行姑息治疗，由于临床进展，研究者决定停止受试者的治疗。又5日后，受试者因疾病进展死亡，未进行尸检。对研究药物采取的措施为停止用药。截至报告时间，受试者SAE血小板减少症转归为未恢复。血液学毒性是研究药物ATM-253的潜在风险之一。根据目前信息，结合其时间关联性，考虑事件血小板减少症与研究药物AMT-253有关，与研究程序无关。受试者基础肿瘤疾病为可能的混杂因素。该事件为非预期事件。于2024年10月1日收到了研究者的首次报告。于2024年10月17日收到了研究者的随访报告，更新信息包括：受试者死亡日期和SAE描述。于2024年11月05日收到研究者的质疑回复和随访报告，更新信息包括：既往用药，SAE名称（由“thrombocytopenia”更新为“pancytopenia”），实验室检查和SAE描述。`
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 构建请求体
	requestBody := RequestBody{
		// 此处以qwen-plus为例，可按需更换模型名称。模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
		Model: "qwen-max-2025-01-25",
		Messages: []Message{
			{
				Role:    "system",
				Content: "你是一个翻译人员，负责将文本翻译成医药方面的专业术语",
			},
			{
				Role:    "user",
				Content: prompt + string(mBytes),
			},
		},
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	// 创建 POST 请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	// 若没有配置环境变量，请用百炼API Key将下行替换为：apiKey := "sk-xxx"
	//apiKey := os.Getenv("QWEN_API_KEY")
	apiKey := "sk-dR0oUlrTNQ"
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 读取响应体
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 打印响应内容
	fmt.Printf("翻译结果\n%s\n", bodyText)

	var result Result

	// 执行反序列化
	err = json.Unmarshal(bodyText, &result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("翻译结果\n%s\n", result.Choices[0].Message.Content)

	var ct map[string]string

	err = json.Unmarshal(bodyText, &ct)
	if err != nil {
		panic(err)
	}

	fmt.Printf("最终解析\n%s\n", ct)

	// 记录结束时间
	end := time.Now()

	// 计算执行时间
	duration := end.Sub(start)

	// 打印执行时间
	fmt.Printf("代码执行时间: %v\n", duration)
}

type Result struct {
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Message Message `json:"message"`
}

// type Message struct {
// 	Content string `json:"content"`
// }
