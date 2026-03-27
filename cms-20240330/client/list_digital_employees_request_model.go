// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListDigitalEmployeesRequest
	GetDisplayName() *string
	SetEmployeeType(v string) *ListDigitalEmployeesRequest
	GetEmployeeType() *string
	SetMaxResults(v int32) *ListDigitalEmployeesRequest
	GetMaxResults() *int32
	SetName(v string) *ListDigitalEmployeesRequest
	GetName() *string
	SetNextToken(v string) *ListDigitalEmployeesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListDigitalEmployeesRequest
	GetResourceGroupId() *string
	SetTags(v []*Tag) *ListDigitalEmployeesRequest
	GetTags() []*Tag
}

type ListDigitalEmployeesRequest struct {
	DisplayName  *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	EmployeeType *string `json:"employeeType,omitempty" xml:"employeeType,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// xxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*Tag  `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListDigitalEmployeesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesRequest) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDigitalEmployeesRequest) GetEmployeeType() *string {
	return s.EmployeeType
}

func (s *ListDigitalEmployeesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDigitalEmployeesRequest) GetName() *string {
	return s.Name
}

func (s *ListDigitalEmployeesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDigitalEmployeesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDigitalEmployeesRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *ListDigitalEmployeesRequest) SetDisplayName(v string) *ListDigitalEmployeesRequest {
	s.DisplayName = &v
	return s
}

func (s *ListDigitalEmployeesRequest) SetEmployeeType(v string) *ListDigitalEmployeesRequest {
	s.EmployeeType = &v
	return s
}

func (s *ListDigitalEmployeesRequest) SetMaxResults(v int32) *ListDigitalEmployeesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDigitalEmployeesRequest) SetName(v string) *ListDigitalEmployeesRequest {
	s.Name = &v
	return s
}

func (s *ListDigitalEmployeesRequest) SetNextToken(v string) *ListDigitalEmployeesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDigitalEmployeesRequest) SetResourceGroupId(v string) *ListDigitalEmployeesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDigitalEmployeesRequest) SetTags(v []*Tag) *ListDigitalEmployeesRequest {
	s.Tags = v
	return s
}

func (s *ListDigitalEmployeesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
