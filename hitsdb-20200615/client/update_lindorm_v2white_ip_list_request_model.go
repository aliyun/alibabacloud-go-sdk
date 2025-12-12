// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2WhiteIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateLindormV2WhiteIpListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateLindormV2WhiteIpListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateLindormV2WhiteIpListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLindormV2WhiteIpListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateLindormV2WhiteIpListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLindormV2WhiteIpListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpdateLindormV2WhiteIpListRequest
	GetSecurityToken() *string
	SetWhiteIpGroupList(v []*UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) *UpdateLindormV2WhiteIpListRequest
	GetWhiteIpGroupList() []*UpdateLindormV2WhiteIpListRequestWhiteIpGroupList
}

type UpdateLindormV2WhiteIpListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-2ze5ipz9zx1e4****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	WhiteIpGroupList []*UpdateLindormV2WhiteIpListRequestWhiteIpGroupList `json:"WhiteIpGroupList,omitempty" xml:"WhiteIpGroupList,omitempty" type:"Repeated"`
}

func (s UpdateLindormV2WhiteIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2WhiteIpListRequest) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2WhiteIpListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLindormV2WhiteIpListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateLindormV2WhiteIpListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLindormV2WhiteIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLindormV2WhiteIpListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLindormV2WhiteIpListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLindormV2WhiteIpListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLindormV2WhiteIpListRequest) GetWhiteIpGroupList() []*UpdateLindormV2WhiteIpListRequestWhiteIpGroupList {
	return s.WhiteIpGroupList
}

func (s *UpdateLindormV2WhiteIpListRequest) SetInstanceId(v string) *UpdateLindormV2WhiteIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetOwnerAccount(v string) *UpdateLindormV2WhiteIpListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetOwnerId(v int64) *UpdateLindormV2WhiteIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetRegionId(v string) *UpdateLindormV2WhiteIpListRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetResourceOwnerAccount(v string) *UpdateLindormV2WhiteIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetResourceOwnerId(v int64) *UpdateLindormV2WhiteIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetSecurityToken(v string) *UpdateLindormV2WhiteIpListRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) SetWhiteIpGroupList(v []*UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) *UpdateLindormV2WhiteIpListRequest {
	s.WhiteIpGroupList = v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequest) Validate() error {
	if s.WhiteIpGroupList != nil {
		for _, item := range s.WhiteIpGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLindormV2WhiteIpListRequestWhiteIpGroupList struct {
	// This parameter is required.
	//
	// example:
	//
	// user001
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.X/24
	WhiteIpList *string `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty"`
}

func (s UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) GetWhiteIpList() *string {
	return s.WhiteIpList
}

func (s *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) SetGroupName(v string) *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) SetWhiteIpList(v string) *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList {
	s.WhiteIpList = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListRequestWhiteIpGroupList) Validate() error {
	return dara.Validate(s)
}
