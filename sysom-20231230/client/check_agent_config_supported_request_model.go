// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAgentConfigSupportedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CheckAgentConfigSupportedRequest
	GetAgentId() *string
	SetCheckType(v string) *CheckAgentConfigSupportedRequest
	GetCheckType() *string
	SetClusters(v string) *CheckAgentConfigSupportedRequest
	GetClusters() *string
	SetConfigId(v string) *CheckAgentConfigSupportedRequest
	GetConfigId() *string
	SetContent(v string) *CheckAgentConfigSupportedRequest
	GetContent() *string
	SetInstances(v string) *CheckAgentConfigSupportedRequest
	GetInstances() *string
}

type CheckAgentConfigSupportedRequest struct {
	// The ID of the component to update. You can obtain the component ID from the ListAgents operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The configuration dimension to check, such as cluster or node.
	//
	// example:
	//
	// instance
	CheckType *string `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c726e7497b7d246d09d8644ed62effacc
	Clusters *string `json:"clusters,omitempty" xml:"clusters,omitempty"`
	// The ID of the component configuration to install.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4d5ecf02-3d01-11f0-8c90-00163e1029a
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
	// The configuration content.
	//
	// example:
	//
	// {\\"code\\": \\"BABA\\", \\"changeRate\\": \\"-1.434%\\", \\"exchange\\": \\"NYSE\\", \\"price\\": \\"112.690\\", \\"changePrice\\": \\"-$1.640\\", \\"currency\\": \\"USD\\", \\"date\\": \\"2026-06-11T20:01 EDT\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The list of instances. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// i-bp10ktsji8vsbd8gco5i,i-bp1163pp7k7njfggvuqp,i-bp12ytprghv6rwzfgypm
	Instances *string `json:"instances,omitempty" xml:"instances,omitempty"`
}

func (s CheckAgentConfigSupportedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentConfigSupportedRequest) GoString() string {
	return s.String()
}

func (s *CheckAgentConfigSupportedRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CheckAgentConfigSupportedRequest) GetCheckType() *string {
	return s.CheckType
}

func (s *CheckAgentConfigSupportedRequest) GetClusters() *string {
	return s.Clusters
}

func (s *CheckAgentConfigSupportedRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *CheckAgentConfigSupportedRequest) GetContent() *string {
	return s.Content
}

func (s *CheckAgentConfigSupportedRequest) GetInstances() *string {
	return s.Instances
}

func (s *CheckAgentConfigSupportedRequest) SetAgentId(v string) *CheckAgentConfigSupportedRequest {
	s.AgentId = &v
	return s
}

func (s *CheckAgentConfigSupportedRequest) SetCheckType(v string) *CheckAgentConfigSupportedRequest {
	s.CheckType = &v
	return s
}

func (s *CheckAgentConfigSupportedRequest) SetClusters(v string) *CheckAgentConfigSupportedRequest {
	s.Clusters = &v
	return s
}

func (s *CheckAgentConfigSupportedRequest) SetConfigId(v string) *CheckAgentConfigSupportedRequest {
	s.ConfigId = &v
	return s
}

func (s *CheckAgentConfigSupportedRequest) SetContent(v string) *CheckAgentConfigSupportedRequest {
	s.Content = &v
	return s
}

func (s *CheckAgentConfigSupportedRequest) SetInstances(v string) *CheckAgentConfigSupportedRequest {
	s.Instances = &v
	return s
}

func (s *CheckAgentConfigSupportedRequest) Validate() error {
	return dara.Validate(s)
}
