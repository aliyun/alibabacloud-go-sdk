// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchAuthenticationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LaunchAuthenticationResponseBody
	GetCode() *string
	SetData(v *LaunchAuthenticationResponseBodyData) *LaunchAuthenticationResponseBody
	GetData() *LaunchAuthenticationResponseBodyData
	SetHttpStatusCode(v int32) *LaunchAuthenticationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *LaunchAuthenticationResponseBody
	GetMessage() *string
	SetParams(v []*string) *LaunchAuthenticationResponseBody
	GetParams() []*string
	SetRequestId(v string) *LaunchAuthenticationResponseBody
	GetRequestId() *string
}

type LaunchAuthenticationResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *LaunchAuthenticationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s LaunchAuthenticationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBody) GetCode() *string {
	return s.Code
}

func (s *LaunchAuthenticationResponseBody) GetData() *LaunchAuthenticationResponseBodyData {
	return s.Data
}

func (s *LaunchAuthenticationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *LaunchAuthenticationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LaunchAuthenticationResponseBody) GetParams() []*string {
	return s.Params
}

func (s *LaunchAuthenticationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LaunchAuthenticationResponseBody) SetCode(v string) *LaunchAuthenticationResponseBody {
	s.Code = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetData(v *LaunchAuthenticationResponseBodyData) *LaunchAuthenticationResponseBody {
	s.Data = v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetHttpStatusCode(v int32) *LaunchAuthenticationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetMessage(v string) *LaunchAuthenticationResponseBody {
	s.Message = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetParams(v []*string) *LaunchAuthenticationResponseBody {
	s.Params = v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetRequestId(v string) *LaunchAuthenticationResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LaunchAuthenticationResponseBodyData struct {
	// Call context environment.
	CallContext *LaunchAuthenticationResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *LaunchAuthenticationResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s LaunchAuthenticationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyData) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyData) GetCallContext() *LaunchAuthenticationResponseBodyDataCallContext {
	return s.CallContext
}

func (s *LaunchAuthenticationResponseBodyData) GetUserContext() *LaunchAuthenticationResponseBodyDataUserContext {
	return s.UserContext
}

func (s *LaunchAuthenticationResponseBodyData) SetCallContext(v *LaunchAuthenticationResponseBodyDataCallContext) *LaunchAuthenticationResponseBodyData {
	s.CallContext = v
	return s
}

func (s *LaunchAuthenticationResponseBodyData) SetUserContext(v *LaunchAuthenticationResponseBodyDataUserContext) *LaunchAuthenticationResponseBodyData {
	s.UserContext = v
	return s
}

func (s *LaunchAuthenticationResponseBodyData) Validate() error {
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

type LaunchAuthenticationResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*LaunchAuthenticationResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s LaunchAuthenticationResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) GetChannelContexts() []*LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetCallType(v string) *LaunchAuthenticationResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetChannelContexts(v []*LaunchAuthenticationResponseBodyDataCallContextChannelContexts) *LaunchAuthenticationResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetInstanceId(v string) *LaunchAuthenticationResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetJobId(v string) *LaunchAuthenticationResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) Validate() error {
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

type LaunchAuthenticationResponseBodyDataCallContextChannelContexts struct {
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
	// The status of the call channel.
	//
	// example:
	//
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// Called party of the call channel.
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
	// The party that initiated the release of the call channel, indicating who first hung up the call.
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
	// The agent ID associated with the channel. This field is empty if the channel belongs to a Customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LaunchAuthenticationResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetCallType(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetDestination(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetJobId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetUserId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type LaunchAuthenticationResponseBodyDataUserContext struct {
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
	// instance ID.
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
	// The most recent time when the agent was reserved. Being reserved means an incoming call will be assigned to the agent shortly. The value is a UNIX timestamp in milliseconds.
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

func (s LaunchAuthenticationResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetBreakCode(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetDeviceId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetExtension(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetHeartbeat(v int64) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetInstanceId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetJobId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetMobile(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetOutboundScenario(v bool) *LaunchAuthenticationResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetReserved(v int64) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetUserId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetUserState(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetWorkMode(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
