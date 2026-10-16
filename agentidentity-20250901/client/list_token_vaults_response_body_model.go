// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTokenVaultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTokenVaultsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTokenVaultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTokenVaultsResponseBody
	GetRequestId() *string
	SetTokenVaults(v []*ListTokenVaultsResponseBodyTokenVaults) *ListTokenVaultsResponseBody
	GetTokenVaults() []*ListTokenVaultsResponseBodyTokenVaults
	SetTotalCount(v int32) *ListTokenVaultsResponseBody
	GetTotalCount() *int32
}

type ListTokenVaultsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenVaults []*ListTokenVaultsResponseBodyTokenVaults `json:"TokenVaults,omitempty" xml:"TokenVaults,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTokenVaultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTokenVaultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTokenVaultsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTokenVaultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTokenVaultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTokenVaultsResponseBody) GetTokenVaults() []*ListTokenVaultsResponseBodyTokenVaults {
	return s.TokenVaults
}

func (s *ListTokenVaultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTokenVaultsResponseBody) SetMaxResults(v int32) *ListTokenVaultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTokenVaultsResponseBody) SetNextToken(v string) *ListTokenVaultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTokenVaultsResponseBody) SetRequestId(v string) *ListTokenVaultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTokenVaultsResponseBody) SetTokenVaults(v []*ListTokenVaultsResponseBodyTokenVaults) *ListTokenVaultsResponseBody {
	s.TokenVaults = v
	return s
}

func (s *ListTokenVaultsResponseBody) SetTotalCount(v int32) *ListTokenVaultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTokenVaultsResponseBody) Validate() error {
	if s.TokenVaults != nil {
		for _, item := range s.TokenVaults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTokenVaultsResponseBodyTokenVaults struct {
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example description
	Description      *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptionConfig *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig `json:"EncryptionConfig,omitempty" xml:"EncryptionConfig,omitempty" type:"Struct"`
	// example:
	//
	// acs:ram::123456:role/AliyunAgentIdentityVaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default
	TokenVaultArn *string `json:"TokenVaultArn,omitempty" xml:"TokenVaultArn,omitempty"`
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
	// example:
	//
	// 2026-05-08T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTokenVaultsResponseBodyTokenVaults) String() string {
	return dara.Prettify(s)
}

func (s ListTokenVaultsResponseBodyTokenVaults) GoString() string {
	return s.String()
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetDescription() *string {
	return s.Description
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetEncryptionConfig() *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig {
	return s.EncryptionConfig
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetTokenVaultArn() *string {
	return s.TokenVaultArn
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *ListTokenVaultsResponseBodyTokenVaults) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetCreateTime(v string) *ListTokenVaultsResponseBodyTokenVaults {
	s.CreateTime = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetDescription(v string) *ListTokenVaultsResponseBodyTokenVaults {
	s.Description = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetEncryptionConfig(v *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) *ListTokenVaultsResponseBodyTokenVaults {
	s.EncryptionConfig = v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetRoleArn(v string) *ListTokenVaultsResponseBodyTokenVaults {
	s.RoleArn = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetTokenVaultArn(v string) *ListTokenVaultsResponseBodyTokenVaults {
	s.TokenVaultArn = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetTokenVaultName(v string) *ListTokenVaultsResponseBodyTokenVaults {
	s.TokenVaultName = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) SetUpdateTime(v string) *ListTokenVaultsResponseBodyTokenVaults {
	s.UpdateTime = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaults) Validate() error {
	if s.EncryptionConfig != nil {
		if err := s.EncryptionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig struct {
	// example:
	//
	// SERVICE_MANAGED_KEY
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// acs:kms:cn-beijing:123456:key/key-bjxxxxxxxx
	KmsKeyArn *string `json:"KmsKeyArn,omitempty" xml:"KmsKeyArn,omitempty"`
}

func (s ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) GoString() string {
	return s.String()
}

func (s *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) GetKeyType() *string {
	return s.KeyType
}

func (s *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) GetKmsKeyArn() *string {
	return s.KmsKeyArn
}

func (s *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) SetKeyType(v string) *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig {
	s.KeyType = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) SetKmsKeyArn(v string) *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig {
	s.KmsKeyArn = &v
	return s
}

func (s *ListTokenVaultsResponseBodyTokenVaultsEncryptionConfig) Validate() error {
	return dara.Validate(s)
}
