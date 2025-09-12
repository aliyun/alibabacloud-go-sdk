// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceIpWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelete(v bool) *UpdateInstanceIpWhiteListRequest
	GetDelete() *bool
	SetGroupName(v string) *UpdateInstanceIpWhiteListRequest
	GetGroupName() *string
	SetInstanceId(v string) *UpdateInstanceIpWhiteListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateInstanceIpWhiteListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateInstanceIpWhiteListRequest
	GetResourceOwnerId() *int64
	SetSecurityIpList(v string) *UpdateInstanceIpWhiteListRequest
	GetSecurityIpList() *string
	SetSecurityToken(v string) *UpdateInstanceIpWhiteListRequest
	GetSecurityToken() *string
}

type UpdateInstanceIpWhiteListRequest struct {
	// Specifies whether to clear all IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// false
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// The name of the IP whitelist. Default value: user.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the instance for which you want to configure a whitelist. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses or CIDR blocks that you want to add to the whitelist.
	//
	// >  If you add 127.0.0.1 to the whitelist, all IP addresses cannot be used to access the Lindorm instance. If you add the CIDR block 192.168.0.0/24 to the whitelist, you can use all IP addresses in the CIDR block to access the Lindorm instance. Separate multiple IP addresses or CIDR blocks with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 106.11.XX.XX/24
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateInstanceIpWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListRequest) GetDelete() *bool {
	return s.Delete
}

func (s *UpdateInstanceIpWhiteListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateInstanceIpWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceIpWhiteListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateInstanceIpWhiteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateInstanceIpWhiteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateInstanceIpWhiteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateInstanceIpWhiteListRequest) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *UpdateInstanceIpWhiteListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateInstanceIpWhiteListRequest) SetDelete(v bool) *UpdateInstanceIpWhiteListRequest {
	s.Delete = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetGroupName(v string) *UpdateInstanceIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetInstanceId(v string) *UpdateInstanceIpWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetOwnerId(v int64) *UpdateInstanceIpWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetResourceOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetResourceOwnerId(v int64) *UpdateInstanceIpWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetSecurityIpList(v string) *UpdateInstanceIpWhiteListRequest {
	s.SecurityIpList = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetSecurityToken(v string) *UpdateInstanceIpWhiteListRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
