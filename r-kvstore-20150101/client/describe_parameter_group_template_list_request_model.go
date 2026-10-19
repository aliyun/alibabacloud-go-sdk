// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeParameterGroupTemplateListRequest
	GetCategory() *string
	SetCharacterType(v string) *DescribeParameterGroupTemplateListRequest
	GetCharacterType() *string
	SetEngineType(v string) *DescribeParameterGroupTemplateListRequest
	GetEngineType() *string
	SetEngineVersion(v string) *DescribeParameterGroupTemplateListRequest
	GetEngineVersion() *string
	SetOwnerAccount(v string) *DescribeParameterGroupTemplateListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterGroupTemplateListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeParameterGroupTemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupTemplateListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParameterGroupTemplateListRequest
	GetSecurityToken() *string
}

type DescribeParameterGroupTemplateListRequest struct {
	// The product category. Valid values:
	//
	// - **standard**: Redis Community Edition.
	//
	// - **enterprise**: Tair (Enhanced Edition).
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The role type of the instance. Valid values:
	//
	// 	- **logic**: cluster logical instance.
	//
	// 	- **db**: cluster DB instance.
	//
	// 	- **proxy**: cluster proxy instance.
	//
	// 	- **cs**: cluster CS instance.
	//
	// 	- **normal**: primary/secondary DB instance.
	//
	// example:
	//
	// db
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The engine type. Valid values:
	//
	// - **redis**: Redis Community Edition or Tair memory-optimized.
	//
	// - **tair_pena**: Tair persistent memory-optimized.
	//
	// - **tair_pdb**: Tair disk-based (ESSD/SSD).
	//
	// This parameter is required.
	//
	// example:
	//
	// redis
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The compatible version. Valid values:
	//
	// 	- Redis Community Edition supports **5.0**, **6.0**, and **7.0**.
	//
	// 	- Tair memory-optimized instances are compatible with Redis 5.0, 6.0, and 7.0 protocols. Set this parameter to **5.0**, **6.0**, or **7.0**.
	//
	// 	- Tair persistent memory-optimized instances are compatible with the Redis 6.0 protocol. Set this parameter to **1.0**.
	//
	// 	- Tair disk-based (ESSD/SSD) instances are compatible with the Redis 6.0 protocol. To create an ESSD-based parameter template, set this parameter to **1.0**. To create an SSD-based parameter template, set this parameter to **2.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeParameterGroupTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupTemplateListRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeParameterGroupTemplateListRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeParameterGroupTemplateListRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeParameterGroupTemplateListRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupTemplateListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterGroupTemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupTemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupTemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupTemplateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParameterGroupTemplateListRequest) SetCategory(v string) *DescribeParameterGroupTemplateListRequest {
	s.Category = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetCharacterType(v string) *DescribeParameterGroupTemplateListRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetEngineType(v string) *DescribeParameterGroupTemplateListRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetEngineVersion(v string) *DescribeParameterGroupTemplateListRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetOwnerAccount(v string) *DescribeParameterGroupTemplateListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetOwnerId(v int64) *DescribeParameterGroupTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) SetSecurityToken(v string) *DescribeParameterGroupTemplateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterGroupTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
