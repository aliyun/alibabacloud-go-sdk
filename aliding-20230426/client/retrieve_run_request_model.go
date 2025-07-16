// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *RetrieveRunRequest
	GetAssistantId() *string
	SetOriginalAssistantId(v string) *RetrieveRunRequest
	GetOriginalAssistantId() *string
	SetRunId(v string) *RetrieveRunRequest
	GetRunId() *string
	SetSourceIdOfOriginalAssistantId(v string) *RetrieveRunRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *RetrieveRunRequest
	GetSourceTypeOfOriginalAssistantId() *string
}

type RetrieveRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// example:
	//
	// assistantId
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// runId123
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
}

func (s RetrieveRunRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRunRequest) GoString() string {
	return s.String()
}

func (s *RetrieveRunRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *RetrieveRunRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *RetrieveRunRequest) GetRunId() *string {
	return s.RunId
}

func (s *RetrieveRunRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *RetrieveRunRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *RetrieveRunRequest) SetAssistantId(v string) *RetrieveRunRequest {
	s.AssistantId = &v
	return s
}

func (s *RetrieveRunRequest) SetOriginalAssistantId(v string) *RetrieveRunRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *RetrieveRunRequest) SetRunId(v string) *RetrieveRunRequest {
	s.RunId = &v
	return s
}

func (s *RetrieveRunRequest) SetSourceIdOfOriginalAssistantId(v string) *RetrieveRunRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *RetrieveRunRequest) SetSourceTypeOfOriginalAssistantId(v string) *RetrieveRunRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *RetrieveRunRequest) Validate() error {
	return dara.Validate(s)
}
