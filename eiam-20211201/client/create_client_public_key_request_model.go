// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmType(v string) *CreateClientPublicKeyRequest
	GetAlgorithmType() *string
	SetApplicationId(v string) *CreateClientPublicKeyRequest
	GetApplicationId() *string
	SetClientToken(v string) *CreateClientPublicKeyRequest
	GetClientToken() *string
	SetInstanceId(v string) *CreateClientPublicKeyRequest
	GetInstanceId() *string
	SetPublicKey(v string) *CreateClientPublicKeyRequest
	GetPublicKey() *string
}

type CreateClientPublicKeyRequest struct {
	// IDaaS的应用ClientpublicKey算法类型
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA-2048
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 应用ClientpublicKey的公钥，只接收 PEM 格式
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmnWMdp9FU3vXljeIqpgR
	//
	// 05E6jEgzIfKsFaLkv+07e2Lg8luTaJh8Q2nkbxdNpTfqBnMMyTgml88WktP45F78
	//
	// La7hQtR3vz0Eu1yA92gXwD5Oob7ay4JYQZ0C80o2tB3FsbXG2jUvr31MNkf/0oBY
	//
	// qUKK5Hnlk1TdrnZ5VkzgLGHKlj+NJHHF/57DbT64C2qpAWeKHAr9umJ8++0hKqG/
	//
	// oRSOpb7oWK4t5c39ulp1j5JJ6cwqrKVCXvsHfWHywOHkcyus+ZNPTQvpwjRnEmRz
	//
	// Vy3NWrjT7JlIa8vS1aUU+FxeFd2MLQzxFxquFLwi05su2faRexaeYwWW6IWAI3tX
	//
	// twxxxxxx
	//
	// -----END PUBLIC KEY-----
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
}

func (s CreateClientPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateClientPublicKeyRequest) GetAlgorithmType() *string {
	return s.AlgorithmType
}

func (s *CreateClientPublicKeyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateClientPublicKeyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateClientPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateClientPublicKeyRequest) GetPublicKey() *string {
	return s.PublicKey
}

func (s *CreateClientPublicKeyRequest) SetAlgorithmType(v string) *CreateClientPublicKeyRequest {
	s.AlgorithmType = &v
	return s
}

func (s *CreateClientPublicKeyRequest) SetApplicationId(v string) *CreateClientPublicKeyRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateClientPublicKeyRequest) SetClientToken(v string) *CreateClientPublicKeyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClientPublicKeyRequest) SetInstanceId(v string) *CreateClientPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateClientPublicKeyRequest) SetPublicKey(v string) *CreateClientPublicKeyRequest {
	s.PublicKey = &v
	return s
}

func (s *CreateClientPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
