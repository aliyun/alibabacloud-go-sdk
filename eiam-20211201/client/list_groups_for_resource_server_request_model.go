// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForResourceServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListGroupsForResourceServerRequest
	GetApplicationId() *string
	SetFilter(v []*ListGroupsForResourceServerRequestFilter) *ListGroupsForResourceServerRequest
	GetFilter() []*ListGroupsForResourceServerRequestFilter
	SetInstanceId(v string) *ListGroupsForResourceServerRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListGroupsForResourceServerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsForResourceServerRequest
	GetNextToken() *string
	SetResourceServerScopeId(v string) *ListGroupsForResourceServerRequest
	GetResourceServerScopeId() *string
}

type ListGroupsForResourceServerRequest struct {
	// The ID of the resource server application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The filter conditions.
	Filter []*ListGroupsForResourceServerRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. You do not need to provide this parameter for the first request. For subsequent requests, set this to the `NextToken` value from the previous response.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the scope.
	//
	// example:
	//
	// ress_nbte4bb3qqqnaq73rlmkqixxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
}

func (s ListGroupsForResourceServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListGroupsForResourceServerRequest) GetFilter() []*ListGroupsForResourceServerRequestFilter {
	return s.Filter
}

func (s *ListGroupsForResourceServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForResourceServerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsForResourceServerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForResourceServerRequest) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListGroupsForResourceServerRequest) SetApplicationId(v string) *ListGroupsForResourceServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetFilter(v []*ListGroupsForResourceServerRequestFilter) *ListGroupsForResourceServerRequest {
	s.Filter = v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetInstanceId(v string) *ListGroupsForResourceServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetMaxResults(v int32) *ListGroupsForResourceServerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetNextToken(v string) *ListGroupsForResourceServerRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetResourceServerScopeId(v string) *ListGroupsForResourceServerRequest {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) Validate() error {
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

type ListGroupsForResourceServerRequestFilter struct {
	// The name of the filter condition. The only valid value is GroupIds.
	//
	// example:
	//
	// GroupIds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values for the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListGroupsForResourceServerRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerRequestFilter) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListGroupsForResourceServerRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListGroupsForResourceServerRequestFilter) SetName(v string) *ListGroupsForResourceServerRequestFilter {
	s.Name = &v
	return s
}

func (s *ListGroupsForResourceServerRequestFilter) SetValue(v []*string) *ListGroupsForResourceServerRequestFilter {
	s.Value = v
	return s
}

func (s *ListGroupsForResourceServerRequestFilter) Validate() error {
	return dara.Validate(s)
}
