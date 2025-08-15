// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredName(v string) *CredentialVerifyIntlRequest
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyIntlRequest
	GetCredType() *string
	SetImageFile(v string) *CredentialVerifyIntlRequest
	GetImageFile() *string
	SetImageUrl(v string) *CredentialVerifyIntlRequest
	GetImageUrl() *string
	SetProductCode(v string) *CredentialVerifyIntlRequest
	GetProductCode() *string
}

type CredentialVerifyIntlRequest struct {
	// Credential name (numeric code):
	//
	// - Starting with 03: Enterprise Qualification
	//
	//   - 0301: Mainland China Business License
	//
	// - Starting with 04, Transaction Voucher
	//
	//   - 0401: Bank Statement
	//
	//   - 0402: Pay Slip
	//
	//   - 0403: Utility Bill
	//
	//   - 0405: Credit Card Statement
	//
	//   - 0499: Others
	//
	// This parameter is required.
	//
	// example:
	//
	// 0301
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// Credential type:
	//
	// - 03: Enterprise Qualification
	//
	// - 04: Transaction Voucher
	//
	// This parameter is required.
	//
	// example:
	//
	// 03
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// Image input stream.
	//
	// > Choose either ImageUrl or ImageFile.
	//
	// example:
	//
	// 无
	ImageFile *string `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	// The URL of the image.
	//
	// > Choose either ImageUrl or ImageFile.
	//
	// example:
	//
	// https://oss-bj01.avic.com/eavic-prod-commodity/pic/commodity/94677ee6-1067-4287-8ff4-6e030ef3a5a8.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Invocation mode:
	//
	// - ANTI_FAKE_CHECK: Image quality and tampering detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 取值：ANTI_FAKE_CHECK
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CredentialVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyIntlRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyIntlRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyIntlRequest) GetImageFile() *string {
	return s.ImageFile
}

func (s *CredentialVerifyIntlRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyIntlRequest) SetCredName(v string) *CredentialVerifyIntlRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyIntlRequest) SetCredType(v string) *CredentialVerifyIntlRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyIntlRequest) SetImageFile(v string) *CredentialVerifyIntlRequest {
	s.ImageFile = &v
	return s
}

func (s *CredentialVerifyIntlRequest) SetImageUrl(v string) *CredentialVerifyIntlRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyIntlRequest) SetProductCode(v string) *CredentialVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
