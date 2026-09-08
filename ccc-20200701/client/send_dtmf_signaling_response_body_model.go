// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDtmfSignalingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendDtmfSignalingResponseBody
	GetCode() *string
	SetData(v *SendDtmfSignalingResponseBodyData) *SendDtmfSignalingResponseBody
	GetData() *SendDtmfSignalingResponseBodyData
	SetHttpStatusCode(v int32) *SendDtmfSignalingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendDtmfSignalingResponseBody
	GetMessage() *string
	SetParams(v []*string) *SendDtmfSignalingResponseBody
	GetParams() []*string
	SetRequestId(v string) *SendDtmfSignalingResponseBody
	GetRequestId() *string
}

type SendDtmfSignalingResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *SendDtmfSignalingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 842399EC-7D32-4472-AD08-9504C3F141FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendDtmfSignalingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingResponseBody) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendDtmfSignalingResponseBody) GetData() *SendDtmfSignalingResponseBodyData {
	return s.Data
}

func (s *SendDtmfSignalingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendDtmfSignalingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendDtmfSignalingResponseBody) GetParams() []*string {
	return s.Params
}

func (s *SendDtmfSignalingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendDtmfSignalingResponseBody) SetCode(v string) *SendDtmfSignalingResponseBody {
	s.Code = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetData(v *SendDtmfSignalingResponseBodyData) *SendDtmfSignalingResponseBody {
	s.Data = v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetHttpStatusCode(v int32) *SendDtmfSignalingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetMessage(v string) *SendDtmfSignalingResponseBody {
	s.Message = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetParams(v []*string) *SendDtmfSignalingResponseBody {
	s.Params = v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetRequestId(v string) *SendDtmfSignalingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendDtmfSignalingResponseBodyData struct {
	// Call context environment.
	CallContext *SendDtmfSignalingResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *SendDtmfSignalingResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s SendDtmfSignalingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyData) GetCallContext() *SendDtmfSignalingResponseBodyDataCallContext {
	return s.CallContext
}

func (s *SendDtmfSignalingResponseBodyData) GetUserContext() *SendDtmfSignalingResponseBodyDataUserContext {
	return s.UserContext
}

func (s *SendDtmfSignalingResponseBodyData) SetCallContext(v *SendDtmfSignalingResponseBodyDataCallContext) *SendDtmfSignalingResponseBodyData {
	s.CallContext = v
	return s
}

func (s *SendDtmfSignalingResponseBodyData) SetUserContext(v *SendDtmfSignalingResponseBodyDataUserContext) *SendDtmfSignalingResponseBodyData {
	s.UserContext = v
	return s
}

func (s *SendDtmfSignalingResponseBodyData) Validate() error {
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

type SendDtmfSignalingResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*SendDtmfSignalingResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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
	// job-6573574060089****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SendDtmfSignalingResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) GetChannelContexts() []*SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetCallType(v string) *SendDtmfSignalingResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetChannelContexts(v []*SendDtmfSignalingResponseBodyDataCallContextChannelContexts) *SendDtmfSignalingResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetInstanceId(v string) *SendDtmfSignalingResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetJobId(v string) *SendDtmfSignalingResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) Validate() error {
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

type SendDtmfSignalingResponseBodyDataCallContextChannelContexts struct {
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
	// ch:user:131888****->8001****:1609234221870:job-6573574060089****
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
	// 8001****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// An auto-increment ID assigned by the system. Customers do not need to concern themselves with this value.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6573574060089****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The originator of the channel.
	//
	// example:
	//
	// 0101234****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the call channel, indicating who first hung up the call.
	//
	// example:
	//
	// 无
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The release reason for the voice channel, indicating why the current voice channel was disconnected. The value corresponds to response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 无
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
	// 1609234222367
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The extension number of the agent associated with the channel.
	//
	// example:
	//
	// 8001****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the channel. This field is empty if the channel belongs to a Customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SendDtmfSignalingResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetCallType(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetDestination(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetJobId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetUserId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type SendDtmfSignalingResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after an agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent rejecting a call). There are no restrictions on Custom-defined status codes; customers can define them according to their business needs.
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
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat was received from the agent, in UNIX timestamp format, in milliseconds.
	//
	// example:
	//
	// 1609234222375
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
	// job-6573574060089****
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
	// The most recent time when the agent was reserved. Being reserved means an incoming call will be assigned to the agent shortly. The value is in UNIX timestamp format, in milliseconds.
	//
	// example:
	//
	// 1609234221864
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

func (s SendDtmfSignalingResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetBreakCode(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetDeviceId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetExtension(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetHeartbeat(v int64) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetInstanceId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetJobId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetMobile(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetOutboundScenario(v bool) *SendDtmfSignalingResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetReserved(v int64) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetUserId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetUserState(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetWorkMode(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
