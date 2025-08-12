// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitoringAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateMonitoringAgentProcessRequest
	GetInstanceId() *string
	SetProcessName(v string) *CreateMonitoringAgentProcessRequest
	GetProcessName() *string
	SetProcessUser(v string) *CreateMonitoringAgentProcessRequest
	GetProcessUser() *string
	SetRegionId(v string) *CreateMonitoringAgentProcessRequest
	GetRegionId() *string
}

type CreateMonitoringAgentProcessRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze51wjtwox01r8g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// java
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The user who launches the process.
	//
	// example:
	//
	// admin
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitoringAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitoringAgentProcessRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMonitoringAgentProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateMonitoringAgentProcessRequest) GetProcessUser() *string {
	return s.ProcessUser
}

func (s *CreateMonitoringAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitoringAgentProcessRequest) SetInstanceId(v string) *CreateMonitoringAgentProcessRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitoringAgentProcessRequest) SetProcessName(v string) *CreateMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateMonitoringAgentProcessRequest) SetProcessUser(v string) *CreateMonitoringAgentProcessRequest {
	s.ProcessUser = &v
	return s
}

func (s *CreateMonitoringAgentProcessRequest) SetRegionId(v string) *CreateMonitoringAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitoringAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}
