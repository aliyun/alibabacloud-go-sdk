// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddListenerWhiteListItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *AddListenerWhiteListItemRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *AddListenerWhiteListItemRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *AddListenerWhiteListItemRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *AddListenerWhiteListItemRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddListenerWhiteListItemRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddListenerWhiteListItemRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddListenerWhiteListItemRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddListenerWhiteListItemRequest
	GetResourceOwnerId() *int64
	SetSourceItems(v string) *AddListenerWhiteListItemRequest
	GetSourceItems() *string
}

type AddListenerWhiteListItemRequest struct {
	// The frontend port that is used by the CLB instance.
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
	// lb-bp1o94dp5i6ea*******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses or CIDR blocks that you want to add to the whitelist.
	//
	// This parameter takes effect when the **AccessControlStatus*	- parameter of the listener is set to **open_white_list**.
	//
	// Separate multiple IP addresses or CIDR blocks with commas (,).
	//
	// You cannot enter **0.0.0.0*	- or **0.0.0.0/0**. To disable access control, you can call the [SetListenerAccessControlStatus](https://help.aliyun.com/document_detail/27599.html) operation to set the value of the **AccessControlStatus*	- parameter to **close**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceItems *string `json:"SourceItems,omitempty" xml:"SourceItems,omitempty"`
}

func (s AddListenerWhiteListItemRequest) String() string {
	return dara.Prettify(s)
}

func (s AddListenerWhiteListItemRequest) GoString() string {
	return s.String()
}

func (s *AddListenerWhiteListItemRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *AddListenerWhiteListItemRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *AddListenerWhiteListItemRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddListenerWhiteListItemRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddListenerWhiteListItemRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddListenerWhiteListItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddListenerWhiteListItemRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddListenerWhiteListItemRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddListenerWhiteListItemRequest) GetSourceItems() *string {
	return s.SourceItems
}

func (s *AddListenerWhiteListItemRequest) SetListenerPort(v int32) *AddListenerWhiteListItemRequest {
	s.ListenerPort = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetListenerProtocol(v string) *AddListenerWhiteListItemRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetLoadBalancerId(v string) *AddListenerWhiteListItemRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetOwnerAccount(v string) *AddListenerWhiteListItemRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetOwnerId(v int64) *AddListenerWhiteListItemRequest {
	s.OwnerId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetRegionId(v string) *AddListenerWhiteListItemRequest {
	s.RegionId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetResourceOwnerAccount(v string) *AddListenerWhiteListItemRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetResourceOwnerId(v int64) *AddListenerWhiteListItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetSourceItems(v string) *AddListenerWhiteListItemRequest {
	s.SourceItems = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) Validate() error {
	return dara.Validate(s)
}
