// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileId(v string) *GetAgentProfileRequest
	GetAgentProfileId() *string
	SetAppIp(v string) *GetAgentProfileRequest
	GetAppIp() *string
	SetInstanceId(v string) *GetAgentProfileRequest
	GetInstanceId() *string
}

type GetAgentProfileRequest struct {
	// Agent profile ID.
	//
	// > You can obtain the agent profile ID of an already created scenario from the ChatbotId response parameter of DescribeScript.
	//
	// This parameter is required.
	//
	// example:
	//
	// d31794e2a51f47d2901b4094d88311d7
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// app_ip (system field, optional)
	//
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0ec0c897-b92c-40e4-9ad7-e6e4f5ce13bb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileRequest) GoString() string {
	return s.String()
}

func (s *GetAgentProfileRequest) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *GetAgentProfileRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *GetAgentProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentProfileRequest) SetAgentProfileId(v string) *GetAgentProfileRequest {
	s.AgentProfileId = &v
	return s
}

func (s *GetAgentProfileRequest) SetAppIp(v string) *GetAgentProfileRequest {
	s.AppIp = &v
	return s
}

func (s *GetAgentProfileRequest) SetInstanceId(v string) *GetAgentProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentProfileRequest) Validate() error {
	return dara.Validate(s)
}
