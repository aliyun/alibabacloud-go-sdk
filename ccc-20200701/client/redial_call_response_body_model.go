// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedialCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RedialCallResponseBody
	GetCode() *string
	SetData(v *RedialCallResponseBodyData) *RedialCallResponseBody
	GetData() *RedialCallResponseBodyData
	SetHttpStatusCode(v int32) *RedialCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RedialCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *RedialCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *RedialCallResponseBody
	GetRequestId() *string
}

type RedialCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RedialCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BF268B34-09C2-43FD-BAC4-5D31EA63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RedialCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedialCallResponseBody) GoString() string {
	return s.String()
}

func (s *RedialCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *RedialCallResponseBody) GetData() *RedialCallResponseBodyData {
	return s.Data
}

func (s *RedialCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RedialCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RedialCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RedialCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedialCallResponseBody) SetCode(v string) *RedialCallResponseBody {
	s.Code = &v
	return s
}

func (s *RedialCallResponseBody) SetData(v *RedialCallResponseBodyData) *RedialCallResponseBody {
	s.Data = v
	return s
}

func (s *RedialCallResponseBody) SetHttpStatusCode(v int32) *RedialCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RedialCallResponseBody) SetMessage(v string) *RedialCallResponseBody {
	s.Message = &v
	return s
}

func (s *RedialCallResponseBody) SetParams(v []*string) *RedialCallResponseBody {
	s.Params = v
	return s
}

func (s *RedialCallResponseBody) SetRequestId(v string) *RedialCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedialCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type RedialCallResponseBodyData struct {
	CallContext *RedialCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// example:
	//
	// 123456789
	ContextId   *int64                                 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	UserContext *RedialCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s RedialCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RedialCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *RedialCallResponseBodyData) GetCallContext() *RedialCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *RedialCallResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *RedialCallResponseBodyData) GetUserContext() *RedialCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *RedialCallResponseBodyData) SetCallContext(v *RedialCallResponseBodyDataCallContext) *RedialCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *RedialCallResponseBodyData) SetContextId(v int64) *RedialCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *RedialCallResponseBodyData) SetUserContext(v *RedialCallResponseBodyDataUserContext) *RedialCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *RedialCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RedialCallResponseBodyDataCallContext struct {
	// example:
	//
	// OUTBOUND
	CallType        *string                                                 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelContexts []*RedialCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// abc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s RedialCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s RedialCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *RedialCallResponseBodyDataCallContext) GetCallType() *string {
	return s.CallType
}

func (s *RedialCallResponseBodyDataCallContext) GetChannelContexts() []*RedialCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *RedialCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RedialCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *RedialCallResponseBodyDataCallContext) SetCallType(v string) *RedialCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContext) SetChannelContexts(v []*RedialCallResponseBodyDataCallContextChannelContexts) *RedialCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *RedialCallResponseBodyDataCallContext) SetInstanceId(v string) *RedialCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContext) SetJobId(v string) *RedialCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContext) Validate() error {
	return dara.Validate(s)
}

type RedialCallResponseBodyDataCallContextChannelContexts struct {
	// example:
	//
	// OUTBOUND
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
	// CREATED
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
	// 139xxxx0501
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// example:
	//
	// 486:USER_BUSY
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// 1609138903315
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 8000****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// example:
	//
	// samzhang@abc
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RedialCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s RedialCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetChannelFlags() *string {
	return s.ChannelFlags
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *RedialCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *RedialCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type RedialCallResponseBodyDataUserContext struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 8000****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// abc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6573574060089****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// False
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// example:
	//
	// samzhang@abc
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	WorkMode  *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s RedialCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s RedialCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *RedialCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *RedialCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RedialCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *RedialCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RedialCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *RedialCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *RedialCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *RedialCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *RedialCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *RedialCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *RedialCallResponseBodyDataUserContext) SetBreakCode(v string) *RedialCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetDeviceId(v string) *RedialCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetExtension(v string) *RedialCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetInstanceId(v string) *RedialCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetJobId(v string) *RedialCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *RedialCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *RedialCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetUserId(v string) *RedialCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetUserState(v string) *RedialCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) SetWorkMode(v string) *RedialCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *RedialCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
