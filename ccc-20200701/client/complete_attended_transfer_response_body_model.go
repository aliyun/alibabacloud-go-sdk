// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteAttendedTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CompleteAttendedTransferResponseBody
	GetCode() *string
	SetData(v *CompleteAttendedTransferResponseBodyData) *CompleteAttendedTransferResponseBody
	GetData() *CompleteAttendedTransferResponseBodyData
	SetHttpStatusCode(v int32) *CompleteAttendedTransferResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CompleteAttendedTransferResponseBody
	GetMessage() *string
	SetParams(v []*string) *CompleteAttendedTransferResponseBody
	GetParams() []*string
	SetRequestId(v string) *CompleteAttendedTransferResponseBody
	GetRequestId() *string
}

type CompleteAttendedTransferResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *CompleteAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CompleteAttendedTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompleteAttendedTransferResponseBody) GetData() *CompleteAttendedTransferResponseBodyData {
	return s.Data
}

func (s *CompleteAttendedTransferResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CompleteAttendedTransferResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompleteAttendedTransferResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CompleteAttendedTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteAttendedTransferResponseBody) SetCode(v string) *CompleteAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetData(v *CompleteAttendedTransferResponseBodyData) *CompleteAttendedTransferResponseBody {
	s.Data = v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetHttpStatusCode(v int32) *CompleteAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetMessage(v string) *CompleteAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetParams(v []*string) *CompleteAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetRequestId(v string) *CompleteAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompleteAttendedTransferResponseBodyData struct {
	// Call context environment.
	CallContext *CompleteAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto increment ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 103652
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *CompleteAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CompleteAttendedTransferResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyData) GetCallContext() *CompleteAttendedTransferResponseBodyDataCallContext {
	return s.CallContext
}

func (s *CompleteAttendedTransferResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *CompleteAttendedTransferResponseBodyData) GetUserContext() *CompleteAttendedTransferResponseBodyDataUserContext {
	return s.UserContext
}

func (s *CompleteAttendedTransferResponseBodyData) SetCallContext(v *CompleteAttendedTransferResponseBodyDataCallContext) *CompleteAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) SetContextId(v int64) *CompleteAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) SetUserContext(v *CompleteAttendedTransferResponseBodyDataUserContext) *CompleteAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) Validate() error {
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

type CompleteAttendedTransferResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s CompleteAttendedTransferResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetChannelContexts() []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) Validate() error {
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

type CompleteAttendedTransferResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
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
	// The voice channel ID.
	//
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-653821410368****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The status of the voice channel.
	//
	// example:
	//
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// The callee of the voice channel.
	//
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// An auto increment ID assigned by the system. Customers do not need to follow this.
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
	// The party that initiated the hang-up of the voice channel, indicating who first terminated the call.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The hang-up reason for the voice channel, indicating why the current voice channel was disconnected. The value corresponds to response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the hang-up reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the voice channel. In inbound scenarios, the associated skill group ID is determined by the skill group configured in the IVR transfer-to-agent module. In outbound scenarios, the associated skill group ID is the first skill group that the agent signed into.
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
	// The agent ID associated with the channel. This field is empty if the channel belongs to a customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type CompleteAttendedTransferResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent rejecting a call). There are no restrictions on Custom-defined status codes; customers can define them according to their business needs.
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
	// The time when the last heartbeat was received from the agent, formatted as a UNIX timestamp in milliseconds.
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
	// Call ID during the call state.
	//
	// example:
	//
	// job-65382141036853491
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
	// The most recent time when the agent was reserved. Being reserved means an incoming call will soon be assigned to this agent. The format is a UNIX timestamp in milliseconds.
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

func (s CompleteAttendedTransferResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
