// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeParameterTemplatesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeParameterTemplatesRequest
	GetEngineVersion() *string
	SetOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeParameterTemplatesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetResourceOwnerId() *int64
	SetRole(v string) *DescribeParameterTemplatesRequest
	GetRole() *string
}

type DescribeParameterTemplatesRequest struct {
	// The database engine of the instance. Set the value to **MongoDB**.
	//
	// This parameter is required.
	//
	// example:
	//
	// mongodb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// 	- **5.0**
	//
	// 	- **4.4**
	//
	// 	- **4.2**
	//
	// 	- **4.0**
	//
	// 	- **3.4**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the instance. Valid values:
	//
	// 1. db:  a shard node.
	//
	// 1. cs:  a Configserver node.
	//
	// 1. mongos:  a mongos node.
	//
	// 1. normal: a replica set node.
	//
	// 1. physical: a standalone node.
	//
	// default: normal
	//
	// example:
	//
	// normal
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeParameterTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterTemplatesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterTemplatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterTemplatesRequest) GetRole() *string {
	return s.Role
}

func (s *DescribeParameterTemplatesRequest) SetEngine(v string) *DescribeParameterTemplatesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRegionId(v string) *DescribeParameterTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRole(v string) *DescribeParameterTemplatesRequest {
	s.Role = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
