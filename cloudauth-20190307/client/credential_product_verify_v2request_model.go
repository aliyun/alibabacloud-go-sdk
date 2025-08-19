// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialProductVerifyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCredName(v string) *CredentialProductVerifyV2Request
	GetCredName() *string
	SetCredType(v string) *CredentialProductVerifyV2Request
	GetCredType() *string
	SetImageFile(v string) *CredentialProductVerifyV2Request
	GetImageFile() *string
	SetImageUrl(v string) *CredentialProductVerifyV2Request
	GetImageUrl() *string
	SetMerchantId(v string) *CredentialProductVerifyV2Request
	GetMerchantId() *string
	SetProductCode(v string) *CredentialProductVerifyV2Request
	GetProductCode() *string
}

type CredentialProductVerifyV2Request struct {
	// Credential name: Only supports value 0501 (product image).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0501
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// Credential type: Only supports value 05 (product image).
	//
	// This parameter is required.
	//
	// example:
	//
	// 05
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// InputStream object of the image.
	//
	// example:
	//
	// https://aliyundoc.com/picture*****.jpeg
	ImageFile *string `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	// URL of the image.
	//
	// example:
	//
	// https://aliyundoc.com/picture*****.jpeg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Merchant ID.
	//
	// example:
	//
	// 无。
	MerchantId *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	// Invocation mode:
	//
	// Only supports value ANTI_FAKE_CHECK.
	//
	// This parameter is required.
	//
	// example:
	//
	// ANTI_FAKE_CHECK
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CredentialProductVerifyV2Request) String() string {
	return dara.Prettify(s)
}

func (s CredentialProductVerifyV2Request) GoString() string {
	return s.String()
}

func (s *CredentialProductVerifyV2Request) GetCredName() *string {
	return s.CredName
}

func (s *CredentialProductVerifyV2Request) GetCredType() *string {
	return s.CredType
}

func (s *CredentialProductVerifyV2Request) GetImageFile() *string {
	return s.ImageFile
}

func (s *CredentialProductVerifyV2Request) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialProductVerifyV2Request) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialProductVerifyV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialProductVerifyV2Request) SetCredName(v string) *CredentialProductVerifyV2Request {
	s.CredName = &v
	return s
}

func (s *CredentialProductVerifyV2Request) SetCredType(v string) *CredentialProductVerifyV2Request {
	s.CredType = &v
	return s
}

func (s *CredentialProductVerifyV2Request) SetImageFile(v string) *CredentialProductVerifyV2Request {
	s.ImageFile = &v
	return s
}

func (s *CredentialProductVerifyV2Request) SetImageUrl(v string) *CredentialProductVerifyV2Request {
	s.ImageUrl = &v
	return s
}

func (s *CredentialProductVerifyV2Request) SetMerchantId(v string) *CredentialProductVerifyV2Request {
	s.MerchantId = &v
	return s
}

func (s *CredentialProductVerifyV2Request) SetProductCode(v string) *CredentialProductVerifyV2Request {
	s.ProductCode = &v
	return s
}

func (s *CredentialProductVerifyV2Request) Validate() error {
	return dara.Validate(s)
}
