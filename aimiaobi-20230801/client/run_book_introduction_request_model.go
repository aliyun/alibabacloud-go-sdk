// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookIntroductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *RunBookIntroductionRequest
	GetDocId() *string
	SetKeyPointPrompt(v string) *RunBookIntroductionRequest
	GetKeyPointPrompt() *string
	SetSessionId(v string) *RunBookIntroductionRequest
	GetSessionId() *string
	SetSummaryPrompt(v string) *RunBookIntroductionRequest
	GetSummaryPrompt() *string
	SetWorkspaceId(v string) *RunBookIntroductionRequest
	GetWorkspaceId() *string
}

type RunBookIntroductionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3YQRatoe8phnpIsIE6z7DTPknhG8Fj
	DocId          *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	KeyPointPrompt *string `json:"KeyPointPrompt,omitempty" xml:"KeyPointPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0f56f98a-f2d8-47ec-98e9-1cbdcffa9539
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SummaryPrompt *string `json:"SummaryPrompt,omitempty" xml:"SummaryPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-vtmox6g2bhq2qv5c
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunBookIntroductionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionRequest) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunBookIntroductionRequest) GetKeyPointPrompt() *string {
	return s.KeyPointPrompt
}

func (s *RunBookIntroductionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunBookIntroductionRequest) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *RunBookIntroductionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunBookIntroductionRequest) SetDocId(v string) *RunBookIntroductionRequest {
	s.DocId = &v
	return s
}

func (s *RunBookIntroductionRequest) SetKeyPointPrompt(v string) *RunBookIntroductionRequest {
	s.KeyPointPrompt = &v
	return s
}

func (s *RunBookIntroductionRequest) SetSessionId(v string) *RunBookIntroductionRequest {
	s.SessionId = &v
	return s
}

func (s *RunBookIntroductionRequest) SetSummaryPrompt(v string) *RunBookIntroductionRequest {
	s.SummaryPrompt = &v
	return s
}

func (s *RunBookIntroductionRequest) SetWorkspaceId(v string) *RunBookIntroductionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunBookIntroductionRequest) Validate() error {
	return dara.Validate(s)
}
