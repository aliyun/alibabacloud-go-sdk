// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocIntroductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanCache(v bool) *RunDocIntroductionRequest
	GetCleanCache() *bool
	SetDocId(v string) *RunDocIntroductionRequest
	GetDocId() *string
	SetIntroductionPrompt(v string) *RunDocIntroductionRequest
	GetIntroductionPrompt() *string
	SetKeyPointPrompt(v string) *RunDocIntroductionRequest
	GetKeyPointPrompt() *string
	SetModelName(v string) *RunDocIntroductionRequest
	GetModelName() *string
	SetSessionId(v string) *RunDocIntroductionRequest
	GetSessionId() *string
	SetSummaryPrompt(v string) *RunDocIntroductionRequest
	GetSummaryPrompt() *string
	SetWorkspaceId(v string) *RunDocIntroductionRequest
	GetWorkspaceId() *string
	SetReferenceContent(v string) *RunDocIntroductionRequest
	GetReferenceContent() *string
}

type RunDocIntroductionRequest struct {
	CleanCache *bool `json:"CleanCache,omitempty" xml:"CleanCache,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DocId              *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	IntroductionPrompt *string `json:"IntroductionPrompt,omitempty" xml:"IntroductionPrompt,omitempty"`
	KeyPointPrompt     *string `json:"KeyPointPrompt,omitempty" xml:"KeyPointPrompt,omitempty"`
	ModelName          *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a3b5eb35-6b28-4cf9-ac09-1dec25ab4df6
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SummaryPrompt *string `json:"SummaryPrompt,omitempty" xml:"SummaryPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ReferenceContent *string `json:"referenceContent,omitempty" xml:"referenceContent,omitempty"`
}

func (s RunDocIntroductionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionRequest) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionRequest) GetCleanCache() *bool {
	return s.CleanCache
}

func (s *RunDocIntroductionRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunDocIntroductionRequest) GetIntroductionPrompt() *string {
	return s.IntroductionPrompt
}

func (s *RunDocIntroductionRequest) GetKeyPointPrompt() *string {
	return s.KeyPointPrompt
}

func (s *RunDocIntroductionRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocIntroductionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocIntroductionRequest) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *RunDocIntroductionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocIntroductionRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunDocIntroductionRequest) SetCleanCache(v bool) *RunDocIntroductionRequest {
	s.CleanCache = &v
	return s
}

func (s *RunDocIntroductionRequest) SetDocId(v string) *RunDocIntroductionRequest {
	s.DocId = &v
	return s
}

func (s *RunDocIntroductionRequest) SetIntroductionPrompt(v string) *RunDocIntroductionRequest {
	s.IntroductionPrompt = &v
	return s
}

func (s *RunDocIntroductionRequest) SetKeyPointPrompt(v string) *RunDocIntroductionRequest {
	s.KeyPointPrompt = &v
	return s
}

func (s *RunDocIntroductionRequest) SetModelName(v string) *RunDocIntroductionRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocIntroductionRequest) SetSessionId(v string) *RunDocIntroductionRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocIntroductionRequest) SetSummaryPrompt(v string) *RunDocIntroductionRequest {
	s.SummaryPrompt = &v
	return s
}

func (s *RunDocIntroductionRequest) SetWorkspaceId(v string) *RunDocIntroductionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocIntroductionRequest) SetReferenceContent(v string) *RunDocIntroductionRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunDocIntroductionRequest) Validate() error {
	return dara.Validate(s)
}
