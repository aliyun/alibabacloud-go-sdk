// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpgradeProxyRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpgradeProxyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeProxyRequest
	GetOwnerId() *int64
	SetProxyInstanceIds(v string) *UpgradeProxyRequest
	GetProxyInstanceIds() *string
	SetResourceOwnerAccount(v string) *UpgradeProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeProxyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpgradeProxyRequest
	GetSecurityToken() *string
	SetSwitchTimeMode(v int32) *UpgradeProxyRequest
	GetSwitchTimeMode() *int32
}

type UpgradeProxyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1vgja77wl7pb****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of the proxy nodes to upgrade, separated by commas (,). This parameter is valid only for Classic Edition instances.
	//
	// example:
	//
	// r-bp1vgja77wl7pb****-proxy-0
	ProxyInstanceIds     *string `json:"ProxyInstanceIds,omitempty" xml:"ProxyInstanceIds,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The time when the upgrade is performed. Valid values:
	//
	// 	- **Immediately**: The upgrade is immediately performed. This is the default value.
	//
	// 	- **MaintainTime**: The upgrade is performed within the O&M window of the instance.
	//
	// > You can call [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) to modify the O&M window of a Redis instance.
	//
	// example:
	//
	// Immediately
	SwitchTimeMode *int32 `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s UpgradeProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeProxyRequest) GoString() string {
	return s.String()
}

func (s *UpgradeProxyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeProxyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeProxyRequest) GetProxyInstanceIds() *string {
	return s.ProxyInstanceIds
}

func (s *UpgradeProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeProxyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpgradeProxyRequest) GetSwitchTimeMode() *int32 {
	return s.SwitchTimeMode
}

func (s *UpgradeProxyRequest) SetInstanceId(v string) *UpgradeProxyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeProxyRequest) SetOwnerAccount(v string) *UpgradeProxyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeProxyRequest) SetOwnerId(v int64) *UpgradeProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeProxyRequest) SetProxyInstanceIds(v string) *UpgradeProxyRequest {
	s.ProxyInstanceIds = &v
	return s
}

func (s *UpgradeProxyRequest) SetResourceOwnerAccount(v string) *UpgradeProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeProxyRequest) SetResourceOwnerId(v int64) *UpgradeProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeProxyRequest) SetSecurityToken(v string) *UpgradeProxyRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeProxyRequest) SetSwitchTimeMode(v int32) *UpgradeProxyRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeProxyRequest) Validate() error {
	return dara.Validate(s)
}
