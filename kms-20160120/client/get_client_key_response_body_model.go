// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAapName(v string) *GetClientKeyResponseBody
	GetAapName() *string
	SetClientKeyId(v string) *GetClientKeyResponseBody
	GetClientKeyId() *string
	SetCreateTime(v string) *GetClientKeyResponseBody
	GetCreateTime() *string
	SetKeyAlgorithm(v string) *GetClientKeyResponseBody
	GetKeyAlgorithm() *string
	SetKeyOrigin(v string) *GetClientKeyResponseBody
	GetKeyOrigin() *string
	SetNotAfter(v string) *GetClientKeyResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *GetClientKeyResponseBody
	GetNotBefore() *string
	SetPublicKeyData(v string) *GetClientKeyResponseBody
	GetPublicKeyData() *string
	SetRequestId(v string) *GetClientKeyResponseBody
	GetRequestId() *string
}

type GetClientKeyResponseBody struct {
	// The name of the application access point (AAP).
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
	// The ID of the client key.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
	// The time when the client key was created.
	//
	// example:
	//
	// 2023-08-31T09:14:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private key algorithm of the client key.
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The provider of the client key.
	//
	// Currently, only Key Management Service (KMS) is supported. The value is fixed as KMS_PROVIDED.
	//
	// example:
	//
	// KMS_PROVIDED
	KeyOrigin *string `json:"KeyOrigin,omitempty" xml:"KeyOrigin,omitempty"`
	// The end of the validity period of the client key.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the client key.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The content of the public key of the client key.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIDcjCCAlqgAwIBAgIQT/sAVRxwYp54mrw****-----END CERTIFICATE-----
	PublicKeyData *string `json:"PublicKeyData,omitempty" xml:"PublicKeyData,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 63d849a6-045b-4a57-ad9f-c5f756cea9e9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientKeyResponseBody) GetAapName() *string {
	return s.AapName
}

func (s *GetClientKeyResponseBody) GetClientKeyId() *string {
	return s.ClientKeyId
}

func (s *GetClientKeyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetClientKeyResponseBody) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *GetClientKeyResponseBody) GetKeyOrigin() *string {
	return s.KeyOrigin
}

func (s *GetClientKeyResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetClientKeyResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetClientKeyResponseBody) GetPublicKeyData() *string {
	return s.PublicKeyData
}

func (s *GetClientKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientKeyResponseBody) SetAapName(v string) *GetClientKeyResponseBody {
	s.AapName = &v
	return s
}

func (s *GetClientKeyResponseBody) SetClientKeyId(v string) *GetClientKeyResponseBody {
	s.ClientKeyId = &v
	return s
}

func (s *GetClientKeyResponseBody) SetCreateTime(v string) *GetClientKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetClientKeyResponseBody) SetKeyAlgorithm(v string) *GetClientKeyResponseBody {
	s.KeyAlgorithm = &v
	return s
}

func (s *GetClientKeyResponseBody) SetKeyOrigin(v string) *GetClientKeyResponseBody {
	s.KeyOrigin = &v
	return s
}

func (s *GetClientKeyResponseBody) SetNotAfter(v string) *GetClientKeyResponseBody {
	s.NotAfter = &v
	return s
}

func (s *GetClientKeyResponseBody) SetNotBefore(v string) *GetClientKeyResponseBody {
	s.NotBefore = &v
	return s
}

func (s *GetClientKeyResponseBody) SetPublicKeyData(v string) *GetClientKeyResponseBody {
	s.PublicKeyData = &v
	return s
}

func (s *GetClientKeyResponseBody) SetRequestId(v string) *GetClientKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
