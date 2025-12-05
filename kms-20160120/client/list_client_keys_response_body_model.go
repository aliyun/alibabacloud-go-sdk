// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientKeys(v []*ListClientKeysResponseBodyClientKeys) *ListClientKeysResponseBody
	GetClientKeys() []*ListClientKeysResponseBodyClientKeys
	SetRequestId(v string) *ListClientKeysResponseBody
	GetRequestId() *string
}

type ListClientKeysResponseBody struct {
	// A list of client keys.
	ClientKeys []*ListClientKeysResponseBodyClientKeys `json:"ClientKeys,omitempty" xml:"ClientKeys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2312e45f-b2fa-4c34-ad94-3eca50932916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientKeysResponseBody) GetClientKeys() []*ListClientKeysResponseBodyClientKeys {
	return s.ClientKeys
}

func (s *ListClientKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientKeysResponseBody) SetClientKeys(v []*ListClientKeysResponseBodyClientKeys) *ListClientKeysResponseBody {
	s.ClientKeys = v
	return s
}

func (s *ListClientKeysResponseBody) SetRequestId(v string) *ListClientKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientKeysResponseBody) Validate() error {
	if s.ClientKeys != nil {
		for _, item := range s.ClientKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClientKeysResponseBodyClientKeys struct {
	// The name of the AAP.
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
	// Currently, only KMS is supported. The value is fixed as KMS_PROVIDED.
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
	// The public key of the client key.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIDcjCCAlqgAwIBAgIQT/sAVRxwYp54mrw****-----END CERTIFICATE-----
	PublicKeyData *string `json:"PublicKeyData,omitempty" xml:"PublicKeyData,omitempty"`
}

func (s ListClientKeysResponseBodyClientKeys) String() string {
	return dara.Prettify(s)
}

func (s ListClientKeysResponseBodyClientKeys) GoString() string {
	return s.String()
}

func (s *ListClientKeysResponseBodyClientKeys) GetAapName() *string {
	return s.AapName
}

func (s *ListClientKeysResponseBodyClientKeys) GetClientKeyId() *string {
	return s.ClientKeyId
}

func (s *ListClientKeysResponseBodyClientKeys) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClientKeysResponseBodyClientKeys) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *ListClientKeysResponseBodyClientKeys) GetKeyOrigin() *string {
	return s.KeyOrigin
}

func (s *ListClientKeysResponseBodyClientKeys) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListClientKeysResponseBodyClientKeys) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListClientKeysResponseBodyClientKeys) GetPublicKeyData() *string {
	return s.PublicKeyData
}

func (s *ListClientKeysResponseBodyClientKeys) SetAapName(v string) *ListClientKeysResponseBodyClientKeys {
	s.AapName = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetClientKeyId(v string) *ListClientKeysResponseBodyClientKeys {
	s.ClientKeyId = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetCreateTime(v string) *ListClientKeysResponseBodyClientKeys {
	s.CreateTime = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetKeyAlgorithm(v string) *ListClientKeysResponseBodyClientKeys {
	s.KeyAlgorithm = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetKeyOrigin(v string) *ListClientKeysResponseBodyClientKeys {
	s.KeyOrigin = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetNotAfter(v string) *ListClientKeysResponseBodyClientKeys {
	s.NotAfter = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetNotBefore(v string) *ListClientKeysResponseBodyClientKeys {
	s.NotBefore = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) SetPublicKeyData(v string) *ListClientKeysResponseBodyClientKeys {
	s.PublicKeyData = &v
	return s
}

func (s *ListClientKeysResponseBodyClientKeys) Validate() error {
	return dara.Validate(s)
}
