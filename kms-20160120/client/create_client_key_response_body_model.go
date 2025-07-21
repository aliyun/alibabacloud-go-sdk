// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientKeyId(v string) *CreateClientKeyResponseBody
	GetClientKeyId() *string
	SetKeyAlgorithm(v string) *CreateClientKeyResponseBody
	GetKeyAlgorithm() *string
	SetNotAfter(v string) *CreateClientKeyResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *CreateClientKeyResponseBody
	GetNotBefore() *string
	SetPrivateKeyData(v string) *CreateClientKeyResponseBody
	GetPrivateKeyData() *string
	SetRequestId(v string) *CreateClientKeyResponseBody
	GetRequestId() *string
}

type CreateClientKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
	// The ID of the client key.
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The private key of the client key.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The algorithm that is used to encrypt the private key of the client key. Currently, only RSA_2048 is supported.
	//
	// example:
	//
	// MIIJqwIBAzCCCXcGCSqGSIb3DQEHAaCCCWgEgglkMIIJYDCCBBcGCSqGSIb3DQEHBqCCBAgwgg******
	PrivateKeyData *string `json:"PrivateKeyData,omitempty" xml:"PrivateKeyData,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// Specify the time in the ISO 8601 standard. The time must be in UTC. The time must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >
	//
	// 	- If you do not configure NotBefore, the default value is the time when the client key was created.
	//
	// 	- If you configure NotBefore, you must configure NotAfter.
	//
	// example:
	//
	// 2312e45f-b2fa-4c34-ad94-3eca50932916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClientKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientKeyResponseBody) GetClientKeyId() *string {
	return s.ClientKeyId
}

func (s *CreateClientKeyResponseBody) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *CreateClientKeyResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *CreateClientKeyResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *CreateClientKeyResponseBody) GetPrivateKeyData() *string {
	return s.PrivateKeyData
}

func (s *CreateClientKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientKeyResponseBody) SetClientKeyId(v string) *CreateClientKeyResponseBody {
	s.ClientKeyId = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetKeyAlgorithm(v string) *CreateClientKeyResponseBody {
	s.KeyAlgorithm = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetNotAfter(v string) *CreateClientKeyResponseBody {
	s.NotAfter = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetNotBefore(v string) *CreateClientKeyResponseBody {
	s.NotBefore = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetPrivateKeyData(v string) *CreateClientKeyResponseBody {
	s.PrivateKeyData = &v
	return s
}

func (s *CreateClientKeyResponseBody) SetRequestId(v string) *CreateClientKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
