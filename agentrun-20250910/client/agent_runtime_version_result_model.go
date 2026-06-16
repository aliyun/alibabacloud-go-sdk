// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeVersionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgentRuntimeVersionResult
	GetCode() *string
	SetData(v *AgentRuntimeVersion) *AgentRuntimeVersionResult
	GetData() *AgentRuntimeVersion
	SetRequestId(v string) *AgentRuntimeVersionResult
	GetRequestId() *string
}

type AgentRuntimeVersionResult struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Detailed information about the agent runtime version
	Data *AgentRuntimeVersion `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request identifier used for troubleshooting
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AgentRuntimeVersionResult) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeVersionResult) GoString() string {
	return s.String()
}

func (s *AgentRuntimeVersionResult) GetCode() *string {
	return s.Code
}

func (s *AgentRuntimeVersionResult) GetData() *AgentRuntimeVersion {
	return s.Data
}

func (s *AgentRuntimeVersionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRuntimeVersionResult) SetCode(v string) *AgentRuntimeVersionResult {
	s.Code = &v
	return s
}

func (s *AgentRuntimeVersionResult) SetData(v *AgentRuntimeVersion) *AgentRuntimeVersionResult {
	s.Data = v
	return s
}

func (s *AgentRuntimeVersionResult) SetRequestId(v string) *AgentRuntimeVersionResult {
	s.RequestId = &v
	return s
}

func (s *AgentRuntimeVersionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
