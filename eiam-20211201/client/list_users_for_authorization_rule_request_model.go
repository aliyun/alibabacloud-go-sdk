// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *ListUsersForAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetFilter(v []*ListUsersForAuthorizationRuleRequestFilter) *ListUsersForAuthorizationRuleRequest
	GetFilter() []*ListUsersForAuthorizationRuleRequestFilter
	SetInstanceId(v string) *ListUsersForAuthorizationRuleRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListUsersForAuthorizationRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersForAuthorizationRuleRequest
	GetNextToken() *string
}

type ListUsersForAuthorizationRuleRequest struct {
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 过滤条件
	Filter []*ListUsersForAuthorizationRuleRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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

func (s ListUsersForAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListUsersForAuthorizationRuleRequest) GetFilter() []*ListUsersForAuthorizationRuleRequestFilter {
	return s.Filter
}

func (s *ListUsersForAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForAuthorizationRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersForAuthorizationRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersForAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *ListUsersForAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListUsersForAuthorizationRuleRequest) SetFilter(v []*ListUsersForAuthorizationRuleRequestFilter) *ListUsersForAuthorizationRuleRequest {
	s.Filter = v
	return s
}

func (s *ListUsersForAuthorizationRuleRequest) SetInstanceId(v string) *ListUsersForAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForAuthorizationRuleRequest) SetMaxResults(v int32) *ListUsersForAuthorizationRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForAuthorizationRuleRequest) SetNextToken(v string) *ListUsersForAuthorizationRuleRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersForAuthorizationRuleRequest) Validate() error {
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

type ListUsersForAuthorizationRuleRequestFilter struct {
	// 过滤条件名称。
	//
	// example:
	//
	// UserId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 过滤条件值。
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListUsersForAuthorizationRuleRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForAuthorizationRuleRequestFilter) GoString() string {
	return s.String()
}

func (s *ListUsersForAuthorizationRuleRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListUsersForAuthorizationRuleRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListUsersForAuthorizationRuleRequestFilter) SetName(v string) *ListUsersForAuthorizationRuleRequestFilter {
	s.Name = &v
	return s
}

func (s *ListUsersForAuthorizationRuleRequestFilter) SetValue(v []*string) *ListUsersForAuthorizationRuleRequestFilter {
	s.Value = v
	return s
}

func (s *ListUsersForAuthorizationRuleRequestFilter) Validate() error {
	return dara.Validate(s)
}
