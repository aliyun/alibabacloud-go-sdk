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
	// The error code returned when the validation fails. An empty string is returned when the validation is successful.
	//
	// example:
	//
	// KeyNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The description of the failure cause and remediation suggestions returned when the validation fails. An empty string is returned when the validation is successful.
	//
	// example:
	//
	// The specified parameter KMS keyId is not found.
	Hint *string `json:"hint,omitempty" xml:"hint,omitempty"`
	// The KMS key identifier actually used by the probe object. When the validation is successful, this corresponds to the customer master key (CMK) specified in the request.
	//
	// example:
	//
	// key-1234567890abcdef
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	// The server-side encryption method actually used by the probe object. Returns KMS when the validation is successful.
	//
	// example:
	//
	// KMS
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" xml:"serverSideEncryption,omitempty"`
	// Indicates whether the validation is successful. A value of true indicates that the write probe succeeded and the SSE-KMS configuration of the object meets expectations. A value of false indicates that the validation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
