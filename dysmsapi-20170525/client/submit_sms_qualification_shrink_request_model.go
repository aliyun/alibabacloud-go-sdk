// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmsQualificationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminIDCardExpDate(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminIDCardExpDate() *string
	SetAdminIDCardFrontFace(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminIDCardFrontFace() *string
	SetAdminIDCardNo(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminIDCardNo() *string
	SetAdminIDCardPic(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminIDCardPic() *string
	SetAdminIDCardType(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminIDCardType() *string
	SetAdminName(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminName() *string
	SetAdminPhoneNo(v string) *SubmitSmsQualificationShrinkRequest
	GetAdminPhoneNo() *string
	SetBusinessLicensePicsShrink(v string) *SubmitSmsQualificationShrinkRequest
	GetBusinessLicensePicsShrink() *string
	SetBussinessLicenseExpDate(v string) *SubmitSmsQualificationShrinkRequest
	GetBussinessLicenseExpDate() *string
	SetCertifyCode(v string) *SubmitSmsQualificationShrinkRequest
	GetCertifyCode() *string
	SetCompanyName(v string) *SubmitSmsQualificationShrinkRequest
	GetCompanyName() *string
	SetCompanyType(v string) *SubmitSmsQualificationShrinkRequest
	GetCompanyType() *string
	SetLegalPersonIDCardNo(v string) *SubmitSmsQualificationShrinkRequest
	GetLegalPersonIDCardNo() *string
	SetLegalPersonIDCardType(v string) *SubmitSmsQualificationShrinkRequest
	GetLegalPersonIDCardType() *string
	SetLegalPersonIdCardBackSide(v string) *SubmitSmsQualificationShrinkRequest
	GetLegalPersonIdCardBackSide() *string
	SetLegalPersonIdCardEffTime(v string) *SubmitSmsQualificationShrinkRequest
	GetLegalPersonIdCardEffTime() *string
	SetLegalPersonIdCardFrontSide(v string) *SubmitSmsQualificationShrinkRequest
	GetLegalPersonIdCardFrontSide() *string
	SetLegalPersonName(v string) *SubmitSmsQualificationShrinkRequest
	GetLegalPersonName() *string
	SetOrganizationCode(v string) *SubmitSmsQualificationShrinkRequest
	GetOrganizationCode() *string
	SetOtherFilesShrink(v string) *SubmitSmsQualificationShrinkRequest
	GetOtherFilesShrink() *string
	SetOwnerId(v int64) *SubmitSmsQualificationShrinkRequest
	GetOwnerId() *int64
	SetQualificationName(v string) *SubmitSmsQualificationShrinkRequest
	GetQualificationName() *string
	SetRemark(v string) *SubmitSmsQualificationShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *SubmitSmsQualificationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitSmsQualificationShrinkRequest
	GetResourceOwnerId() *int64
	SetUseBySelf(v bool) *SubmitSmsQualificationShrinkRequest
	GetUseBySelf() *bool
	SetWhetherShare(v bool) *SubmitSmsQualificationShrinkRequest
	GetWhetherShare() *bool
}

type SubmitSmsQualificationShrinkRequest struct {
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
	BusinessLicensePicsShrink *string `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty"`
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
	OtherFilesShrink *string `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s SubmitSmsQualificationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmsQualificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminIDCardExpDate() *string {
	return s.AdminIDCardExpDate
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminIDCardFrontFace() *string {
	return s.AdminIDCardFrontFace
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminIDCardNo() *string {
	return s.AdminIDCardNo
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminIDCardPic() *string {
	return s.AdminIDCardPic
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminIDCardType() *string {
	return s.AdminIDCardType
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminName() *string {
	return s.AdminName
}

func (s *SubmitSmsQualificationShrinkRequest) GetAdminPhoneNo() *string {
	return s.AdminPhoneNo
}

func (s *SubmitSmsQualificationShrinkRequest) GetBusinessLicensePicsShrink() *string {
	return s.BusinessLicensePicsShrink
}

func (s *SubmitSmsQualificationShrinkRequest) GetBussinessLicenseExpDate() *string {
	return s.BussinessLicenseExpDate
}

func (s *SubmitSmsQualificationShrinkRequest) GetCertifyCode() *string {
	return s.CertifyCode
}

func (s *SubmitSmsQualificationShrinkRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *SubmitSmsQualificationShrinkRequest) GetCompanyType() *string {
	return s.CompanyType
}

func (s *SubmitSmsQualificationShrinkRequest) GetLegalPersonIDCardNo() *string {
	return s.LegalPersonIDCardNo
}

func (s *SubmitSmsQualificationShrinkRequest) GetLegalPersonIDCardType() *string {
	return s.LegalPersonIDCardType
}

func (s *SubmitSmsQualificationShrinkRequest) GetLegalPersonIdCardBackSide() *string {
	return s.LegalPersonIdCardBackSide
}

func (s *SubmitSmsQualificationShrinkRequest) GetLegalPersonIdCardEffTime() *string {
	return s.LegalPersonIdCardEffTime
}

func (s *SubmitSmsQualificationShrinkRequest) GetLegalPersonIdCardFrontSide() *string {
	return s.LegalPersonIdCardFrontSide
}

func (s *SubmitSmsQualificationShrinkRequest) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *SubmitSmsQualificationShrinkRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *SubmitSmsQualificationShrinkRequest) GetOtherFilesShrink() *string {
	return s.OtherFilesShrink
}

func (s *SubmitSmsQualificationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitSmsQualificationShrinkRequest) GetQualificationName() *string {
	return s.QualificationName
}

func (s *SubmitSmsQualificationShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *SubmitSmsQualificationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitSmsQualificationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitSmsQualificationShrinkRequest) GetUseBySelf() *bool {
	return s.UseBySelf
}

func (s *SubmitSmsQualificationShrinkRequest) GetWhetherShare() *bool {
	return s.WhetherShare
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminIDCardExpDate(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminIDCardExpDate = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminIDCardFrontFace(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminIDCardFrontFace = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminIDCardNo(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminIDCardNo = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminIDCardPic(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminIDCardPic = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminIDCardType(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminIDCardType = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminName(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminName = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetAdminPhoneNo(v string) *SubmitSmsQualificationShrinkRequest {
	s.AdminPhoneNo = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetBusinessLicensePicsShrink(v string) *SubmitSmsQualificationShrinkRequest {
	s.BusinessLicensePicsShrink = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetBussinessLicenseExpDate(v string) *SubmitSmsQualificationShrinkRequest {
	s.BussinessLicenseExpDate = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetCertifyCode(v string) *SubmitSmsQualificationShrinkRequest {
	s.CertifyCode = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetCompanyName(v string) *SubmitSmsQualificationShrinkRequest {
	s.CompanyName = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetCompanyType(v string) *SubmitSmsQualificationShrinkRequest {
	s.CompanyType = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetLegalPersonIDCardNo(v string) *SubmitSmsQualificationShrinkRequest {
	s.LegalPersonIDCardNo = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetLegalPersonIDCardType(v string) *SubmitSmsQualificationShrinkRequest {
	s.LegalPersonIDCardType = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetLegalPersonIdCardBackSide(v string) *SubmitSmsQualificationShrinkRequest {
	s.LegalPersonIdCardBackSide = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetLegalPersonIdCardEffTime(v string) *SubmitSmsQualificationShrinkRequest {
	s.LegalPersonIdCardEffTime = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetLegalPersonIdCardFrontSide(v string) *SubmitSmsQualificationShrinkRequest {
	s.LegalPersonIdCardFrontSide = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetLegalPersonName(v string) *SubmitSmsQualificationShrinkRequest {
	s.LegalPersonName = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetOrganizationCode(v string) *SubmitSmsQualificationShrinkRequest {
	s.OrganizationCode = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetOtherFilesShrink(v string) *SubmitSmsQualificationShrinkRequest {
	s.OtherFilesShrink = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetOwnerId(v int64) *SubmitSmsQualificationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetQualificationName(v string) *SubmitSmsQualificationShrinkRequest {
	s.QualificationName = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetRemark(v string) *SubmitSmsQualificationShrinkRequest {
	s.Remark = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetResourceOwnerAccount(v string) *SubmitSmsQualificationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetResourceOwnerId(v int64) *SubmitSmsQualificationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetUseBySelf(v bool) *SubmitSmsQualificationShrinkRequest {
	s.UseBySelf = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) SetWhetherShare(v bool) *SubmitSmsQualificationShrinkRequest {
	s.WhetherShare = &v
	return s
}

func (s *SubmitSmsQualificationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
