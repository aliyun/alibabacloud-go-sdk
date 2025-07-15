// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchToConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SwitchToConferenceResponseBody
	GetCode() *string
	SetData(v *SwitchToConferenceResponseBodyData) *SwitchToConferenceResponseBody
	GetData() *SwitchToConferenceResponseBodyData
	SetHttpStatusCode(v int32) *SwitchToConferenceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SwitchToConferenceResponseBody
	GetMessage() *string
	SetParams(v []*string) *SwitchToConferenceResponseBody
	GetParams() []*string
	SetRequestId(v string) *SwitchToConferenceResponseBody
	GetRequestId() *string
}

type SwitchToConferenceResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SwitchToConferenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// D9C96A73-09C9-5E2A-8CDB-85EC0BC246DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchToConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SwitchToConferenceResponseBody) GetData() *SwitchToConferenceResponseBodyData {
	return s.Data
}

func (s *SwitchToConferenceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SwitchToConferenceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SwitchToConferenceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *SwitchToConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchToConferenceResponseBody) SetCode(v string) *SwitchToConferenceResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchToConferenceResponseBody) SetData(v *SwitchToConferenceResponseBodyData) *SwitchToConferenceResponseBody {
	s.Data = v
	return s
}

func (s *SwitchToConferenceResponseBody) SetHttpStatusCode(v int32) *SwitchToConferenceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SwitchToConferenceResponseBody) SetMessage(v string) *SwitchToConferenceResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchToConferenceResponseBody) SetParams(v []*string) *SwitchToConferenceResponseBody {
	s.Params = v
	return s
}

func (s *SwitchToConferenceResponseBody) SetRequestId(v string) *SwitchToConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchToConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}

type SwitchToConferenceResponseBodyData struct {
	CallContext *SwitchToConferenceResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *SwitchToConferenceResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s SwitchToConferenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceResponseBodyData) GetCallContext() *SwitchToConferenceResponseBodyDataCallContext {
	return s.CallContext
}

func (s *SwitchToConferenceResponseBodyData) GetUserContext() *SwitchToConferenceResponseBodyDataUserContext {
	return s.UserContext
}

func (s *SwitchToConferenceResponseBodyData) SetCallContext(v *SwitchToConferenceResponseBodyDataCallContext) *SwitchToConferenceResponseBodyData {
	s.CallContext = v
	return s
}

func (s *SwitchToConferenceResponseBodyData) SetUserContext(v *SwitchToConferenceResponseBodyDataUserContext) *SwitchToConferenceResponseBodyData {
	s.UserContext = v
	return s
}

func (s *SwitchToConferenceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SwitchToConferenceResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                         `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*SwitchToConferenceResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SwitchToConferenceResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *SwitchToConferenceResponseBodyDataCallContext) GetChannelContexts() []*SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *SwitchToConferenceResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchToConferenceResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *SwitchToConferenceResponseBodyDataCallContext) SetCallType(v string) *SwitchToConferenceResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContext) SetChannelContexts(v []*SwitchToConferenceResponseBodyDataCallContextChannelContexts) *SwitchToConferenceResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContext) SetInstanceId(v string) *SwitchToConferenceResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContext) SetJobId(v string) *SwitchToConferenceResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type SwitchToConferenceResponseBodyDataCallContextChannelContexts struct {
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
	// ch:user:131888****->8001****:1609225718294:job-6538214103685****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// NONE
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// example:
	//
	// 8001****
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
	// 1318888****
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
	// 1609255716900
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

func (s SwitchToConferenceResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetCallType(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetDestination(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetJobId(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) SetUserId(v string) *SwitchToConferenceResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type SwitchToConferenceResponseBodyDataUserContext struct {
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
	// false
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// agent@ccc-test
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
	// 1609136956378
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

func (s SwitchToConferenceResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *SwitchToConferenceResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetBreakCode(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetDeviceId(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetDeviceState(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetExtension(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetHeartbeat(v int64) *SwitchToConferenceResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetInstanceId(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetJobId(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetMobile(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetOutboundScenario(v bool) *SwitchToConferenceResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *SwitchToConferenceResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetUserId(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetUserState(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) SetWorkMode(v string) *SwitchToConferenceResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *SwitchToConferenceResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
