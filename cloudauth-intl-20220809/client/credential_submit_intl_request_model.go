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
	// Base64 encoding of the image. If you choose to upload the photo this way, please check the photo size and avoid uploading overly large photos.
	//
	// example:
	//
	// base64
	CredentialOcrPictureBase64 *string `json:"CredentialOcrPictureBase64,omitempty" xml:"CredentialOcrPictureBase64,omitempty"`
	// Image URL, accessible via HTTP or HTTPS on the public network.
	//
	// example:
	//
	// https://***
	CredentialOcrPictureUrl *string `json:"CredentialOcrPictureUrl,omitempty" xml:"CredentialOcrPictureUrl,omitempty"`
	// Credential type:
	//
	// - 02: Vehicle registration certificate
	//
	// This parameter is required.
	//
	// example:
	//
	// 02
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Whether to enable tampering detection
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// A unique business identifier defined on the merchant side, used for troubleshooting issues later. Supports a combination of letters and digits, with a maximum length of 32 characters. Ensure uniqueness.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Extraction type:
	//
	// - 0201: Thai vehicle registration certificate
	//
	// This parameter is required.
	//
	// example:
	//
	// 0201
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// The product solution to be integrated. Value: CREDENTIAL_RECOGNITION.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREDENTIAL_RECOGNITION
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Your custom authentication scenario ID, used for querying related records by entering this scenario ID in the console later. Supports a combination of 10 characters, digits, or underscores.
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
