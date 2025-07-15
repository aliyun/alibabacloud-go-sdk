// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartConferenceResponseBody
	GetCode() *string
	SetData(v *StartConferenceResponseBodyData) *StartConferenceResponseBody
	GetData() *StartConferenceResponseBodyData
	SetHttpStatusCode(v int32) *StartConferenceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartConferenceResponseBody
	GetMessage() *string
	SetParams(v []*string) *StartConferenceResponseBody
	GetParams() []*string
	SetRequestId(v string) *StartConferenceResponseBody
	GetRequestId() *string
}

type StartConferenceResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *StartConferenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// FDD327D1-AB8A-596B-883F-F63582A73F1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *StartConferenceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartConferenceResponseBody) GetData() *StartConferenceResponseBodyData {
	return s.Data
}

func (s *StartConferenceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartConferenceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartConferenceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *StartConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartConferenceResponseBody) SetCode(v string) *StartConferenceResponseBody {
	s.Code = &v
	return s
}

func (s *StartConferenceResponseBody) SetData(v *StartConferenceResponseBodyData) *StartConferenceResponseBody {
	s.Data = v
	return s
}

func (s *StartConferenceResponseBody) SetHttpStatusCode(v int32) *StartConferenceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartConferenceResponseBody) SetMessage(v string) *StartConferenceResponseBody {
	s.Message = &v
	return s
}

func (s *StartConferenceResponseBody) SetParams(v []*string) *StartConferenceResponseBody {
	s.Params = v
	return s
}

func (s *StartConferenceResponseBody) SetRequestId(v string) *StartConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartConferenceResponseBodyData struct {
	CallContext *StartConferenceResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *StartConferenceResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s StartConferenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartConferenceResponseBodyData) GetCallContext() *StartConferenceResponseBodyDataCallContext {
	return s.CallContext
}

func (s *StartConferenceResponseBodyData) GetUserContext() *StartConferenceResponseBodyDataUserContext {
	return s.UserContext
}

func (s *StartConferenceResponseBodyData) SetCallContext(v *StartConferenceResponseBodyDataCallContext) *StartConferenceResponseBodyData {
	s.CallContext = v
	return s
}

func (s *StartConferenceResponseBodyData) SetUserContext(v *StartConferenceResponseBodyDataUserContext) *StartConferenceResponseBodyData {
	s.UserContext = v
	return s
}

func (s *StartConferenceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type StartConferenceResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                      `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*StartConferenceResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StartConferenceResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *StartConferenceResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *StartConferenceResponseBodyDataCallContext) GetChannelContexts() []*StartConferenceResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *StartConferenceResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartConferenceResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *StartConferenceResponseBodyDataCallContext) SetCallType(v string) *StartConferenceResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContext) SetChannelContexts(v []*StartConferenceResponseBodyDataCallContextChannelContexts) *StartConferenceResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *StartConferenceResponseBodyDataCallContext) SetInstanceId(v string) *StartConferenceResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContext) SetJobId(v string) *StartConferenceResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type StartConferenceResponseBodyDataCallContextChannelContexts struct {
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
	// ch:user:131888****->8001****:1609225718294:job-65700074013925376
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// example:
	//
	// 8001****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// job-6573574060089****
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

func (s StartConferenceResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetCallType(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetDestination(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetJobId(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) SetUserId(v string) *StartConferenceResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *StartConferenceResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type StartConferenceResponseBodyDataUserContext struct {
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
	// 1609136956378
	Heartbeat *int64 `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	// example:
	//
	// 1609136956378
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

func (s StartConferenceResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *StartConferenceResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *StartConferenceResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *StartConferenceResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *StartConferenceResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *StartConferenceResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *StartConferenceResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartConferenceResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *StartConferenceResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *StartConferenceResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *StartConferenceResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *StartConferenceResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *StartConferenceResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *StartConferenceResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *StartConferenceResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *StartConferenceResponseBodyDataUserContext) SetBreakCode(v string) *StartConferenceResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetDeviceId(v string) *StartConferenceResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetDeviceState(v string) *StartConferenceResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetExtension(v string) *StartConferenceResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetHeartbeat(v int64) *StartConferenceResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetInstanceId(v string) *StartConferenceResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetJobId(v string) *StartConferenceResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetMobile(v string) *StartConferenceResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetOutboundScenario(v bool) *StartConferenceResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetReserved(v int64) *StartConferenceResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *StartConferenceResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetUserId(v string) *StartConferenceResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetUserState(v string) *StartConferenceResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) SetWorkMode(v string) *StartConferenceResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *StartConferenceResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
