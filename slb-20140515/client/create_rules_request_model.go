// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *CreateRulesRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *CreateRulesRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *CreateRulesRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *CreateRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRulesRequest
	GetResourceOwnerId() *int64
	SetRuleList(v string) *CreateRulesRequest
	GetRuleList() *string
}

type CreateRulesRequest struct {
	// The frontend listener port that is used by the SLB instance.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the SLB instance.
	//
	// > This parameter is required if the same port is used by listeners that use different protocols.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1ca0zt07t934w******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The forwarding rules that you want to create. You can create at most 10 forwarding rules in each call. Each forwarding rule contains the following parameters:
	//
	// 	- **RuleName**: Required. The value must be of the STRING type. The name of the forwarding rule. The name must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_). Forwarding rule names must be unique within the same listener.
	//
	// 	- **Domain**: Optional. The value must be a string. The domain name that is associated with the forwarding rule. You must specify this parameter or the **URL*	- parameter.
	//
	// 	- **Url**: Optional. The value must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The value must be a string. The URL cannot be only a forward slash (/). However, it must start with a forward slash (/). You must specify this parameter or the **Domain*	- parameter.
	//
	// 	- **VServerGroupId**: Required. The value must be a string. The ID of the vServer group to be specified in the forwarding rule.
	//
	// >  You must specify at least one between the `Domain` and `URL` parameters. You can also specify both. The combination of `Domain` and `Url` must be unique within the same listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"RuleName":"Rule2","Domain":"test.com","VServerGroupId":"rsp-bp114ni******"}]
	RuleList *string `json:"RuleList,omitempty" xml:"RuleList,omitempty"`
}

func (s CreateRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateRulesRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateRulesRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *CreateRulesRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRulesRequest) GetRuleList() *string {
	return s.RuleList
}

func (s *CreateRulesRequest) SetListenerPort(v int32) *CreateRulesRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateRulesRequest) SetListenerProtocol(v string) *CreateRulesRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateRulesRequest) SetLoadBalancerId(v string) *CreateRulesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateRulesRequest) SetOwnerAccount(v string) *CreateRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRulesRequest) SetOwnerId(v int64) *CreateRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRulesRequest) SetRegionId(v string) *CreateRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRulesRequest) SetResourceOwnerAccount(v string) *CreateRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRulesRequest) SetResourceOwnerId(v int64) *CreateRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRulesRequest) SetRuleList(v string) *CreateRulesRequest {
	s.RuleList = &v
	return s
}

func (s *CreateRulesRequest) Validate() error {
	return dara.Validate(s)
}
