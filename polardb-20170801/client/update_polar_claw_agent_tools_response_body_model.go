// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentToolsResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentToolsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpdatePolarClawAgentToolsResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdatePolarClawAgentToolsResponseBody
	GetMessage() *string
	SetOk(v bool) *UpdatePolarClawAgentToolsResponseBody
	GetOk() *bool
	SetRequestId(v string) *UpdatePolarClawAgentToolsResponseBody
	GetRequestId() *string
	SetTools(v *UpdatePolarClawAgentToolsResponseBodyTools) *UpdatePolarClawAgentToolsResponseBody
	GetTools() *UpdatePolarClawAgentToolsResponseBodyTools
}

type UpdatePolarClawAgentToolsResponseBody struct {
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
	// The status code.
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
	// The updated tool configuration.
	Tools *UpdatePolarClawAgentToolsResponseBodyTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Struct"`
}

func (s UpdatePolarClawAgentToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentToolsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolarClawAgentToolsResponseBody) GetTools() *UpdatePolarClawAgentToolsResponseBodyTools {
	return s.Tools
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetAgentId(v string) *UpdatePolarClawAgentToolsResponseBody {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetApplicationId(v string) *UpdatePolarClawAgentToolsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetCode(v int32) *UpdatePolarClawAgentToolsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetMessage(v string) *UpdatePolarClawAgentToolsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetOk(v bool) *UpdatePolarClawAgentToolsResponseBody {
	s.Ok = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetRequestId(v string) *UpdatePolarClawAgentToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) SetTools(v *UpdatePolarClawAgentToolsResponseBodyTools) *UpdatePolarClawAgentToolsResponseBody {
	s.Tools = v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBody) Validate() error {
	if s.Tools != nil {
		if err := s.Tools.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawAgentToolsResponseBodyTools struct {
	// The list of explicitly allowed tools.
	Allow []*string `json:"Allow,omitempty" xml:"Allow,omitempty" type:"Repeated"`
	// The list of additionally allowed tools.
	AlsoAllow []*string `json:"AlsoAllow,omitempty" xml:"AlsoAllow,omitempty" type:"Repeated"`
	// The list of denied tools.
	Deny []*string `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Repeated"`
	// The tool profile.
	//
	// example:
	//
	// coding
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdatePolarClawAgentToolsResponseBodyTools) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentToolsResponseBodyTools) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) GetAllow() []*string {
	return s.Allow
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) GetAlsoAllow() []*string {
	return s.AlsoAllow
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) GetDeny() []*string {
	return s.Deny
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) GetProfile() *string {
	return s.Profile
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) SetAllow(v []*string) *UpdatePolarClawAgentToolsResponseBodyTools {
	s.Allow = v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) SetAlsoAllow(v []*string) *UpdatePolarClawAgentToolsResponseBodyTools {
	s.AlsoAllow = v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) SetDeny(v []*string) *UpdatePolarClawAgentToolsResponseBodyTools {
	s.Deny = v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) SetProfile(v string) *UpdatePolarClawAgentToolsResponseBodyTools {
	s.Profile = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponseBodyTools) Validate() error {
	return dara.Validate(s)
}
