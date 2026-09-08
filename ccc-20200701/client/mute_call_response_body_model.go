// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MuteCallResponseBody
	GetCode() *string
	SetData(v *MuteCallResponseBodyData) *MuteCallResponseBody
	GetData() *MuteCallResponseBodyData
	SetHttpStatusCode(v int32) *MuteCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *MuteCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *MuteCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *MuteCallResponseBody
	GetRequestId() *string
}

type MuteCallResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *MuteCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// A275B008-A25B-494D-AB53-93CE253815B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MuteCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MuteCallResponseBody) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *MuteCallResponseBody) GetData() *MuteCallResponseBodyData {
	return s.Data
}

func (s *MuteCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *MuteCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MuteCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *MuteCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MuteCallResponseBody) SetCode(v string) *MuteCallResponseBody {
	s.Code = &v
	return s
}

func (s *MuteCallResponseBody) SetData(v *MuteCallResponseBodyData) *MuteCallResponseBody {
	s.Data = v
	return s
}

func (s *MuteCallResponseBody) SetHttpStatusCode(v int32) *MuteCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MuteCallResponseBody) SetMessage(v string) *MuteCallResponseBody {
	s.Message = &v
	return s
}

func (s *MuteCallResponseBody) SetParams(v []*string) *MuteCallResponseBody {
	s.Params = v
	return s
}

func (s *MuteCallResponseBody) SetRequestId(v string) *MuteCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MuteCallResponseBodyData struct {
	// Call context environment.
	CallContext *MuteCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *MuteCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s MuteCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MuteCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyData) GetCallContext() *MuteCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *MuteCallResponseBodyData) GetUserContext() *MuteCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *MuteCallResponseBodyData) SetCallContext(v *MuteCallResponseBodyDataCallContext) *MuteCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *MuteCallResponseBodyData) SetUserContext(v *MuteCallResponseBodyDataUserContext) *MuteCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *MuteCallResponseBodyData) Validate() error {
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

type MuteCallResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*MuteCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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
	// job-6581536084722****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s MuteCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s MuteCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *MuteCallResponseBodyDataCallContext) GetChannelContexts() []*MuteCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *MuteCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MuteCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *MuteCallResponseBodyDataCallContext) SetCallType(v string) *MuteCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) SetChannelContexts(v []*MuteCallResponseBodyDataCallContextChannelContexts) *MuteCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) SetInstanceId(v string) *MuteCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) SetJobId(v string) *MuteCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) Validate() error {
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

type MuteCallResponseBodyDataCallContextChannelContexts struct {
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
	// [MUTED]
	ChannelFlags *string `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// ch:user:1318888****->8001****:1609253204816:job-6581536084722****
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
	// 8001****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// An auto-increment ID generated by the system. Customers do not need to concern themselves with it.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6581536084722****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The originator of the channel.
	//
	// example:
	//
	// 1318888****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the call channel, indicating who hung up first.
	//
	// example:
	//
	// 无
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The release reason for the channel, indicating why the current channel was disconnected. The value comes from the response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 无
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the channel. In inbound scenarios, the associated skill group ID is determined by the skill group configured in the IVR transfer-to-agent module. In outbound scenarios, the associated skill group ID is the first skill group that the agent signed into.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The UNIX timestamp of the most recent status change of the channel, in milliseconds.
	//
	// example:
	//
	// 1609253212511
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The extension number of the agent associated with the channel.
	//
	// example:
	//
	// 8001****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the channel. This field is empty for a customer\\"s channel.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s MuteCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s MuteCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type MuteCallResponseBodyDataUserContext struct {
	// Break status code, which can be either system-defined or customer-defined. System-defined break codes include: Warm-up (temporary break state after the agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on customer-defined status codes; customers can define them according to their business needs.
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
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat was received from the agent, in UNIX timestamp format, in milliseconds.
	//
	// example:
	//
	// 1609253205896
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
	// job-6581536084722****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent\\"s personal phone number.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the agent is in outbound-only mode.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// The time when the agent was most recently reserved. Being reserved means an incoming call will soon be assigned to the agent, in UNIX timestamp format, in milliseconds.
	//
	// example:
	//
	// 1609253204811
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

func (s MuteCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s MuteCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *MuteCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MuteCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *MuteCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *MuteCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MuteCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *MuteCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *MuteCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *MuteCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *MuteCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *MuteCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *MuteCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *MuteCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *MuteCallResponseBodyDataUserContext) SetBreakCode(v string) *MuteCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetDeviceId(v string) *MuteCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetExtension(v string) *MuteCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetHeartbeat(v int64) *MuteCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetInstanceId(v string) *MuteCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetJobId(v string) *MuteCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetMobile(v string) *MuteCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *MuteCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetReserved(v int64) *MuteCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *MuteCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetUserId(v string) *MuteCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetUserState(v string) *MuteCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetWorkMode(v string) *MuteCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
