// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMultiDocIntroductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocIds(v []*string) *RunMultiDocIntroductionRequest
	GetDocIds() []*string
	SetKeyPointPrompt(v string) *RunMultiDocIntroductionRequest
	GetKeyPointPrompt() *string
	SetModelName(v string) *RunMultiDocIntroductionRequest
	GetModelName() *string
	SetSessionId(v string) *RunMultiDocIntroductionRequest
	GetSessionId() *string
	SetSummaryPrompt(v string) *RunMultiDocIntroductionRequest
	GetSummaryPrompt() *string
	SetWorkspaceId(v string) *RunMultiDocIntroductionRequest
	GetWorkspaceId() *string
}

type RunMultiDocIntroductionRequest struct {
	// This parameter is required.
	DocIds         []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	KeyPointPrompt *string   `json:"KeyPointPrompt,omitempty" xml:"KeyPointPrompt,omitempty"`
	ModelName      *string   `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
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

func (s RunMultiDocIntroductionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionRequest) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionRequest) GetDocIds() []*string {
	return s.DocIds
}

func (s *RunMultiDocIntroductionRequest) GetKeyPointPrompt() *string {
	return s.KeyPointPrompt
}

func (s *RunMultiDocIntroductionRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunMultiDocIntroductionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunMultiDocIntroductionRequest) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *RunMultiDocIntroductionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunMultiDocIntroductionRequest) SetDocIds(v []*string) *RunMultiDocIntroductionRequest {
	s.DocIds = v
	return s
}

func (s *RunMultiDocIntroductionRequest) SetKeyPointPrompt(v string) *RunMultiDocIntroductionRequest {
	s.KeyPointPrompt = &v
	return s
}

func (s *RunMultiDocIntroductionRequest) SetModelName(v string) *RunMultiDocIntroductionRequest {
	s.ModelName = &v
	return s
}

func (s *RunMultiDocIntroductionRequest) SetSessionId(v string) *RunMultiDocIntroductionRequest {
	s.SessionId = &v
	return s
}

func (s *RunMultiDocIntroductionRequest) SetSummaryPrompt(v string) *RunMultiDocIntroductionRequest {
	s.SummaryPrompt = &v
	return s
}

func (s *RunMultiDocIntroductionRequest) SetWorkspaceId(v string) *RunMultiDocIntroductionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunMultiDocIntroductionRequest) Validate() error {
	return dara.Validate(s)
}
