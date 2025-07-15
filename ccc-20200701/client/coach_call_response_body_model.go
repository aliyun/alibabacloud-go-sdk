// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCoachCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CoachCallResponseBody
	GetCode() *string
	SetData(v *CoachCallResponseBodyData) *CoachCallResponseBody
	GetData() *CoachCallResponseBodyData
	SetHttpStatusCode(v int32) *CoachCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CoachCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *CoachCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *CoachCallResponseBody
	GetRequestId() *string
}

type CoachCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CoachCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CoachCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBody) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *CoachCallResponseBody) GetData() *CoachCallResponseBodyData {
	return s.Data
}

func (s *CoachCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CoachCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CoachCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CoachCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CoachCallResponseBody) SetCode(v string) *CoachCallResponseBody {
	s.Code = &v
	return s
}

func (s *CoachCallResponseBody) SetData(v *CoachCallResponseBodyData) *CoachCallResponseBody {
	s.Data = v
	return s
}

func (s *CoachCallResponseBody) SetHttpStatusCode(v int32) *CoachCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CoachCallResponseBody) SetMessage(v string) *CoachCallResponseBody {
	s.Message = &v
	return s
}

func (s *CoachCallResponseBody) SetParams(v []*string) *CoachCallResponseBody {
	s.Params = v
	return s
}

func (s *CoachCallResponseBody) SetRequestId(v string) *CoachCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CoachCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type CoachCallResponseBodyData struct {
	CallContext *CoachCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *CoachCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CoachCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyData) GetCallContext() *CoachCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *CoachCallResponseBodyData) GetUserContext() *CoachCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *CoachCallResponseBodyData) SetCallContext(v *CoachCallResponseBodyDataCallContext) *CoachCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CoachCallResponseBodyData) SetUserContext(v *CoachCallResponseBodyDataUserContext) *CoachCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *CoachCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CoachCallResponseBodyDataCallContext struct {
	// example:
	//
	// COACH
	CallType        *string                                                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*CoachCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CoachCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *CoachCallResponseBodyDataCallContext) GetChannelContexts() []*CoachCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *CoachCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CoachCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallResponseBodyDataCallContext) SetCallType(v string) *CoachCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetChannelContexts(v []*CoachCallResponseBodyDataCallContextChannelContexts) *CoachCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetInstanceId(v string) *CoachCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetJobId(v string) *CoachCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type CoachCallResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// COACH
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// COACHING
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

func (s CoachCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetIndex() *int32 {
	return s.Index
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type CoachCallResponseBodyDataUserContext struct {
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
	// UNREGISTERED
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
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
	// 1609136956370
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

func (s CoachCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *CoachCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CoachCallResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *CoachCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *CoachCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *CoachCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CoachCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *CoachCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *CoachCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *CoachCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *CoachCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *CoachCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *CoachCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CoachCallResponseBodyDataUserContext) SetBreakCode(v string) *CoachCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetDeviceId(v string) *CoachCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetDeviceState(v string) *CoachCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetExtension(v string) *CoachCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetHeartbeat(v int64) *CoachCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetInstanceId(v string) *CoachCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetJobId(v string) *CoachCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetMobile(v string) *CoachCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *CoachCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetReserved(v int64) *CoachCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CoachCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUserId(v string) *CoachCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUserState(v string) *CoachCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetWorkMode(v string) *CoachCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
