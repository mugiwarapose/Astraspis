package model

type E2bField struct {
	CaseIndentifier     CaseIndentifier      `bson:"caseIndentifier"`               //C.1.1
	SenderHeldDocuments []SenderHeldDocument `bson:"senderHeldDocuments,omitempty"` //C.1.6.1.r 发送者提交的文件
	OtherCaseCode       []LinkedCode         `bson:"otherCaseCode"`                 //C.1.9.1.r -其他病例识别码
	LinkedReportCode    []string             `bson:"linkedReportCode"`              //C.1.10.r -与本报告相关的其他报告识别码
	Reporters           []E2bReporter        `bson:"reporters"`                     //C.2.r 报告者

}

type CaseIndentifier struct {
	//ICH
	SenderUniqueId               string `bson:"senderUniqueId,omitempty"`               //C.1.1
	CreateDate                   string `bson:"createDate,omitempty"`                   //C.1.2
	ReportType                   string `bson:"reportType,omitempty"`                   //C.1.3
	FirstReceiveDate             string `bson:"firstReceiveDate,omitempty"`             //C.1.4
	VisitReceiveDate             string `bson:"visitReceiveDate,omitempty"`             //C.1.5
	AdditionalDocumentsAvailable string `bson:"additionalDocumentsAvailable,omitempty"` //C.1.6.1
	LocalCriteria                string `bson:"localCriteria,omitempty"`                //C.1.7 此病例是否满足加速报告的当地标准？
	GlobalUniqueCaseId           string `bson:"globalUniqueCaseId,omitempty"`           //C.1.8.1 全球唯一病例识别码
	CaseFirstSender              int    `bson:"caseFirstSender,omitempty"`              //C.1.8.2 病例的首个发送者
	OtherCaseCodeAvailable       string `bson:"otherCaseCodeAvailable,omitempty"`       //其他病例识别码是否可用

	NullificationOrAmendment       int    `bson:"nullificationOrAmendment,omitempty"`       //C.1.11 报告作废/修正
	NullificationOrAmendmentReason string `bson:"nullificationOrAmendmentReason,omitempty"` //C.1.12 报告作废/修正原因

	//MFDS
	OtherStudiesType int `bson:"otherStudiesType,omitempty"` //C.5.4.KR.1

	//R2
	Country string `bson:"nullificationOrAmendmentReason,omitempty"` //A.1.1发生国家

	//CN区域字段
	ReportSource     string `bson:"reportSource,omitempty"`     // C.1.CN.1 报告来源
	ReportCategory   string `bson:"reportCategory,omitempty"`   // C.1.CN.2 报告分类
	HolderIdentifier string `bson:"holderIdentifier,omitempty"` // C.1.CN.3 持有人标识

	//FDA
	FdaCombinationProductReportIndicator string `bson:"fdaCombinationProductReportIndicator,omitempty"` //FDA.C.1.12
	FdaLocalCriteriaType                 string `bson:"fdaLocalCriteriaType,omitempty"`                 // FDA.C.1.7.1 Local Criteria Report Type

}

type SenderHeldDocument struct {
	FileId         string `bson:"fileId,omitempty"`         //文件在OSS中的ID
	Title          string `bson:"title,omitempty"`          //文件名,附件内容另行处理
	MediaType      string `bson:"mediaType,omitempty"`      //Cotent-Type
	Representation string `bson:"representation,omitempty"` //编码方式-一般是B64
	Compress       string `bson:"compress,omitempty"`       //压缩方式 一般是DF
}

type LinkedCode struct {
	Code   string `bson:"code,omitempty"`   //C.1.9.1.r.1 病例识别码的来源
	Source string `bson:"source,omitempty"` //C.1.9.1.r.2 病例识别码
}

type E2bReporter struct {
	Title                     string `bson:"title,omitempty"`                     //C.2.r.1.1 -报告者的称呼
	TitleNullFlavor           string `bson:"titleNullFlavor,omitempty"`           //C.2.r.1.1 -报告者的称呼 NF
	GivenName                 string `bson:"givenName,omitempty"`                 //C.2.r.1.2 -报告者的名字
	GivenNameNullFlavor       string `bson:"givenNameNullFlavor,omitempty"`       //C.2.r.1.2 -报告者的名字NF
	MiddleName                string `bson:"middleName,omitempty"`                //C.2.r.1.3 -报告者的中间名字
	MiddleNameNullFlavor      string `bson:"middleNameNullFlavor,omitempty"`      //C.2.r.1.3 -报告者的中间名字NF
	FamilyName                string `bson:"familyName,omitempty"`                //C.2.r.1.4 -报告者的姓氏
	FamilyNameNullFlavor      string `bson:"familyNameNullFlavor,omitempty"`      //C.2.r.1.4 -报告者的姓氏NF
	Organisation              string `bson:"organisation,omitempty"`              //C.2.r.2.1 -报告者所在机构NF
	OrganisationNullFlavor    string `bson:"organisationNullFlavor,omitempty"`    //C.2.r.2.1 -报告者所在机构
	Department                string `bson:"department,omitempty"`                //C.2.r.2.2 -报告者所在部门
	DepartmentNullFlavor      string `bson:"departmentNullFlavor,omitempty"`      //C.2.r.2.2 -报告者所在部门NF
	StreetAddress             string `bson:"streetAddress,omitempty"`             //C.2.r.2.3  -报告者所在街道地址
	StreetAddressNullFlavor   string `bson:"streetAddressNullFlavor,omitempty"`   //C.2.r.2.3  -报告者所在街道地址NF
	City                      string `bson:"city,omitempty"`                      //C.2.r.2.4 -报告者所在城市NF
	CityNullFlavor            string `bson:"cityNullFlavor,omitempty"`            //C.2.r.2.4 -报告者所在城市
	ProvinceOrState           string `bson:"provinceOrState,omitempty"`           //C.2.r.2.5 -报告者所在州或省
	ProvinceOrStateNullFlavor string `bson:"provinceOrStateNullFlavor,omitempty"` //C.2.r.2.5 -报告者所在州或省NF
	PostCode                  string `bson:"postCode,omitempty"`                  //C.2.r.2.6 -报告者所在地区的邮政编码
	PostCodeNullFlavor        string `bson:"postCodeNullFlavor,omitempty"`        //C.2.r.2.6 -报告者所在地区的邮政编码NF
	Telephone                 string `bson:"telephone,omitempty"`                 //C.2.r.2.7 -报告者的电话号码
	TelephoneNullFlavor       string `bson:"telephoneNullFlavor,omitempty"`       //C.2.r.2.7 -报告者的电话号码NF
	Mail                      string `bson:"mail,omitempty"`                      //C.2.r.CN.1 -电子邮箱
	CountryCode               string `bson:"countryCode,omitempty"`               //C.2.r.3  -报告者的国家代码
	Qualification             int    `bson:"qualification,omitempty"`             //C.2.r.4  -资质
	QualificationNullFlavor   string `bson:"qualificationNullFlavor,omitempty"`   //C.2.r.4  -资质NF
	PrimarySource             bool   `bson:"primarySource,omitempty"`             //C.2.r.2.1 -基于监管目的的主要来源
	//MFDS
	OtherMedicalExpertsClassification int `bson:"otherMedicalExpertsClassification,omitempty"` //C.2.r.4.KR.1

	//FDA
	FdaMail           string `bson:"fdaMail,omitempty"`           //FDA.C.2.r.2.8电子邮箱
	FdaMailNullFlavor string `bson:"fdaMailNullFlavor,omitempty"` //FDA.C.2.r.2.8-电子邮箱 NF
}

type E2bSender struct {
}
