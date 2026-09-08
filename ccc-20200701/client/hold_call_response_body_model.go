// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHoldCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HoldCallResponseBody
	GetCode() *string
	SetData(v *HoldCallResponseBodyData) *HoldCallResponseBody
	GetData() *HoldCallResponseBodyData
	SetHttpStatusCode(v int32) *HoldCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *HoldCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *HoldCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *HoldCallResponseBody
	GetRequestId() *string
}

type HoldCallResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *HoldCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 174F7777-2F6C-4F10-B889-C698E26C1AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HoldCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HoldCallResponseBody) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *HoldCallResponseBody) GetData() *HoldCallResponseBodyData {
	return s.Data
}

func (s *HoldCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *HoldCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HoldCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *HoldCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HoldCallResponseBody) SetCode(v string) *HoldCallResponseBody {
	s.Code = &v
	return s
}

func (s *HoldCallResponseBody) SetData(v *HoldCallResponseBodyData) *HoldCallResponseBody {
	s.Data = v
	return s
}

func (s *HoldCallResponseBody) SetHttpStatusCode(v int32) *HoldCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HoldCallResponseBody) SetMessage(v string) *HoldCallResponseBody {
	s.Message = &v
	return s
}

func (s *HoldCallResponseBody) SetParams(v []*string) *HoldCallResponseBody {
	s.Params = v
	return s
}

func (s *HoldCallResponseBody) SetRequestId(v string) *HoldCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HoldCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HoldCallResponseBodyData struct {
	// Call context environment.
	CallContext *HoldCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *HoldCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s HoldCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s HoldCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyData) GetCallContext() *HoldCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *HoldCallResponseBodyData) GetUserContext() *HoldCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *HoldCallResponseBodyData) SetCallContext(v *HoldCallResponseBodyDataCallContext) *HoldCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *HoldCallResponseBodyData) SetUserContext(v *HoldCallResponseBodyDataUserContext) *HoldCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *HoldCallResponseBodyData) Validate() error {
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

type HoldCallResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*HoldCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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
	// job-6582589278232****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s HoldCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s HoldCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *HoldCallResponseBodyDataCallContext) GetChannelContexts() []*HoldCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *HoldCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HoldCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *HoldCallResponseBodyDataCallContext) SetCallType(v string) *HoldCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) SetChannelContexts(v []*HoldCallResponseBodyDataCallContextChannelContexts) *HoldCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) SetInstanceId(v string) *HoldCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) SetJobId(v string) *HoldCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) Validate() error {
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

type HoldCallResponseBodyDataCallContextChannelContexts struct {
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
	// 8001****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6582589278232****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The calling party of the voice channel.
	//
	// example:
	//
	// 1318888****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the call channel, indicating who first initiated the hang-up for this call channel.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The release reason of the call channel, indicating why the current call channel was released. The value comes from the response codes defined in the SIP protocol. Customers can refer to the SIP protocol to analyze the release reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the channel. In inbound scenarios, the associated skill group ID is determined by the skill group configured in the IVR transfer-to-agent module. In outbound scenarios, the associated skill group ID is the ID of the first skill group the agent signed into.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The UNIX timestamp (in milliseconds) of the most recent status change of the channel.
	//
	// example:
	//
	// 1609255716900
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The extension number of the agent associated with the channel.
	//
	// example:
	//
	// 8001****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the voice channel. This field is empty if the channel belongs to a Customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s HoldCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s HoldCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type HoldCallResponseBodyDataUserContext struct {
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
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat was received from the agent, formatted as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1609255716908
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
	// job-6582589278232****
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
	// The time when the agent was most recently reserved. Being reserved means an incoming call will soon be assigned to this agent. The format is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1609255715822
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

func (s HoldCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s HoldCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *HoldCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *HoldCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *HoldCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *HoldCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HoldCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *HoldCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *HoldCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *HoldCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *HoldCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *HoldCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *HoldCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *HoldCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *HoldCallResponseBodyDataUserContext) SetBreakCode(v string) *HoldCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetDeviceId(v string) *HoldCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetExtension(v string) *HoldCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetHeartbeat(v int64) *HoldCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetInstanceId(v string) *HoldCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetJobId(v string) *HoldCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetMobile(v string) *HoldCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *HoldCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetReserved(v int64) *HoldCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *HoldCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetUserId(v string) *HoldCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetUserState(v string) *HoldCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetWorkMode(v string) *HoldCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
