// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListAuthorizationRulesRequestFilter) *ListAuthorizationRulesRequest
	GetFilter() []*ListAuthorizationRulesRequestFilter
	SetInstanceId(v string) *ListAuthorizationRulesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListAuthorizationRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesRequest
	GetNextToken() *string
}

type ListAuthorizationRulesRequest struct {
	// The filter conditions.
	Filter []*ListAuthorizationRulesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// - If you do not specify this parameter, the default value is 20.
	//
	// - The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that identifies the start of the next page of results.
	//
	// - If you do not specify this parameter, the system returns the first page of results.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAuthorizationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesRequest) GetFilter() []*ListAuthorizationRulesRequestFilter {
	return s.Filter
}

func (s *ListAuthorizationRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesRequest) SetFilter(v []*ListAuthorizationRulesRequestFilter) *ListAuthorizationRulesRequest {
	s.Filter = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetInstanceId(v string) *ListAuthorizationRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetMaxResults(v int32) *ListAuthorizationRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetNextToken(v string) *ListAuthorizationRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesRequest) Validate() error {
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

type ListAuthorizationRulesRequestFilter struct {
	// The name of the field to filter. Valid values:
	//
	// - AuthorizationRuleId: the ID of the authorization rule.
	//
	// - AuthorizationRuleNameStartWith: the leftmost characters of the authorization rule name.
	//
	// example:
	//
	// AuthorizationRuleId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of values for the field to filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListAuthorizationRulesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListAuthorizationRulesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListAuthorizationRulesRequestFilter) SetName(v string) *ListAuthorizationRulesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListAuthorizationRulesRequestFilter) SetValue(v []*string) *ListAuthorizationRulesRequestFilter {
	s.Value = v
	return s
}

func (s *ListAuthorizationRulesRequestFilter) Validate() error {
	return dara.Validate(s)
}
