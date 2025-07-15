// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListParameterVersionsRequest
	GetMaxResults() *int32
	SetName(v string) *ListParameterVersionsRequest
	GetName() *string
	SetNextToken(v string) *ListParameterVersionsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListParameterVersionsRequest
	GetRegionId() *string
	SetShareType(v string) *ListParameterVersionsRequest
	GetShareType() *string
}

type ListParameterVersionsRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListParameterVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListParameterVersionsRequest) GetName() *string {
	return s.Name
}

func (s *ListParameterVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListParameterVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListParameterVersionsRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListParameterVersionsRequest) SetMaxResults(v int32) *ListParameterVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParameterVersionsRequest) SetName(v string) *ListParameterVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsRequest) SetNextToken(v string) *ListParameterVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListParameterVersionsRequest) SetRegionId(v string) *ListParameterVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListParameterVersionsRequest) SetShareType(v string) *ListParameterVersionsRequest {
	s.ShareType = &v
	return s
}

func (s *ListParameterVersionsRequest) Validate() error {
	return dara.Validate(s)
}
