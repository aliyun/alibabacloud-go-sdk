// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AnswerCallResponseBody
	GetCode() *string
	SetData(v *AnswerCallResponseBodyData) *AnswerCallResponseBody
	GetData() *AnswerCallResponseBodyData
	SetHttpStatusCode(v int32) *AnswerCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AnswerCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *AnswerCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *AnswerCallResponseBody
	GetRequestId() *string
}

type AnswerCallResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *AnswerCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s AnswerCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *AnswerCallResponseBody) GetData() *AnswerCallResponseBodyData {
	return s.Data
}

func (s *AnswerCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AnswerCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AnswerCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *AnswerCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnswerCallResponseBody) SetCode(v string) *AnswerCallResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerCallResponseBody) SetData(v *AnswerCallResponseBodyData) *AnswerCallResponseBody {
	s.Data = v
	return s
}

func (s *AnswerCallResponseBody) SetHttpStatusCode(v int32) *AnswerCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AnswerCallResponseBody) SetMessage(v string) *AnswerCallResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerCallResponseBody) SetParams(v []*string) *AnswerCallResponseBody {
	s.Params = v
	return s
}

func (s *AnswerCallResponseBody) SetRequestId(v string) *AnswerCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AnswerCallResponseBodyData struct {
	// Call context environment.
	CallContext *AnswerCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto increment ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 103655
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *AnswerCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s AnswerCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyData) GetCallContext() *AnswerCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *AnswerCallResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *AnswerCallResponseBodyData) GetUserContext() *AnswerCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *AnswerCallResponseBodyData) SetCallContext(v *AnswerCallResponseBodyDataCallContext) *AnswerCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *AnswerCallResponseBodyData) SetContextId(v int64) *AnswerCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *AnswerCallResponseBodyData) SetUserContext(v *AnswerCallResponseBodyDataUserContext) *AnswerCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *AnswerCallResponseBodyData) Validate() error {
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

type AnswerCallResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// INBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of channels.
	ChannelContexts []*AnswerCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s AnswerCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *AnswerCallResponseBodyDataCallContext) GetChannelContexts() []*AnswerCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *AnswerCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AnswerCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *AnswerCallResponseBodyDataCallContext) SetCallType(v string) *AnswerCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContext) SetChannelContexts(v []*AnswerCallResponseBodyDataCallContextChannelContexts) *AnswerCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *AnswerCallResponseBodyDataCallContext) SetInstanceId(v string) *AnswerCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContext) SetJobId(v string) *AnswerCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContext) Validate() error {
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

type AnswerCallResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// INBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// The ingest endpoint ID.
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
	// Channel-associated data.
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
	// The order in which the channel was created during the call procedure.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The calling party of the ingest endpoint.
	//
	// example:
	//
	// 0830019****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the voice channel, indicating who hung up first.
	//
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The hang-up reason for the ingest endpoint, indicating why the current ingest endpoint was disconnected. The value corresponds to the response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the hang-up reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the ingest endpoint. In inbound scenarios, the associated skill group ID is determined by the skill group configured in the IVR transfer-to-agent module. In outbound scenarios, the associated skill group ID is the first skill group that the agent signed into.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The UNIX timestamp of the most recent status change of the ingest endpoint, in milliseconds.
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
	// The agent ID associated with the ingest endpoint. This field is empty if the ingest endpoint belongs to a customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AnswerCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetChannelVariables() *string {
	return s.ChannelVariables
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetIndex() *int64 {
	return s.Index
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetChannelVariables(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ChannelVariables = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetIndex(v int64) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type AnswerCallResponseBodyDataUserContext struct {
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
	// The most recent time the agent was reserved. Being reserved means an incoming call will soon be assigned to the agent. The value is formatted as a UNIX timestamp in milliseconds.
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

func (s AnswerCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *AnswerCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AnswerCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *AnswerCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *AnswerCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AnswerCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *AnswerCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *AnswerCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *AnswerCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *AnswerCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *AnswerCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *AnswerCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *AnswerCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *AnswerCallResponseBodyDataUserContext) SetBreakCode(v string) *AnswerCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetDeviceId(v string) *AnswerCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetExtension(v string) *AnswerCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetHeartbeat(v int64) *AnswerCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetInstanceId(v string) *AnswerCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetJobId(v string) *AnswerCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetMobile(v string) *AnswerCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *AnswerCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetReserved(v int64) *AnswerCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *AnswerCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetUserId(v string) *AnswerCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetUserState(v string) *AnswerCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetWorkMode(v string) *AnswerCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
