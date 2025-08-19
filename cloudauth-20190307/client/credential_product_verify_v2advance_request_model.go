// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCredentialProductVerifyV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredName(v string) *CredentialProductVerifyV2AdvanceRequest
	GetCredName() *string
	SetCredType(v string) *CredentialProductVerifyV2AdvanceRequest
	GetCredType() *string
	SetImageFileObject(v io.Reader) *CredentialProductVerifyV2AdvanceRequest
	GetImageFileObject() io.Reader
	SetImageUrl(v string) *CredentialProductVerifyV2AdvanceRequest
	GetImageUrl() *string
	SetMerchantId(v string) *CredentialProductVerifyV2AdvanceRequest
	GetMerchantId() *string
	SetProductCode(v string) *CredentialProductVerifyV2AdvanceRequest
	GetProductCode() *string
}

type CredentialProductVerifyV2AdvanceRequest struct {
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
	ImageFileObject io.Reader `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
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

func (s CredentialProductVerifyV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialProductVerifyV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *CredentialProductVerifyV2AdvanceRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialProductVerifyV2AdvanceRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialProductVerifyV2AdvanceRequest) GetImageFileObject() io.Reader {
	return s.ImageFileObject
}

func (s *CredentialProductVerifyV2AdvanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialProductVerifyV2AdvanceRequest) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialProductVerifyV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialProductVerifyV2AdvanceRequest) SetCredName(v string) *CredentialProductVerifyV2AdvanceRequest {
	s.CredName = &v
	return s
}

func (s *CredentialProductVerifyV2AdvanceRequest) SetCredType(v string) *CredentialProductVerifyV2AdvanceRequest {
	s.CredType = &v
	return s
}

func (s *CredentialProductVerifyV2AdvanceRequest) SetImageFileObject(v io.Reader) *CredentialProductVerifyV2AdvanceRequest {
	s.ImageFileObject = v
	return s
}

func (s *CredentialProductVerifyV2AdvanceRequest) SetImageUrl(v string) *CredentialProductVerifyV2AdvanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialProductVerifyV2AdvanceRequest) SetMerchantId(v string) *CredentialProductVerifyV2AdvanceRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialProductVerifyV2AdvanceRequest) SetProductCode(v string) *CredentialProductVerifyV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialProductVerifyV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
