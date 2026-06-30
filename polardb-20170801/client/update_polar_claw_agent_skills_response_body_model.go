// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentSkillsResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentSkillsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpdatePolarClawAgentSkillsResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdatePolarClawAgentSkillsResponseBody
	GetMessage() *string
	SetOk(v bool) *UpdatePolarClawAgentSkillsResponseBody
	GetOk() *bool
	SetRequestId(v string) *UpdatePolarClawAgentSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []*string) *UpdatePolarClawAgentSkillsResponseBody
	GetSkills() []*string
}

type UpdatePolarClawAgentSkillsResponseBody struct {
	// Agent ID
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The return code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The updated skill list.
	Skills []*string `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
}

func (s UpdatePolarClawAgentSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolarClawAgentSkillsResponseBody) GetSkills() []*string {
	return s.Skills
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetAgentId(v string) *UpdatePolarClawAgentSkillsResponseBody {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetApplicationId(v string) *UpdatePolarClawAgentSkillsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetCode(v int32) *UpdatePolarClawAgentSkillsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetMessage(v string) *UpdatePolarClawAgentSkillsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetOk(v bool) *UpdatePolarClawAgentSkillsResponseBody {
	s.Ok = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetRequestId(v string) *UpdatePolarClawAgentSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) SetSkills(v []*string) *UpdatePolarClawAgentSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
