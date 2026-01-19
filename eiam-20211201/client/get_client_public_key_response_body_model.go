// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientPublicKey(v *GetClientPublicKeyResponseBodyClientPublicKey) *GetClientPublicKeyResponseBody
	GetClientPublicKey() *GetClientPublicKeyResponseBodyClientPublicKey
	SetRequestId(v string) *GetClientPublicKeyResponseBody
	GetRequestId() *string
}

type GetClientPublicKeyResponseBody struct {
	ClientPublicKey *GetClientPublicKeyResponseBodyClientPublicKey `json:"ClientPublicKey,omitempty" xml:"ClientPublicKey,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientPublicKeyResponseBody) GetClientPublicKey() *GetClientPublicKeyResponseBodyClientPublicKey {
	return s.ClientPublicKey
}

func (s *GetClientPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientPublicKeyResponseBody) SetClientPublicKey(v *GetClientPublicKeyResponseBodyClientPublicKey) *GetClientPublicKeyResponseBody {
	s.ClientPublicKey = v
	return s
}

func (s *GetClientPublicKeyResponseBody) SetRequestId(v string) *GetClientPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientPublicKeyResponseBody) Validate() error {
	if s.ClientPublicKey != nil {
		if err := s.ClientPublicKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClientPublicKeyResponseBodyClientPublicKey struct {
	// IDaaS EIAM 应用公私钥对算法类型 rsa2048、ecc256
	//
	// example:
	//
	// RSA-2048
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	// IDaaS EIAM 应用Id
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM 应用公私钥对Id
	//
	// example:
	//
	// KEYEqDnDJhztiEAwSin7MZoxGcihzCAuxxxx
	ClientPublicKeyId *string `json:"ClientPublicKeyId,omitempty" xml:"ClientPublicKeyId,omitempty"`
	// IDaaS EIAM 应用公私钥对创建时间
	//
	// example:
	//
	// 1731305755000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1771305755000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// IDaaS EIAM 应用当前是否为首要使用的公私钥对
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// IDaaS EIAM 应用公钥
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
	// IDaaS EIAM 应用公私钥对状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetClientPublicKeyResponseBodyClientPublicKey) String() string {
	return dara.Prettify(s)
}

func (s GetClientPublicKeyResponseBodyClientPublicKey) GoString() string {
	return s.String()
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetAlgorithmType() *string {
	return s.AlgorithmType
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetPrimary() *bool {
	return s.Primary
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetPublicKey() *string {
	return s.PublicKey
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) GetStatus() *string {
	return s.Status
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetAlgorithmType(v string) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.AlgorithmType = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetApplicationId(v string) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.ApplicationId = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetClientPublicKeyId(v string) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.ClientPublicKeyId = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetCreateTime(v int64) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.CreateTime = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetInstanceId(v string) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.InstanceId = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetLastUsedTime(v int64) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.LastUsedTime = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetPrimary(v bool) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.Primary = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetPublicKey(v string) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.PublicKey = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) SetStatus(v string) *GetClientPublicKeyResponseBodyClientPublicKey {
	s.Status = &v
	return s
}

func (s *GetClientPublicKeyResponseBodyClientPublicKey) Validate() error {
	return dara.Validate(s)
}
