// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAudioNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SubmitAudioNoteRequest
	GetAgentId() *string
	SetBizId(v string) *SubmitAudioNoteRequest
	GetBizId() *string
	SetFilePath(v string) *SubmitAudioNoteRequest
	GetFilePath() *string
	SetLlmModelId(v int64) *SubmitAudioNoteRequest
	GetLlmModelId() *int64
}

type SubmitAudioNoteRequest struct {
	// The ID of the notes agent. Specify the ID of a published recording notes agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The custom task ID defined by the business side. This ID is used to associate external business records during callbacks or troubleshooting.
	//
	// example:
	//
	// biz-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The storage path of the recording file in OSS. Use the FilePath returned by the upload address retrieval operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// audio-note/100/uuid/test.wav
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The ID of the LLM model used for notes inference. If this parameter is not specified, the default model configuration of the agent is used.
	//
	// example:
	//
	// 88
	LlmModelId *int64 `json:"LlmModelId,omitempty" xml:"LlmModelId,omitempty"`
}

func (s SubmitAudioNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAudioNoteRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioNoteRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SubmitAudioNoteRequest) GetBizId() *string {
	return s.BizId
}

func (s *SubmitAudioNoteRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *SubmitAudioNoteRequest) GetLlmModelId() *int64 {
	return s.LlmModelId
}

func (s *SubmitAudioNoteRequest) SetAgentId(v string) *SubmitAudioNoteRequest {
	s.AgentId = &v
	return s
}

func (s *SubmitAudioNoteRequest) SetBizId(v string) *SubmitAudioNoteRequest {
	s.BizId = &v
	return s
}

func (s *SubmitAudioNoteRequest) SetFilePath(v string) *SubmitAudioNoteRequest {
	s.FilePath = &v
	return s
}

func (s *SubmitAudioNoteRequest) SetLlmModelId(v int64) *SubmitAudioNoteRequest {
	s.LlmModelId = &v
	return s
}

func (s *SubmitAudioNoteRequest) Validate() error {
	return dara.Validate(s)
}
