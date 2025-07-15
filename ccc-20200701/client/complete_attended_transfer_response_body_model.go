// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteAttendedTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CompleteAttendedTransferResponseBody
	GetCode() *string
	SetData(v *CompleteAttendedTransferResponseBodyData) *CompleteAttendedTransferResponseBody
	GetData() *CompleteAttendedTransferResponseBodyData
	SetHttpStatusCode(v int32) *CompleteAttendedTransferResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CompleteAttendedTransferResponseBody
	GetMessage() *string
	SetParams(v []*string) *CompleteAttendedTransferResponseBody
	GetParams() []*string
	SetRequestId(v string) *CompleteAttendedTransferResponseBody
	GetRequestId() *string
}

type CompleteAttendedTransferResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CompleteAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CompleteAttendedTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompleteAttendedTransferResponseBody) GetData() *CompleteAttendedTransferResponseBodyData {
	return s.Data
}

func (s *CompleteAttendedTransferResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CompleteAttendedTransferResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompleteAttendedTransferResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CompleteAttendedTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteAttendedTransferResponseBody) SetCode(v string) *CompleteAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetData(v *CompleteAttendedTransferResponseBodyData) *CompleteAttendedTransferResponseBody {
	s.Data = v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetHttpStatusCode(v int32) *CompleteAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetMessage(v string) *CompleteAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetParams(v []*string) *CompleteAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetRequestId(v string) *CompleteAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) Validate() error {
	return dara.Validate(s)
}

type CompleteAttendedTransferResponseBodyData struct {
	CallContext *CompleteAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// example:
	//
	// 103652
	ContextId   *int64                                               `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	UserContext *CompleteAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CompleteAttendedTransferResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyData) GetCallContext() *CompleteAttendedTransferResponseBodyDataCallContext {
	return s.CallContext
}

func (s *CompleteAttendedTransferResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *CompleteAttendedTransferResponseBodyData) GetUserContext() *CompleteAttendedTransferResponseBodyDataUserContext {
	return s.UserContext
}

func (s *CompleteAttendedTransferResponseBodyData) SetCallContext(v *CompleteAttendedTransferResponseBodyDataCallContext) *CompleteAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) SetContextId(v int64) *CompleteAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) SetUserContext(v *CompleteAttendedTransferResponseBodyDataUserContext) *CompleteAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CompleteAttendedTransferResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CompleteAttendedTransferResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetChannelContexts() []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type CompleteAttendedTransferResponseBodyDataCallContextChannelContexts struct {
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
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 10
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
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

func (s CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type CompleteAttendedTransferResponseBodyDataUserContext struct {
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
	// job-65382141036853491
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

func (s CompleteAttendedTransferResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
