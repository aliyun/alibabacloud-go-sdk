// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *QueryAiVoiceAgentDetailNewRequest
	GetAgentId() *string
	SetBranchId(v string) *QueryAiVoiceAgentDetailNewRequest
	GetBranchId() *string
	SetVersionId(v string) *QueryAiVoiceAgentDetailNewRequest
	GetVersionId() *string
}

type QueryAiVoiceAgentDetailNewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 12
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// example:
	//
	// 21
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewRequest) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *QueryAiVoiceAgentDetailNewRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *QueryAiVoiceAgentDetailNewRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *QueryAiVoiceAgentDetailNewRequest) SetAgentId(v string) *QueryAiVoiceAgentDetailNewRequest {
	s.AgentId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewRequest) SetBranchId(v string) *QueryAiVoiceAgentDetailNewRequest {
	s.BranchId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewRequest) SetVersionId(v string) *QueryAiVoiceAgentDetailNewRequest {
	s.VersionId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewRequest) Validate() error {
	return dara.Validate(s)
}
