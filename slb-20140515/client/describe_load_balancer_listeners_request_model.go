// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeLoadBalancerListenersRequest
	GetDescription() *string
	SetListenerPort(v int32) *DescribeLoadBalancerListenersRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *DescribeLoadBalancerListenersRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v []*string) *DescribeLoadBalancerListenersRequest
	GetLoadBalancerId() []*string
	SetMaxResults(v int32) *DescribeLoadBalancerListenersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLoadBalancerListenersRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeLoadBalancerListenersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerListenersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLoadBalancerListenersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerListenersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerListenersRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeLoadBalancerListenersRequestTag) *DescribeLoadBalancerListenersRequest
	GetTag() []*DescribeLoadBalancerListenersRequestTag
}

type DescribeLoadBalancerListenersRequest struct {
	// The description of the listener.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// HTTPS_443
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The frontend port that is used by the CLB instance.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The protocol used by the listener. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// example:
	//
	// http
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The IDs of the CLB instances. You can specify up to 10 instance IDs.
	//
	// >  If you do not use the SDK to call this operation, use the LoadBalancerId.N parameter.
	//
	// example:
	//
	// lb-123wrwer
	LoadBalancerId []*string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty" type:"Repeated"`
	// The number of entries to return in each call.
	//
	// Valid values: **1*	- to **100**. If you do not specify this parameter, the default value **20*	- is used.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query and no subsequent queries are to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the most recent region list.
	//
	// >  If the endpoint of the selected region is slb.aliyuncs.com, the `RegionId` parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeLoadBalancerListenersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerListenersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerListenersRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerListenersRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeLoadBalancerListenersRequest) GetLoadBalancerId() []*string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerListenersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLoadBalancerListenersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLoadBalancerListenersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerListenersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerListenersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerListenersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerListenersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerListenersRequest) GetTag() []*DescribeLoadBalancerListenersRequestTag {
	return s.Tag
}

func (s *DescribeLoadBalancerListenersRequest) SetDescription(v string) *DescribeLoadBalancerListenersRequest {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetListenerPort(v int32) *DescribeLoadBalancerListenersRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetListenerProtocol(v string) *DescribeLoadBalancerListenersRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetLoadBalancerId(v []*string) *DescribeLoadBalancerListenersRequest {
	s.LoadBalancerId = v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetMaxResults(v int32) *DescribeLoadBalancerListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetNextToken(v string) *DescribeLoadBalancerListenersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetOwnerAccount(v string) *DescribeLoadBalancerListenersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetOwnerId(v int64) *DescribeLoadBalancerListenersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetRegionId(v string) *DescribeLoadBalancerListenersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerListenersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerListenersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetTag(v []*DescribeLoadBalancerListenersRequestTag) *DescribeLoadBalancerListenersRequest {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerListenersRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1 to 20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLoadBalancerListenersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLoadBalancerListenersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLoadBalancerListenersRequestTag) SetKey(v string) *DescribeLoadBalancerListenersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequestTag) SetValue(v string) *DescribeLoadBalancerListenersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequestTag) Validate() error {
	return dara.Validate(s)
}
