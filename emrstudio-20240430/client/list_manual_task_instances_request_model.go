// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListManualTaskInstancesRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListManualTaskInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListManualTaskInstancesRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListManualTaskInstancesRequest
	GetSearchVal() *string
	SetStartTime(v string) *ListManualTaskInstancesRequest
	GetStartTime() *string
	SetStatus(v string) *ListManualTaskInstancesRequest
	GetStatus() *string
	SetWorkspaceId(v string) *ListManualTaskInstancesRequest
	GetWorkspaceId() *string
}

type ListManualTaskInstancesRequest struct {
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListManualTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManualTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListManualTaskInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListManualTaskInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManualTaskInstancesRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListManualTaskInstancesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListManualTaskInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListManualTaskInstancesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListManualTaskInstancesRequest) SetEndTime(v string) *ListManualTaskInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetMaxResults(v int32) *ListManualTaskInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetNextToken(v string) *ListManualTaskInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetSearchVal(v string) *ListManualTaskInstancesRequest {
	s.SearchVal = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetStartTime(v string) *ListManualTaskInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetStatus(v string) *ListManualTaskInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetWorkspaceId(v string) *ListManualTaskInstancesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListManualTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
