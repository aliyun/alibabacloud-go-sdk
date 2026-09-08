// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCoachCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CoachCallResponseBody
	GetCode() *string
	SetData(v *CoachCallResponseBodyData) *CoachCallResponseBody
	GetData() *CoachCallResponseBodyData
	SetHttpStatusCode(v int32) *CoachCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CoachCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *CoachCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *CoachCallResponseBody
	GetRequestId() *string
}

type CoachCallResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *CoachCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CoachCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBody) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *CoachCallResponseBody) GetData() *CoachCallResponseBodyData {
	return s.Data
}

func (s *CoachCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CoachCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CoachCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CoachCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CoachCallResponseBody) SetCode(v string) *CoachCallResponseBody {
	s.Code = &v
	return s
}

func (s *CoachCallResponseBody) SetData(v *CoachCallResponseBodyData) *CoachCallResponseBody {
	s.Data = v
	return s
}

func (s *CoachCallResponseBody) SetHttpStatusCode(v int32) *CoachCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CoachCallResponseBody) SetMessage(v string) *CoachCallResponseBody {
	s.Message = &v
	return s
}

func (s *CoachCallResponseBody) SetParams(v []*string) *CoachCallResponseBody {
	s.Params = v
	return s
}

func (s *CoachCallResponseBody) SetRequestId(v string) *CoachCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CoachCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CoachCallResponseBodyData struct {
	// Call context environment.
	CallContext *CoachCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *CoachCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CoachCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyData) GetCallContext() *CoachCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *CoachCallResponseBodyData) GetUserContext() *CoachCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *CoachCallResponseBodyData) SetCallContext(v *CoachCallResponseBodyDataCallContext) *CoachCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CoachCallResponseBodyData) SetUserContext(v *CoachCallResponseBodyDataUserContext) *CoachCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *CoachCallResponseBodyData) Validate() error {
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

type CoachCallResponseBodyDataCallContext struct {
	// Call type.
	//
	// example:
	//
	// COACH
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*CoachCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s CoachCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *CoachCallResponseBodyDataCallContext) GetChannelContexts() []*CoachCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *CoachCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CoachCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallResponseBodyDataCallContext) SetCallType(v string) *CoachCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetChannelContexts(v []*CoachCallResponseBodyDataCallContextChannelContexts) *CoachCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetInstanceId(v string) *CoachCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetJobId(v string) *CoachCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) Validate() error {
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

type CoachCallResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// COACH
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Channel flags.
	//
	// example:
	//
	// COACHING
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
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// The called party of the call channel.
	//
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// An auto-incremented ID assigned by the system. Customers do not need to concern themselves with this value.
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
	// The originator of the channel.
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
	// The reason for releasing the channel, indicating why the current channel was disconnected. The value corresponds to response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The skill group ID associated with the channel. In inbound scenarios, the associated skill group ID is determined by the agent transfer module configured in the IVR. In outbound scenarios, the associated skill group ID is the first skill group ID that the agent signed into.
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
	// The extension number of the agent associated with the channel.
	//
	// example:
	//
	// 8032****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// The agent ID associated with the channel. This field is empty for a Customer channel.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CoachCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type CoachCallResponseBodyDataUserContext struct {
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
	// The status of the SIP phone device. If the SIP phone is not registered, the status is UNREGISTERED (unregistered). If the SIP phone was previously registered but is currently offline, the status is OFFLINE (offline). If the SIP phone is registered and online, the status is ONLINE (online).
	//
	// example:
	//
	// UNREGISTERED
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// The agent\\"s extension number.
	//
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The UNIX timestamp in milliseconds of the last heartbeat received from the agent.
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
	// The UNIX timestamp in milliseconds when the agent was most recently reserved. Being reserved means an incoming call will be assigned to the agent shortly.
	//
	// example:
	//
	// 1609136956370
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

func (s CoachCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *CoachCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CoachCallResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *CoachCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *CoachCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *CoachCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CoachCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *CoachCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *CoachCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *CoachCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *CoachCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *CoachCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *CoachCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CoachCallResponseBodyDataUserContext) SetBreakCode(v string) *CoachCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetDeviceId(v string) *CoachCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetDeviceState(v string) *CoachCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetExtension(v string) *CoachCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetHeartbeat(v int64) *CoachCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetInstanceId(v string) *CoachCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetJobId(v string) *CoachCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetMobile(v string) *CoachCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *CoachCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetReserved(v int64) *CoachCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CoachCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUserId(v string) *CoachCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUserState(v string) *CoachCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetWorkMode(v string) *CoachCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
