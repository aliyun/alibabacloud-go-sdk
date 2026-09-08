// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAttendedTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelAttendedTransferResponseBody
	GetCode() *string
	SetData(v *CancelAttendedTransferResponseBodyData) *CancelAttendedTransferResponseBody
	GetData() *CancelAttendedTransferResponseBodyData
	SetHttpStatusCode(v int32) *CancelAttendedTransferResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelAttendedTransferResponseBody
	GetMessage() *string
	SetParams(v []*string) *CancelAttendedTransferResponseBody
	GetParams() []*string
	SetRequestId(v string) *CancelAttendedTransferResponseBody
	GetRequestId() *string
}

type CancelAttendedTransferResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *CancelAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CancelAttendedTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAttendedTransferResponseBody) GetData() *CancelAttendedTransferResponseBodyData {
	return s.Data
}

func (s *CancelAttendedTransferResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelAttendedTransferResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAttendedTransferResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CancelAttendedTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAttendedTransferResponseBody) SetCode(v string) *CancelAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetData(v *CancelAttendedTransferResponseBodyData) *CancelAttendedTransferResponseBody {
	s.Data = v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetHttpStatusCode(v int32) *CancelAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetMessage(v string) *CancelAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetParams(v []*string) *CancelAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetRequestId(v string) *CancelAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelAttendedTransferResponseBodyData struct {
	// Call context environment.
	CallContext *CancelAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto increment ID. The Customer does not need to concern themselves with this.
	//
	// example:
	//
	// 103656
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *CancelAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CancelAttendedTransferResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyData) GetCallContext() *CancelAttendedTransferResponseBodyDataCallContext {
	return s.CallContext
}

func (s *CancelAttendedTransferResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *CancelAttendedTransferResponseBodyData) GetUserContext() *CancelAttendedTransferResponseBodyDataUserContext {
	return s.UserContext
}

func (s *CancelAttendedTransferResponseBodyData) SetCallContext(v *CancelAttendedTransferResponseBodyDataCallContext) *CancelAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CancelAttendedTransferResponseBodyData) SetContextId(v int64) *CancelAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyData) SetUserContext(v *CancelAttendedTransferResponseBodyDataUserContext) *CancelAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

func (s *CancelAttendedTransferResponseBodyData) Validate() error {
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

type CancelAttendedTransferResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*CancelAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s CancelAttendedTransferResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) GetChannelContexts() []*CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *CancelAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*CancelAttendedTransferResponseBodyDataCallContextChannelContexts) *CancelAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *CancelAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *CancelAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) Validate() error {
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

type CancelAttendedTransferResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// CONSULTANT
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Channel flags.
	//
	// example:
	//
	// MONITORING
	ChannelFlags *string `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-653821410368****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The status of the call channel.
	//
	// example:
	//
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// The callee of the call channel.
	//
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// An auto-increment ID generated by the system. Customers do not need to follow it.
	//
	// example:
	//
	// 10
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The originator of the voice channel.
	//
	// example:
	//
	// 0830019****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the call channel, indicating who first hung up the call.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The reason for releasing the channel, indicating why the current channel was disconnected. The value is a response code defined in the SIP protocol. Customers can refer to the SIP protocol to analyze the disconnection reason.
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
	// The extension number of the agent associated with the voice channel.
	//
	// example:
	//
	// 8032****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the voice channel. This field is empty for customer voice channels.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type CancelAttendedTransferResponseBodyDataUserContext struct {
	// Break status code, which can be either system-defined or customer-defined. System-defined break codes include: Warm-up (temporary break state after the agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent rejecting a call). There are no restrictions on customer-defined status codes; customers can define them as needed for their business.
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Device ID, which is the identity ID of a browser-based Web Real-Time Communication (WebRTC) softphone or a physical phone device. Only one type of device can be registered at a time.
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
	// The time when the last heartbeat was received from the agent, in UNIX timestamp format, in milliseconds.
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
	// The most recent time the agent was reserved. Being reserved means an incoming call will soon be assigned to the agent, in UNIX timestamp format, in milliseconds.
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
	// TALKING
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s CancelAttendedTransferResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *CancelAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
