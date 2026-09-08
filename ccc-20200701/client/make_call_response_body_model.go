// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MakeCallResponseBody
	GetCode() *string
	SetData(v *MakeCallResponseBodyData) *MakeCallResponseBody
	GetData() *MakeCallResponseBodyData
	SetHttpStatusCode(v int32) *MakeCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *MakeCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *MakeCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *MakeCallResponseBody
	GetRequestId() *string
}

type MakeCallResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *MakeCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MakeCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *MakeCallResponseBody) GetData() *MakeCallResponseBodyData {
	return s.Data
}

func (s *MakeCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *MakeCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MakeCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *MakeCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MakeCallResponseBody) SetCode(v string) *MakeCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakeCallResponseBody) SetData(v *MakeCallResponseBodyData) *MakeCallResponseBody {
	s.Data = v
	return s
}

func (s *MakeCallResponseBody) SetHttpStatusCode(v int32) *MakeCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MakeCallResponseBody) SetMessage(v string) *MakeCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakeCallResponseBody) SetParams(v []*string) *MakeCallResponseBody {
	s.Params = v
	return s
}

func (s *MakeCallResponseBody) SetRequestId(v string) *MakeCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MakeCallResponseBodyData struct {
	// Call context environment.
	CallContext *MakeCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// System auto-increment ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 123456
	ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Agent context environment.
	UserContext *MakeCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s MakeCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyData) GetCallContext() *MakeCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *MakeCallResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *MakeCallResponseBodyData) GetUserContext() *MakeCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *MakeCallResponseBodyData) SetCallContext(v *MakeCallResponseBodyDataCallContext) *MakeCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *MakeCallResponseBodyData) SetContextId(v int64) *MakeCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *MakeCallResponseBodyData) SetUserContext(v *MakeCallResponseBodyDataUserContext) *MakeCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *MakeCallResponseBodyData) Validate() error {
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

type MakeCallResponseBodyDataCallContext struct {
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
	ChannelContexts []*MakeCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
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

func (s MakeCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *MakeCallResponseBodyDataCallContext) GetCallVariables() *string {
	return s.CallVariables
}

func (s *MakeCallResponseBodyDataCallContext) GetChannelContexts() []*MakeCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *MakeCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MakeCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *MakeCallResponseBodyDataCallContext) SetCallType(v string) *MakeCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetCallVariables(v string) *MakeCallResponseBodyDataCallContext {
	s.CallVariables = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetChannelContexts(v []*MakeCallResponseBodyDataCallContextChannelContexts) *MakeCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetInstanceId(v string) *MakeCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetJobId(v string) *MakeCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) Validate() error {
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

type MakeCallResponseBodyDataCallContextChannelContexts struct {
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
	// NONE
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// Callee of the call channel.
	//
	// example:
	//
	// 8001****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// Call ID.
	//
	// example:
	//
	// job-6570007401392****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Media type. The default value is AUDIO. Other valid values include VIDEO.
	//
	// example:
	//
	// Audio
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
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
	// The reason for releasing the channel. This value indicates why the current channel was disconnected and corresponds to a response code defined in the SIP protocol. Customers should refer to the SIP protocol to analyze the disconnection reason.
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
	// The agent ID associated with the channel. This field is empty for customer channels.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s MakeCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetMediaType() *string {
	return s.MediaType
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetMediaType(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.MediaType = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type MakeCallResponseBodyDataUserContext struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent call rejection). There are no restrictions on Custom-defined status codes, and Customers can define them according to their business needs.
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
	// Indicates whether the agent is in outbound-only mode.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
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

func (s MakeCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *MakeCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MakeCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *MakeCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MakeCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *MakeCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *MakeCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *MakeCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *MakeCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *MakeCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *MakeCallResponseBodyDataUserContext) SetBreakCode(v string) *MakeCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetDeviceId(v string) *MakeCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetExtension(v string) *MakeCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetInstanceId(v string) *MakeCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetJobId(v string) *MakeCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *MakeCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *MakeCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetUserId(v string) *MakeCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetUserState(v string) *MakeCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetWorkMode(v string) *MakeCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
