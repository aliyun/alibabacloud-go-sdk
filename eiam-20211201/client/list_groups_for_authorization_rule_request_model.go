// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *ListGroupsForAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetFilter(v []*ListGroupsForAuthorizationRuleRequestFilter) *ListGroupsForAuthorizationRuleRequest
	GetFilter() []*ListGroupsForAuthorizationRuleRequestFilter
	SetInstanceId(v string) *ListGroupsForAuthorizationRuleRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListGroupsForAuthorizationRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsForAuthorizationRuleRequest
	GetNextToken() *string
}

type ListGroupsForAuthorizationRuleRequest struct {
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 过滤条件
	Filter []*ListGroupsForAuthorizationRuleRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListGroupsForAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListGroupsForAuthorizationRuleRequest) GetFilter() []*ListGroupsForAuthorizationRuleRequestFilter {
	return s.Filter
}

func (s *ListGroupsForAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForAuthorizationRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsForAuthorizationRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *ListGroupsForAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequest) SetFilter(v []*ListGroupsForAuthorizationRuleRequestFilter) *ListGroupsForAuthorizationRuleRequest {
	s.Filter = v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequest) SetInstanceId(v string) *ListGroupsForAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequest) SetMaxResults(v int32) *ListGroupsForAuthorizationRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequest) SetNextToken(v string) *ListGroupsForAuthorizationRuleRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForAuthorizationRuleRequestFilter struct {
	// 过滤条件名称。
	//
	// example:
	//
	// GroupId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 过滤条件值。
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListGroupsForAuthorizationRuleRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForAuthorizationRuleRequestFilter) GoString() string {
	return s.String()
}

func (s *ListGroupsForAuthorizationRuleRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListGroupsForAuthorizationRuleRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListGroupsForAuthorizationRuleRequestFilter) SetName(v string) *ListGroupsForAuthorizationRuleRequestFilter {
	s.Name = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequestFilter) SetValue(v []*string) *ListGroupsForAuthorizationRuleRequestFilter {
	s.Value = v
	return s
}

func (s *ListGroupsForAuthorizationRuleRequestFilter) Validate() error {
	return dara.Validate(s)
}
