// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentForClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *UpgradeAgentForClusterRequest
	GetXDebugId() *string
	SetAgentId(v string) *UpgradeAgentForClusterRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UpgradeAgentForClusterRequest
	GetAgentVersion() *string
	SetClusterId(v string) *UpgradeAgentForClusterRequest
	GetClusterId() *string
	SetXSysomInvokeSource(v string) *UpgradeAgentForClusterRequest
	GetXSysomInvokeSource() *string
}

type UpgradeAgentForClusterRequest struct {
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
	// > This must be the ID of an ACK cluster.
	//
	// example:
	//
	// c1c187fd513cb41a19876bac0e6b05212
	ClusterId          *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s UpgradeAgentForClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *UpgradeAgentForClusterRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpgradeAgentForClusterRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UpgradeAgentForClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeAgentForClusterRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *UpgradeAgentForClusterRequest) SetXDebugId(v string) *UpgradeAgentForClusterRequest {
	s.XDebugId = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetAgentId(v string) *UpgradeAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetAgentVersion(v string) *UpgradeAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetClusterId(v string) *UpgradeAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetXSysomInvokeSource(v string) *UpgradeAgentForClusterRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) Validate() error {
	return dara.Validate(s)
}
