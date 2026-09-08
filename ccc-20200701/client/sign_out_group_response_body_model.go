// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SignOutGroupResponseBody
	GetCode() *string
	SetData(v *SignOutGroupResponseBodyData) *SignOutGroupResponseBody
	GetData() *SignOutGroupResponseBodyData
	SetHttpStatusCode(v int32) *SignOutGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SignOutGroupResponseBody
	GetMessage() *string
	SetParams(v []*string) *SignOutGroupResponseBody
	GetParams() []*string
	SetRequestId(v string) *SignOutGroupResponseBody
	GetRequestId() *string
}

type SignOutGroupResponseBody struct {
	// Response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *SignOutGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of response parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SignOutGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SignOutGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SignOutGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *SignOutGroupResponseBody) GetData() *SignOutGroupResponseBodyData {
	return s.Data
}

func (s *SignOutGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SignOutGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SignOutGroupResponseBody) GetParams() []*string {
	return s.Params
}

func (s *SignOutGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SignOutGroupResponseBody) SetCode(v string) *SignOutGroupResponseBody {
	s.Code = &v
	return s
}

func (s *SignOutGroupResponseBody) SetData(v *SignOutGroupResponseBodyData) *SignOutGroupResponseBody {
	s.Data = v
	return s
}

func (s *SignOutGroupResponseBody) SetHttpStatusCode(v int32) *SignOutGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SignOutGroupResponseBody) SetMessage(v string) *SignOutGroupResponseBody {
	s.Message = &v
	return s
}

func (s *SignOutGroupResponseBody) SetParams(v []*string) *SignOutGroupResponseBody {
	s.Params = v
	return s
}

func (s *SignOutGroupResponseBody) SetRequestId(v string) *SignOutGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignOutGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SignOutGroupResponseBodyData struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break status after the agent is published but before becoming idle), RingingTimeout (break caused by ringing timeout), and RejectCall (break caused by the agent rejecting a call). There are no restrictions on Custom-defined status codes; customers can define them according to their business needs.
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Device ID, which is the identity of a browser-based Web Real-Time Communication (WebRTC) softphone or a physical phone device. Only one type of device can be registered at a time.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Agent extension number.
	//
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Time of the last received agent heartbeat, formatted as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1609136956378
	Heartbeat *int64 `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent\\"s personal phone number.
	//
	// example:
	//
	// 1324730****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the agent is in outbound-only mode.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// The UNIX timestamp (in milliseconds) indicating the most recent time the agent was reserved. Being reserved means an incoming call will be assigned to the agent shortly.
	//
	// example:
	//
	// 1609136956378
	Reserved *int64 `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// List of skill group IDs that the agent has signed into.
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Agent status.
	//
	// example:
	//
	// BREAK
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s SignOutGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SignOutGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *SignOutGroupResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *SignOutGroupResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SignOutGroupResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *SignOutGroupResponseBodyData) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *SignOutGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SignOutGroupResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SignOutGroupResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *SignOutGroupResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *SignOutGroupResponseBodyData) GetReserved() *int64 {
	return s.Reserved
}

func (s *SignOutGroupResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *SignOutGroupResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *SignOutGroupResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *SignOutGroupResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SignOutGroupResponseBodyData) SetBreakCode(v string) *SignOutGroupResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetDeviceId(v string) *SignOutGroupResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetExtension(v string) *SignOutGroupResponseBodyData {
	s.Extension = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetHeartbeat(v int64) *SignOutGroupResponseBodyData {
	s.Heartbeat = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetInstanceId(v string) *SignOutGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetJobId(v string) *SignOutGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetMobile(v string) *SignOutGroupResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetOutboundScenario(v bool) *SignOutGroupResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetReserved(v int64) *SignOutGroupResponseBodyData {
	s.Reserved = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetSignedSkillGroupIdList(v []*string) *SignOutGroupResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *SignOutGroupResponseBodyData) SetUserId(v string) *SignOutGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetUserState(v string) *SignOutGroupResponseBodyData {
	s.UserState = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetWorkMode(v string) *SignOutGroupResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *SignOutGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
