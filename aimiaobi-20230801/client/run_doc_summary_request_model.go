// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanCache(v bool) *RunDocSummaryRequest
	GetCleanCache() *bool
	SetDocId(v string) *RunDocSummaryRequest
	GetDocId() *string
	SetModelName(v string) *RunDocSummaryRequest
	GetModelName() *string
	SetQuery(v string) *RunDocSummaryRequest
	GetQuery() *string
	SetRecommendContent(v string) *RunDocSummaryRequest
	GetRecommendContent() *string
	SetSessionId(v string) *RunDocSummaryRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunDocSummaryRequest
	GetWorkspaceId() *string
}

type RunDocSummaryRequest struct {
	CleanCache *bool `json:"CleanCache,omitempty" xml:"CleanCache,omitempty"`
	// example:
	//
	// 12345
	DocId            *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	ModelName        *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Query            *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RecommendContent *string `json:"RecommendContent,omitempty" xml:"RecommendContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0f56f98a-f2d8-47ec-98e9-1cbdcffa9539
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunDocSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryRequest) GoString() string {
	return s.String()
}

func (s *RunDocSummaryRequest) GetCleanCache() *bool {
	return s.CleanCache
}

func (s *RunDocSummaryRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunDocSummaryRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocSummaryRequest) GetQuery() *string {
	return s.Query
}

func (s *RunDocSummaryRequest) GetRecommendContent() *string {
	return s.RecommendContent
}

func (s *RunDocSummaryRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocSummaryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocSummaryRequest) SetCleanCache(v bool) *RunDocSummaryRequest {
	s.CleanCache = &v
	return s
}

func (s *RunDocSummaryRequest) SetDocId(v string) *RunDocSummaryRequest {
	s.DocId = &v
	return s
}

func (s *RunDocSummaryRequest) SetModelName(v string) *RunDocSummaryRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocSummaryRequest) SetQuery(v string) *RunDocSummaryRequest {
	s.Query = &v
	return s
}

func (s *RunDocSummaryRequest) SetRecommendContent(v string) *RunDocSummaryRequest {
	s.RecommendContent = &v
	return s
}

func (s *RunDocSummaryRequest) SetSessionId(v string) *RunDocSummaryRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocSummaryRequest) SetWorkspaceId(v string) *RunDocSummaryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocSummaryRequest) Validate() error {
	return dara.Validate(s)
}
