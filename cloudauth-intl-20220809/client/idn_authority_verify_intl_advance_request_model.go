// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iIdnAuthorityVerifyIntlAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBirthDate(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetBirthDate() *string
	SetEmail(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetEmail() *string
	SetFullName(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetFullName() *string
	SetIdNumber(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetIdNumber() *string
	SetMerchantBizId(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetMerchantUserId() *string
	SetMobile(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetMobile() *string
	SetProductCode(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetProductCode() *string
	SetSceneCode(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetSceneCode() *string
	SetSourceFacePicture(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureFileObject(v io.Reader) *IdnAuthorityVerifyIntlAdvanceRequest
	GetSourceFacePictureFileObject() io.Reader
	SetSourceFacePictureUrl(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetSourceFacePictureUrl() *string
	SetTimestamp(v string) *IdnAuthorityVerifyIntlAdvanceRequest
	GetTimestamp() *string
}

type IdnAuthorityVerifyIntlAdvanceRequest struct {
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
	// The merchant-side custom business unique identifier, which is used for subsequent troubleshooting. The value can be a combination of letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// dso932dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID, or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize the value of this field in advance, for example, by hashing the value.
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
	// The product solution to use. Set the value to IDN_META.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDN_META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The custom authentication scenario ID. You can use this scenario ID to query related records in the console. The value can be a combination of letters, digits, or underscores with a maximum length of 10 characters.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The Base64-encoded facial photo.
	//
	// > **Note**
	//
	// - If you use this method to pass the ID photo, check the photo size and do not pass an excessively large photo.
	//
	// - Specify one of the following parameters: SourceFacePicture, SourceFacePictureUrl, or SourceFacePictureFile.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The file stream of the facial photo.
	//
	// example:
	//
	// InputStream
	SourceFacePictureFileObject io.Reader `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
	// The URL of the facial photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// This parameter is required.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s IdnAuthorityVerifyIntlAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s IdnAuthorityVerifyIntlAdvanceRequest) GoString() string {
	return s.String()
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetBirthDate() *string {
	return s.BirthDate
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetEmail() *string {
	return s.Email
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetFullName() *string {
	return s.FullName
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetIdNumber() *string {
	return s.IdNumber
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetMobile() *string {
	return s.Mobile
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetSourceFacePictureFileObject() io.Reader {
	return s.SourceFacePictureFileObject
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetBirthDate(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.BirthDate = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetEmail(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.Email = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetFullName(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.FullName = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetIdNumber(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.IdNumber = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetMerchantBizId(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetMerchantUserId(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetMobile(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetProductCode(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetSceneCode(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.SceneCode = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetSourceFacePicture(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetSourceFacePictureFileObject(v io.Reader) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.SourceFacePictureFileObject = v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetSourceFacePictureUrl(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) SetTimestamp(v string) *IdnAuthorityVerifyIntlAdvanceRequest {
	s.Timestamp = &v
	return s
}

func (s *IdnAuthorityVerifyIntlAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
