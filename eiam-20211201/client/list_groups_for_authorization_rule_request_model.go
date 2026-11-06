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
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The filter conditions.
	Filter []*ListGroupsForAuthorizationRuleRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of records to return per page.
	//
	// - If this parameter is not specified, the default value is 20.
	//
	// - The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to identify the starting position of the next page.
	//
	// - If this parameter is not specified, the query starts from the first page by default.
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
	// The name of the filter field. Valid values:
	//
	// - GroupId: the group ID
	//
	// example:
	//
	// GroupId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter field values.
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
