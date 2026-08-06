// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCatalogKmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKmsKeyId(v string) *VerifyCatalogKmsRequest
	GetKmsKeyId() *string
}

type VerifyCatalogKmsRequest struct {
	// The ID of the KMS customer master key (CMK) to be validated. The server uses this key to perform an SSE-KMS write probe.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-1234567890abcdef
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
}

func (s VerifyCatalogKmsRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCatalogKmsRequest) GoString() string {
	return s.String()
}

func (s *VerifyCatalogKmsRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *VerifyCatalogKmsRequest) SetKmsKeyId(v string) *VerifyCatalogKmsRequest {
	s.KmsKeyId = &v
	return s
}

func (s *VerifyCatalogKmsRequest) Validate() error {
	return dara.Validate(s)
}
