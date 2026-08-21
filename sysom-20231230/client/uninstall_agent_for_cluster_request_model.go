// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentForClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *UninstallAgentForClusterRequest
	GetXDebugId() *string
	SetAgentId(v string) *UninstallAgentForClusterRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UninstallAgentForClusterRequest
	GetAgentVersion() *string
	SetClusterId(v string) *UninstallAgentForClusterRequest
	GetClusterId() *string
	SetXSysomInvokeSource(v string) *UninstallAgentForClusterRequest
	GetXSysomInvokeSource() *string
}

type UninstallAgentForClusterRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The component ID.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// The component version.
	//
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// The cluster ID.
	//
	// > This cluster ID must be the ID of an ACK cluster.
	//
	// example:
	//
	// c822f83bb45994ddbac9326b4c2f04f35
	ClusterId          *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s UninstallAgentForClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *UninstallAgentForClusterRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UninstallAgentForClusterRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UninstallAgentForClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UninstallAgentForClusterRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *UninstallAgentForClusterRequest) SetXDebugId(v string) *UninstallAgentForClusterRequest {
	s.XDebugId = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetAgentId(v string) *UninstallAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetAgentVersion(v string) *UninstallAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetClusterId(v string) *UninstallAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetXSysomInvokeSource(v string) *UninstallAgentForClusterRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *UninstallAgentForClusterRequest) Validate() error {
	return dara.Validate(s)
}
