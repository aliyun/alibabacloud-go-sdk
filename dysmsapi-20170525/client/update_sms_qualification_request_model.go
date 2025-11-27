// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsQualificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminIDCardExpDate(v string) *UpdateSmsQualificationRequest
	GetAdminIDCardExpDate() *string
	SetAdminIDCardFrontFace(v string) *UpdateSmsQualificationRequest
	GetAdminIDCardFrontFace() *string
	SetAdminIDCardNo(v string) *UpdateSmsQualificationRequest
	GetAdminIDCardNo() *string
	SetAdminIDCardPic(v string) *UpdateSmsQualificationRequest
	GetAdminIDCardPic() *string
	SetAdminIDCardType(v string) *UpdateSmsQualificationRequest
	GetAdminIDCardType() *string
	SetAdminName(v string) *UpdateSmsQualificationRequest
	GetAdminName() *string
	SetAdminPhoneNo(v string) *UpdateSmsQualificationRequest
	GetAdminPhoneNo() *string
	SetBusinessLicensePics(v []*UpdateSmsQualificationRequestBusinessLicensePics) *UpdateSmsQualificationRequest
	GetBusinessLicensePics() []*UpdateSmsQualificationRequestBusinessLicensePics
	SetBussinessLicenseExpDate(v string) *UpdateSmsQualificationRequest
	GetBussinessLicenseExpDate() *string
	SetCertifyCode(v string) *UpdateSmsQualificationRequest
	GetCertifyCode() *string
	SetCompanyName(v string) *UpdateSmsQualificationRequest
	GetCompanyName() *string
	SetLegalPersonIDCardNo(v string) *UpdateSmsQualificationRequest
	GetLegalPersonIDCardNo() *string
	SetLegalPersonIDCardType(v string) *UpdateSmsQualificationRequest
	GetLegalPersonIDCardType() *string
	SetLegalPersonIdCardBackSide(v string) *UpdateSmsQualificationRequest
	GetLegalPersonIdCardBackSide() *string
	SetLegalPersonIdCardEffTime(v string) *UpdateSmsQualificationRequest
	GetLegalPersonIdCardEffTime() *string
	SetLegalPersonIdCardFrontSide(v string) *UpdateSmsQualificationRequest
	GetLegalPersonIdCardFrontSide() *string
	SetLegalPersonName(v string) *UpdateSmsQualificationRequest
	GetLegalPersonName() *string
	SetOrderId(v int64) *UpdateSmsQualificationRequest
	GetOrderId() *int64
	SetOtherFiles(v []*UpdateSmsQualificationRequestOtherFiles) *UpdateSmsQualificationRequest
	GetOtherFiles() []*UpdateSmsQualificationRequestOtherFiles
	SetOwnerId(v int64) *UpdateSmsQualificationRequest
	GetOwnerId() *int64
	SetQualificationGroupId(v int64) *UpdateSmsQualificationRequest
	GetQualificationGroupId() *int64
	SetResourceOwnerAccount(v string) *UpdateSmsQualificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsQualificationRequest
	GetResourceOwnerId() *int64
}

type UpdateSmsQualificationRequest struct {
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
	BusinessLicensePics []*UpdateSmsQualificationRequestBusinessLicensePics `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty" type:"Repeated"`
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
	OtherFiles []*UpdateSmsQualificationRequestOtherFiles `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty" type:"Repeated"`
	OwnerId    *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s UpdateSmsQualificationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsQualificationRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsQualificationRequest) GetAdminIDCardExpDate() *string {
	return s.AdminIDCardExpDate
}

func (s *UpdateSmsQualificationRequest) GetAdminIDCardFrontFace() *string {
	return s.AdminIDCardFrontFace
}

func (s *UpdateSmsQualificationRequest) GetAdminIDCardNo() *string {
	return s.AdminIDCardNo
}

func (s *UpdateSmsQualificationRequest) GetAdminIDCardPic() *string {
	return s.AdminIDCardPic
}

func (s *UpdateSmsQualificationRequest) GetAdminIDCardType() *string {
	return s.AdminIDCardType
}

func (s *UpdateSmsQualificationRequest) GetAdminName() *string {
	return s.AdminName
}

func (s *UpdateSmsQualificationRequest) GetAdminPhoneNo() *string {
	return s.AdminPhoneNo
}

func (s *UpdateSmsQualificationRequest) GetBusinessLicensePics() []*UpdateSmsQualificationRequestBusinessLicensePics {
	return s.BusinessLicensePics
}

func (s *UpdateSmsQualificationRequest) GetBussinessLicenseExpDate() *string {
	return s.BussinessLicenseExpDate
}

func (s *UpdateSmsQualificationRequest) GetCertifyCode() *string {
	return s.CertifyCode
}

func (s *UpdateSmsQualificationRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *UpdateSmsQualificationRequest) GetLegalPersonIDCardNo() *string {
	return s.LegalPersonIDCardNo
}

func (s *UpdateSmsQualificationRequest) GetLegalPersonIDCardType() *string {
	return s.LegalPersonIDCardType
}

func (s *UpdateSmsQualificationRequest) GetLegalPersonIdCardBackSide() *string {
	return s.LegalPersonIdCardBackSide
}

func (s *UpdateSmsQualificationRequest) GetLegalPersonIdCardEffTime() *string {
	return s.LegalPersonIdCardEffTime
}

func (s *UpdateSmsQualificationRequest) GetLegalPersonIdCardFrontSide() *string {
	return s.LegalPersonIdCardFrontSide
}

func (s *UpdateSmsQualificationRequest) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *UpdateSmsQualificationRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpdateSmsQualificationRequest) GetOtherFiles() []*UpdateSmsQualificationRequestOtherFiles {
	return s.OtherFiles
}

func (s *UpdateSmsQualificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsQualificationRequest) GetQualificationGroupId() *int64 {
	return s.QualificationGroupId
}

func (s *UpdateSmsQualificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsQualificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsQualificationRequest) SetAdminIDCardExpDate(v string) *UpdateSmsQualificationRequest {
	s.AdminIDCardExpDate = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetAdminIDCardFrontFace(v string) *UpdateSmsQualificationRequest {
	s.AdminIDCardFrontFace = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetAdminIDCardNo(v string) *UpdateSmsQualificationRequest {
	s.AdminIDCardNo = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetAdminIDCardPic(v string) *UpdateSmsQualificationRequest {
	s.AdminIDCardPic = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetAdminIDCardType(v string) *UpdateSmsQualificationRequest {
	s.AdminIDCardType = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetAdminName(v string) *UpdateSmsQualificationRequest {
	s.AdminName = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetAdminPhoneNo(v string) *UpdateSmsQualificationRequest {
	s.AdminPhoneNo = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetBusinessLicensePics(v []*UpdateSmsQualificationRequestBusinessLicensePics) *UpdateSmsQualificationRequest {
	s.BusinessLicensePics = v
	return s
}

func (s *UpdateSmsQualificationRequest) SetBussinessLicenseExpDate(v string) *UpdateSmsQualificationRequest {
	s.BussinessLicenseExpDate = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetCertifyCode(v string) *UpdateSmsQualificationRequest {
	s.CertifyCode = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetCompanyName(v string) *UpdateSmsQualificationRequest {
	s.CompanyName = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetLegalPersonIDCardNo(v string) *UpdateSmsQualificationRequest {
	s.LegalPersonIDCardNo = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetLegalPersonIDCardType(v string) *UpdateSmsQualificationRequest {
	s.LegalPersonIDCardType = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetLegalPersonIdCardBackSide(v string) *UpdateSmsQualificationRequest {
	s.LegalPersonIdCardBackSide = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetLegalPersonIdCardEffTime(v string) *UpdateSmsQualificationRequest {
	s.LegalPersonIdCardEffTime = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetLegalPersonIdCardFrontSide(v string) *UpdateSmsQualificationRequest {
	s.LegalPersonIdCardFrontSide = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetLegalPersonName(v string) *UpdateSmsQualificationRequest {
	s.LegalPersonName = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetOrderId(v int64) *UpdateSmsQualificationRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetOtherFiles(v []*UpdateSmsQualificationRequestOtherFiles) *UpdateSmsQualificationRequest {
	s.OtherFiles = v
	return s
}

func (s *UpdateSmsQualificationRequest) SetOwnerId(v int64) *UpdateSmsQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetQualificationGroupId(v int64) *UpdateSmsQualificationRequest {
	s.QualificationGroupId = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetResourceOwnerAccount(v string) *UpdateSmsQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsQualificationRequest) SetResourceOwnerId(v int64) *UpdateSmsQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsQualificationRequest) Validate() error {
	if s.BusinessLicensePics != nil {
		for _, item := range s.BusinessLicensePics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OtherFiles != nil {
		for _, item := range s.OtherFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSmsQualificationRequestBusinessLicensePics struct {
	// 证件图片标识的osskey
	//
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
	// 企业证件类型，businessLicense:营业执照;organizationCodeLicense:组织机构代码证;taxRegistrationLicense:税务登记证;socialCreditLicense:社会信用代码证书;newStyleBusinessLicense:三证合一;signLegalLicense:签名归属方的事业单位法人证书;otherLicense:其他类型执照证书
	//
	// example:
	//
	// businessLicense
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateSmsQualificationRequestBusinessLicensePics) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsQualificationRequestBusinessLicensePics) GoString() string {
	return s.String()
}

func (s *UpdateSmsQualificationRequestBusinessLicensePics) GetLicensePic() *string {
	return s.LicensePic
}

func (s *UpdateSmsQualificationRequestBusinessLicensePics) GetType() *string {
	return s.Type
}

func (s *UpdateSmsQualificationRequestBusinessLicensePics) SetLicensePic(v string) *UpdateSmsQualificationRequestBusinessLicensePics {
	s.LicensePic = &v
	return s
}

func (s *UpdateSmsQualificationRequestBusinessLicensePics) SetType(v string) *UpdateSmsQualificationRequestBusinessLicensePics {
	s.Type = &v
	return s
}

func (s *UpdateSmsQualificationRequestBusinessLicensePics) Validate() error {
	return dara.Validate(s)
}

type UpdateSmsQualificationRequestOtherFiles struct {
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
}

func (s UpdateSmsQualificationRequestOtherFiles) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsQualificationRequestOtherFiles) GoString() string {
	return s.String()
}

func (s *UpdateSmsQualificationRequestOtherFiles) GetLicensePic() *string {
	return s.LicensePic
}

func (s *UpdateSmsQualificationRequestOtherFiles) SetLicensePic(v string) *UpdateSmsQualificationRequestOtherFiles {
	s.LicensePic = &v
	return s
}

func (s *UpdateSmsQualificationRequestOtherFiles) Validate() error {
	return dara.Validate(s)
}
