// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MonitorCallResponseBody
	GetCode() *string
	SetData(v *MonitorCallResponseBodyData) *MonitorCallResponseBody
	GetData() *MonitorCallResponseBodyData
	SetHttpStatusCode(v int32) *MonitorCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *MonitorCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *MonitorCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *MonitorCallResponseBody
	GetRequestId() *string
}

type MonitorCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MonitorCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MonitorCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallResponseBody) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *MonitorCallResponseBody) GetData() *MonitorCallResponseBodyData {
	return s.Data
}

func (s *MonitorCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *MonitorCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MonitorCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *MonitorCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorCallResponseBody) SetCode(v string) *MonitorCallResponseBody {
	s.Code = &v
	return s
}

func (s *MonitorCallResponseBody) SetData(v *MonitorCallResponseBodyData) *MonitorCallResponseBody {
	s.Data = v
	return s
}

func (s *MonitorCallResponseBody) SetHttpStatusCode(v int32) *MonitorCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MonitorCallResponseBody) SetMessage(v string) *MonitorCallResponseBody {
	s.Message = &v
	return s
}

func (s *MonitorCallResponseBody) SetParams(v []*string) *MonitorCallResponseBody {
	s.Params = v
	return s
}

func (s *MonitorCallResponseBody) SetRequestId(v string) *MonitorCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonitorCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type MonitorCallResponseBodyData struct {
	CallContext *MonitorCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *MonitorCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s MonitorCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyData) GetCallContext() *MonitorCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *MonitorCallResponseBodyData) GetUserContext() *MonitorCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *MonitorCallResponseBodyData) SetCallContext(v *MonitorCallResponseBodyDataCallContext) *MonitorCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *MonitorCallResponseBodyData) SetUserContext(v *MonitorCallResponseBodyDataUserContext) *MonitorCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *MonitorCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type MonitorCallResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*MonitorCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s MonitorCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *MonitorCallResponseBodyDataCallContext) GetChannelContexts() []*MonitorCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *MonitorCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MonitorCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *MonitorCallResponseBodyDataCallContext) SetCallType(v string) *MonitorCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) SetChannelContexts(v []*MonitorCallResponseBodyDataCallContextChannelContexts) *MonitorCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) SetInstanceId(v string) *MonitorCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) SetJobId(v string) *MonitorCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type MonitorCallResponseBodyDataCallContextChannelContexts struct {
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

func (s MonitorCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type MonitorCallResponseBodyDataUserContext struct {
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

func (s MonitorCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *MonitorCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MonitorCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *MonitorCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *MonitorCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MonitorCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *MonitorCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *MonitorCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *MonitorCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *MonitorCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *MonitorCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *MonitorCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *MonitorCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *MonitorCallResponseBodyDataUserContext) SetBreakCode(v string) *MonitorCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetDeviceId(v string) *MonitorCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetExtension(v string) *MonitorCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetHeartbeat(v int64) *MonitorCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetInstanceId(v string) *MonitorCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetJobId(v string) *MonitorCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetMobile(v string) *MonitorCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *MonitorCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetReserved(v int64) *MonitorCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *MonitorCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetUserId(v string) *MonitorCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetUserState(v string) *MonitorCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetWorkMode(v string) *MonitorCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
