// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2WhiteIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteGroup(v bool) *ModifyLindormV2WhiteIpListRequest
	GetDeleteGroup() *bool
	SetGroupName(v string) *ModifyLindormV2WhiteIpListRequest
	GetGroupName() *string
	SetInstanceId(v string) *ModifyLindormV2WhiteIpListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyLindormV2WhiteIpListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLindormV2WhiteIpListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLindormV2WhiteIpListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLindormV2WhiteIpListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLindormV2WhiteIpListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyLindormV2WhiteIpListRequest
	GetSecurityToken() *string
	SetWhiteIpList(v string) *ModifyLindormV2WhiteIpListRequest
	GetWhiteIpList() *string
}

type ModifyLindormV2WhiteIpListRequest struct {
	DeleteGroup *bool `json:"DeleteGroup,omitempty" xml:"DeleteGroup,omitempty"`
	// This parameter is required.
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	WhiteIpList *string `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty"`
}

func (s ModifyLindormV2WhiteIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2WhiteIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2WhiteIpListRequest) GetDeleteGroup() *bool {
	return s.DeleteGroup
}

func (s *ModifyLindormV2WhiteIpListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyLindormV2WhiteIpListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyLindormV2WhiteIpListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLindormV2WhiteIpListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLindormV2WhiteIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLindormV2WhiteIpListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLindormV2WhiteIpListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLindormV2WhiteIpListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyLindormV2WhiteIpListRequest) GetWhiteIpList() *string {
	return s.WhiteIpList
}

func (s *ModifyLindormV2WhiteIpListRequest) SetDeleteGroup(v bool) *ModifyLindormV2WhiteIpListRequest {
	s.DeleteGroup = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetGroupName(v string) *ModifyLindormV2WhiteIpListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetInstanceId(v string) *ModifyLindormV2WhiteIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetOwnerAccount(v string) *ModifyLindormV2WhiteIpListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetOwnerId(v int64) *ModifyLindormV2WhiteIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetRegionId(v string) *ModifyLindormV2WhiteIpListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetResourceOwnerAccount(v string) *ModifyLindormV2WhiteIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetResourceOwnerId(v int64) *ModifyLindormV2WhiteIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetSecurityToken(v string) *ModifyLindormV2WhiteIpListRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetWhiteIpList(v string) *ModifyLindormV2WhiteIpListRequest {
	s.WhiteIpList = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) Validate() error {
	return dara.Validate(s)
}
