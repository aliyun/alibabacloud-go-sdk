// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWritingStylesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWritingStylesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWritingStylesRequest
	GetNextToken() *string
	SetScene(v string) *ListWritingStylesRequest
	GetScene() *string
	SetWorkspaceId(v string) *ListWritingStylesRequest
	GetWorkspaceId() *string
}

type ListWritingStylesRequest struct {
	// The maximum number of results to return.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Pagination token
	//
	// example:
	//
	// 下一页token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Filters by writing scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// media
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// [The workspace ID.](https://help.aliyun.com/document_detail/2782167.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWritingStylesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWritingStylesRequest) GoString() string {
	return s.String()
}

func (s *ListWritingStylesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWritingStylesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWritingStylesRequest) GetScene() *string {
	return s.Scene
}

func (s *ListWritingStylesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWritingStylesRequest) SetMaxResults(v int32) *ListWritingStylesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWritingStylesRequest) SetNextToken(v string) *ListWritingStylesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWritingStylesRequest) SetScene(v string) *ListWritingStylesRequest {
	s.Scene = &v
	return s
}

func (s *ListWritingStylesRequest) SetWorkspaceId(v string) *ListWritingStylesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWritingStylesRequest) Validate() error {
	return dara.Validate(s)
}
