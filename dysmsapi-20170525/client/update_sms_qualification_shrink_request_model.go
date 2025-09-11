// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsQualificationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminIDCardExpDate(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminIDCardExpDate() *string
	SetAdminIDCardFrontFace(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminIDCardFrontFace() *string
	SetAdminIDCardNo(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminIDCardNo() *string
	SetAdminIDCardPic(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminIDCardPic() *string
	SetAdminIDCardType(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminIDCardType() *string
	SetAdminName(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminName() *string
	SetAdminPhoneNo(v string) *UpdateSmsQualificationShrinkRequest
	GetAdminPhoneNo() *string
	SetBusinessLicensePicsShrink(v string) *UpdateSmsQualificationShrinkRequest
	GetBusinessLicensePicsShrink() *string
	SetBussinessLicenseExpDate(v string) *UpdateSmsQualificationShrinkRequest
	GetBussinessLicenseExpDate() *string
	SetCertifyCode(v string) *UpdateSmsQualificationShrinkRequest
	GetCertifyCode() *string
	SetCompanyName(v string) *UpdateSmsQualificationShrinkRequest
	GetCompanyName() *string
	SetLegalPersonIDCardNo(v string) *UpdateSmsQualificationShrinkRequest
	GetLegalPersonIDCardNo() *string
	SetLegalPersonIDCardType(v string) *UpdateSmsQualificationShrinkRequest
	GetLegalPersonIDCardType() *string
	SetLegalPersonIdCardBackSide(v string) *UpdateSmsQualificationShrinkRequest
	GetLegalPersonIdCardBackSide() *string
	SetLegalPersonIdCardEffTime(v string) *UpdateSmsQualificationShrinkRequest
	GetLegalPersonIdCardEffTime() *string
	SetLegalPersonIdCardFrontSide(v string) *UpdateSmsQualificationShrinkRequest
	GetLegalPersonIdCardFrontSide() *string
	SetLegalPersonName(v string) *UpdateSmsQualificationShrinkRequest
	GetLegalPersonName() *string
	SetOrderId(v int64) *UpdateSmsQualificationShrinkRequest
	GetOrderId() *int64
	SetOtherFilesShrink(v string) *UpdateSmsQualificationShrinkRequest
	GetOtherFilesShrink() *string
	SetOwnerId(v int64) *UpdateSmsQualificationShrinkRequest
	GetOwnerId() *int64
	SetQualificationGroupId(v int64) *UpdateSmsQualificationShrinkRequest
	GetQualificationGroupId() *int64
	SetResourceOwnerAccount(v string) *UpdateSmsQualificationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsQualificationShrinkRequest
	GetResourceOwnerId() *int64
}

type UpdateSmsQualificationShrinkRequest struct {
	// 经办人身份证有效期，格式示例2023-01-01~2033-01-01
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	AdminIDCardExpDate *string `json:"AdminIDCardExpDate,omitempty" xml:"AdminIDCardExpDate,omitempty"`
	// 经办人身份证照片国徽面
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardFrontFace *string `json:"AdminIDCardFrontFace,omitempty" xml:"AdminIDCardFrontFace,omitempty"`
	// 经办人身份证号码
	//
	// example:
	//
	// 511391********5123
	AdminIDCardNo *string `json:"AdminIDCardNo,omitempty" xml:"AdminIDCardNo,omitempty"`
	// 经办人身份证照片人像面
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardPic *string `json:"AdminIDCardPic,omitempty" xml:"AdminIDCardPic,omitempty"`
	// 管理员身份证类型。identityCard:中国居民身份证;passport:护照;homeReturnPermit:港澳居民来往内地通行证;TaiwanCompatriotPermit:台湾居民来往大陆通行证;residencePermit:港澳台居民居住证";other:其他
	//
	// example:
	//
	// identityCard
	AdminIDCardType *string `json:"AdminIDCardType,omitempty" xml:"AdminIDCardType,omitempty"`
	// 经办人姓名
	//
	// example:
	//
	// 示例值
	AdminName *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	// 经办人手机号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 137********
	AdminPhoneNo *string `json:"AdminPhoneNo,omitempty" xml:"AdminPhoneNo,omitempty"`
	// 企业证件信息
	BusinessLicensePicsShrink *string `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty"`
	// 企业营业时间开始和结束字符串，格式示例2023-01-01~2033-01-01
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
	// example:
	//
	// 示例值示例值示例值
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// 法人身份证号码
	//
	// example:
	//
	// 511391********5123
	LegalPersonIDCardNo *string `json:"LegalPersonIDCardNo,omitempty" xml:"LegalPersonIDCardNo,omitempty"`
	// 法人身份证类型。identityCard:中国居民身份证;passport:护照;homeReturnPermit:港澳居民来往内地通行证;TaiwanCompatriotPermit:台湾居民来往大陆通行证;residencePermit:港澳台居民居住证";other:其他
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
	// 法人身份证有效期，格式示例2023-01-01~2033-01-01
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	LegalPersonIdCardEffTime *string `json:"LegalPersonIdCardEffTime,omitempty" xml:"LegalPersonIdCardEffTime,omitempty"`
	// 法人身份照片证国徽面
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardFrontSide *string `json:"LegalPersonIdCardFrontSide,omitempty" xml:"LegalPersonIdCardFrontSide,omitempty"`
	// 法人姓名
	//
	// example:
	//
	// 示例值示例值
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// 工单ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2001*****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 更多资料
	OtherFilesShrink *string `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 资质组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000****
	QualificationGroupId *int64  `json:"QualificationGroupId,omitempty" xml:"QualificationGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateSmsQualificationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsQualificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminIDCardExpDate() *string {
	return s.AdminIDCardExpDate
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminIDCardFrontFace() *string {
	return s.AdminIDCardFrontFace
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminIDCardNo() *string {
	return s.AdminIDCardNo
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminIDCardPic() *string {
	return s.AdminIDCardPic
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminIDCardType() *string {
	return s.AdminIDCardType
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminName() *string {
	return s.AdminName
}

func (s *UpdateSmsQualificationShrinkRequest) GetAdminPhoneNo() *string {
	return s.AdminPhoneNo
}

func (s *UpdateSmsQualificationShrinkRequest) GetBusinessLicensePicsShrink() *string {
	return s.BusinessLicensePicsShrink
}

func (s *UpdateSmsQualificationShrinkRequest) GetBussinessLicenseExpDate() *string {
	return s.BussinessLicenseExpDate
}

func (s *UpdateSmsQualificationShrinkRequest) GetCertifyCode() *string {
	return s.CertifyCode
}

func (s *UpdateSmsQualificationShrinkRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *UpdateSmsQualificationShrinkRequest) GetLegalPersonIDCardNo() *string {
	return s.LegalPersonIDCardNo
}

func (s *UpdateSmsQualificationShrinkRequest) GetLegalPersonIDCardType() *string {
	return s.LegalPersonIDCardType
}

func (s *UpdateSmsQualificationShrinkRequest) GetLegalPersonIdCardBackSide() *string {
	return s.LegalPersonIdCardBackSide
}

func (s *UpdateSmsQualificationShrinkRequest) GetLegalPersonIdCardEffTime() *string {
	return s.LegalPersonIdCardEffTime
}

func (s *UpdateSmsQualificationShrinkRequest) GetLegalPersonIdCardFrontSide() *string {
	return s.LegalPersonIdCardFrontSide
}

func (s *UpdateSmsQualificationShrinkRequest) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *UpdateSmsQualificationShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpdateSmsQualificationShrinkRequest) GetOtherFilesShrink() *string {
	return s.OtherFilesShrink
}

func (s *UpdateSmsQualificationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsQualificationShrinkRequest) GetQualificationGroupId() *int64 {
	return s.QualificationGroupId
}

func (s *UpdateSmsQualificationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsQualificationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminIDCardExpDate(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminIDCardExpDate = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminIDCardFrontFace(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminIDCardFrontFace = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminIDCardNo(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminIDCardNo = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminIDCardPic(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminIDCardPic = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminIDCardType(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminIDCardType = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminName(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminName = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetAdminPhoneNo(v string) *UpdateSmsQualificationShrinkRequest {
	s.AdminPhoneNo = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetBusinessLicensePicsShrink(v string) *UpdateSmsQualificationShrinkRequest {
	s.BusinessLicensePicsShrink = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetBussinessLicenseExpDate(v string) *UpdateSmsQualificationShrinkRequest {
	s.BussinessLicenseExpDate = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetCertifyCode(v string) *UpdateSmsQualificationShrinkRequest {
	s.CertifyCode = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetCompanyName(v string) *UpdateSmsQualificationShrinkRequest {
	s.CompanyName = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetLegalPersonIDCardNo(v string) *UpdateSmsQualificationShrinkRequest {
	s.LegalPersonIDCardNo = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetLegalPersonIDCardType(v string) *UpdateSmsQualificationShrinkRequest {
	s.LegalPersonIDCardType = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetLegalPersonIdCardBackSide(v string) *UpdateSmsQualificationShrinkRequest {
	s.LegalPersonIdCardBackSide = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetLegalPersonIdCardEffTime(v string) *UpdateSmsQualificationShrinkRequest {
	s.LegalPersonIdCardEffTime = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetLegalPersonIdCardFrontSide(v string) *UpdateSmsQualificationShrinkRequest {
	s.LegalPersonIdCardFrontSide = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetLegalPersonName(v string) *UpdateSmsQualificationShrinkRequest {
	s.LegalPersonName = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetOrderId(v int64) *UpdateSmsQualificationShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetOtherFilesShrink(v string) *UpdateSmsQualificationShrinkRequest {
	s.OtherFilesShrink = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetOwnerId(v int64) *UpdateSmsQualificationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetQualificationGroupId(v int64) *UpdateSmsQualificationShrinkRequest {
	s.QualificationGroupId = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetResourceOwnerAccount(v string) *UpdateSmsQualificationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) SetResourceOwnerId(v int64) *UpdateSmsQualificationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsQualificationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
