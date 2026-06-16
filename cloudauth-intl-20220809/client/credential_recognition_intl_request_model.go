// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlRequest
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlRequest
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialRecognitionIntlRequest
	GetDocType() *string
	SetFraudCheck(v string) *CredentialRecognitionIntlRequest
	GetFraudCheck() *string
	SetOcrArea(v string) *CredentialRecognitionIntlRequest
	GetOcrArea() *string
	SetProductCode(v string) *CredentialRecognitionIntlRequest
	GetProductCode() *string
}

type CredentialRecognitionIntlRequest struct {
	// The Base64-encoded image. If you choose to pass in the image by using IdOcrPictureBase64 (Base64-encoded photo), check the photo size and do not pass in an excessively large photo.
	//
	// example:
	//
	// base64
	CredentialOcrPictureBase64 *string `json:"CredentialOcrPictureBase64,omitempty" xml:"CredentialOcrPictureBase64,omitempty"`
	// The URL of the image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***
	CredentialOcrPictureUrl *string `json:"CredentialOcrPictureUrl,omitempty" xml:"CredentialOcrPictureUrl,omitempty"`
	// The credential type. Valid values:
	//
	// - 01: transaction credential (including electronic bill images for water, electricity, gas, credit cards, and other types).
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to enable tampering detection. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// The extraction type. Valid values:
	//
	// - 0101: electronic bill address and name module (extracts the address and name module through intelligent analysis).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0101
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// The product solution to use. Set this to CREDENTIAL_RECOGNITION.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREDENTIAL_RECOGNITION
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CredentialRecognitionIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlRequest) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlRequest) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialRecognitionIntlRequest) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialRecognitionIntlRequest) GetDocType() *string {
	return s.DocType
}

func (s *CredentialRecognitionIntlRequest) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialRecognitionIntlRequest) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialRecognitionIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialRecognitionIntlRequest) SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlRequest {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlRequest {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetDocType(v string) *CredentialRecognitionIntlRequest {
	s.DocType = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetFraudCheck(v string) *CredentialRecognitionIntlRequest {
	s.FraudCheck = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetOcrArea(v string) *CredentialRecognitionIntlRequest {
	s.OcrArea = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetProductCode(v string) *CredentialRecognitionIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) Validate() error {
	return dara.Validate(s)
}
