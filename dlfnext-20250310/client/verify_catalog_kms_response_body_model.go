// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCatalogKmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *VerifyCatalogKmsResponseBody
	GetErrorCode() *string
	SetHint(v string) *VerifyCatalogKmsResponseBody
	GetHint() *string
	SetKmsKeyId(v string) *VerifyCatalogKmsResponseBody
	GetKmsKeyId() *string
	SetServerSideEncryption(v string) *VerifyCatalogKmsResponseBody
	GetServerSideEncryption() *string
	SetSuccess(v bool) *VerifyCatalogKmsResponseBody
	GetSuccess() *bool
}

type VerifyCatalogKmsResponseBody struct {
	ErrorCode            *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	Hint                 *string `json:"hint,omitempty" xml:"hint,omitempty"`
	KmsKeyId             *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" xml:"serverSideEncryption,omitempty"`
	Success              *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyCatalogKmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCatalogKmsResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCatalogKmsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *VerifyCatalogKmsResponseBody) GetHint() *string {
	return s.Hint
}

func (s *VerifyCatalogKmsResponseBody) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *VerifyCatalogKmsResponseBody) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *VerifyCatalogKmsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyCatalogKmsResponseBody) SetErrorCode(v string) *VerifyCatalogKmsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *VerifyCatalogKmsResponseBody) SetHint(v string) *VerifyCatalogKmsResponseBody {
	s.Hint = &v
	return s
}

func (s *VerifyCatalogKmsResponseBody) SetKmsKeyId(v string) *VerifyCatalogKmsResponseBody {
	s.KmsKeyId = &v
	return s
}

func (s *VerifyCatalogKmsResponseBody) SetServerSideEncryption(v string) *VerifyCatalogKmsResponseBody {
	s.ServerSideEncryption = &v
	return s
}

func (s *VerifyCatalogKmsResponseBody) SetSuccess(v bool) *VerifyCatalogKmsResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyCatalogKmsResponseBody) Validate() error {
	return dara.Validate(s)
}
