// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlindTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BlindTransferResponseBody
	GetCode() *string
	SetData(v *BlindTransferResponseBodyData) *BlindTransferResponseBody
	GetData() *BlindTransferResponseBodyData
	SetHttpStatusCode(v int32) *BlindTransferResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BlindTransferResponseBody
	GetMessage() *string
	SetParams(v []*string) *BlindTransferResponseBody
	GetParams() []*string
	SetRequestId(v string) *BlindTransferResponseBody
	GetRequestId() *string
}

type BlindTransferResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *BlindTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s BlindTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferResponseBody) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBody) GetCode() *string {
	return s.Code
}

func (s *BlindTransferResponseBody) GetData() *BlindTransferResponseBodyData {
	return s.Data
}

func (s *BlindTransferResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BlindTransferResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BlindTransferResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BlindTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BlindTransferResponseBody) SetCode(v string) *BlindTransferResponseBody {
	s.Code = &v
	return s
}

func (s *BlindTransferResponseBody) SetData(v *BlindTransferResponseBodyData) *BlindTransferResponseBody {
	s.Data = v
	return s
}

func (s *BlindTransferResponseBody) SetHttpStatusCode(v int32) *BlindTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BlindTransferResponseBody) SetMessage(v string) *BlindTransferResponseBody {
	s.Message = &v
	return s
}

func (s *BlindTransferResponseBody) SetParams(v []*string) *BlindTransferResponseBody {
	s.Params = v
	return s
}

func (s *BlindTransferResponseBody) SetRequestId(v string) *BlindTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlindTransferResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BlindTransferResponseBodyData struct {
	// Call context environment.
	CallContext *BlindTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto increment ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 103654
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *BlindTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s BlindTransferResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyData) GetCallContext() *BlindTransferResponseBodyDataCallContext {
	return s.CallContext
}

func (s *BlindTransferResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *BlindTransferResponseBodyData) GetUserContext() *BlindTransferResponseBodyDataUserContext {
	return s.UserContext
}

func (s *BlindTransferResponseBodyData) SetCallContext(v *BlindTransferResponseBodyDataCallContext) *BlindTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *BlindTransferResponseBodyData) SetContextId(v int64) *BlindTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *BlindTransferResponseBodyData) SetUserContext(v *BlindTransferResponseBodyDataUserContext) *BlindTransferResponseBodyData {
	s.UserContext = v
	return s
}

func (s *BlindTransferResponseBodyData) Validate() error {
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

type BlindTransferResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*BlindTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s BlindTransferResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *BlindTransferResponseBodyDataCallContext) GetChannelContexts() []*BlindTransferResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *BlindTransferResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BlindTransferResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *BlindTransferResponseBodyDataCallContext) SetCallType(v string) *BlindTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) SetChannelContexts(v []*BlindTransferResponseBodyDataCallContextChannelContexts) *BlindTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) SetInstanceId(v string) *BlindTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) SetJobId(v string) *BlindTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) Validate() error {
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

type BlindTransferResponseBodyDataCallContextChannelContexts struct {
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
	// Channel ID.
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
	// Call ID.
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
	// The reason for releasing the voice channel, indicating why the current voice channel was disconnected. The value is derived from the response codes defined in the SIP protocol. Customers can refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The UNIX timestamp (in milliseconds) of the most recent status change of the channel.
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
	// The agent ID associated with the call channel. This field is empty for a Customer\\"s call channel.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BlindTransferResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type BlindTransferResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on Custom-defined status codes; customers can define them according to their business needs.
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

func (s BlindTransferResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *BlindTransferResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BlindTransferResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *BlindTransferResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *BlindTransferResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BlindTransferResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *BlindTransferResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *BlindTransferResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *BlindTransferResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *BlindTransferResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *BlindTransferResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *BlindTransferResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *BlindTransferResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *BlindTransferResponseBodyDataUserContext) SetBreakCode(v string) *BlindTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetDeviceId(v string) *BlindTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetExtension(v string) *BlindTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *BlindTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetInstanceId(v string) *BlindTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetJobId(v string) *BlindTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetMobile(v string) *BlindTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *BlindTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetReserved(v int64) *BlindTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *BlindTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetUserId(v string) *BlindTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetUserState(v string) *BlindTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetWorkMode(v string) *BlindTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
