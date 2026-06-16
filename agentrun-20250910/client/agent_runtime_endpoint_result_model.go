// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeEndpointResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgentRuntimeEndpointResult
	GetCode() *string
	SetData(v *AgentRuntimeEndpoint) *AgentRuntimeEndpointResult
	GetData() *AgentRuntimeEndpoint
	SetRequestId(v string) *AgentRuntimeEndpointResult
	GetRequestId() *string
}

type AgentRuntimeEndpointResult struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Detailed information about the agent runtime endpoint
	Data *AgentRuntimeEndpoint `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request identifier used for troubleshooting
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AgentRuntimeEndpointResult) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeEndpointResult) GoString() string {
	return s.String()
}

func (s *AgentRuntimeEndpointResult) GetCode() *string {
	return s.Code
}

func (s *AgentRuntimeEndpointResult) GetData() *AgentRuntimeEndpoint {
	return s.Data
}

func (s *AgentRuntimeEndpointResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRuntimeEndpointResult) SetCode(v string) *AgentRuntimeEndpointResult {
	s.Code = &v
	return s
}

func (s *AgentRuntimeEndpointResult) SetData(v *AgentRuntimeEndpoint) *AgentRuntimeEndpointResult {
	s.Data = v
	return s
}

func (s *AgentRuntimeEndpointResult) SetRequestId(v string) *AgentRuntimeEndpointResult {
	s.RequestId = &v
	return s
}

func (s *AgentRuntimeEndpointResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
