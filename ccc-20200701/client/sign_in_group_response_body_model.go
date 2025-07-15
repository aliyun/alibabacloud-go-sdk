// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SignInGroupResponseBody
	GetCode() *string
	SetData(v *SignInGroupResponseBodyData) *SignInGroupResponseBody
	GetData() *SignInGroupResponseBodyData
	SetHttpStatusCode(v int32) *SignInGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SignInGroupResponseBody
	GetMessage() *string
	SetParams(v []*string) *SignInGroupResponseBody
	GetParams() []*string
	SetRequestId(v string) *SignInGroupResponseBody
	GetRequestId() *string
}

type SignInGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SignInGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SignInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SignInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SignInGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *SignInGroupResponseBody) GetData() *SignInGroupResponseBodyData {
	return s.Data
}

func (s *SignInGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SignInGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SignInGroupResponseBody) GetParams() []*string {
	return s.Params
}

func (s *SignInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SignInGroupResponseBody) SetCode(v string) *SignInGroupResponseBody {
	s.Code = &v
	return s
}

func (s *SignInGroupResponseBody) SetData(v *SignInGroupResponseBodyData) *SignInGroupResponseBody {
	s.Data = v
	return s
}

func (s *SignInGroupResponseBody) SetHttpStatusCode(v int32) *SignInGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SignInGroupResponseBody) SetMessage(v string) *SignInGroupResponseBody {
	s.Message = &v
	return s
}

func (s *SignInGroupResponseBody) SetParams(v []*string) *SignInGroupResponseBody {
	s.Params = v
	return s
}

func (s *SignInGroupResponseBody) SetRequestId(v string) *SignInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignInGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type SignInGroupResponseBodyData struct {
	// example:
	//
	// Warm-up
	BreakCode    *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	ChatDeviceId *string `json:"ChatDeviceId,omitempty" xml:"ChatDeviceId,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// false
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// READY
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s SignInGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SignInGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *SignInGroupResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *SignInGroupResponseBodyData) GetChatDeviceId() *string {
	return s.ChatDeviceId
}

func (s *SignInGroupResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SignInGroupResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *SignInGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SignInGroupResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SignInGroupResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *SignInGroupResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *SignInGroupResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *SignInGroupResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *SignInGroupResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SignInGroupResponseBodyData) SetBreakCode(v string) *SignInGroupResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetChatDeviceId(v string) *SignInGroupResponseBodyData {
	s.ChatDeviceId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetDeviceId(v string) *SignInGroupResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetExtension(v string) *SignInGroupResponseBodyData {
	s.Extension = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetInstanceId(v string) *SignInGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetJobId(v string) *SignInGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetOutboundScenario(v bool) *SignInGroupResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetSignedSkillGroupIdList(v []*string) *SignInGroupResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *SignInGroupResponseBodyData) SetUserId(v string) *SignInGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetUserState(v string) *SignInGroupResponseBodyData {
	s.UserState = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetWorkMode(v string) *SignInGroupResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *SignInGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
