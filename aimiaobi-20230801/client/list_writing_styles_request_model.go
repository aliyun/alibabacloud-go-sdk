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
	// example:
	//
	// 100
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// media
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
