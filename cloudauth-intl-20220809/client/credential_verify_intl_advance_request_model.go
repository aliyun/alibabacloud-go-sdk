// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCredentialVerifyIntlAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredName(v string) *CredentialVerifyIntlAdvanceRequest
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyIntlAdvanceRequest
	GetCredType() *string
	SetImageFileObject(v io.Reader) *CredentialVerifyIntlAdvanceRequest
	GetImageFileObject() io.Reader
	SetImageUrl(v string) *CredentialVerifyIntlAdvanceRequest
	GetImageUrl() *string
	SetProductCode(v string) *CredentialVerifyIntlAdvanceRequest
	GetProductCode() *string
}

type CredentialVerifyIntlAdvanceRequest struct {
	// The credential name (specified as a numeric code). Valid values:
	//
	// - Codes starting with 03: enterprise qualification
	//
	//   - 0301: business license issued in the Chinese mainland
	//
	// - Codes starting with 04: transaction voucher
	//
	//   - 0401: bank statement
	//
	//   - 0402: payslip
	//
	//   - 0403: utility bill
	//
	//   - 0405: credit card statement
	//
	//   - 0499: other.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0301
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// The credential type. Valid values:
	//
	// - 03: enterprise qualification
	//
	// - 04: transaction voucher.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// The image input stream.
	//
	// > Specify either ImageUrl or ImageFile.
	//
	// example:
	//
	// 无
	ImageFileObject io.Reader `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	// The URL of the image.
	//
	// > Specify either ImageUrl or ImageFile.
	//
	// example:
	//
	// https://oss-bj01.avic.com/eavic-prod-commodity/pic/commodity/94677ee6-1067-4287-8ff4-6e030ef3a5a8.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The call mode. Valid values:
	//
	// - ANTI_FAKE_CHECK: image quality and tampering detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 取值：ANTI_FAKE_CHECK
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CredentialVerifyIntlAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyIntlAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyIntlAdvanceRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyIntlAdvanceRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyIntlAdvanceRequest) GetImageFileObject() io.Reader {
	return s.ImageFileObject
}

func (s *CredentialVerifyIntlAdvanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyIntlAdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyIntlAdvanceRequest) SetCredName(v string) *CredentialVerifyIntlAdvanceRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyIntlAdvanceRequest) SetCredType(v string) *CredentialVerifyIntlAdvanceRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyIntlAdvanceRequest) SetImageFileObject(v io.Reader) *CredentialVerifyIntlAdvanceRequest {
	s.ImageFileObject = v
	return s
}

func (s *CredentialVerifyIntlAdvanceRequest) SetImageUrl(v string) *CredentialVerifyIntlAdvanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyIntlAdvanceRequest) SetProductCode(v string) *CredentialVerifyIntlAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyIntlAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
