// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *ListAuthorizationResourcesRequest
	GetAuthorizationRuleId() *string
	SetFilter(v []*ListAuthorizationResourcesRequestFilter) *ListAuthorizationResourcesRequest
	GetFilter() []*ListAuthorizationResourcesRequestFilter
	SetInstanceId(v string) *ListAuthorizationResourcesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListAuthorizationResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationResourcesRequest
	GetNextToken() *string
}

type ListAuthorizationResourcesRequest struct {
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The filter conditions.
	Filter []*ListAuthorizationResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page.
	//
	// - Default value: 20.
	//
	// - Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that identifies the start position of the next page.
	//
	// - If you do not specify this parameter, the query starts from the first page.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAuthorizationResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListAuthorizationResourcesRequest) GetFilter() []*ListAuthorizationResourcesRequestFilter {
	return s.Filter
}

func (s *ListAuthorizationResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationResourcesRequest) SetAuthorizationRuleId(v string) *ListAuthorizationResourcesRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationResourcesRequest) SetFilter(v []*ListAuthorizationResourcesRequestFilter) *ListAuthorizationResourcesRequest {
	s.Filter = v
	return s
}

func (s *ListAuthorizationResourcesRequest) SetInstanceId(v string) *ListAuthorizationResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationResourcesRequest) SetMaxResults(v int32) *ListAuthorizationResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationResourcesRequest) SetNextToken(v string) *ListAuthorizationResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationResourcesRequest) Validate() error {
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

type ListAuthorizationResourcesRequestFilter struct {
	// The name of the filter field. Valid values:
	//
	// - AuthorizationResourceEntityType: the resource entity type associated with the authorization resource.
	//
	// - AuthorizationResourceEntityId: the resource entity ID associated with the authorization resource.
	//
	// example:
	//
	// AuthorizationResourceEntityId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the filter field.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListAuthorizationResourcesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListAuthorizationResourcesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListAuthorizationResourcesRequestFilter) SetName(v string) *ListAuthorizationResourcesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListAuthorizationResourcesRequestFilter) SetValue(v []*string) *ListAuthorizationResourcesRequestFilter {
	s.Value = v
	return s
}

func (s *ListAuthorizationResourcesRequestFilter) Validate() error {
	return dara.Validate(s)
}
