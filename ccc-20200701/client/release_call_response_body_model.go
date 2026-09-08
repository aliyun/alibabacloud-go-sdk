// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseCallResponseBody
	GetCode() *string
	SetData(v *ReleaseCallResponseBodyData) *ReleaseCallResponseBody
	GetData() *ReleaseCallResponseBodyData
	SetHttpStatusCode(v int32) *ReleaseCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReleaseCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *ReleaseCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *ReleaseCallResponseBody
	GetRequestId() *string
}

type ReleaseCallResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ReleaseCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ReleaseCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseCallResponseBody) GetData() *ReleaseCallResponseBodyData {
	return s.Data
}

func (s *ReleaseCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleaseCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ReleaseCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseCallResponseBody) SetCode(v string) *ReleaseCallResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseCallResponseBody) SetData(v *ReleaseCallResponseBodyData) *ReleaseCallResponseBody {
	s.Data = v
	return s
}

func (s *ReleaseCallResponseBody) SetHttpStatusCode(v int32) *ReleaseCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseCallResponseBody) SetMessage(v string) *ReleaseCallResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseCallResponseBody) SetParams(v []*string) *ReleaseCallResponseBody {
	s.Params = v
	return s
}

func (s *ReleaseCallResponseBody) SetRequestId(v string) *ReleaseCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseCallResponseBodyData struct {
	// Call context environment.
	CallContext *ReleaseCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto-increment ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 123456
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *ReleaseCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s ReleaseCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyData) GetCallContext() *ReleaseCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *ReleaseCallResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *ReleaseCallResponseBodyData) GetUserContext() *ReleaseCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *ReleaseCallResponseBodyData) SetCallContext(v *ReleaseCallResponseBodyDataCallContext) *ReleaseCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *ReleaseCallResponseBodyData) SetContextId(v int64) *ReleaseCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *ReleaseCallResponseBodyData) SetUserContext(v *ReleaseCallResponseBodyDataUserContext) *ReleaseCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *ReleaseCallResponseBodyData) Validate() error {
	if s.CallContext != nil {
		if err := s.CallContext.Validate(); err != nil {
			return err
		}
	}
	if s.UserContext != nil {
		if err := s.UserContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseCallResponseBodyDataCallContext struct {
	// List of call channels.
	ChannelContexts []*ReleaseCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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
}

func (s ReleaseCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyDataCallContext) GetChannelContexts() []*ReleaseCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *ReleaseCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *ReleaseCallResponseBodyDataCallContext) SetChannelContexts(v []*ReleaseCallResponseBodyDataCallContextChannelContexts) *ReleaseCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContext) SetInstanceId(v string) *ReleaseCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContext) SetJobId(v string) *ReleaseCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContext) Validate() error {
	if s.ChannelContexts != nil {
		for _, item := range s.ChannelContexts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReleaseCallResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-653821410368****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The status of the voice channel.
	//
	// example:
	//
	// CREATED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// Channel-associated data
	//
	// example:
	//
	// a=b;c=d;
	ChannelVariables *string `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	// The callee of the voice channel.
	//
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The calling party of the call channel.
	//
	// example:
	//
	// 0830019****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the hang-up of the call channel, indicating who first terminated the call.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The reason for releasing the voice channel, indicating why the current voice channel was disconnected. The value is derived from the response codes defined in the SIP protocol. Customers can refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The UNIX timestamp of the most recent status change of the channel, in milliseconds.
	//
	// example:
	//
	// 1609138903315
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The extension number of the agent associated with the channel.
	//
	// example:
	//
	// 8032****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the call channel. This field is empty if the call channel belongs to a Customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ReleaseCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetChannelVariables() *string {
	return s.ChannelVariables
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetChannelVariables(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ChannelVariables = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type ReleaseCallResponseBodyDataUserContext struct {
	// Break status code, which can be either system-defined or customer-defined. System-defined break codes include: Warm-up (temporary break state after agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on customer-defined status codes; customers can define them according to their business needs.
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
	// Indicates whether the agent is in outbound-only mode.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
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

func (s ReleaseCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ReleaseCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ReleaseCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *ReleaseCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *ReleaseCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ReleaseCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ReleaseCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *ReleaseCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *ReleaseCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ReleaseCallResponseBodyDataUserContext) SetBreakCode(v string) *ReleaseCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetDeviceId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetExtension(v string) *ReleaseCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetInstanceId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetJobId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *ReleaseCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *ReleaseCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetUserId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetUserState(v string) *ReleaseCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetWorkMode(v string) *ReleaseCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
