// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsForResourceServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListOrganizationalUnitsForResourceServerRequest
	GetApplicationId() *string
	SetFilter(v []*ListOrganizationalUnitsForResourceServerRequestFilter) *ListOrganizationalUnitsForResourceServerRequest
	GetFilter() []*ListOrganizationalUnitsForResourceServerRequestFilter
	SetInstanceId(v string) *ListOrganizationalUnitsForResourceServerRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListOrganizationalUnitsForResourceServerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOrganizationalUnitsForResourceServerRequest
	GetNextToken() *string
	SetResourceServerScopeId(v string) *ListOrganizationalUnitsForResourceServerRequest
	GetResourceServerScopeId() *string
}

type ListOrganizationalUnitsForResourceServerRequest struct {
	// The resource server application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The list of filter conditions.
	Filter []*ListOrganizationalUnitsForResourceServerRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The instance ID.
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
	// The pagination token. To retrieve the next page of results, set this parameter to the NextToken value from the previous response.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource server scope.
	//
	// example:
	//
	// ress_nbte4bb3qqqnaq73rlmkqixxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
}

func (s ListOrganizationalUnitsForResourceServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForResourceServerRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForResourceServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListOrganizationalUnitsForResourceServerRequest) GetFilter() []*ListOrganizationalUnitsForResourceServerRequestFilter {
	return s.Filter
}

func (s *ListOrganizationalUnitsForResourceServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitsForResourceServerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOrganizationalUnitsForResourceServerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOrganizationalUnitsForResourceServerRequest) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListOrganizationalUnitsForResourceServerRequest) SetApplicationId(v string) *ListOrganizationalUnitsForResourceServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequest) SetFilter(v []*ListOrganizationalUnitsForResourceServerRequestFilter) *ListOrganizationalUnitsForResourceServerRequest {
	s.Filter = v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequest) SetInstanceId(v string) *ListOrganizationalUnitsForResourceServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequest) SetMaxResults(v int32) *ListOrganizationalUnitsForResourceServerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequest) SetNextToken(v string) *ListOrganizationalUnitsForResourceServerRequest {
	s.NextToken = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequest) SetResourceServerScopeId(v string) *ListOrganizationalUnitsForResourceServerRequest {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequest) Validate() error {
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

type ListOrganizationalUnitsForResourceServerRequestFilter struct {
	// The filter key.
	//
	// example:
	//
	// OrganizationalUnitIds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListOrganizationalUnitsForResourceServerRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForResourceServerRequestFilter) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForResourceServerRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListOrganizationalUnitsForResourceServerRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListOrganizationalUnitsForResourceServerRequestFilter) SetName(v string) *ListOrganizationalUnitsForResourceServerRequestFilter {
	s.Name = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequestFilter) SetValue(v []*string) *ListOrganizationalUnitsForResourceServerRequestFilter {
	s.Value = v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerRequestFilter) Validate() error {
	return dara.Validate(s)
}
