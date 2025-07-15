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
	// example:
	//
	// OK
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AnswerCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type AnswerCallResponseBodyData struct {
	CallContext *AnswerCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// example:
	//
	// 103655
	ContextId   *int64                                 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
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
	return dara.Validate(s)
}

type AnswerCallResponseBodyDataCallContext struct {
	// example:
	//
	// INBOUND
	CallType        *string                                                 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*AnswerCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	return dara.Validate(s)
}

type AnswerCallResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// INBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-653821410368****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// ANSWERED
	ChannelState     *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	ChannelVariables *string `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Index       *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0830019****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// example:
	//
	// 404 - No destination
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// 1609138903315
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 8032****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
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
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 1609136956378
	Heartbeat *int64 `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1324730****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// 1609136956378
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
