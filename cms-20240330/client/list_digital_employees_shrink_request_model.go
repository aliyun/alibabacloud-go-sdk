// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListDigitalEmployeesShrinkRequest
	GetDisplayName() *string
	SetEmployeeType(v string) *ListDigitalEmployeesShrinkRequest
	GetEmployeeType() *string
	SetMaxResults(v int32) *ListDigitalEmployeesShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListDigitalEmployeesShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListDigitalEmployeesShrinkRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListDigitalEmployeesShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *ListDigitalEmployeesShrinkRequest
	GetTagsShrink() *string
}

type ListDigitalEmployeesShrinkRequest struct {
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
	TagsShrink      *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListDigitalEmployeesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDigitalEmployeesShrinkRequest) GetEmployeeType() *string {
	return s.EmployeeType
}

func (s *ListDigitalEmployeesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDigitalEmployeesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListDigitalEmployeesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDigitalEmployeesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDigitalEmployeesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListDigitalEmployeesShrinkRequest) SetDisplayName(v string) *ListDigitalEmployeesShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) SetEmployeeType(v string) *ListDigitalEmployeesShrinkRequest {
	s.EmployeeType = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) SetMaxResults(v int32) *ListDigitalEmployeesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) SetName(v string) *ListDigitalEmployeesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) SetNextToken(v string) *ListDigitalEmployeesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) SetResourceGroupId(v string) *ListDigitalEmployeesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) SetTagsShrink(v string) *ListDigitalEmployeesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListDigitalEmployeesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
