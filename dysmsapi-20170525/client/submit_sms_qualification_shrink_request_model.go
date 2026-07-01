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
	// The administrator\\"s ID card validity period. Format: YYYY-MM-DD~YYYY-MM-DD.
	//
	// > If the ID card has a long-term validity period, set the end date to 2099-12-31.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	AdminIDCardExpDate *string `json:"AdminIDCardExpDate,omitempty" xml:"AdminIDCardExpDate,omitempty"`
	// The front photo of the administrator\\"s ID card (national emblem side). Only jpg, png, gif, and jpeg formats are supported. The image must not exceed 5 MB. Specify the file path uploaded to OSS. The file name must not contain Chinese characters or special characters. For upload instructions, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// 	Notice:
	//
	// Color originals do not require a stamp. If you upload a photocopy or black-and-white photo, stamp the photocopy with the company seal and take a photo to upload.
	//
	// .
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardFrontFace *string `json:"AdminIDCardFrontFace,omitempty" xml:"AdminIDCardFrontFace,omitempty"`
	// The administrator\\"s ID card number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 511391********5123
	AdminIDCardNo *string `json:"AdminIDCardNo,omitempty" xml:"AdminIDCardNo,omitempty"`
	// The back photo of the administrator\\"s ID card (portrait side). Only jpg, png, gif, and jpeg formats are supported. The image must not exceed 5 MB. Specify the file path uploaded to OSS. The file name must not contain Chinese characters or special characters. For upload instructions, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// 	Notice:
	//
	// Color originals do not require a stamp. If you upload a photocopy or black-and-white photo, stamp the photocopy with the company seal and take a photo to upload.
	//
	// .
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/111.png
	AdminIDCardPic *string `json:"AdminIDCardPic,omitempty" xml:"AdminIDCardPic,omitempty"`
	// The administrator\\"s ID card type. Valid values:
	//
	// - identityCard: ID card.
	//
	// - passport: passport.
	//
	// - homeReturnPermit: Hong Kong/Macao resident travel permit to mainland.
	//
	// - TaiwanCompatriotPermit: Taiwan resident travel permit to mainland.
	//
	// - residencePermit: Hong Kong/Macao/Taiwan resident residence permit.
	//
	// - other: other.
	//
	// This parameter is required.
	//
	// example:
	//
	// identityCard
	AdminIDCardType *string `json:"AdminIDCardType,omitempty" xml:"AdminIDCardType,omitempty"`
	// The administrator\\"s name. Maximum length: 50 characters. **Under the current [SMS signature real-name requirements](https://help.aliyun.com/document_detail/2873145.html), if the same administrator applies for qualifications for multiple different enterprises, carrier registration will fail. Ensure one administrator per enterprise to improve the registration success rate.**
	//
	// > The administrator (also called the handler) is the person who logs on to the Alibaba Cloud account and manages SMS services. This person typically manages qualifications, signatures, and templates under this Alibaba Cloud account and performs SMS sending operations. This person\\"s phone number must be able to receive verification codes. The administrator does not have to be the Alibaba Cloud account administrator and can be the same person as the legal representative.
	//
	// This parameter is required.
	//
	// example:
	//
	// 李华
	AdminName *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	// The administrator\\"s phone number. Format: +/+86/0086/86 or a phone number without any prefix, such as 1390000****.
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****1234
	AdminPhoneNo *string `json:"AdminPhoneNo,omitempty" xml:"AdminPhoneNo,omitempty"`
	// The business license information. This parameter is required when the qualification purpose `UseBySelf` is set to `false` (third-party use).
	//
	// > - Based on carrier real-name registration regulatory requirements, we strongly recommend that you provide the relevant field information. Otherwise, the probability of "review rejection or carrier registration failure" increases significantly.
	BusinessLicensePicsShrink *string `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty"`
	// The business license validity period. Format: YYYY-MM-DD~YYYY-MM-DD.
	//
	// > If the certificate has a long-term validity period, set the end date to 2099-12-31.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	BussinessLicenseExpDate *string `json:"BussinessLicenseExpDate,omitempty" xml:"BussinessLicenseExpDate,omitempty"`
	// The phone verification code. Call the [RequiredPhoneCode](~~RequiredPhoneCode~~) operation with the **administrator\\"s phone number**, and then enter the SMS verification code received.
	//
	// > You can use [ValidPhoneCode](~~ValidPhoneCode~~) to verify whether the SMS verification code is correct before passing it in.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	CertifyCode *string `json:"CertifyCode,omitempty" xml:"CertifyCode,omitempty"`
	// The enterprise name. Only the middle dot `·`, Chinese brackets `【】（）`, English parentheses `()`, and `spaces` are supported as symbols. Other symbols or pure digits are not allowed. Maximum length: 150 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云云通信有限公司
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The enterprise type. Valid values:
	//
	// - COMPANY: enterprise.
	//
	// - NON_PROFIT_ORGANIZATION: government agency or public institution.
	//
	// This parameter is required.
	//
	// example:
	//
	// COMPANY
	CompanyType *string `json:"CompanyType,omitempty" xml:"CompanyType,omitempty"`
	// The legal representative\\"s ID card number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 511391********5123
	LegalPersonIDCardNo *string `json:"LegalPersonIDCardNo,omitempty" xml:"LegalPersonIDCardNo,omitempty"`
	// The legal representative\\"s ID card type. Valid values:
	//
	// - identityCard: ID card.
	//
	// - passport: passport.
	//
	// - homeReturnPermit: Hong Kong/Macao resident travel permit to mainland.
	//
	// - TaiwanCompatriotPermit: Taiwan resident travel permit to mainland.
	//
	// - residencePermit: Hong Kong/Macao/Taiwan resident residence permit.
	//
	// - other: other.
	//
	// This parameter is required.
	//
	// example:
	//
	// identityCard
	LegalPersonIDCardType *string `json:"LegalPersonIDCardType,omitempty" xml:"LegalPersonIDCardType,omitempty"`
	// The back photo of the legal representative\\"s ID card (portrait side). Only jpg, png, gif, and jpeg formats are supported. The image must not exceed 5 MB. Specify the file path uploaded to OSS. The file name must not contain Chinese characters or special characters. For upload instructions, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// > The system verifies the legal representative\\"s name and ID number you provide. If verification fails, you must upload photos of the legal representative\\"s ID card.
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardBackSide *string `json:"LegalPersonIdCardBackSide,omitempty" xml:"LegalPersonIdCardBackSide,omitempty"`
	// The legal representative\\"s ID card validity period. Format: YYYY-MM-DD~YYYY-MM-DD.
	//
	// > If the ID card has a long-term validity period, set the end date to 2099-12-31.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	LegalPersonIdCardEffTime *string `json:"LegalPersonIdCardEffTime,omitempty" xml:"LegalPersonIdCardEffTime,omitempty"`
	// The front photo of the legal representative\\"s ID card (national emblem side). Only jpg, png, gif, and jpeg formats are supported. The image must not exceed 5 MB. Specify the file path uploaded to OSS. The file name must not contain Chinese characters or special characters. For upload instructions, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	//
	// > The system verifies the legal representative\\"s name and ID number you provide. If verification fails, you must upload photos of the legal representative\\"s ID card.
	//
	// example:
	//
	// 123456/111.png
	LegalPersonIdCardFrontSide *string `json:"LegalPersonIdCardFrontSide,omitempty" xml:"LegalPersonIdCardFrontSide,omitempty"`
	// The legal representative\\"s name. Maximum length: 50 characters.
	//
	// > - If the organization certificate does not contain legal representative information but includes a person in charge or chief representative, prepare the ID card photos of the corresponding person in charge or chief representative listed on the certificate.
	//
	// > - If the organization certificate contains neither legal representative information nor any person in charge, prepare the name and ID card photos of the primary business contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// 李华
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// The unified social credit code. Maximum length: 150 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 910X********0012
	OrganizationCode *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	// Additional materials. If you have other supporting documents, notes, or photos, upload them here.
	OtherFilesShrink *string `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The qualification name, used to manage and distinguish multiple qualifications you apply for. It does not appear in SMS content. The name must be unique among your existing qualifications. Only Chinese characters, English letters, or combinations with digits are supported. Symbols or pure digits are not supported. Maximum length: 100 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云云通信有限公司资质李华
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// Remarks. If you have additional information to provide or notes for the qualification verification reviewer, add a description here.
	//
	// example:
	//
	// 无
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The purpose of the qualification application. Valid values:
	//
	// - **true**: **Self-use**. The entity that owns the signature is the same as the entity verified for this account.
	//
	// - **false**: **Third-party use**. The entity that owns the signature is different from the entity verified for this account.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	UseBySelf *bool `json:"UseBySelf,omitempty" xml:"UseBySelf,omitempty"`
	// Qualification authorization. Specifies whether to share the qualification with other cloud communication products (such as domestic voice services and domestic number privacy protection). Sharing is available only when you apply for a **self-use qualification*	- and the qualification information **matches the enterprise information verified for the current Alibaba Cloud account**. Otherwise, this setting has no effect. Valid values:
	//
	// - true: Agree. Your qualification information can be referenced during the qualification verification process of other cloud communication products, eliminating redundant verification.
	//
	// - false: Disagree.
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
