// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMultiDocIntroductionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocIdsShrink(v string) *RunMultiDocIntroductionShrinkRequest
	GetDocIdsShrink() *string
	SetKeyPointPrompt(v string) *RunMultiDocIntroductionShrinkRequest
	GetKeyPointPrompt() *string
	SetModelName(v string) *RunMultiDocIntroductionShrinkRequest
	GetModelName() *string
	SetSessionId(v string) *RunMultiDocIntroductionShrinkRequest
	GetSessionId() *string
	SetSummaryPrompt(v string) *RunMultiDocIntroductionShrinkRequest
	GetSummaryPrompt() *string
	SetWorkspaceId(v string) *RunMultiDocIntroductionShrinkRequest
	GetWorkspaceId() *string
}

type RunMultiDocIntroductionShrinkRequest struct {
	// This parameter is required.
	DocIdsShrink   *string `json:"DocIds,omitempty" xml:"DocIds,omitempty"`
	KeyPointPrompt *string `json:"KeyPointPrompt,omitempty" xml:"KeyPointPrompt,omitempty"`
	ModelName      *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 75bf82fa-b71b-45d7-ae40-0b00e496cd9e
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SummaryPrompt *string `json:"SummaryPrompt,omitempty" xml:"SummaryPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunMultiDocIntroductionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionShrinkRequest) GetDocIdsShrink() *string {
	return s.DocIdsShrink
}

func (s *RunMultiDocIntroductionShrinkRequest) GetKeyPointPrompt() *string {
	return s.KeyPointPrompt
}

func (s *RunMultiDocIntroductionShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunMultiDocIntroductionShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunMultiDocIntroductionShrinkRequest) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *RunMultiDocIntroductionShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunMultiDocIntroductionShrinkRequest) SetDocIdsShrink(v string) *RunMultiDocIntroductionShrinkRequest {
	s.DocIdsShrink = &v
	return s
}

func (s *RunMultiDocIntroductionShrinkRequest) SetKeyPointPrompt(v string) *RunMultiDocIntroductionShrinkRequest {
	s.KeyPointPrompt = &v
	return s
}

func (s *RunMultiDocIntroductionShrinkRequest) SetModelName(v string) *RunMultiDocIntroductionShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *RunMultiDocIntroductionShrinkRequest) SetSessionId(v string) *RunMultiDocIntroductionShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunMultiDocIntroductionShrinkRequest) SetSummaryPrompt(v string) *RunMultiDocIntroductionShrinkRequest {
	s.SummaryPrompt = &v
	return s
}

func (s *RunMultiDocIntroductionShrinkRequest) SetWorkspaceId(v string) *RunMultiDocIntroductionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunMultiDocIntroductionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
