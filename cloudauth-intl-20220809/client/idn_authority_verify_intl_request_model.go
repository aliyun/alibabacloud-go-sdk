// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdnAuthorityVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBirthDate(v string) *IdnAuthorityVerifyIntlRequest
	GetBirthDate() *string
	SetEmail(v string) *IdnAuthorityVerifyIntlRequest
	GetEmail() *string
	SetFullName(v string) *IdnAuthorityVerifyIntlRequest
	GetFullName() *string
	SetIdNumber(v string) *IdnAuthorityVerifyIntlRequest
	GetIdNumber() *string
	SetMerchantBizId(v string) *IdnAuthorityVerifyIntlRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *IdnAuthorityVerifyIntlRequest
	GetMerchantUserId() *string
	SetMobile(v string) *IdnAuthorityVerifyIntlRequest
	GetMobile() *string
	SetProductCode(v string) *IdnAuthorityVerifyIntlRequest
	GetProductCode() *string
	SetSceneCode(v string) *IdnAuthorityVerifyIntlRequest
	GetSceneCode() *string
	SetSourceFacePicture(v string) *IdnAuthorityVerifyIntlRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureFile(v string) *IdnAuthorityVerifyIntlRequest
	GetSourceFacePictureFile() *string
	SetSourceFacePictureUrl(v string) *IdnAuthorityVerifyIntlRequest
	GetSourceFacePictureUrl() *string
	SetTimestamp(v string) *IdnAuthorityVerifyIntlRequest
	GetTimestamp() *string
}

type IdnAuthorityVerifyIntlRequest struct {
	// The date of birth.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000-01-01
	BirthDate *string `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	// The email address of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// a*@gmail.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The full name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ray
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// The ID card number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3612841***47001
	IdNumber *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	// The custom unique business identifier on the merchant side, which is used for subsequent troubleshooting. The value can be up to 32 characters in length and can contain letters and digits. Make sure that the value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// dso932dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or other identifier that identifies a specific user, such as a mobile phone number or email address. Desensitize the value of this field in advance, such as by hashing the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The Indonesian mobile phone number. The number must start with +62, followed by 9 to 11 digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// +6281293671234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The product solution to be integrated. Valid value: IDN_META.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDN_META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The custom authentication scenario ID, which is used to query related records by entering this scenario ID in the console later. The value can be up to 10 characters in length and can contain letters, digits, and underscores.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The Base64-encoded portrait photo.
	//
	// Note
	//
	// - If you choose this method to pass the ID photo, check the photo size and do not pass oversized photos.
	//
	// - Specify only one of the following parameters: SourceFacePicture, SourceFacePictureUrl, and SourceFacePictureFile.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The file stream of the portrait photo.
	//
	// example:
	//
	// InputStream
	SourceFacePictureFile *string `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
	// The URL of the portrait photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// This parameter is required.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s IdnAuthorityVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s IdnAuthorityVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *IdnAuthorityVerifyIntlRequest) GetBirthDate() *string {
	return s.BirthDate
}

func (s *IdnAuthorityVerifyIntlRequest) GetEmail() *string {
	return s.Email
}

func (s *IdnAuthorityVerifyIntlRequest) GetFullName() *string {
	return s.FullName
}

func (s *IdnAuthorityVerifyIntlRequest) GetIdNumber() *string {
	return s.IdNumber
}

func (s *IdnAuthorityVerifyIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *IdnAuthorityVerifyIntlRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *IdnAuthorityVerifyIntlRequest) GetMobile() *string {
	return s.Mobile
}

func (s *IdnAuthorityVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *IdnAuthorityVerifyIntlRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *IdnAuthorityVerifyIntlRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *IdnAuthorityVerifyIntlRequest) GetSourceFacePictureFile() *string {
	return s.SourceFacePictureFile
}

func (s *IdnAuthorityVerifyIntlRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *IdnAuthorityVerifyIntlRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *IdnAuthorityVerifyIntlRequest) SetBirthDate(v string) *IdnAuthorityVerifyIntlRequest {
	s.BirthDate = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetEmail(v string) *IdnAuthorityVerifyIntlRequest {
	s.Email = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetFullName(v string) *IdnAuthorityVerifyIntlRequest {
	s.FullName = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetIdNumber(v string) *IdnAuthorityVerifyIntlRequest {
	s.IdNumber = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetMerchantBizId(v string) *IdnAuthorityVerifyIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetMerchantUserId(v string) *IdnAuthorityVerifyIntlRequest {
	s.MerchantUserId = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetMobile(v string) *IdnAuthorityVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetProductCode(v string) *IdnAuthorityVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetSceneCode(v string) *IdnAuthorityVerifyIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetSourceFacePicture(v string) *IdnAuthorityVerifyIntlRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetSourceFacePictureFile(v string) *IdnAuthorityVerifyIntlRequest {
	s.SourceFacePictureFile = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetSourceFacePictureUrl(v string) *IdnAuthorityVerifyIntlRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) SetTimestamp(v string) *IdnAuthorityVerifyIntlRequest {
	s.Timestamp = &v
	return s
}

func (s *IdnAuthorityVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
