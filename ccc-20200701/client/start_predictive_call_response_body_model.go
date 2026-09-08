// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPredictiveCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartPredictiveCallResponseBody
	GetCode() *string
	SetData(v *StartPredictiveCallResponseBodyData) *StartPredictiveCallResponseBody
	GetData() *StartPredictiveCallResponseBodyData
	SetHttpStatusCode(v int32) *StartPredictiveCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartPredictiveCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *StartPredictiveCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *StartPredictiveCallResponseBody
	GetRequestId() *string
}

type StartPredictiveCallResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *StartPredictiveCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 26A34338-5CD9-4C95-A7A6-5BDCE76C6B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPredictiveCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartPredictiveCallResponseBody) GetData() *StartPredictiveCallResponseBodyData {
	return s.Data
}

func (s *StartPredictiveCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartPredictiveCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartPredictiveCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *StartPredictiveCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPredictiveCallResponseBody) SetCode(v string) *StartPredictiveCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetData(v *StartPredictiveCallResponseBodyData) *StartPredictiveCallResponseBody {
	s.Data = v
	return s
}

func (s *StartPredictiveCallResponseBody) SetHttpStatusCode(v int32) *StartPredictiveCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetMessage(v string) *StartPredictiveCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetParams(v []*string) *StartPredictiveCallResponseBody {
	s.Params = v
	return s
}

func (s *StartPredictiveCallResponseBody) SetRequestId(v string) *StartPredictiveCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPredictiveCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartPredictiveCallResponseBodyData struct {
	// Call context environment.
	CallContext *StartPredictiveCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *StartPredictiveCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s StartPredictiveCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyData) GetCallContext() *StartPredictiveCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *StartPredictiveCallResponseBodyData) GetUserContext() *StartPredictiveCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *StartPredictiveCallResponseBodyData) SetCallContext(v *StartPredictiveCallResponseBodyDataCallContext) *StartPredictiveCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *StartPredictiveCallResponseBodyData) SetUserContext(v *StartPredictiveCallResponseBodyDataUserContext) *StartPredictiveCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *StartPredictiveCallResponseBodyData) Validate() error {
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

type StartPredictiveCallResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*StartPredictiveCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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
	// job-6570007401392****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StartPredictiveCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetChannelContexts() []*StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetCallType(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetChannelContexts(v []*StartPredictiveCallResponseBodyDataCallContextChannelContexts) *StartPredictiveCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetInstanceId(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetJobId(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) Validate() error {
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

type StartPredictiveCallResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Flags of the voice channel.
	//
	// example:
	//
	// []
	ChannelFlags *string `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// ch:user:131888****->8001****:1609225718294:job-6570007401392****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The status of the voice channel.
	//
	// example:
	//
	// NONE
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// Callee of the call channel.
	//
	// example:
	//
	// 8001****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6570007401392****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The originator of the call channel.
	//
	// example:
	//
	// 1318888****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the disconnection of the voice channel, indicating who first hung up the call.
	//
	// example:
	//
	// 无
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The release reason for the voice channel, indicating why the current voice channel was disconnected. The value comes from the response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 无
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The UNIX timestamp indicating the most recent status change of the channel, in milliseconds.
	//
	// example:
	//
	// 1609225718295
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The extension number of the agent associated with the channel.
	//
	// example:
	//
	// 8001****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the call channel. This field is empty if the channel belongs to a Customer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartPredictiveCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type StartPredictiveCallResponseBodyDataUserContext struct {
	// Break code, which can be either system-defined or customer-defined. System-defined break codes include: Warm-up (temporary break state after an agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on customer-defined break codes; customers can define them as needed for their business.
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
	// Device status.
	//
	// example:
	//
	// ONLINE
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// Agent extension number.
	//
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat from the agent was received, in UNIX timestamp format, in milliseconds.
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
	// job-6570007401392****
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
	// The timestamp in milliseconds of the most recent time the agent was reserved. Being reserved means an incoming call will soon be assigned to this agent. The format is a UNIX timestamp in milliseconds.
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
	// READY
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s StartPredictiveCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetBreakCode(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetDeviceId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetDeviceState(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetExtension(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetHeartbeat(v int64) *StartPredictiveCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetInstanceId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetJobId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetMobile(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *StartPredictiveCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetReserved(v int64) *StartPredictiveCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *StartPredictiveCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetUserId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetUserState(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetWorkMode(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
