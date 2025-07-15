// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetrieveCallResponseBody
	GetCode() *string
	SetData(v *RetrieveCallResponseBodyData) *RetrieveCallResponseBody
	GetData() *RetrieveCallResponseBodyData
	SetHttpStatusCode(v int32) *RetrieveCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RetrieveCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *RetrieveCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *RetrieveCallResponseBody
	GetRequestId() *string
}

type RetrieveCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RetrieveCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s RetrieveCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetrieveCallResponseBody) GetData() *RetrieveCallResponseBodyData {
	return s.Data
}

func (s *RetrieveCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RetrieveCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetrieveCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RetrieveCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveCallResponseBody) SetCode(v string) *RetrieveCallResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveCallResponseBody) SetData(v *RetrieveCallResponseBodyData) *RetrieveCallResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveCallResponseBody) SetHttpStatusCode(v int32) *RetrieveCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetrieveCallResponseBody) SetMessage(v string) *RetrieveCallResponseBody {
	s.Message = &v
	return s
}

func (s *RetrieveCallResponseBody) SetParams(v []*string) *RetrieveCallResponseBody {
	s.Params = v
	return s
}

func (s *RetrieveCallResponseBody) SetRequestId(v string) *RetrieveCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type RetrieveCallResponseBodyData struct {
	CallContext *RetrieveCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *RetrieveCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s RetrieveCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyData) GetCallContext() *RetrieveCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *RetrieveCallResponseBodyData) GetUserContext() *RetrieveCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *RetrieveCallResponseBodyData) SetCallContext(v *RetrieveCallResponseBodyDataCallContext) *RetrieveCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *RetrieveCallResponseBodyData) SetUserContext(v *RetrieveCallResponseBodyDataUserContext) *RetrieveCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *RetrieveCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RetrieveCallResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                   `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*RetrieveCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s RetrieveCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *RetrieveCallResponseBodyDataCallContext) GetChannelContexts() []*RetrieveCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *RetrieveCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RetrieveCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *RetrieveCallResponseBodyDataCallContext) SetCallType(v string) *RetrieveCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) SetChannelContexts(v []*RetrieveCallResponseBodyDataCallContextChannelContexts) *RetrieveCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) SetInstanceId(v string) *RetrieveCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) SetJobId(v string) *RetrieveCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type RetrieveCallResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-653821410368****
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

func (s RetrieveCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type RetrieveCallResponseBodyDataUserContext struct {
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
	// BREAK
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s RetrieveCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *RetrieveCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RetrieveCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *RetrieveCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *RetrieveCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RetrieveCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *RetrieveCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *RetrieveCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *RetrieveCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *RetrieveCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *RetrieveCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *RetrieveCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *RetrieveCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *RetrieveCallResponseBodyDataUserContext) SetBreakCode(v string) *RetrieveCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetDeviceId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetExtension(v string) *RetrieveCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetHeartbeat(v int64) *RetrieveCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetInstanceId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetJobId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetMobile(v string) *RetrieveCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *RetrieveCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetReserved(v int64) *RetrieveCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *RetrieveCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetUserId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetUserState(v string) *RetrieveCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetWorkMode(v string) *RetrieveCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
