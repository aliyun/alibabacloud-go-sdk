// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PollUserStatusResponseBody
	GetCode() *string
	SetData(v *PollUserStatusResponseBodyData) *PollUserStatusResponseBody
	GetData() *PollUserStatusResponseBodyData
	SetHttpStatusCode(v int32) *PollUserStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PollUserStatusResponseBody
	GetMessage() *string
	SetParams(v []*string) *PollUserStatusResponseBody
	GetParams() []*string
	SetRequestId(v string) *PollUserStatusResponseBody
	GetRequestId() *string
}

type PollUserStatusResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *PollUserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s PollUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *PollUserStatusResponseBody) GetData() *PollUserStatusResponseBodyData {
	return s.Data
}

func (s *PollUserStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PollUserStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PollUserStatusResponseBody) GetParams() []*string {
	return s.Params
}

func (s *PollUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PollUserStatusResponseBody) SetCode(v string) *PollUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *PollUserStatusResponseBody) SetData(v *PollUserStatusResponseBodyData) *PollUserStatusResponseBody {
	s.Data = v
	return s
}

func (s *PollUserStatusResponseBody) SetHttpStatusCode(v int32) *PollUserStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PollUserStatusResponseBody) SetMessage(v string) *PollUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *PollUserStatusResponseBody) SetParams(v []*string) *PollUserStatusResponseBody {
	s.Params = v
	return s
}

func (s *PollUserStatusResponseBody) SetRequestId(v string) *PollUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *PollUserStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PollUserStatusResponseBodyData struct {
	// Call context environment.
	CallContext *PollUserStatusResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Session context environment.
	ChatContexts []*PollUserStatusResponseBodyDataChatContexts `json:"ChatContexts,omitempty" xml:"ChatContexts,omitempty" type:"Repeated"`
	// System auto increment ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 103655
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *PollUserStatusResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s PollUserStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyData) GetCallContext() *PollUserStatusResponseBodyDataCallContext {
	return s.CallContext
}

func (s *PollUserStatusResponseBodyData) GetChatContexts() []*PollUserStatusResponseBodyDataChatContexts {
	return s.ChatContexts
}

func (s *PollUserStatusResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *PollUserStatusResponseBodyData) GetUserContext() *PollUserStatusResponseBodyDataUserContext {
	return s.UserContext
}

func (s *PollUserStatusResponseBodyData) SetCallContext(v *PollUserStatusResponseBodyDataCallContext) *PollUserStatusResponseBodyData {
	s.CallContext = v
	return s
}

func (s *PollUserStatusResponseBodyData) SetChatContexts(v []*PollUserStatusResponseBodyDataChatContexts) *PollUserStatusResponseBodyData {
	s.ChatContexts = v
	return s
}

func (s *PollUserStatusResponseBodyData) SetContextId(v int64) *PollUserStatusResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *PollUserStatusResponseBodyData) SetUserContext(v *PollUserStatusResponseBodyDataUserContext) *PollUserStatusResponseBodyData {
	s.UserContext = v
	return s
}

func (s *PollUserStatusResponseBodyData) Validate() error {
	if s.CallContext != nil {
		if err := s.CallContext.Validate(); err != nil {
			return err
		}
	}
	if s.ChatContexts != nil {
		for _, item := range s.ChatContexts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserContext != nil {
		if err := s.UserContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PollUserStatusResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Ingest endpoint data. Custom data passed through SIP signaling.
	//
	// example:
	//
	// a=b;c=d
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// List of call channels.
	ChannelContexts []*PollUserStatusResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s PollUserStatusResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *PollUserStatusResponseBodyDataCallContext) GetCallVariables() *string {
	return s.CallVariables
}

func (s *PollUserStatusResponseBodyDataCallContext) GetChannelContexts() []*PollUserStatusResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *PollUserStatusResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PollUserStatusResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *PollUserStatusResponseBodyDataCallContext) SetCallType(v string) *PollUserStatusResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetCallVariables(v string) *PollUserStatusResponseBodyDataCallContext {
	s.CallVariables = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetChannelContexts(v []*PollUserStatusResponseBodyDataCallContextChannelContexts) *PollUserStatusResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetInstanceId(v string) *PollUserStatusResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetJobId(v string) *PollUserStatusResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) Validate() error {
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

type PollUserStatusResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Call channel flags.
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
	// CREATED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// Channel-associated data.
	//
	// example:
	//
	// 123
	ChannelVariables *string `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	// Callee of the call channel.
	//
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// Auto-incremented system ID. Customers do not need to concern themselves with this value.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
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
	// The hang-up reason for the call channel, indicating why the current call channel was disconnected. The value comes from the response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the hang-up reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the channel. In inbound scenarios, the associated skill group ID is determined by the skill group configured in the IVR transfer-to-agent module. In outbound scenarios, the associated skill group ID is the first skill group the agent signed into.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The UNIX timestamp indicating the most recent status change of the channel, in milliseconds.
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
	// The agent ID associated with the voice channel. This field is empty if the channel belongs to a customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PollUserStatusResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetChannelVariables() *string {
	return s.ChannelVariables
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetCallType(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelVariables(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelVariables = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetDestination(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetJobId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetUserId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type PollUserStatusResponseBodyDataChatContexts struct {
	// Ingest endpoint data.
	//
	// example:
	//
	// 无
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// Session type.
	//
	// example:
	//
	// INBOUND
	ChatType *string `json:"ChatType,omitempty" xml:"ChatType,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call job ID.
	//
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Information about the agents and customers participating in this chat session.
	Members []*PollUserStatusResponseBodyDataChatContextsMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s PollUserStatusResponseBodyDataChatContexts) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyDataChatContexts) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataChatContexts) GetCallVariables() *string {
	return s.CallVariables
}

func (s *PollUserStatusResponseBodyDataChatContexts) GetChatType() *string {
	return s.ChatType
}

func (s *PollUserStatusResponseBodyDataChatContexts) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PollUserStatusResponseBodyDataChatContexts) GetJobId() *string {
	return s.JobId
}

func (s *PollUserStatusResponseBodyDataChatContexts) GetMembers() []*PollUserStatusResponseBodyDataChatContextsMembers {
	return s.Members
}

func (s *PollUserStatusResponseBodyDataChatContexts) SetCallVariables(v string) *PollUserStatusResponseBodyDataChatContexts {
	s.CallVariables = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContexts) SetChatType(v string) *PollUserStatusResponseBodyDataChatContexts {
	s.ChatType = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContexts) SetInstanceId(v string) *PollUserStatusResponseBodyDataChatContexts {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContexts) SetJobId(v string) *PollUserStatusResponseBodyDataChatContexts {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContexts) SetMembers(v []*PollUserStatusResponseBodyDataChatContextsMembers) *PollUserStatusResponseBodyDataChatContexts {
	s.Members = v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContexts) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PollUserStatusResponseBodyDataChatContextsMembers struct {
	// System auto increment ID. Customers do not need to concern themselves with this value.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The party that initiated the hang-up.
	//
	// example:
	//
	// User
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The hang-up reason for the call channel, indicating why the current call channel was terminated. The value is derived from the response codes defined in the SIP protocol. Customers should refer to the SIP protocol for analysis of the hang-up reason.
	//
	// example:
	//
	// 无
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// Skill group ID.
	//
	// example:
	//
	// a527a80a-75cf-448f-aba0-fca37c61946d
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The status of the agent in this session.
	//
	// Enumeration values:
	//
	// - RINGING
	//
	// - TALKING
	//
	// - RELEASED (End)
	//
	// - MISSED (Missed)
	//
	// example:
	//
	// RINGING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// User type.
	//
	// Enumeration values:
	//
	// - CUSTOMER: Customer
	//
	// - AGENT: Agent
	//
	// example:
	//
	// Manager
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s PollUserStatusResponseBodyDataChatContextsMembers) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyDataChatContextsMembers) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetIndex() *int32 {
	return s.Index
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetStatus() *string {
	return s.Status
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetUserId() *string {
	return s.UserId
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) GetUserType() *string {
	return s.UserType
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetIndex(v int32) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.Index = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetReleaseInitiator(v string) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.ReleaseInitiator = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetReleaseReason(v string) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.ReleaseReason = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetSkillGroupId(v string) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.SkillGroupId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetStatus(v string) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.Status = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetUserId(v string) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.UserId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) SetUserType(v string) *PollUserStatusResponseBodyDataChatContextsMembers {
	s.UserType = &v
	return s
}

func (s *PollUserStatusResponseBodyDataChatContextsMembers) Validate() error {
	return dara.Validate(s)
}

type PollUserStatusResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Customer-defined. System-defined break codes include: Warm-up (temporary break state after an agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on Customer-defined status codes, and customers can define them according to their business needs.
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
	// Agent extension number.
	//
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat was received from the agent, in Unix timestamp format, in milliseconds.
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
	// Chat sessions currently being handled or already handled by the agent
	ParallelJobList []*PollUserStatusResponseBodyDataUserContextParallelJobList `json:"ParallelJobList,omitempty" xml:"ParallelJobList,omitempty" type:"Repeated"`
	// The UNIX timestamp (in milliseconds) indicating when the agent was most recently reserved. Being reserved means an incoming call will soon be assigned to the agent.
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
	// BREAK
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s PollUserStatusResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *PollUserStatusResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *PollUserStatusResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *PollUserStatusResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *PollUserStatusResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PollUserStatusResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *PollUserStatusResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *PollUserStatusResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *PollUserStatusResponseBodyDataUserContext) GetParallelJobList() []*PollUserStatusResponseBodyDataUserContextParallelJobList {
	return s.ParallelJobList
}

func (s *PollUserStatusResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *PollUserStatusResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *PollUserStatusResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *PollUserStatusResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *PollUserStatusResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *PollUserStatusResponseBodyDataUserContext) SetBreakCode(v string) *PollUserStatusResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetDeviceId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetExtension(v string) *PollUserStatusResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetHeartbeat(v int64) *PollUserStatusResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetInstanceId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetJobId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetMobile(v string) *PollUserStatusResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetOutboundScenario(v bool) *PollUserStatusResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetParallelJobList(v []*PollUserStatusResponseBodyDataUserContextParallelJobList) *PollUserStatusResponseBodyDataUserContext {
	s.ParallelJobList = v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetReserved(v int64) *PollUserStatusResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *PollUserStatusResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetUserId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetUserState(v string) *PollUserStatusResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetWorkMode(v string) *PollUserStatusResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) Validate() error {
	if s.ParallelJobList != nil {
		for _, item := range s.ParallelJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PollUserStatusResponseBodyDataUserContextParallelJobList struct {
	// Session ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent\\"s status in this session.
	//
	// ### Enumeration values:
	//
	// - RINGING
	//
	// - TALKING
	//
	// - RELEASED (End)
	//
	// - MISSED (Missed)
	//
	// example:
	//
	// RINGING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the status changed.
	//
	// example:
	//
	// 1609136956378
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s PollUserStatusResponseBodyDataUserContextParallelJobList) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponseBodyDataUserContextParallelJobList) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) GetJobId() *string {
	return s.JobId
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) GetStatus() *string {
	return s.Status
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) SetJobId(v string) *PollUserStatusResponseBodyDataUserContextParallelJobList {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) SetStatus(v string) *PollUserStatusResponseBodyDataUserContextParallelJobList {
	s.Status = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) SetTimestamp(v int64) *PollUserStatusResponseBodyDataUserContextParallelJobList {
	s.Timestamp = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContextParallelJobList) Validate() error {
	return dara.Validate(s)
}
