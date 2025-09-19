// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentAccessKey(v string) *InstallCloudMonitorRequest
	GetAgentAccessKey() *string
	SetAgentSecretKey(v string) *InstallCloudMonitorRequest
	GetAgentSecretKey() *string
	SetArgusVersion(v string) *InstallCloudMonitorRequest
	GetArgusVersion() *string
	SetInstanceIdList(v []*string) *InstallCloudMonitorRequest
	GetInstanceIdList() []*string
	SetUuidList(v []*string) *InstallCloudMonitorRequest
	GetUuidList() []*string
}

type InstallCloudMonitorRequest struct {
	// The AccessKey ID that is required to install the CloudMonitor agent. You can call the [DescribeMonitoringAgentAccessKey](https://help.aliyun.com/document_detail/114948.html) operation to query the AccessKey ID.
	//
	// > This parameter is required only when you install the CloudMonitor agent on servers that are not deployed on Alibaba Cloud.
	//
	// example:
	//
	// usY*****R_U
	AgentAccessKey *string `json:"AgentAccessKey,omitempty" xml:"AgentAccessKey,omitempty"`
	// The AccessKey secret that is required to install the CloudMonitor agent. You can call the [DescribeMonitoringAgentAccessKey](https://help.aliyun.com/document_detail/114948.html) operation to query the AccessKey secret.
	//
	// > This parameter is required only when you install the CloudMonitor agent on servers that are not deployed on Alibaba Cloud.
	//
	// example:
	//
	// UCxF2R1sIO90XlU9****
	AgentSecretKey *string `json:"AgentSecretKey,omitempty" xml:"AgentSecretKey,omitempty"`
	// The version of the CloudMonitor agent that you want to install on the servers. For more information about the latest version of the CloudMonitor agent, see [Overview](https://help.aliyun.com/document_detail/183431.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.5.6
	ArgusVersion *string `json:"ArgusVersion,omitempty" xml:"ArgusVersion,omitempty"`
	// The IDs of the servers on which you want to install the CloudMonitor agent. Separate multiple IDs with commas (,).
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The UUIDs of the servers on which you want to install the CloudMonitor agent. Separate multiple UUIDs with commas (,).
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s InstallCloudMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudMonitorRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorRequest) GetAgentAccessKey() *string {
	return s.AgentAccessKey
}

func (s *InstallCloudMonitorRequest) GetAgentSecretKey() *string {
	return s.AgentSecretKey
}

func (s *InstallCloudMonitorRequest) GetArgusVersion() *string {
	return s.ArgusVersion
}

func (s *InstallCloudMonitorRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *InstallCloudMonitorRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *InstallCloudMonitorRequest) SetAgentAccessKey(v string) *InstallCloudMonitorRequest {
	s.AgentAccessKey = &v
	return s
}

func (s *InstallCloudMonitorRequest) SetAgentSecretKey(v string) *InstallCloudMonitorRequest {
	s.AgentSecretKey = &v
	return s
}

func (s *InstallCloudMonitorRequest) SetArgusVersion(v string) *InstallCloudMonitorRequest {
	s.ArgusVersion = &v
	return s
}

func (s *InstallCloudMonitorRequest) SetInstanceIdList(v []*string) *InstallCloudMonitorRequest {
	s.InstanceIdList = v
	return s
}

func (s *InstallCloudMonitorRequest) SetUuidList(v []*string) *InstallCloudMonitorRequest {
	s.UuidList = v
	return s
}

func (s *InstallCloudMonitorRequest) Validate() error {
	return dara.Validate(s)
}
