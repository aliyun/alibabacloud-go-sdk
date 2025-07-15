// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExecutionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *ListResourceExecutionStatusRequest
	GetExecutionId() *string
	SetMaxResults(v int32) *ListResourceExecutionStatusRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceExecutionStatusRequest
	GetNextToken() *string
	SetRegionId(v string) *ListResourceExecutionStatusRequest
	GetRegionId() *string
}

type ListResourceExecutionStatusRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxxxxxxxxxxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListResourceExecutionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExecutionStatusRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListResourceExecutionStatusRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceExecutionStatusRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceExecutionStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceExecutionStatusRequest) SetExecutionId(v string) *ListResourceExecutionStatusRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetMaxResults(v int32) *ListResourceExecutionStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetNextToken(v string) *ListResourceExecutionStatusRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetRegionId(v string) *ListResourceExecutionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) Validate() error {
	return dara.Validate(s)
}
