// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBargeInCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BargeInCallResponseBody
	GetCode() *string
	SetData(v *BargeInCallResponseBodyData) *BargeInCallResponseBody
	GetData() *BargeInCallResponseBodyData
	SetHttpStatusCode(v int32) *BargeInCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BargeInCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *BargeInCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *BargeInCallResponseBody
	GetRequestId() *string
}

type BargeInCallResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *BargeInCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s BargeInCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallResponseBody) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BargeInCallResponseBody) GetData() *BargeInCallResponseBodyData {
	return s.Data
}

func (s *BargeInCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BargeInCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BargeInCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BargeInCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BargeInCallResponseBody) SetCode(v string) *BargeInCallResponseBody {
	s.Code = &v
	return s
}

func (s *BargeInCallResponseBody) SetData(v *BargeInCallResponseBodyData) *BargeInCallResponseBody {
	s.Data = v
	return s
}

func (s *BargeInCallResponseBody) SetHttpStatusCode(v int32) *BargeInCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BargeInCallResponseBody) SetMessage(v string) *BargeInCallResponseBody {
	s.Message = &v
	return s
}

func (s *BargeInCallResponseBody) SetParams(v []*string) *BargeInCallResponseBody {
	s.Params = v
	return s
}

func (s *BargeInCallResponseBody) SetRequestId(v string) *BargeInCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BargeInCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BargeInCallResponseBodyData struct {
	// Call context environment.
	CallContext *BargeInCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *BargeInCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s BargeInCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyData) GetCallContext() *BargeInCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *BargeInCallResponseBodyData) GetUserContext() *BargeInCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *BargeInCallResponseBodyData) SetCallContext(v *BargeInCallResponseBodyDataCallContext) *BargeInCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *BargeInCallResponseBodyData) SetUserContext(v *BargeInCallResponseBodyDataUserContext) *BargeInCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *BargeInCallResponseBodyData) Validate() error {
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

type BargeInCallResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// BARGE
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*BargeInCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s BargeInCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *BargeInCallResponseBodyDataCallContext) GetChannelContexts() []*BargeInCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *BargeInCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BargeInCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *BargeInCallResponseBodyDataCallContext) SetCallType(v string) *BargeInCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) SetChannelContexts(v []*BargeInCallResponseBodyDataCallContextChannelContexts) *BargeInCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) SetInstanceId(v string) *BargeInCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) SetJobId(v string) *BargeInCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) Validate() error {
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

type BargeInCallResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// BARGE
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Channel flag.
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
	// The status of the call channel.
	//
	// example:
	//
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// Callee of the call channel.
	//
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// An auto-increment ID assigned by the System. Customers do not need to concern themselves with this value.
	//
	// example:
	//
	// 10
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Call ID.
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
	// The party that initiated the release of the call channel, indicating who first hung up the call.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The hang-up reason for the voice channel, indicating why the current voice channel was disconnected. The value corresponds to a response code defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the hang-up reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the voice channel. In inbound scenarios, the associated skill group ID is determined by the skill group configured in the IVR-to-agent module. In outbound scenarios, the associated skill group ID is the first skill group that the agent signed into.
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

func (s BargeInCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type BargeInCallResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after an agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on Custom-defined status codes; customers can define them according to their business needs.
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
	// The time when the agent was most recently reserved. Being reserved means an incoming call will be assigned to the agent shortly. The value is a UNIX timestamp in milliseconds.
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

func (s BargeInCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *BargeInCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BargeInCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *BargeInCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *BargeInCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BargeInCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *BargeInCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *BargeInCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *BargeInCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *BargeInCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *BargeInCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *BargeInCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *BargeInCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *BargeInCallResponseBodyDataUserContext) SetBreakCode(v string) *BargeInCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetDeviceId(v string) *BargeInCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetExtension(v string) *BargeInCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetHeartbeat(v int64) *BargeInCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetInstanceId(v string) *BargeInCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetJobId(v string) *BargeInCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetMobile(v string) *BargeInCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *BargeInCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetReserved(v int64) *BargeInCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *BargeInCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetUserId(v string) *BargeInCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetUserState(v string) *BargeInCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetWorkMode(v string) *BargeInCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
