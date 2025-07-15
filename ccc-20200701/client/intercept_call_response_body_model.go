// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterceptCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InterceptCallResponseBody
	GetCode() *string
	SetData(v *InterceptCallResponseBodyData) *InterceptCallResponseBody
	GetData() *InterceptCallResponseBodyData
	SetHttpStatusCode(v int32) *InterceptCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InterceptCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *InterceptCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *InterceptCallResponseBody
	GetRequestId() *string
}

type InterceptCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InterceptCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s InterceptCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallResponseBody) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *InterceptCallResponseBody) GetData() *InterceptCallResponseBodyData {
	return s.Data
}

func (s *InterceptCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InterceptCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InterceptCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *InterceptCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InterceptCallResponseBody) SetCode(v string) *InterceptCallResponseBody {
	s.Code = &v
	return s
}

func (s *InterceptCallResponseBody) SetData(v *InterceptCallResponseBodyData) *InterceptCallResponseBody {
	s.Data = v
	return s
}

func (s *InterceptCallResponseBody) SetHttpStatusCode(v int32) *InterceptCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InterceptCallResponseBody) SetMessage(v string) *InterceptCallResponseBody {
	s.Message = &v
	return s
}

func (s *InterceptCallResponseBody) SetParams(v []*string) *InterceptCallResponseBody {
	s.Params = v
	return s
}

func (s *InterceptCallResponseBody) SetRequestId(v string) *InterceptCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterceptCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type InterceptCallResponseBodyData struct {
	CallContext *InterceptCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *InterceptCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s InterceptCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyData) GetCallContext() *InterceptCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *InterceptCallResponseBodyData) GetUserContext() *InterceptCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *InterceptCallResponseBodyData) SetCallContext(v *InterceptCallResponseBodyDataCallContext) *InterceptCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *InterceptCallResponseBodyData) SetUserContext(v *InterceptCallResponseBodyDataUserContext) *InterceptCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *InterceptCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type InterceptCallResponseBodyDataCallContext struct {
	// example:
	//
	// INTERCEPT
	CallType        *string                                                    `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*InterceptCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s InterceptCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *InterceptCallResponseBodyDataCallContext) GetChannelContexts() []*InterceptCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *InterceptCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InterceptCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *InterceptCallResponseBodyDataCallContext) SetCallType(v string) *InterceptCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) SetChannelContexts(v []*InterceptCallResponseBodyDataCallContextChannelContexts) *InterceptCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) SetInstanceId(v string) *InterceptCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) SetJobId(v string) *InterceptCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type InterceptCallResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// INTERCEPT
	CallType     *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
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

func (s InterceptCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type InterceptCallResponseBodyDataUserContext struct {
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

func (s InterceptCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *InterceptCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *InterceptCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *InterceptCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *InterceptCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InterceptCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *InterceptCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *InterceptCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *InterceptCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *InterceptCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *InterceptCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *InterceptCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *InterceptCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *InterceptCallResponseBodyDataUserContext) SetBreakCode(v string) *InterceptCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetDeviceId(v string) *InterceptCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetExtension(v string) *InterceptCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetHeartbeat(v int64) *InterceptCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetInstanceId(v string) *InterceptCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetJobId(v string) *InterceptCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetMobile(v string) *InterceptCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *InterceptCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetReserved(v int64) *InterceptCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *InterceptCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetUserId(v string) *InterceptCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetUserState(v string) *InterceptCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetWorkMode(v string) *InterceptCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
