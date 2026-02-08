// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentProfileTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileTemplateId(v string) *GetAgentProfileTemplateRequest
	GetAgentProfileTemplateId() *string
	SetAppIp(v string) *GetAgentProfileTemplateRequest
	GetAppIp() *string
}

type GetAgentProfileTemplateRequest struct {
	// Agent Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// app_ip (system field; optional)
	//
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
}

func (s GetAgentProfileTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAgentProfileTemplateRequest) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *GetAgentProfileTemplateRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *GetAgentProfileTemplateRequest) SetAgentProfileTemplateId(v string) *GetAgentProfileTemplateRequest {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *GetAgentProfileTemplateRequest) SetAppIp(v string) *GetAgentProfileTemplateRequest {
	s.AppIp = &v
	return s
}

func (s *GetAgentProfileTemplateRequest) Validate() error {
	return dara.Validate(s)
}
