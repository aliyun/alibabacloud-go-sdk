// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RebootProxyRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RebootProxyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RebootProxyRequest
	GetOwnerId() *int64
	SetProxyList(v string) *RebootProxyRequest
	GetProxyList() *string
	SetProxyNodeList(v string) *RebootProxyRequest
	GetProxyNodeList() *string
	SetRebootMode(v string) *RebootProxyRequest
	GetRebootMode() *string
	SetResourceOwnerAccount(v string) *RebootProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebootProxyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RebootProxyRequest
	GetSecurityToken() *string
}

type RebootProxyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1ymwrhvq79zj****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of the proxy nodes that you want to restart or rebuild. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 90310144,90310144
	ProxyList *string `json:"ProxyList,omitempty" xml:"ProxyList,omitempty"`
	// The IDs of the hosts on which the proxy nodes are deployed. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 36912280,36912282
	ProxyNodeList *string `json:"ProxyNodeList,omitempty" xml:"ProxyNodeList,omitempty"`
	// The type of operation that you want to perform. Valid values:
	//
	// 	- **reboot**: restarts the proxy node.
	//
	// 	- **rebuild**: rebuilds the proxy node.
	//
	// example:
	//
	// reboot
	RebootMode           *string `json:"RebootMode,omitempty" xml:"RebootMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RebootProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootProxyRequest) GoString() string {
	return s.String()
}

func (s *RebootProxyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootProxyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RebootProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebootProxyRequest) GetProxyList() *string {
	return s.ProxyList
}

func (s *RebootProxyRequest) GetProxyNodeList() *string {
	return s.ProxyNodeList
}

func (s *RebootProxyRequest) GetRebootMode() *string {
	return s.RebootMode
}

func (s *RebootProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebootProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebootProxyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RebootProxyRequest) SetInstanceId(v string) *RebootProxyRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootProxyRequest) SetOwnerAccount(v string) *RebootProxyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebootProxyRequest) SetOwnerId(v int64) *RebootProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *RebootProxyRequest) SetProxyList(v string) *RebootProxyRequest {
	s.ProxyList = &v
	return s
}

func (s *RebootProxyRequest) SetProxyNodeList(v string) *RebootProxyRequest {
	s.ProxyNodeList = &v
	return s
}

func (s *RebootProxyRequest) SetRebootMode(v string) *RebootProxyRequest {
	s.RebootMode = &v
	return s
}

func (s *RebootProxyRequest) SetResourceOwnerAccount(v string) *RebootProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebootProxyRequest) SetResourceOwnerId(v int64) *RebootProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebootProxyRequest) SetSecurityToken(v string) *RebootProxyRequest {
	s.SecurityToken = &v
	return s
}

func (s *RebootProxyRequest) Validate() error {
	return dara.Validate(s)
}
