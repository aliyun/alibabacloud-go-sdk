// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSecretValueResponseBody
	GetRequestId() *string
	SetSecretId(v string) *CreateSecretValueResponseBody
	GetSecretId() *string
}

type CreateSecretValueResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5C2B60C6-03CE-5F31-A8A4-13FCBEBACA1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ciphertext ID.
	//
	// example:
	//
	// st-xxxxxxxxxx
	SecretId *string `json:"secretId,omitempty" xml:"secretId,omitempty"`
}

func (s CreateSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretValueResponseBody) GetSecretId() *string {
	return s.SecretId
}

func (s *CreateSecretValueResponseBody) SetRequestId(v string) *CreateSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretValueResponseBody) SetSecretId(v string) *CreateSecretValueResponseBody {
	s.SecretId = &v
	return s
}

func (s *CreateSecretValueResponseBody) Validate() error {
	return dara.Validate(s)
}
