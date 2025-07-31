// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetDBClusterId() *string
	SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetGlobalSecurityGroupId() *string
	SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest
	GetResourceOwnerId() *int64
}

type ModifyGlobalSecurityIPGroupRelationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-2ze6069764423m0l
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the IP whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// g-u0qdtybfvxhaxrrhk4n7
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetDBClusterId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetRegionId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
