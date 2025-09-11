// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmsQualificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminIDCardExpDate(v string) *SubmitSmsQualificationRequest
	GetAdminIDCardExpDate() *string
	SetAdminIDCardFrontFace(v string) *SubmitSmsQualificationRequest
	GetAdminIDCardFrontFace() *string
	SetAdminIDCardNo(v string) *SubmitSmsQualificationRequest
	GetAdminIDCardNo() *string
	SetAdminIDCardPic(v string) *SubmitSmsQualificationRequest
	GetAdminIDCardPic() *string
	SetAdminIDCardType(v string) *SubmitSmsQualificationRequest
	GetAdminIDCardType() *string
	SetAdminName(v string) *SubmitSmsQualificationRequest
	GetAdminName() *string
	SetAdminPhoneNo(v string) *SubmitSmsQualificationRequest
	GetAdminPhoneNo() *string
	SetBusinessLicensePics(v []*SubmitSmsQualificationRequestBusinessLicensePics) *SubmitSmsQualificationRequest
	GetBusinessLicensePics() []*SubmitSmsQualificationRequestBusinessLicensePics
	SetBussinessLicenseExpDate(v string) *SubmitSmsQualificationRequest
	GetBussinessLicenseExpDate() *string
	SetCertifyCode(v string) *SubmitSmsQualificationRequest
	GetCertifyCode() *string
	SetCompanyName(v string) *SubmitSmsQualificationRequest
	GetCompanyName() *string
	SetCompanyType(v string) *SubmitSmsQualificationRequest
	GetCompanyType() *string
	SetLegalPersonIDCardNo(v string) *SubmitSmsQualificationRequest
	GetLegalPersonIDCardNo() *string
	SetLegalPersonIDCardType(v string) *SubmitSmsQualificationRequest
	GetLegalPersonIDCardType() *string
	SetLegalPersonIdCardBackSide(v string) *SubmitSmsQualificationRequest
	GetLegalPersonIdCardBackSide() *string
	SetLegalPersonIdCardEffTime(v string) *SubmitSmsQualificationRequest
	GetLegalPersonIdCardEffTime() *string
	SetLegalPersonIdCardFrontSide(v string) *SubmitSmsQualificationRequest
	GetLegalPersonIdCardFrontSide() *string
	SetLegalPersonName(v string) *SubmitSmsQualificationRequest
	GetLegalPersonName() *string
	SetOrganizationCode(v string) *SubmitSmsQualificationRequest
	GetOrganizationCode() *string
	SetOtherFiles(v []*SubmitSmsQualificationRequestOtherFiles) *SubmitSmsQualificationRequest
	GetOtherFiles() []*SubmitSmsQualificationRequestOtherFiles
	SetOwnerId(v int64) *SubmitSmsQualificationRequest
	GetOwnerId() *int64
	SetQualificationName(v string) *SubmitSmsQualificationRequest
	GetQualificationName() *string
	SetRemark(v string) *SubmitSmsQualificationRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *SubmitSmsQualificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitSmsQualificationRequest
	GetResourceOwnerId() *int64
	SetUseBySelf(v bool) *SubmitSmsQualificationRequest
	GetUseBySelf() *bool
	SetWhetherShare(v bool) *SubmitSmsQualificationRequest
	GetWhetherShare() *bool
}

type SubmitSmsQualificationRequest struct {
	// 经办人身份证有效期，格式示例2023-01-01~2033-01-01
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	AdminIDCardExpDate *string `json:"AdminIDCardExpDate,omitempty" xml:"AdminIDCardExpDate,omitempty"`
	// 经办人身份证照片国徽面
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardFrontFace *string `json:"AdminIDCardFrontFace,omitempty" xml:"AdminIDCardFrontFace,omitempty"`
	// 经办人身份证号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 511391********5123
	AdminIDCardNo *string `json:"AdminIDCardNo,omitempty" xml:"AdminIDCardNo,omitempty"`
	// 经办人身份证照片人像面
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardPic *string `json:"AdminIDCardPic,omitempty" xml:"AdminIDCardPic,omitempty"`
	// 管理员身份证类型。identityCard:中国居民身份证;passport:护照;homeReturnPermit:港澳居民来往内地通行证;TaiwanCompatriotPermit:台湾居民来往大陆通行证;residencePermit:港澳台居民居住证";other:其他
	//
	// This parameter is required.
	//
	// example:
	//
	// identityCard
	AdminIDCardType *string `json:"AdminIDCardType,omitempty" xml:"AdminIDCardType,omitempty"`
	// 经办人姓名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	AdminName *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	// 经办人手机号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****1234
	AdminPhoneNo *string `json:"AdminPhoneNo,omitempty" xml:"AdminPhoneNo,omitempty"`
	// 企业营业证件信息，非大客户必填
	BusinessLicensePics []*SubmitSmsQualificationRequestBusinessLicensePics `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty" type:"Repeated"`
	// 企业营业时间开始和结束字符串，格式示例2023-01-01~2033-01-01
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	BussinessLicenseExpDate *string `json:"BussinessLicenseExpDate,omitempty" xml:"BussinessLicenseExpDate,omitempty"`
	// 手机号验证码
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	CertifyCode *string `json:"CertifyCode,omitempty" xml:"CertifyCode,omitempty"`
	// 公司名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// 企业类型, COMPANY:公司;NON_PROFIT_ORGANIZATION:政府或者事业单位
	//
	// This parameter is required.
	//
	// example:
	//
	// COMPANY
	CompanyType *string `json:"CompanyType,omitempty" xml:"CompanyType,omitempty"`
	// 法人身份证号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 511391********5123
	LegalPersonIDCardNo *string `json:"LegalPersonIDCardNo,omitempty" xml:"LegalPersonIDCardNo,omitempty"`
	// 法人身份证类型。identityCard:中国居民身份证;passport:护照;homeReturnPermit:港澳居民来往内地通行证;TaiwanCompatriotPermit:台湾居民来往大陆通行证;residencePermit:港澳台居民居住证";other:其他
	//
	// This parameter is required.
	//
	// example:
	//
	// identityCard
	LegalPersonIDCardType *string `json:"LegalPersonIDCardType,omitempty" xml:"LegalPersonIDCardType,omitempty"`
	// 法人身份证照片人像面
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardBackSide *string `json:"LegalPersonIdCardBackSide,omitempty" xml:"LegalPersonIdCardBackSide,omitempty"`
	// 法人身份证有效期
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	LegalPersonIdCardEffTime *string `json:"LegalPersonIdCardEffTime,omitempty" xml:"LegalPersonIdCardEffTime,omitempty"`
	// 法人身份证照片国徽面
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardFrontSide *string `json:"LegalPersonIdCardFrontSide,omitempty" xml:"LegalPersonIdCardFrontSide,omitempty"`
	// 法人姓名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// 社会统一信用代码
	//
	// This parameter is required.
	//
	// example:
	//
	// 910X********0012
	OrganizationCode *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	// 更多资料
	OtherFiles []*SubmitSmsQualificationRequestOtherFiles `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty" type:"Repeated"`
	OwnerId    *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 资质名称,名称不能重复
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值示例值
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 资质是自用还是他用，true：自用，false：他用
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	UseBySelf *bool `json:"UseBySelf,omitempty" xml:"UseBySelf,omitempty"`
	// 是否同意与其他业务线共享
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	WhetherShare *bool `json:"WhetherShare,omitempty" xml:"WhetherShare,omitempty"`
}

func (s SubmitSmsQualificationRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmsQualificationRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmsQualificationRequest) GetAdminIDCardExpDate() *string {
	return s.AdminIDCardExpDate
}

func (s *SubmitSmsQualificationRequest) GetAdminIDCardFrontFace() *string {
	return s.AdminIDCardFrontFace
}

func (s *SubmitSmsQualificationRequest) GetAdminIDCardNo() *string {
	return s.AdminIDCardNo
}

func (s *SubmitSmsQualificationRequest) GetAdminIDCardPic() *string {
	return s.AdminIDCardPic
}

func (s *SubmitSmsQualificationRequest) GetAdminIDCardType() *string {
	return s.AdminIDCardType
}

func (s *SubmitSmsQualificationRequest) GetAdminName() *string {
	return s.AdminName
}

func (s *SubmitSmsQualificationRequest) GetAdminPhoneNo() *string {
	return s.AdminPhoneNo
}

func (s *SubmitSmsQualificationRequest) GetBusinessLicensePics() []*SubmitSmsQualificationRequestBusinessLicensePics {
	return s.BusinessLicensePics
}

func (s *SubmitSmsQualificationRequest) GetBussinessLicenseExpDate() *string {
	return s.BussinessLicenseExpDate
}

func (s *SubmitSmsQualificationRequest) GetCertifyCode() *string {
	return s.CertifyCode
}

func (s *SubmitSmsQualificationRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *SubmitSmsQualificationRequest) GetCompanyType() *string {
	return s.CompanyType
}

func (s *SubmitSmsQualificationRequest) GetLegalPersonIDCardNo() *string {
	return s.LegalPersonIDCardNo
}

func (s *SubmitSmsQualificationRequest) GetLegalPersonIDCardType() *string {
	return s.LegalPersonIDCardType
}

func (s *SubmitSmsQualificationRequest) GetLegalPersonIdCardBackSide() *string {
	return s.LegalPersonIdCardBackSide
}

func (s *SubmitSmsQualificationRequest) GetLegalPersonIdCardEffTime() *string {
	return s.LegalPersonIdCardEffTime
}

func (s *SubmitSmsQualificationRequest) GetLegalPersonIdCardFrontSide() *string {
	return s.LegalPersonIdCardFrontSide
}

func (s *SubmitSmsQualificationRequest) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *SubmitSmsQualificationRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *SubmitSmsQualificationRequest) GetOtherFiles() []*SubmitSmsQualificationRequestOtherFiles {
	return s.OtherFiles
}

func (s *SubmitSmsQualificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitSmsQualificationRequest) GetQualificationName() *string {
	return s.QualificationName
}

func (s *SubmitSmsQualificationRequest) GetRemark() *string {
	return s.Remark
}

func (s *SubmitSmsQualificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitSmsQualificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitSmsQualificationRequest) GetUseBySelf() *bool {
	return s.UseBySelf
}

func (s *SubmitSmsQualificationRequest) GetWhetherShare() *bool {
	return s.WhetherShare
}

func (s *SubmitSmsQualificationRequest) SetAdminIDCardExpDate(v string) *SubmitSmsQualificationRequest {
	s.AdminIDCardExpDate = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetAdminIDCardFrontFace(v string) *SubmitSmsQualificationRequest {
	s.AdminIDCardFrontFace = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetAdminIDCardNo(v string) *SubmitSmsQualificationRequest {
	s.AdminIDCardNo = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetAdminIDCardPic(v string) *SubmitSmsQualificationRequest {
	s.AdminIDCardPic = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetAdminIDCardType(v string) *SubmitSmsQualificationRequest {
	s.AdminIDCardType = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetAdminName(v string) *SubmitSmsQualificationRequest {
	s.AdminName = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetAdminPhoneNo(v string) *SubmitSmsQualificationRequest {
	s.AdminPhoneNo = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetBusinessLicensePics(v []*SubmitSmsQualificationRequestBusinessLicensePics) *SubmitSmsQualificationRequest {
	s.BusinessLicensePics = v
	return s
}

func (s *SubmitSmsQualificationRequest) SetBussinessLicenseExpDate(v string) *SubmitSmsQualificationRequest {
	s.BussinessLicenseExpDate = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetCertifyCode(v string) *SubmitSmsQualificationRequest {
	s.CertifyCode = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetCompanyName(v string) *SubmitSmsQualificationRequest {
	s.CompanyName = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetCompanyType(v string) *SubmitSmsQualificationRequest {
	s.CompanyType = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetLegalPersonIDCardNo(v string) *SubmitSmsQualificationRequest {
	s.LegalPersonIDCardNo = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetLegalPersonIDCardType(v string) *SubmitSmsQualificationRequest {
	s.LegalPersonIDCardType = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetLegalPersonIdCardBackSide(v string) *SubmitSmsQualificationRequest {
	s.LegalPersonIdCardBackSide = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetLegalPersonIdCardEffTime(v string) *SubmitSmsQualificationRequest {
	s.LegalPersonIdCardEffTime = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetLegalPersonIdCardFrontSide(v string) *SubmitSmsQualificationRequest {
	s.LegalPersonIdCardFrontSide = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetLegalPersonName(v string) *SubmitSmsQualificationRequest {
	s.LegalPersonName = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetOrganizationCode(v string) *SubmitSmsQualificationRequest {
	s.OrganizationCode = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetOtherFiles(v []*SubmitSmsQualificationRequestOtherFiles) *SubmitSmsQualificationRequest {
	s.OtherFiles = v
	return s
}

func (s *SubmitSmsQualificationRequest) SetOwnerId(v int64) *SubmitSmsQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetQualificationName(v string) *SubmitSmsQualificationRequest {
	s.QualificationName = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetRemark(v string) *SubmitSmsQualificationRequest {
	s.Remark = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetResourceOwnerAccount(v string) *SubmitSmsQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetResourceOwnerId(v int64) *SubmitSmsQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetUseBySelf(v bool) *SubmitSmsQualificationRequest {
	s.UseBySelf = &v
	return s
}

func (s *SubmitSmsQualificationRequest) SetWhetherShare(v bool) *SubmitSmsQualificationRequest {
	s.WhetherShare = &v
	return s
}

func (s *SubmitSmsQualificationRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitSmsQualificationRequestBusinessLicensePics struct {
	// 营业证件图片标识的osskey
	//
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
	// 营业证件类型，businessLicense:营业执照;organizationCodeLicense:组织机构代码证;taxRegistrationLicense:税务登记证;socialCreditLicense:社会信用代码证书;newStyleBusinessLicense:三证合一;signLegalLicense:签名归属方的事业单位法人证书;otherLicense:其他类型执照证书
	//
	// example:
	//
	// businessLicense
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSmsQualificationRequestBusinessLicensePics) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmsQualificationRequestBusinessLicensePics) GoString() string {
	return s.String()
}

func (s *SubmitSmsQualificationRequestBusinessLicensePics) GetLicensePic() *string {
	return s.LicensePic
}

func (s *SubmitSmsQualificationRequestBusinessLicensePics) GetType() *string {
	return s.Type
}

func (s *SubmitSmsQualificationRequestBusinessLicensePics) SetLicensePic(v string) *SubmitSmsQualificationRequestBusinessLicensePics {
	s.LicensePic = &v
	return s
}

func (s *SubmitSmsQualificationRequestBusinessLicensePics) SetType(v string) *SubmitSmsQualificationRequestBusinessLicensePics {
	s.Type = &v
	return s
}

func (s *SubmitSmsQualificationRequestBusinessLicensePics) Validate() error {
	return dara.Validate(s)
}

type SubmitSmsQualificationRequestOtherFiles struct {
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
}

func (s SubmitSmsQualificationRequestOtherFiles) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmsQualificationRequestOtherFiles) GoString() string {
	return s.String()
}

func (s *SubmitSmsQualificationRequestOtherFiles) GetLicensePic() *string {
	return s.LicensePic
}

func (s *SubmitSmsQualificationRequestOtherFiles) SetLicensePic(v string) *SubmitSmsQualificationRequestOtherFiles {
	s.LicensePic = &v
	return s
}

func (s *SubmitSmsQualificationRequestOtherFiles) Validate() error {
	return dara.Validate(s)
}
