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
	// This parameter is required.
	//
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
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
