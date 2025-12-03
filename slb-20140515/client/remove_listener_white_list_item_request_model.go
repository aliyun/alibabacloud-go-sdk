// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveListenerWhiteListItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *RemoveListenerWhiteListItemRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *RemoveListenerWhiteListItemRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *RemoveListenerWhiteListItemRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *RemoveListenerWhiteListItemRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveListenerWhiteListItemRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveListenerWhiteListItemRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveListenerWhiteListItemRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveListenerWhiteListItemRequest
	GetResourceOwnerId() *int64
	SetSourceItems(v string) *RemoveListenerWhiteListItemRequest
	GetSourceItems() *string
}

type RemoveListenerWhiteListItemRequest struct {
	// The listening port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the CLB instance.
	//
	// >  This parameter is required when listeners that use different protocols listen on the same port.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-8vb86hxixo8lvsja8****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the Classic Load Balancer (CLB) instance is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of IP addresses or CIDR blocks that you want to remove from the whitelist. Separate multiple IP addresses or CIDR blocks with commas (,).
	//
	// >  If all IP addresses are removed from the whitelist, the listener does not forward requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceItems *string `json:"SourceItems,omitempty" xml:"SourceItems,omitempty"`
}

func (s RemoveListenerWhiteListItemRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveListenerWhiteListItemRequest) GoString() string {
	return s.String()
}

func (s *RemoveListenerWhiteListItemRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *RemoveListenerWhiteListItemRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *RemoveListenerWhiteListItemRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveListenerWhiteListItemRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveListenerWhiteListItemRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveListenerWhiteListItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveListenerWhiteListItemRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveListenerWhiteListItemRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveListenerWhiteListItemRequest) GetSourceItems() *string {
	return s.SourceItems
}

func (s *RemoveListenerWhiteListItemRequest) SetListenerPort(v int32) *RemoveListenerWhiteListItemRequest {
	s.ListenerPort = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetListenerProtocol(v string) *RemoveListenerWhiteListItemRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetLoadBalancerId(v string) *RemoveListenerWhiteListItemRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetOwnerAccount(v string) *RemoveListenerWhiteListItemRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetOwnerId(v int64) *RemoveListenerWhiteListItemRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetRegionId(v string) *RemoveListenerWhiteListItemRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetResourceOwnerAccount(v string) *RemoveListenerWhiteListItemRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetResourceOwnerId(v int64) *RemoveListenerWhiteListItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetSourceItems(v string) *RemoveListenerWhiteListItemRequest {
	s.SourceItems = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) Validate() error {
	return dara.Validate(s)
}
