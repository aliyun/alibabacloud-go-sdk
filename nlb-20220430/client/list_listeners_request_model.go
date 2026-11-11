// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerIds(v []*string) *ListListenersRequest
	GetListenerIds() []*string
	SetListenerProtocol(v string) *ListListenersRequest
	GetListenerProtocol() *string
	SetLoadBalancerIds(v []*string) *ListListenersRequest
	GetLoadBalancerIds() []*string
	SetMaxResults(v int32) *ListListenersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListListenersRequest
	GetRegionId() *string
	SetSecSensorEnabled(v string) *ListListenersRequest
	GetSecSensorEnabled() *string
	SetTag(v []*ListListenersRequestTag) *ListListenersRequest
	GetTag() []*ListListenersRequestTag
}

type ListListenersRequest struct {
	// The instance IDs. You can specify up to 20 instance IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The listener protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	//
	// example:
	//
	// TCPSSL
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The IDs of the Network Load Balancer (NLB) instances. You can specify up to 20 instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The number of entries to return per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Valid values:
	//
	// - If this is the first query or no next query is required, you do not need to specify this parameter.
	//
	// - If a next query is required, the value is the NextToken value returned from the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable second-level monitoring. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	SecSensorEnabled *string `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The list of tags.
	Tag []*ListListenersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListListenersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) GetListenerIds() []*string {
	return s.ListenerIds
}

func (s *ListListenersRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *ListListenersRequest) GetLoadBalancerIds() []*string {
	return s.LoadBalancerIds
}

func (s *ListListenersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListListenersRequest) GetSecSensorEnabled() *string {
	return s.SecSensorEnabled
}

func (s *ListListenersRequest) GetTag() []*ListListenersRequestTag {
	return s.Tag
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetListenerProtocol(v string) *ListListenersRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenersRequest) SetRegionId(v string) *ListListenersRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenersRequest) SetSecSensorEnabled(v string) *ListListenersRequest {
	s.SecSensorEnabled = &v
	return s
}

func (s *ListListenersRequest) SetTag(v []*ListListenersRequestTag) *ListListenersRequest {
	s.Tag = v
	return s
}

func (s *ListListenersRequest) Validate() error {
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

type ListListenersRequestTag struct {
	// The tag key of the instance. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListListenersRequestTag) GoString() string {
	return s.String()
}

func (s *ListListenersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListListenersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListListenersRequestTag) SetKey(v string) *ListListenersRequestTag {
	s.Key = &v
	return s
}

func (s *ListListenersRequestTag) SetValue(v string) *ListListenersRequestTag {
	s.Value = &v
	return s
}

func (s *ListListenersRequestTag) Validate() error {
	return dara.Validate(s)
}
