// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateMonitorAgentProcessRequest
	GetInstanceId() *string
	SetProcessName(v string) *CreateMonitorAgentProcessRequest
	GetProcessName() *string
	SetProcessUser(v string) *CreateMonitorAgentProcessRequest
	GetProcessUser() *string
	SetRegionId(v string) *CreateMonitorAgentProcessRequest
	GetRegionId() *string
}

type CreateMonitorAgentProcessRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze2d6j5uhg20x47****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The process name.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliYunDun
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The user who launches the process.
	//
	// example:
	//
	// admin
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorAgentProcessRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMonitorAgentProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateMonitorAgentProcessRequest) GetProcessUser() *string {
	return s.ProcessUser
}

func (s *CreateMonitorAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorAgentProcessRequest) SetInstanceId(v string) *CreateMonitorAgentProcessRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitorAgentProcessRequest) SetProcessName(v string) *CreateMonitorAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateMonitorAgentProcessRequest) SetProcessUser(v string) *CreateMonitorAgentProcessRequest {
	s.ProcessUser = &v
	return s
}

func (s *CreateMonitorAgentProcessRequest) SetRegionId(v string) *CreateMonitorAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}
