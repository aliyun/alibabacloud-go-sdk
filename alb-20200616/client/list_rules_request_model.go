// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListRulesRequest
	GetDirection() *string
	SetListenerIds(v []*string) *ListRulesRequest
	GetListenerIds() []*string
	SetLoadBalancerIds(v []*string) *ListRulesRequest
	GetLoadBalancerIds() []*string
	SetMaxResults(v int32) *ListRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRulesRequest
	GetNextToken() *string
	SetRuleIds(v []*string) *ListRulesRequest
	GetRuleIds() []*string
	SetTag(v []*ListRulesRequestTag) *ListRulesRequest
	GetTag() []*ListRulesRequestTag
}

type ListRulesRequest struct {
	// The direction to which the forwarding rule is applied. Valid values:
	//
	// 	- **Request*	- (default): The forwarding rule is applied to the client requests received by ALB.
	//
	// 	- **Response**: The forwarding rule is applied to the responses returned by backend servers.
	//
	// > You cannot set this parameter to Response if you use basic ALB instances.
	//
	// example:
	//
	// Request
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The Application Load Balancer (ALB) instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// Valid values: **1 to 100**.
	//
	// Default value: **20**. If you do not specify this parameter, the default value is used.
	//
	// > This parameter is optional.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The starting point of the current query. If you do not specify this parameter, the query starts from the beginning.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The forwarding rules.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The tag.
	Tag []*ListRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListRulesRequest) GetListenerIds() []*string {
	return s.ListenerIds
}

func (s *ListRulesRequest) GetLoadBalancerIds() []*string {
	return s.LoadBalancerIds
}

func (s *ListRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRulesRequest) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *ListRulesRequest) GetTag() []*ListRulesRequestTag {
	return s.Tag
}

func (s *ListRulesRequest) SetDirection(v string) *ListRulesRequest {
	s.Direction = &v
	return s
}

func (s *ListRulesRequest) SetListenerIds(v []*string) *ListRulesRequest {
	s.ListenerIds = v
	return s
}

func (s *ListRulesRequest) SetLoadBalancerIds(v []*string) *ListRulesRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListRulesRequest) SetMaxResults(v int32) *ListRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRulesRequest) SetNextToken(v string) *ListRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRulesRequest) SetRuleIds(v []*string) *ListRulesRequest {
	s.RuleIds = v
	return s
}

func (s *ListRulesRequest) SetTag(v []*ListRulesRequestTag) *ListRulesRequest {
	s.Tag = v
	return s
}

func (s *ListRulesRequest) Validate() error {
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

type ListRulesRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListRulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListRulesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListRulesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListRulesRequestTag) SetKey(v string) *ListRulesRequestTag {
	s.Key = &v
	return s
}

func (s *ListRulesRequestTag) SetValue(v string) *ListRulesRequestTag {
	s.Value = &v
	return s
}

func (s *ListRulesRequestTag) Validate() error {
	return dara.Validate(s)
}
