// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchSurveyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LaunchSurveyResponseBody
	GetCode() *string
	SetData(v *LaunchSurveyResponseBodyData) *LaunchSurveyResponseBody
	GetData() *LaunchSurveyResponseBodyData
	SetHttpStatusCode(v int32) *LaunchSurveyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *LaunchSurveyResponseBody
	GetMessage() *string
	SetParams(v []*string) *LaunchSurveyResponseBody
	GetParams() []*string
	SetRequestId(v string) *LaunchSurveyResponseBody
	GetRequestId() *string
}

type LaunchSurveyResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *LaunchSurveyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// AF1E5957-5276-48FF-A6E6-347166A4ADCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LaunchSurveyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBody) GetCode() *string {
	return s.Code
}

func (s *LaunchSurveyResponseBody) GetData() *LaunchSurveyResponseBodyData {
	return s.Data
}

func (s *LaunchSurveyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *LaunchSurveyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LaunchSurveyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *LaunchSurveyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LaunchSurveyResponseBody) SetCode(v string) *LaunchSurveyResponseBody {
	s.Code = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetData(v *LaunchSurveyResponseBodyData) *LaunchSurveyResponseBody {
	s.Data = v
	return s
}

func (s *LaunchSurveyResponseBody) SetHttpStatusCode(v int32) *LaunchSurveyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetMessage(v string) *LaunchSurveyResponseBody {
	s.Message = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetParams(v []*string) *LaunchSurveyResponseBody {
	s.Params = v
	return s
}

func (s *LaunchSurveyResponseBody) SetRequestId(v string) *LaunchSurveyResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchSurveyResponseBody) Validate() error {
	return dara.Validate(s)
}

type LaunchSurveyResponseBodyData struct {
	CallContext *LaunchSurveyResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// example:
	//
	// 102323
	ContextId   *int64                                   `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	UserContext *LaunchSurveyResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s LaunchSurveyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyResponseBodyData) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyData) GetCallContext() *LaunchSurveyResponseBodyDataCallContext {
	return s.CallContext
}

func (s *LaunchSurveyResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *LaunchSurveyResponseBodyData) GetUserContext() *LaunchSurveyResponseBodyDataUserContext {
	return s.UserContext
}

func (s *LaunchSurveyResponseBodyData) SetCallContext(v *LaunchSurveyResponseBodyDataCallContext) *LaunchSurveyResponseBodyData {
	s.CallContext = v
	return s
}

func (s *LaunchSurveyResponseBodyData) SetContextId(v int64) *LaunchSurveyResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *LaunchSurveyResponseBodyData) SetUserContext(v *LaunchSurveyResponseBodyDataUserContext) *LaunchSurveyResponseBodyData {
	s.UserContext = v
	return s
}

func (s *LaunchSurveyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type LaunchSurveyResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                   `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*LaunchSurveyResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6580466654649****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s LaunchSurveyResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *LaunchSurveyResponseBodyDataCallContext) GetChannelContexts() []*LaunchSurveyResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *LaunchSurveyResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LaunchSurveyResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetCallType(v string) *LaunchSurveyResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetChannelContexts(v []*LaunchSurveyResponseBodyDataCallContextChannelContexts) *LaunchSurveyResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetInstanceId(v string) *LaunchSurveyResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetJobId(v string) *LaunchSurveyResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type LaunchSurveyResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// MONITORING
	ChannelFlags *string `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-653821410368****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// example:
	//
	// 1318888****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// job-6580466654649****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1318888****
	Originator       *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ReleaseReason    *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// 1609250655922
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 8001****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LaunchSurveyResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetCallType(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetDestination(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetJobId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetUserId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type LaunchSurveyResponseBodyDataUserContext struct {
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 1609250656122
	Heartbeat *int64 `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6580466654649****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// 1609250655090
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// TALKING
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s LaunchSurveyResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *LaunchSurveyResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetBreakCode(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetDeviceId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetExtension(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetHeartbeat(v int64) *LaunchSurveyResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetInstanceId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetJobId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetMobile(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetOutboundScenario(v bool) *LaunchSurveyResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetReserved(v int64) *LaunchSurveyResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *LaunchSurveyResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetUserId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetUserState(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetWorkMode(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
