// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientPublicKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientPublicKeys(v []*ListClientPublicKeysResponseBodyClientPublicKeys) *ListClientPublicKeysResponseBody
	GetClientPublicKeys() []*ListClientPublicKeysResponseBodyClientPublicKeys
	SetMaxResults(v int32) *ListClientPublicKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListClientPublicKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListClientPublicKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListClientPublicKeysResponseBody
	GetTotalCount() *int64
}

type ListClientPublicKeysResponseBody struct {
	ClientPublicKeys []*ListClientPublicKeysResponseBodyClientPublicKeys `json:"ClientPublicKeys,omitempty" xml:"ClientPublicKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientPublicKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientPublicKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientPublicKeysResponseBody) GetClientPublicKeys() []*ListClientPublicKeysResponseBodyClientPublicKeys {
	return s.ClientPublicKeys
}

func (s *ListClientPublicKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListClientPublicKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClientPublicKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientPublicKeysResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListClientPublicKeysResponseBody) SetClientPublicKeys(v []*ListClientPublicKeysResponseBodyClientPublicKeys) *ListClientPublicKeysResponseBody {
	s.ClientPublicKeys = v
	return s
}

func (s *ListClientPublicKeysResponseBody) SetMaxResults(v int32) *ListClientPublicKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListClientPublicKeysResponseBody) SetNextToken(v string) *ListClientPublicKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClientPublicKeysResponseBody) SetRequestId(v string) *ListClientPublicKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientPublicKeysResponseBody) SetTotalCount(v int64) *ListClientPublicKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClientPublicKeysResponseBody) Validate() error {
	if s.ClientPublicKeys != nil {
		for _, item := range s.ClientPublicKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClientPublicKeysResponseBodyClientPublicKeys struct {
	// IDaaS EIAM 应用ClientPublicKey的算法类型 rsa2048、ecc256
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
	// IDaaS EIAM 应用ClientPublicKey的Id
	//
	// example:
	//
	// KEYEqDnDJhztiEAwSin7MZoxGcihzCAuxxxx
	ClientPublicKeyId *string `json:"ClientPublicKeyId,omitempty" xml:"ClientPublicKeyId,omitempty"`
	// IDaaS EIAM 应用ClientPublicKey的创建时间
	//
	// example:
	//
	// 1722006052000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1762006052000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// IDaaS EIAM 应用当前是否为首要使用的应用ClientPublicKey的
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// IDaaS EIAM 应用ClientPublicKey的公钥
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
	// IDaaS EIAM 应用ClientPublicKey的状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListClientPublicKeysResponseBodyClientPublicKeys) String() string {
	return dara.Prettify(s)
}

func (s ListClientPublicKeysResponseBodyClientPublicKeys) GoString() string {
	return s.String()
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetAlgorithmType() *string {
	return s.AlgorithmType
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetPrimary() *bool {
	return s.Primary
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) GetStatus() *string {
	return s.Status
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetAlgorithmType(v string) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.AlgorithmType = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetApplicationId(v string) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.ApplicationId = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetClientPublicKeyId(v string) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.ClientPublicKeyId = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetCreateTime(v int64) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.CreateTime = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetInstanceId(v string) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.InstanceId = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetLastUsedTime(v int64) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.LastUsedTime = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetPrimary(v bool) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.Primary = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetPublicKey(v string) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.PublicKey = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) SetStatus(v string) *ListClientPublicKeysResponseBodyClientPublicKeys {
	s.Status = &v
	return s
}

func (s *ListClientPublicKeysResponseBodyClientPublicKeys) Validate() error {
	return dara.Validate(s)
}
