// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupSupportParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeParameterGroupSupportParamRequest
	GetCategory() *string
	SetEngineType(v string) *DescribeParameterGroupSupportParamRequest
	GetEngineType() *string
	SetEngineVersion(v string) *DescribeParameterGroupSupportParamRequest
	GetEngineVersion() *string
	SetOwnerAccount(v string) *DescribeParameterGroupSupportParamRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterGroupSupportParamRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeParameterGroupSupportParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupSupportParamRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParameterGroupSupportParamRequest
	GetSecurityToken() *string
}

type DescribeParameterGroupSupportParamRequest struct {
	// The product category. Valid values:
	//
	// 	- **standard**: Open-source Redis.
	//
	// 	- **enterprise**: Tair (Redis Enterprise Edition).
	//
	// example:
	//
	// standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **redis**: Open-source Redis or memory-optimized Tair.
	//
	// 	- **tair_pena**: Persistent memory-optimized Tair.
	//
	// 	- **tair_pdb**: Disk-optimized Tair (ESSD/SSD).
	//
	// example:
	//
	// redis
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The compatible engine version. Valid values:
	//
	// 	- Open-source Redis supports **5.0**, **6.0**, and **7.0**.
	//
	// 	- Memory-optimized Tair is compatible with the Redis 5.0 and Redis 6.0 protocols. Set the value to **5.0*	- or **6.0**.
	//
	// 	- Persistent memory-optimized Tair is compatible with the Redis 6.0 protocol. Set the value to **1.0**.
	//
	// 	- Disk-optimized Tair (ESSD/SSD) is compatible with the Redis 6.0 protocol. Set the value to **1.0*	- for ESSD-based parameter templates or **2.0*	- for SSD-based parameter templates.
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

func (s DescribeParameterGroupSupportParamRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupSupportParamRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupSupportParamRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeParameterGroupSupportParamRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeParameterGroupSupportParamRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupSupportParamRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterGroupSupportParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupSupportParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupSupportParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupSupportParamRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParameterGroupSupportParamRequest) SetCategory(v string) *DescribeParameterGroupSupportParamRequest {
	s.Category = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetEngineType(v string) *DescribeParameterGroupSupportParamRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetEngineVersion(v string) *DescribeParameterGroupSupportParamRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetOwnerAccount(v string) *DescribeParameterGroupSupportParamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetOwnerId(v int64) *DescribeParameterGroupSupportParamRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupSupportParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupSupportParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) SetSecurityToken(v string) *DescribeParameterGroupSupportParamRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterGroupSupportParamRequest) Validate() error {
	return dara.Validate(s)
}
