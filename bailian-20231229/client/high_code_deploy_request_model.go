// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHighCodeDeployRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentDesc(v string) *HighCodeDeployRequest
	GetAgentDesc() *string
	SetAgentId(v string) *HighCodeDeployRequest
	GetAgentId() *string
	SetAgentName(v string) *HighCodeDeployRequest
	GetAgentName() *string
	SetSourceCodeName(v string) *HighCodeDeployRequest
	GetSourceCodeName() *string
	SetSourceCodeOssUrl(v string) *HighCodeDeployRequest
	GetSourceCodeOssUrl() *string
	SetTelemetryEnabled(v bool) *HighCodeDeployRequest
	GetTelemetryEnabled() *bool
}

type HighCodeDeployRequest struct {
	// example:
	//
	// 智能客服agent
	AgentDesc *string `json:"agentDesc,omitempty" xml:"agentDesc,omitempty"`
	// example:
	//
	// 2000013
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// mma_test_02
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// example:
	//
	// xxxxx.whl
	SourceCodeName *string `json:"sourceCodeName,omitempty" xml:"sourceCodeName,omitempty"`
	// example:
	//
	// https://bailian-application-developer-code.oss-cn-beijing.aliyuncs.com/xxxxx.whl
	SourceCodeOssUrl *string `json:"sourceCodeOssUrl,omitempty" xml:"sourceCodeOssUrl,omitempty"`
	// example:
	//
	// 0
	TelemetryEnabled *bool `json:"telemetryEnabled,omitempty" xml:"telemetryEnabled,omitempty"`
}

func (s HighCodeDeployRequest) String() string {
	return dara.Prettify(s)
}

func (s HighCodeDeployRequest) GoString() string {
	return s.String()
}

func (s *HighCodeDeployRequest) GetAgentDesc() *string {
	return s.AgentDesc
}

func (s *HighCodeDeployRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *HighCodeDeployRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *HighCodeDeployRequest) GetSourceCodeName() *string {
	return s.SourceCodeName
}

func (s *HighCodeDeployRequest) GetSourceCodeOssUrl() *string {
	return s.SourceCodeOssUrl
}

func (s *HighCodeDeployRequest) GetTelemetryEnabled() *bool {
	return s.TelemetryEnabled
}

func (s *HighCodeDeployRequest) SetAgentDesc(v string) *HighCodeDeployRequest {
	s.AgentDesc = &v
	return s
}

func (s *HighCodeDeployRequest) SetAgentId(v string) *HighCodeDeployRequest {
	s.AgentId = &v
	return s
}

func (s *HighCodeDeployRequest) SetAgentName(v string) *HighCodeDeployRequest {
	s.AgentName = &v
	return s
}

func (s *HighCodeDeployRequest) SetSourceCodeName(v string) *HighCodeDeployRequest {
	s.SourceCodeName = &v
	return s
}

func (s *HighCodeDeployRequest) SetSourceCodeOssUrl(v string) *HighCodeDeployRequest {
	s.SourceCodeOssUrl = &v
	return s
}

func (s *HighCodeDeployRequest) SetTelemetryEnabled(v bool) *HighCodeDeployRequest {
	s.TelemetryEnabled = &v
	return s
}

func (s *HighCodeDeployRequest) Validate() error {
	return dara.Validate(s)
}
