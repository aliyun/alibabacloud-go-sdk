// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *ListApplicationsForAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetFilter(v []*ListApplicationsForAuthorizationRuleRequestFilter) *ListApplicationsForAuthorizationRuleRequest
	GetFilter() []*ListApplicationsForAuthorizationRuleRequestFilter
	SetInstanceId(v string) *ListApplicationsForAuthorizationRuleRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListApplicationsForAuthorizationRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationsForAuthorizationRuleRequest
	GetNextToken() *string
}

type ListApplicationsForAuthorizationRuleRequest struct {
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 过滤条件
	Filter []*ListApplicationsForAuthorizationRuleRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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

func (s ListApplicationsForAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListApplicationsForAuthorizationRuleRequest) GetFilter() []*ListApplicationsForAuthorizationRuleRequestFilter {
	return s.Filter
}

func (s *ListApplicationsForAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForAuthorizationRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsForAuthorizationRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsForAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *ListApplicationsForAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequest) SetFilter(v []*ListApplicationsForAuthorizationRuleRequestFilter) *ListApplicationsForAuthorizationRuleRequest {
	s.Filter = v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequest) SetInstanceId(v string) *ListApplicationsForAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequest) SetMaxResults(v int32) *ListApplicationsForAuthorizationRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequest) SetNextToken(v string) *ListApplicationsForAuthorizationRuleRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequest) Validate() error {
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

type ListApplicationsForAuthorizationRuleRequestFilter struct {
	// 过滤条件名称。
	//
	// example:
	//
	// ApplicationId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 过滤条件值。
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListApplicationsForAuthorizationRuleRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForAuthorizationRuleRequestFilter) GoString() string {
	return s.String()
}

func (s *ListApplicationsForAuthorizationRuleRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListApplicationsForAuthorizationRuleRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListApplicationsForAuthorizationRuleRequestFilter) SetName(v string) *ListApplicationsForAuthorizationRuleRequestFilter {
	s.Name = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequestFilter) SetValue(v []*string) *ListApplicationsForAuthorizationRuleRequestFilter {
	s.Value = v
	return s
}

func (s *ListApplicationsForAuthorizationRuleRequestFilter) Validate() error {
	return dara.Validate(s)
}
