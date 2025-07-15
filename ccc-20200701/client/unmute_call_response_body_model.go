// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmuteCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnmuteCallResponseBody
	GetCode() *string
	SetData(v *UnmuteCallResponseBodyData) *UnmuteCallResponseBody
	GetData() *UnmuteCallResponseBodyData
	SetHttpStatusCode(v int32) *UnmuteCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnmuteCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *UnmuteCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *UnmuteCallResponseBody
	GetRequestId() *string
}

type UnmuteCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UnmuteCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UnmuteCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallResponseBody) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnmuteCallResponseBody) GetData() *UnmuteCallResponseBodyData {
	return s.Data
}

func (s *UnmuteCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnmuteCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnmuteCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UnmuteCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnmuteCallResponseBody) SetCode(v string) *UnmuteCallResponseBody {
	s.Code = &v
	return s
}

func (s *UnmuteCallResponseBody) SetData(v *UnmuteCallResponseBodyData) *UnmuteCallResponseBody {
	s.Data = v
	return s
}

func (s *UnmuteCallResponseBody) SetHttpStatusCode(v int32) *UnmuteCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnmuteCallResponseBody) SetMessage(v string) *UnmuteCallResponseBody {
	s.Message = &v
	return s
}

func (s *UnmuteCallResponseBody) SetParams(v []*string) *UnmuteCallResponseBody {
	s.Params = v
	return s
}

func (s *UnmuteCallResponseBody) SetRequestId(v string) *UnmuteCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnmuteCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnmuteCallResponseBodyData struct {
	CallContext *UnmuteCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *UnmuteCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s UnmuteCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyData) GetCallContext() *UnmuteCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *UnmuteCallResponseBodyData) GetUserContext() *UnmuteCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *UnmuteCallResponseBodyData) SetCallContext(v *UnmuteCallResponseBodyDataCallContext) *UnmuteCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *UnmuteCallResponseBodyData) SetUserContext(v *UnmuteCallResponseBodyDataUserContext) *UnmuteCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *UnmuteCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type UnmuteCallResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*UnmuteCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UnmuteCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *UnmuteCallResponseBodyDataCallContext) GetChannelContexts() []*UnmuteCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *UnmuteCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnmuteCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *UnmuteCallResponseBodyDataCallContext) SetCallType(v string) *UnmuteCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) SetChannelContexts(v []*UnmuteCallResponseBodyDataCallContextChannelContexts) *UnmuteCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) SetInstanceId(v string) *UnmuteCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) SetJobId(v string) *UnmuteCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type UnmuteCallResponseBodyDataCallContextChannelContexts struct {
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
	// ch:user:1390501****->8032****:1609138902226:job-6538214103685****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// CREATED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 1
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

func (s UnmuteCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type UnmuteCallResponseBodyDataUserContext struct {
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
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
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

func (s UnmuteCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *UnmuteCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnmuteCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *UnmuteCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *UnmuteCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnmuteCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *UnmuteCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *UnmuteCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *UnmuteCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *UnmuteCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *UnmuteCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *UnmuteCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *UnmuteCallResponseBodyDataUserContext) SetBreakCode(v string) *UnmuteCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetDeviceId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetExtension(v string) *UnmuteCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetHeartbeat(v int64) *UnmuteCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetInstanceId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetJobId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetMobile(v string) *UnmuteCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *UnmuteCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *UnmuteCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetUserId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetUserState(v string) *UnmuteCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetWorkMode(v string) *UnmuteCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
