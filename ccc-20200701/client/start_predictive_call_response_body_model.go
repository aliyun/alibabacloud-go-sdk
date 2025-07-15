// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPredictiveCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartPredictiveCallResponseBody
	GetCode() *string
	SetData(v *StartPredictiveCallResponseBodyData) *StartPredictiveCallResponseBody
	GetData() *StartPredictiveCallResponseBodyData
	SetHttpStatusCode(v int32) *StartPredictiveCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartPredictiveCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *StartPredictiveCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *StartPredictiveCallResponseBody
	GetRequestId() *string
}

type StartPredictiveCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *StartPredictiveCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 26A34338-5CD9-4C95-A7A6-5BDCE76C6B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPredictiveCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartPredictiveCallResponseBody) GetData() *StartPredictiveCallResponseBodyData {
	return s.Data
}

func (s *StartPredictiveCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartPredictiveCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartPredictiveCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *StartPredictiveCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPredictiveCallResponseBody) SetCode(v string) *StartPredictiveCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetData(v *StartPredictiveCallResponseBodyData) *StartPredictiveCallResponseBody {
	s.Data = v
	return s
}

func (s *StartPredictiveCallResponseBody) SetHttpStatusCode(v int32) *StartPredictiveCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetMessage(v string) *StartPredictiveCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetParams(v []*string) *StartPredictiveCallResponseBody {
	s.Params = v
	return s
}

func (s *StartPredictiveCallResponseBody) SetRequestId(v string) *StartPredictiveCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPredictiveCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartPredictiveCallResponseBodyData struct {
	CallContext *StartPredictiveCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *StartPredictiveCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s StartPredictiveCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyData) GetCallContext() *StartPredictiveCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *StartPredictiveCallResponseBodyData) GetUserContext() *StartPredictiveCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *StartPredictiveCallResponseBodyData) SetCallContext(v *StartPredictiveCallResponseBodyDataCallContext) *StartPredictiveCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *StartPredictiveCallResponseBodyData) SetUserContext(v *StartPredictiveCallResponseBodyDataUserContext) *StartPredictiveCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *StartPredictiveCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type StartPredictiveCallResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                          `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*StartPredictiveCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6570007401392****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StartPredictiveCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetChannelContexts() []*StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPredictiveCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetCallType(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetChannelContexts(v []*StartPredictiveCallResponseBodyDataCallContextChannelContexts) *StartPredictiveCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetInstanceId(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetJobId(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type StartPredictiveCallResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// OUTBOUND
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// []
	ChannelFlags *string `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	// example:
	//
	// ch:user:131888****->8001****:1609225718294:job-6570007401392****
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
	// job-6570007401392****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1318888****
	Originator       *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ReleaseReason    *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// 1609225718295
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

func (s StartPredictiveCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type StartPredictiveCallResponseBodyDataUserContext struct {
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
	// ONLINE
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// example:
	//
	// 8001****
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
	// job-6570007401392****
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
	// READY
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s StartPredictiveCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *StartPredictiveCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetBreakCode(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetDeviceId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetDeviceState(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetExtension(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetHeartbeat(v int64) *StartPredictiveCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetInstanceId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetJobId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetMobile(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *StartPredictiveCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetReserved(v int64) *StartPredictiveCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *StartPredictiveCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetUserId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetUserState(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetWorkMode(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
