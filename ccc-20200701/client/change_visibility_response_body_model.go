// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVisibilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeVisibilityResponseBody
	GetCode() *string
	SetData(v *ChangeVisibilityResponseBodyData) *ChangeVisibilityResponseBody
	GetData() *ChangeVisibilityResponseBodyData
	SetHttpStatusCode(v int32) *ChangeVisibilityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ChangeVisibilityResponseBody
	GetMessage() *string
	SetParams(v []*string) *ChangeVisibilityResponseBody
	GetParams() []*string
	SetRequestId(v string) *ChangeVisibilityResponseBody
	GetRequestId() *string
}

type ChangeVisibilityResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ChangeVisibilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Internal service issue. Detail:.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The response parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 24BE19E8-BF7D-4992-A35E-15EBA874F2E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeVisibilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeVisibilityResponseBody) GetData() *ChangeVisibilityResponseBodyData {
	return s.Data
}

func (s *ChangeVisibilityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChangeVisibilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeVisibilityResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ChangeVisibilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeVisibilityResponseBody) SetCode(v string) *ChangeVisibilityResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeVisibilityResponseBody) SetData(v *ChangeVisibilityResponseBodyData) *ChangeVisibilityResponseBody {
	s.Data = v
	return s
}

func (s *ChangeVisibilityResponseBody) SetHttpStatusCode(v int32) *ChangeVisibilityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeVisibilityResponseBody) SetMessage(v string) *ChangeVisibilityResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeVisibilityResponseBody) SetParams(v []*string) *ChangeVisibilityResponseBody {
	s.Params = v
	return s
}

func (s *ChangeVisibilityResponseBody) SetRequestId(v string) *ChangeVisibilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeVisibilityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeVisibilityResponseBodyData struct {
	// The break code.
	//
	// example:
	//
	// 会议
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// The device ID. This can be the ID of a browser-based WebRTC softphone or a physical phone. Only one device can be registered at a time.
	//
	// example:
	//
	// Yealink SIP-T23G 44.84.203.6
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The agent\\"s extension number.
	//
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// szpczf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call ID. If populated, this field indicates that the agent is in a call.
	//
	// example:
	//
	// job-330557290544431104
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the agent is in outbound-only mode.
	//
	// example:
	//
	// False
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// The IDs of the skill groups to which the agent is signed in.
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// The agent ID.
	//
	// example:
	//
	// sam@szpczf
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The agent state.
	//
	// example:
	//
	// Ready
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// The work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ChangeVisibilityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ChangeVisibilityResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ChangeVisibilityResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ChangeVisibilityResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeVisibilityResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ChangeVisibilityResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ChangeVisibilityResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ChangeVisibilityResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ChangeVisibilityResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *ChangeVisibilityResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ChangeVisibilityResponseBodyData) SetBreakCode(v string) *ChangeVisibilityResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetDeviceId(v string) *ChangeVisibilityResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetExtension(v string) *ChangeVisibilityResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetInstanceId(v string) *ChangeVisibilityResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetJobId(v string) *ChangeVisibilityResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetOutboundScenario(v bool) *ChangeVisibilityResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ChangeVisibilityResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetUserId(v string) *ChangeVisibilityResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetUserState(v string) *ChangeVisibilityResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) SetWorkMode(v string) *ChangeVisibilityResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ChangeVisibilityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
