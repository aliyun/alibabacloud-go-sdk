// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPptTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCareerId(v int32) *ListPptTemplatesRequest
	GetCareerId() *int32
	SetColourId(v int32) *ListPptTemplatesRequest
	GetColourId() *int32
	SetMaxResults(v int32) *ListPptTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPptTemplatesRequest
	GetNextToken() *string
	SetSceneId(v int32) *ListPptTemplatesRequest
	GetSceneId() *int32
	SetStyleId(v int32) *ListPptTemplatesRequest
	GetStyleId() *int32
	SetWorkspaceId(v string) *ListPptTemplatesRequest
	GetWorkspaceId() *string
}

type ListPptTemplatesRequest struct {
	// example:
	//
	// 1
	CareerId *int32 `json:"CareerId,omitempty" xml:"CareerId,omitempty"`
	// example:
	//
	// 1
	ColourId   *int32 `json:"ColourId,omitempty" xml:"ColourId,omitempty"`
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// +CBOXvu2YLxC6DOua8Qupg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 7
	SceneId *int32 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 1
	StyleId *int32 `json:"StyleId,omitempty" xml:"StyleId,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPptTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPptTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPptTemplatesRequest) GetCareerId() *int32 {
	return s.CareerId
}

func (s *ListPptTemplatesRequest) GetColourId() *int32 {
	return s.ColourId
}

func (s *ListPptTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPptTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPptTemplatesRequest) GetSceneId() *int32 {
	return s.SceneId
}

func (s *ListPptTemplatesRequest) GetStyleId() *int32 {
	return s.StyleId
}

func (s *ListPptTemplatesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListPptTemplatesRequest) SetCareerId(v int32) *ListPptTemplatesRequest {
	s.CareerId = &v
	return s
}

func (s *ListPptTemplatesRequest) SetColourId(v int32) *ListPptTemplatesRequest {
	s.ColourId = &v
	return s
}

func (s *ListPptTemplatesRequest) SetMaxResults(v int32) *ListPptTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPptTemplatesRequest) SetNextToken(v string) *ListPptTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPptTemplatesRequest) SetSceneId(v int32) *ListPptTemplatesRequest {
	s.SceneId = &v
	return s
}

func (s *ListPptTemplatesRequest) SetStyleId(v int32) *ListPptTemplatesRequest {
	s.StyleId = &v
	return s
}

func (s *ListPptTemplatesRequest) SetWorkspaceId(v string) *ListPptTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListPptTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
