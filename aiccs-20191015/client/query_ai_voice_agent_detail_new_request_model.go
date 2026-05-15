// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *QueryAiVoiceAgentDetailNewRequest
	GetAgentId() *int64
	SetBranchId(v int64) *QueryAiVoiceAgentDetailNewRequest
	GetBranchId() *int64
	SetVersionId(v int64) *QueryAiVoiceAgentDetailNewRequest
	GetVersionId() *int64
}

type QueryAiVoiceAgentDetailNewRequest struct {
	// example:
	//
	// 12345678
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 12
	BranchId *int64 `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// example:
	//
	// 21
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewRequest) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *QueryAiVoiceAgentDetailNewRequest) GetBranchId() *int64 {
	return s.BranchId
}

func (s *QueryAiVoiceAgentDetailNewRequest) GetVersionId() *int64 {
	return s.VersionId
}

func (s *QueryAiVoiceAgentDetailNewRequest) SetAgentId(v int64) *QueryAiVoiceAgentDetailNewRequest {
	s.AgentId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewRequest) SetBranchId(v int64) *QueryAiVoiceAgentDetailNewRequest {
	s.BranchId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewRequest) SetVersionId(v int64) *QueryAiVoiceAgentDetailNewRequest {
	s.VersionId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewRequest) Validate() error {
	return dara.Validate(s)
}
