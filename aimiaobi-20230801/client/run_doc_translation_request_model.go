// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocTranslationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanCache(v bool) *RunDocTranslationRequest
	GetCleanCache() *bool
	SetDocId(v string) *RunDocTranslationRequest
	GetDocId() *string
	SetModelName(v string) *RunDocTranslationRequest
	GetModelName() *string
	SetRecommendContent(v string) *RunDocTranslationRequest
	GetRecommendContent() *string
	SetSessionId(v string) *RunDocTranslationRequest
	GetSessionId() *string
	SetTransType(v string) *RunDocTranslationRequest
	GetTransType() *string
	SetWorkspaceId(v string) *RunDocTranslationRequest
	GetWorkspaceId() *string
}

type RunDocTranslationRequest struct {
	CleanCache *bool `json:"CleanCache,omitempty" xml:"CleanCache,omitempty"`
	// example:
	//
	// 12345
	DocId            *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	ModelName        *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	RecommendContent *string `json:"RecommendContent,omitempty" xml:"RecommendContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2e6b3987-f743-4d4c-8326-d9c41a6af3ee
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// toChinese
	//
	// toEnglish
	TransType *string `json:"TransType,omitempty" xml:"TransType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunDocTranslationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationRequest) GoString() string {
	return s.String()
}

func (s *RunDocTranslationRequest) GetCleanCache() *bool {
	return s.CleanCache
}

func (s *RunDocTranslationRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunDocTranslationRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocTranslationRequest) GetRecommendContent() *string {
	return s.RecommendContent
}

func (s *RunDocTranslationRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocTranslationRequest) GetTransType() *string {
	return s.TransType
}

func (s *RunDocTranslationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocTranslationRequest) SetCleanCache(v bool) *RunDocTranslationRequest {
	s.CleanCache = &v
	return s
}

func (s *RunDocTranslationRequest) SetDocId(v string) *RunDocTranslationRequest {
	s.DocId = &v
	return s
}

func (s *RunDocTranslationRequest) SetModelName(v string) *RunDocTranslationRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocTranslationRequest) SetRecommendContent(v string) *RunDocTranslationRequest {
	s.RecommendContent = &v
	return s
}

func (s *RunDocTranslationRequest) SetSessionId(v string) *RunDocTranslationRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocTranslationRequest) SetTransType(v string) *RunDocTranslationRequest {
	s.TransType = &v
	return s
}

func (s *RunDocTranslationRequest) SetWorkspaceId(v string) *RunDocTranslationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocTranslationRequest) Validate() error {
	return dara.Validate(s)
}
