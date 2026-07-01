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
	BusinessLicensePics []*UpdateSmsQualificationRequestBusinessLicensePics `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty" type:"Repeated"`
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
	OtherFiles []*UpdateSmsQualificationRequestOtherFiles `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty" type:"Repeated"`
	OwnerId    *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// Business license image. Only jpg, png, gif, and jpeg image formats are supported, and the image must not exceed 5 MB. Please provide the path of the file uploaded to OSS. The file name to be uploaded must not contain Chinese characters or special characters. For upload operations, see [Upload Files via OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// 	Notice:
	//
	// No stamp is required for color originals of the certificate. If you upload a photocopy or black-and-white photo, you must affix the enterprise red seal to the photocopy and take a photo to upload.
	//
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
	// Business license type. Valid values:
	//
	// - socialCreditLicense: Social credit code certificate.
	//
	// - businessLicense: Enterprise business license.
	//
	// - signLegalLicense: Public institution legal person certificate.
	//
	// - otherLicense: Other.
	//
	// Choose one to upload. The certificate must contain: enterprise name, unified social credit code, and certificate validity period.
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
	// Additional material file. Only png, jpg, jpeg, doc, docx, and pdf formats are supported, and the file must not exceed 5 MB. Please provide the path of the file uploaded to OSS. The file name to be uploaded must not contain Chinese characters or special characters. For upload operations, see [Upload Files via OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
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
