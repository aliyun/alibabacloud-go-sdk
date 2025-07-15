// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateAttendedTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitiateAttendedTransferResponseBody
	GetCode() *string
	SetData(v *InitiateAttendedTransferResponseBodyData) *InitiateAttendedTransferResponseBody
	GetData() *InitiateAttendedTransferResponseBodyData
	SetHttpStatusCode(v int32) *InitiateAttendedTransferResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InitiateAttendedTransferResponseBody
	GetMessage() *string
	SetParams(v []*string) *InitiateAttendedTransferResponseBody
	GetParams() []*string
	SetRequestId(v string) *InitiateAttendedTransferResponseBody
	GetRequestId() *string
}

type InitiateAttendedTransferResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InitiateAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s InitiateAttendedTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitiateAttendedTransferResponseBody) GetData() *InitiateAttendedTransferResponseBodyData {
	return s.Data
}

func (s *InitiateAttendedTransferResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InitiateAttendedTransferResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitiateAttendedTransferResponseBody) GetParams() []*string {
	return s.Params
}

func (s *InitiateAttendedTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitiateAttendedTransferResponseBody) SetCode(v string) *InitiateAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetData(v *InitiateAttendedTransferResponseBodyData) *InitiateAttendedTransferResponseBody {
	s.Data = v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetHttpStatusCode(v int32) *InitiateAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetMessage(v string) *InitiateAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetParams(v []*string) *InitiateAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetRequestId(v string) *InitiateAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) Validate() error {
	return dara.Validate(s)
}

type InitiateAttendedTransferResponseBodyData struct {
	CallContext *InitiateAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// example:
	//
	// 103655
	ContextId   *int64                                               `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	UserContext *InitiateAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s InitiateAttendedTransferResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyData) GetCallContext() *InitiateAttendedTransferResponseBodyDataCallContext {
	return s.CallContext
}

func (s *InitiateAttendedTransferResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *InitiateAttendedTransferResponseBodyData) GetUserContext() *InitiateAttendedTransferResponseBodyDataUserContext {
	return s.UserContext
}

func (s *InitiateAttendedTransferResponseBodyData) SetCallContext(v *InitiateAttendedTransferResponseBodyDataCallContext) *InitiateAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) SetContextId(v int64) *InitiateAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) SetUserContext(v *InitiateAttendedTransferResponseBodyDataUserContext) *InitiateAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type InitiateAttendedTransferResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s InitiateAttendedTransferResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetChannelContexts() []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type InitiateAttendedTransferResponseBodyDataCallContextChannelContexts struct {
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
	// ch:user:139xxxx0501->80326034:1609138902226:job-6538214103685****
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

func (s InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type InitiateAttendedTransferResponseBodyDataUserContext struct {
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

func (s InitiateAttendedTransferResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
