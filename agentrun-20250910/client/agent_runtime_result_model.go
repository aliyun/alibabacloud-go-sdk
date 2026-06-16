// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgentRuntimeResult
	GetCode() *string
	SetData(v *AgentRuntime) *AgentRuntimeResult
	GetData() *AgentRuntime
	SetRequestId(v string) *AgentRuntimeResult
	GetRequestId() *string
}

type AgentRuntimeResult struct {
	// Returns `SUCCESS` if the operation is successful; otherwise, returns an error code such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Detailed information about the agent runtime.
	//
	// example:
	//
	// {}
	Data *AgentRuntime `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request ID for issue tracking.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AgentRuntimeResult) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeResult) GoString() string {
	return s.String()
}

func (s *AgentRuntimeResult) GetCode() *string {
	return s.Code
}

func (s *AgentRuntimeResult) GetData() *AgentRuntime {
	return s.Data
}

func (s *AgentRuntimeResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRuntimeResult) SetCode(v string) *AgentRuntimeResult {
	s.Code = &v
	return s
}

func (s *AgentRuntimeResult) SetData(v *AgentRuntime) *AgentRuntimeResult {
	s.Data = v
	return s
}

func (s *AgentRuntimeResult) SetRequestId(v string) *AgentRuntimeResult {
	s.RequestId = &v
	return s
}

func (s *AgentRuntimeResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
