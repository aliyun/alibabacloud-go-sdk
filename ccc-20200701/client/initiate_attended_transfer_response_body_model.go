// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateAttendedTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitiateAttendedTransferResponseBody
	GetCode() *string
	SetData(v *InitiateAttendedTransferResponseBodyData) *InitiateAttendedTransferResponseBody
	GetData() *InitiateAttendedTransferResponseBodyData
	SetHttpStatusCode(v int32) *InitiateAttendedTransferResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InitiateAttendedTransferResponseBody
	GetMessage() *string
	SetParams(v []*string) *InitiateAttendedTransferResponseBody
	GetParams() []*string
	SetRequestId(v string) *InitiateAttendedTransferResponseBody
	GetRequestId() *string
}

type InitiateAttendedTransferResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *InitiateAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s InitiateAttendedTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitiateAttendedTransferResponseBody) GetData() *InitiateAttendedTransferResponseBodyData {
	return s.Data
}

func (s *InitiateAttendedTransferResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InitiateAttendedTransferResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitiateAttendedTransferResponseBody) GetParams() []*string {
	return s.Params
}

func (s *InitiateAttendedTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitiateAttendedTransferResponseBody) SetCode(v string) *InitiateAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetData(v *InitiateAttendedTransferResponseBodyData) *InitiateAttendedTransferResponseBody {
	s.Data = v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetHttpStatusCode(v int32) *InitiateAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetMessage(v string) *InitiateAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetParams(v []*string) *InitiateAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetRequestId(v string) *InitiateAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitiateAttendedTransferResponseBodyData struct {
	// Call context environment.
	CallContext *InitiateAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto increment ID. Customers do not need to concern themselves with this field.
	//
	// example:
	//
	// 103655
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *InitiateAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s InitiateAttendedTransferResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyData) GetCallContext() *InitiateAttendedTransferResponseBodyDataCallContext {
	return s.CallContext
}

func (s *InitiateAttendedTransferResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *InitiateAttendedTransferResponseBodyData) GetUserContext() *InitiateAttendedTransferResponseBodyDataUserContext {
	return s.UserContext
}

func (s *InitiateAttendedTransferResponseBodyData) SetCallContext(v *InitiateAttendedTransferResponseBodyDataCallContext) *InitiateAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) SetContextId(v int64) *InitiateAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) SetUserContext(v *InitiateAttendedTransferResponseBodyDataUserContext) *InitiateAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) Validate() error {
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

type InitiateAttendedTransferResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s InitiateAttendedTransferResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetChannelContexts() []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) Validate() error {
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

type InitiateAttendedTransferResponseBodyDataCallContextChannelContexts struct {
	// The call type of the call channel.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Channel flags.
	//
	// example:
	//
	// MONITORING
	ChannelFlags *string `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	// Channel ID.
	//
	// example:
	//
	// ch:user:139xxxx0501->80326034:1609138902226:job-6538214103685****
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
	// Records the order in which this channel was created during the call.
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
	// The originator of the channel.
	//
	// example:
	//
	// 0830019****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the call channel, indicating who hung up first.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The hang-up reason for the voice channel, indicating why the current voice channel was disconnected. The value comes from the response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the hang-up reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the voice channel. In inbound scenarios, the associated skill group ID is determined by the agent transfer module configured in the IVR. In outbound scenarios, the associated skill group ID is the first skill group the agent signed into.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The UNIX timestamp of the most recent status change of the voice channel, in milliseconds.
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
	// The agent ID associated with the channel. This field is empty for a Customer channel.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type InitiateAttendedTransferResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after an agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on Custom-defined status codes; customers can define them as needed for their business.
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Device ID, the identity ID of a browser-based Web Real-Time Communication (WebRTC) softphone or a physical phone device. Only one type of device can be registered at a time.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The agent\\"s extension number.
	//
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat was received from the agent, in UNIX timestamp format with millisecond precision.
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
	// The most recent time when the agent was reserved. Being reserved means an incoming call will be assigned to the agent shortly. The value is formatted as a UNIX timestamp in milliseconds.
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

func (s InitiateAttendedTransferResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
