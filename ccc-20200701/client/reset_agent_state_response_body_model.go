// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgentStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetAgentStateResponseBody
	GetCode() *string
	SetData(v *ResetAgentStateResponseBodyData) *ResetAgentStateResponseBody
	GetData() *ResetAgentStateResponseBodyData
	SetHttpStatusCode(v int32) *ResetAgentStateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetAgentStateResponseBody
	GetMessage() *string
	SetParams(v []*string) *ResetAgentStateResponseBody
	GetParams() []*string
	SetRequestId(v string) *ResetAgentStateResponseBody
	GetRequestId() *string
}

type ResetAgentStateResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ResetAgentStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Reserved for future use.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAgentStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAgentStateResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAgentStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetAgentStateResponseBody) GetData() *ResetAgentStateResponseBodyData {
	return s.Data
}

func (s *ResetAgentStateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetAgentStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetAgentStateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ResetAgentStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAgentStateResponseBody) SetCode(v string) *ResetAgentStateResponseBody {
	s.Code = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetData(v *ResetAgentStateResponseBodyData) *ResetAgentStateResponseBody {
	s.Data = v
	return s
}

func (s *ResetAgentStateResponseBody) SetHttpStatusCode(v int32) *ResetAgentStateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetMessage(v string) *ResetAgentStateResponseBody {
	s.Message = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetParams(v []*string) *ResetAgentStateResponseBody {
	s.Params = v
	return s
}

func (s *ResetAgentStateResponseBody) SetRequestId(v string) *ResetAgentStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAgentStateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetAgentStateResponseBodyData struct {
	// The reason code for the agent\\"s break. This can be a system-defined or a custom code. System-defined codes include:
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// The ID of the registered device, such as a browser-based WebRTC softphone or a physical phone. An agent can register only one device at a time.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The agent\\"s extension.
	//
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the call. This parameter is returned only if the agent is on a call.
	//
	// example:
	//
	// 无
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the outbound-only mode is enabled for the agent.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// The IDs of the skill groups that the agent is signed in to.
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// The ID of the agent.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The current state of the agent.
	//
	// example:
	//
	// OFFLINE
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// The agent\\"s work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ResetAgentStateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetAgentStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetAgentStateResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ResetAgentStateResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ResetAgentStateResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ResetAgentStateResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetAgentStateResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ResetAgentStateResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ResetAgentStateResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ResetAgentStateResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ResetAgentStateResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *ResetAgentStateResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ResetAgentStateResponseBodyData) SetBreakCode(v string) *ResetAgentStateResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetDeviceId(v string) *ResetAgentStateResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetExtension(v string) *ResetAgentStateResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetInstanceId(v string) *ResetAgentStateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetJobId(v string) *ResetAgentStateResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetOutboundScenario(v bool) *ResetAgentStateResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ResetAgentStateResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetUserId(v string) *ResetAgentStateResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetUserState(v string) *ResetAgentStateResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetWorkMode(v string) *ResetAgentStateResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
