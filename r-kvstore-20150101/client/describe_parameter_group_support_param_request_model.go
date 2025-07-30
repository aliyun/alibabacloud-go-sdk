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
	// The service category. Valid values:
	//
	// 	- **standard**: ApsaraDB for Redis Community Edition
	//
	// 	- **enterprise**: ApsaraDB for Redis Enhanced Edition (Tair)
	//
	// example:
	//
	// standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **redis**: ApsaraDB for Redis Community Edition instance or Tair DRAM-based instance
	//
	// 	- **tair_pena**: Tair persistent memory-optimized instance
	//
	// 	- **tair_pdb**: Tair ESSD/SSD-based instance
	//
	// example:
	//
	// redis
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The compatible engine version. Valid values:
	//
	// 	- For ApsaraDB for Redis Community Edition instances, set the parameter to **5.0**, **6.0**, or **7.0**.
	//
	// 	- For Tair DRAM-based instances that are compatible with Redis 5.0 or Redis 6.0, set the parameter to **5.0*	- or **6.0**.
	//
	// 	- For Tair persistent memory-optimized instances that are compatible with Redis 6.0, set the parameter to **1.0**.
	//
	// 	- For Tair ESSD-based instances that are compatible with Redis 6.0, set the parameter to **1.0**. For Tair SSD-based instances that are compatible with Redis 6.0, set the parameter to **2.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
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
