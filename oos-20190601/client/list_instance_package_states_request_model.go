// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePackageStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstancePackageStatesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListInstancePackageStatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancePackageStatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInstancePackageStatesRequest
	GetRegionId() *string
	SetTemplateNames(v string) *ListInstancePackageStatesRequest
	GetTemplateNames() *string
}

type ListInstancePackageStatesRequest struct {
	// ECS instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1cpoxxxwxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page size.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Pagination token.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctzxxxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of extension names
	//
	// example:
	//
	// ["template1","template2"]
	TemplateNames *string `json:"TemplateNames,omitempty" xml:"TemplateNames,omitempty"`
}

func (s ListInstancePackageStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePackageStatesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancePackageStatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancePackageStatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancePackageStatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancePackageStatesRequest) GetTemplateNames() *string {
	return s.TemplateNames
}

func (s *ListInstancePackageStatesRequest) SetInstanceId(v string) *ListInstancePackageStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetMaxResults(v int32) *ListInstancePackageStatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetNextToken(v string) *ListInstancePackageStatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetRegionId(v string) *ListInstancePackageStatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetTemplateNames(v string) *ListInstancePackageStatesRequest {
	s.TemplateNames = &v
	return s
}

func (s *ListInstancePackageStatesRequest) Validate() error {
	return dara.Validate(s)
}
