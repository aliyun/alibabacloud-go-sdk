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
	// This parameter is required.
	//
	// example:
	//
	// 0101
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 01
	CredType  *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	ImageFile *string `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	// example:
	//
	// https://oss-bj01.avic.com/eavic-prod-commodity/pic/commodity/94677ee6-1067-4287-8ff4-6e030ef3a5a8.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// This parameter is required.
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
