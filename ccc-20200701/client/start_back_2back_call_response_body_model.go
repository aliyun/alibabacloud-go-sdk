// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBack2BackCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartBack2BackCallResponseBody
	GetCode() *string
	SetData(v *StartBack2BackCallResponseBodyData) *StartBack2BackCallResponseBody
	GetData() *StartBack2BackCallResponseBodyData
	SetHttpStatusCode(v int32) *StartBack2BackCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartBack2BackCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *StartBack2BackCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *StartBack2BackCallResponseBody
	GetRequestId() *string
}

type StartBack2BackCallResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *StartBack2BackCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s StartBack2BackCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartBack2BackCallResponseBody) GetData() *StartBack2BackCallResponseBodyData {
	return s.Data
}

func (s *StartBack2BackCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartBack2BackCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartBack2BackCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *StartBack2BackCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBack2BackCallResponseBody) SetCode(v string) *StartBack2BackCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetData(v *StartBack2BackCallResponseBodyData) *StartBack2BackCallResponseBody {
	s.Data = v
	return s
}

func (s *StartBack2BackCallResponseBody) SetHttpStatusCode(v int32) *StartBack2BackCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetMessage(v string) *StartBack2BackCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetParams(v []*string) *StartBack2BackCallResponseBody {
	s.Params = v
	return s
}

func (s *StartBack2BackCallResponseBody) SetRequestId(v string) *StartBack2BackCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBack2BackCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartBack2BackCallResponseBodyData struct {
	// The call context environment.
	CallContext *StartBack2BackCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// Agent context environment.
	UserContext *StartBack2BackCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s StartBack2BackCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyData) GetCallContext() *StartBack2BackCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *StartBack2BackCallResponseBodyData) GetUserContext() *StartBack2BackCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *StartBack2BackCallResponseBodyData) SetCallContext(v *StartBack2BackCallResponseBodyDataCallContext) *StartBack2BackCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *StartBack2BackCallResponseBodyData) SetUserContext(v *StartBack2BackCallResponseBodyDataUserContext) *StartBack2BackCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *StartBack2BackCallResponseBodyData) Validate() error {
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

type StartBack2BackCallResponseBodyDataCallContext struct {
	// The call type.
	//
	// example:
	//
	// BACK2BACK
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// List of call channels.
	ChannelContexts []*StartBack2BackCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call ID.
	//
	// example:
	//
	// job-1034159089076****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StartBack2BackCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *StartBack2BackCallResponseBodyDataCallContext) GetChannelContexts() []*StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *StartBack2BackCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartBack2BackCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetCallType(v string) *StartBack2BackCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetChannelContexts(v []*StartBack2BackCallResponseBodyDataCallContextChannelContexts) *StartBack2BackCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetInstanceId(v string) *StartBack2BackCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetJobId(v string) *StartBack2BackCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) Validate() error {
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

type StartBack2BackCallResponseBodyDataCallContextChannelContexts struct {
	// The call type of the channel.
	//
	// example:
	//
	// BACK2BACK
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Flags of the voice channel.
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
	// The status of the voice channel.
	//
	// example:
	//
	// NONE
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// The called party of the call channel.
	//
	// example:
	//
	// 1372168****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-1034159089076****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The originator of the voice channel
	//
	// example:
	//
	// 0102157****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The party that initiated the release of the voice channel, indicating who first disconnected the channel.
	//
	// example:
	//
	// 无
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The reason for releasing the voice channel, indicating why the current voice channel was disconnected. The value is derived from the response codes defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the disconnection reason.
	//
	// example:
	//
	// 无
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The UNIX timestamp indicating the most recent status change of the channel, in milliseconds.
	//
	// example:
	//
	// 1618217874062
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
	// 无
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartBack2BackCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type StartBack2BackCallResponseBodyDataUserContext struct {
	// Break status code, which can be either system-defined or customer-defined. System-defined break codes include: Warm-up (temporary break state after the agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on customer-defined status codes; customers can define them as needed for their business.
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Device ID, which is the identity of a browser-based Web Real-Time Communication (WebRTC) softphone or a physical phone device. Only one type of device can be registered at a time.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Device status.
	//
	// example:
	//
	// 无
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// Agent extension number.
	//
	// example:
	//
	// 8020****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The time when the last heartbeat from the agent was received, in Unix timestamp format, in milliseconds.
	//
	// example:
	//
	// 1618217872911
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
	// job-1034159089076****
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
	// The UNIX timestamp (in milliseconds) indicating when the agent was most recently reserved. Being reserved means an incoming call will be assigned to the agent shortly.
	//
	// example:
	//
	// 1618217794599
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
	// OFFLINE
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s StartBack2BackCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *StartBack2BackCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetBreakCode(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetDeviceId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetDeviceState(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetExtension(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetHeartbeat(v int64) *StartBack2BackCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetInstanceId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetJobId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetMobile(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *StartBack2BackCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetReserved(v int64) *StartBack2BackCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *StartBack2BackCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetUserId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetUserState(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetWorkMode(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
