// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetListenerAccessControlStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlStatus(v string) *SetListenerAccessControlStatusRequest
	GetAccessControlStatus() *string
	SetListenerPort(v int32) *SetListenerAccessControlStatusRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *SetListenerAccessControlStatusRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *SetListenerAccessControlStatusRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetListenerAccessControlStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetListenerAccessControlStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetListenerAccessControlStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetListenerAccessControlStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetListenerAccessControlStatusRequest
	GetResourceOwnerId() *int64
}

type SetListenerAccessControlStatusRequest struct {
	// Specifies whether to enable the whitelist. Valid values:
	//
	// 	- **open_white_list**: enables the whitelist.
	//
	// 	- **close**: disables the whitelist.
	//
	// >  After the whitelist is enabled, if no IP address is added to the whitelist, the CLB instance does not distribute network traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// open_white_list
	AccessControlStatus *string `json:"AccessControlStatus,omitempty" xml:"AccessControlStatus,omitempty"`
	// The frontend port that is used by the CLB instance.
	//
	// Valid values: **1 to 65535**.
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
}

func (s SetListenerAccessControlStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetListenerAccessControlStatusRequest) GoString() string {
	return s.String()
}

func (s *SetListenerAccessControlStatusRequest) GetAccessControlStatus() *string {
	return s.AccessControlStatus
}

func (s *SetListenerAccessControlStatusRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetListenerAccessControlStatusRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *SetListenerAccessControlStatusRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetListenerAccessControlStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetListenerAccessControlStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetListenerAccessControlStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetListenerAccessControlStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetListenerAccessControlStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetListenerAccessControlStatusRequest) SetAccessControlStatus(v string) *SetListenerAccessControlStatusRequest {
	s.AccessControlStatus = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetListenerPort(v int32) *SetListenerAccessControlStatusRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetListenerProtocol(v string) *SetListenerAccessControlStatusRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetLoadBalancerId(v string) *SetListenerAccessControlStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetOwnerAccount(v string) *SetListenerAccessControlStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetOwnerId(v int64) *SetListenerAccessControlStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetRegionId(v string) *SetListenerAccessControlStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetResourceOwnerAccount(v string) *SetListenerAccessControlStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetResourceOwnerId(v int64) *SetListenerAccessControlStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) Validate() error {
	return dara.Validate(s)
}
