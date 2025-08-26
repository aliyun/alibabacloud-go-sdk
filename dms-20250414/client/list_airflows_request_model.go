// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAirflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAirflowsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAirflowsRequest
	GetNextToken() *string
	SetPageIndex(v int32) *ListAirflowsRequest
	GetPageIndex() *int32
	SetSkip(v int32) *ListAirflowsRequest
	GetSkip() *int32
	SetWorkspaceId(v string) *ListAirflowsRequest
	GetWorkspaceId() *string
}

type ListAirflowsRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 5
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86302423828****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAirflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowsRequest) GoString() string {
	return s.String()
}

func (s *ListAirflowsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAirflowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAirflowsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListAirflowsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListAirflowsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAirflowsRequest) SetMaxResults(v int32) *ListAirflowsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAirflowsRequest) SetNextToken(v string) *ListAirflowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAirflowsRequest) SetPageIndex(v int32) *ListAirflowsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAirflowsRequest) SetSkip(v int32) *ListAirflowsRequest {
	s.Skip = &v
	return s
}

func (s *ListAirflowsRequest) SetWorkspaceId(v string) *ListAirflowsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAirflowsRequest) Validate() error {
	return dara.Validate(s)
}
