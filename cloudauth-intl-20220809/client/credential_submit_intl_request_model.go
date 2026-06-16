// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialSubmitIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialOcrPictureBase64(v string) *CredentialSubmitIntlRequest
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureUrl(v string) *CredentialSubmitIntlRequest
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialSubmitIntlRequest
	GetDocType() *string
	SetFraudCheck(v string) *CredentialSubmitIntlRequest
	GetFraudCheck() *string
	SetMerchantBizId(v string) *CredentialSubmitIntlRequest
	GetMerchantBizId() *string
	SetOcrArea(v string) *CredentialSubmitIntlRequest
	GetOcrArea() *string
	SetProductCode(v string) *CredentialSubmitIntlRequest
	GetProductCode() *string
	SetSceneCode(v string) *CredentialSubmitIntlRequest
	GetSceneCode() *string
}

type CredentialSubmitIntlRequest struct {
	// The Base64-encoded image. If you use this method to submit a photo, check the photo size and do not submit an excessively large photo.
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
	// - 02: vehicle registration certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to enable tampering detection. Valid values:
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// The merchant-side custom business unique identifier, used for subsequent troubleshooting. The value can be a combination of letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The extraction type. Valid values:
	//
	// - 0201: Thailand vehicle registration certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0201
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// The product solution to use. Set this parameter to CREDENTIAL_RECOGNITION.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREDENTIAL_RECOGNITION
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The custom authentication scenario ID. You can use this scenario ID to query related records in the console. The value can be a combination of letters, digits, or underscores with a maximum length of 10 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123****123
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s CredentialSubmitIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlRequest) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlRequest) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialSubmitIntlRequest) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialSubmitIntlRequest) GetDocType() *string {
	return s.DocType
}

func (s *CredentialSubmitIntlRequest) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialSubmitIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *CredentialSubmitIntlRequest) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialSubmitIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialSubmitIntlRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *CredentialSubmitIntlRequest) SetCredentialOcrPictureBase64(v string) *CredentialSubmitIntlRequest {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetCredentialOcrPictureUrl(v string) *CredentialSubmitIntlRequest {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetDocType(v string) *CredentialSubmitIntlRequest {
	s.DocType = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetFraudCheck(v string) *CredentialSubmitIntlRequest {
	s.FraudCheck = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetMerchantBizId(v string) *CredentialSubmitIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetOcrArea(v string) *CredentialSubmitIntlRequest {
	s.OcrArea = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetProductCode(v string) *CredentialSubmitIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialSubmitIntlRequest) SetSceneCode(v string) *CredentialSubmitIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *CredentialSubmitIntlRequest) Validate() error {
	return dara.Validate(s)
}
