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
	// Validity period of the administrator ID card. Format: YYYY-MM-DD~YYYY-MM-DD.
	//
	// > When the certificate validity period is long-term, the end date can be set to 2099-12-31.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	AdminIDCardExpDate *string `json:"AdminIDCardExpDate,omitempty" xml:"AdminIDCardExpDate,omitempty"`
	// Photo of the front of the administrator\\"s ID card (national emblem side). Only jpg, png, gif, and jpeg image formats are supported, and the image must not exceed 5 MB. Please provide the path of the file uploaded to OSS. The file name to be uploaded must not contain Chinese characters or special characters. For upload operations, see [Upload Files via OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// 	Notice:
	//
	// No stamp is required for color originals of the certificate. If you upload a photocopy or black-and-white photo, you must affix the enterprise red seal to the photocopy and take a photo to upload.
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardFrontFace *string `json:"AdminIDCardFrontFace,omitempty" xml:"AdminIDCardFrontFace,omitempty"`
	// Administrator\\"s ID number.
	//
	// example:
	//
	// 511391********5123
	AdminIDCardNo *string `json:"AdminIDCardNo,omitempty" xml:"AdminIDCardNo,omitempty"`
	// Photo of the back of the administrator\\"s ID card (portrait side). Only jpg, png, gif, and jpeg image formats are supported, and the image must not exceed 5 MB. Please provide the path of the file uploaded to OSS. The file name to be uploaded must not contain Chinese characters or special characters. For upload operations, see [Upload Files via OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// 	Notice:
	//
	//  No stamp is required for color originals of the certificate. If you upload a photocopy or black-and-white photo, you must affix the enterprise red seal to the photocopy and take a photo to upload.
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardPic *string `json:"AdminIDCardPic,omitempty" xml:"AdminIDCardPic,omitempty"`
	// Administrator ID card type. Valid values:
	//
	// - identityCard: ID card.
	//
	// - passport: Passport.
	//
	// - homeReturnPermit: Mainland Travel Permit for Hong Kong and Macao Residents.
	//
	// - TaiwanCompatriotPermit: Mainland Travel Permit for Taiwan Residents.
	//
	// - residencePermit: Residence Permit for Hong Kong, Macao, and Taiwan Residents.
	//
	// - other: Other.
	//
	// example:
	//
	// identityCard
	AdminIDCardType *string `json:"AdminIDCardType,omitempty" xml:"AdminIDCardType,omitempty"`
	// Administrator name.
	//
	// > The administrator (also known as the operator) refers to the person who logs in to the Alibaba Cloud account and manages the SMS service. Generally, this is the operations personnel who manages qualifications, signatures, and templates and sends SMS messages under this Alibaba Cloud account, and whose phone number can receive verification codes. The administrator is not necessarily the administrator of this Alibaba Cloud account. The administrator can be the same person as the enterprise\\"s legal representative.
	//
	// example:
	//
	// 李华
	AdminName *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	// Administrator\\"s mobile phone number. Format: +/+86/0086/86 prefix or a phone number without any prefix, for example, 1390000****.
	//
	// This parameter is required.
	//
	// example:
	//
	// 137********
	AdminPhoneNo *string `json:"AdminPhoneNo,omitempty" xml:"AdminPhoneNo,omitempty"`
	// Enterprise business license information. This parameter is required when the purpose of the qualification to be modified is for use by others.
	BusinessLicensePicsShrink *string `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty"`
	// Validity period of the business license. Format: YYYY-MM-DD~YYYY-MM-DD.
	//
	// > When the certificate validity period is long-term, the end date can be set to 2099-12-31.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	BussinessLicenseExpDate *string `json:"BussinessLicenseExpDate,omitempty" xml:"BussinessLicenseExpDate,omitempty"`
	// Phone verification code. Please call the [RequiredPhoneCode](~~RequiredPhoneCode~~) API and pass in the **administrator\\"s phone number**, then enter the SMS verification code you receive here.
	//
	// > You can use [ValidPhoneCode](~~ValidPhoneCode~~) to verify whether the SMS verification code is correct before passing it in.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	CertifyCode *string `json:"CertifyCode,omitempty" xml:"CertifyCode,omitempty"`
	// Enterprise name. Supported symbols are only the middle dot `·`, the Chinese symbols `【】（）`, the English symbols `()`, and the `space`. Other symbols or pure numbers are not allowed. The length must not exceed 150 characters.
	//
	// example:
	//
	// 阿里云云通信有限公司
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// Legal person\\"s ID number.
	//
	// example:
	//
	// 511391********5123
	LegalPersonIDCardNo *string `json:"LegalPersonIDCardNo,omitempty" xml:"LegalPersonIDCardNo,omitempty"`
	// Legal person ID card type. Valid values:
	//
	// - identityCard: ID card.
	//
	// - passport: Passport.
	//
	// - homeReturnPermit: Mainland Travel Permit for Hong Kong and Macao Residents.
	//
	// - TaiwanCompatriotPermit: Mainland Travel Permit for Taiwan Residents.
	//
	// - residencePermit: Residence Permit for Hong Kong, Macao, and Taiwan Residents.
	//
	// - other: Other.
	//
	// example:
	//
	// identityCard
	LegalPersonIDCardType *string `json:"LegalPersonIDCardType,omitempty" xml:"LegalPersonIDCardType,omitempty"`
	// Photo of the back of the legal representative\\"s ID card (portrait side). Only jpg, png, gif, and jpeg image formats are supported, and the image must not exceed 5 MB. Please provide the path of the file uploaded to OSS. The file name to be uploaded must not contain Chinese characters or special characters. For upload operations, see [Upload Files via OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// > The system will use the legal person name and ID number you provide for verification. If the verification fails, you need to upload a photo of the legal representative\\"s ID card.
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardBackSide *string `json:"LegalPersonIdCardBackSide,omitempty" xml:"LegalPersonIdCardBackSide,omitempty"`
	// Validity period of the legal person ID card. Format: YYYY-MM-DD~YYYY-MM-DD.
	//
	// > When the certificate validity period is long-term, the end date can be set to 2099-12-31.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	LegalPersonIdCardEffTime *string `json:"LegalPersonIdCardEffTime,omitempty" xml:"LegalPersonIdCardEffTime,omitempty"`
	// Photo of the front of the legal representative\\"s ID card (national emblem side). Only jpg, png, gif, and jpeg image formats are supported, and the image must not exceed 5 MB. Please provide the path of the file uploaded to OSS. The file name to be uploaded must not contain Chinese characters or special characters. For upload operations, see [Upload Files via OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// > The system will use the legal person name and ID number you provide for verification. If the verification fails, you need to upload a photo of the legal representative\\"s ID card.
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardFrontSide *string `json:"LegalPersonIdCardFrontSide,omitempty" xml:"LegalPersonIdCardFrontSide,omitempty"`
	// Name of the legal representative.
	//
	// > - If there is no legal representative information on the organization\\"s certificate, but there is information about a person in charge / chief representative or similar, please prepare the ID card photo of the corresponding person in charge or chief representative listed on the certificate.
	//
	// > - If there is no legal representative information on the organization\\"s certificate, and there is no information about any person in charge, please prepare the name and ID card photo of the main business contact person.
	//
	// example:
	//
	// 李华
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// The review order ID. You can obtain the qualifications and their corresponding review order IDs under the current account by calling [Query Qualification List](~~QuerySmsQualificationRecord~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2001*****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Additional materials. If you have other supporting or supplementary materials, photos, etc., you can upload them here.
	OtherFilesShrink *string `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The qualification ID, that is, the ID returned when you [apply for SMS qualification](~~SubmitSmsQualification~~). You can obtain the qualification IDs under the current account by calling [Query Qualification List](~~QuerySmsQualificationRecord~~).
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
