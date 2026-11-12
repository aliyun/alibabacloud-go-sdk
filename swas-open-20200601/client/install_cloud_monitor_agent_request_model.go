// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudMonitorAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InstallCloudMonitorAgentRequest
	GetClientToken() *string
	SetForce(v bool) *InstallCloudMonitorAgentRequest
	GetForce() *bool
	SetInstanceId(v string) *InstallCloudMonitorAgentRequest
	GetInstanceId() *string
	SetRegionId(v string) *InstallCloudMonitorAgentRequest
	GetRegionId() *string
}

type InstallCloudMonitorAgentRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly install the CloudMonitor agent. Valid values:
	//
	// - true (default value): forcibly installs the CloudMonitor agent.
	//
	// - false: does not forcibly install the CloudMonitor agent.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ae7106e68eb4402b0dcbd48a9de****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudMonitorAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudMonitorAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallCloudMonitorAgentRequest) GetForce() *bool {
	return s.Force
}

func (s *InstallCloudMonitorAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstallCloudMonitorAgentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallCloudMonitorAgentRequest) SetClientToken(v string) *InstallCloudMonitorAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) SetForce(v bool) *InstallCloudMonitorAgentRequest {
	s.Force = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) SetInstanceId(v string) *InstallCloudMonitorAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) SetRegionId(v string) *InstallCloudMonitorAgentRequest {
	s.RegionId = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) Validate() error {
	return dara.Validate(s)
}
