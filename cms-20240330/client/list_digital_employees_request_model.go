// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmployeeType(v string) *ListDigitalEmployeesRequest
	GetEmployeeType() *string
	SetMaxResults(v int32) *ListDigitalEmployeesRequest
	GetMaxResults() *int32
	SetName(v string) *ListDigitalEmployeesRequest
	GetName() *string
	SetNextToken(v string) *ListDigitalEmployeesRequest
	GetNextToken() *string
}

type ListDigitalEmployeesRequest struct {
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
}

func (s ListDigitalEmployeesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesRequest) GoString() string {
	return s.String()
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

func (s *ListDigitalEmployeesRequest) Validate() error {
	return dara.Validate(s)
}
