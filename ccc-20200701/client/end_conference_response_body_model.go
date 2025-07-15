// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndConferenceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EndConferenceResponseBody
  GetCode() *string 
  SetData(v *EndConferenceResponseBodyData) *EndConferenceResponseBody
  GetData() *EndConferenceResponseBodyData 
  SetHttpStatusCode(v int32) *EndConferenceResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EndConferenceResponseBody
  GetMessage() *string 
  SetParams(v []*string) *EndConferenceResponseBody
  GetParams() []*string 
  SetRequestId(v string) *EndConferenceResponseBody
  GetRequestId() *string 
}

type EndConferenceResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EndConferenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // 0630E5DF-CEB0-445B-8626-D5C7481181C3
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndConferenceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceResponseBody) GoString() string {
  return s.String()
}

func (s *EndConferenceResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EndConferenceResponseBody) GetData() *EndConferenceResponseBodyData  {
  return s.Data
}

func (s *EndConferenceResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EndConferenceResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EndConferenceResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *EndConferenceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndConferenceResponseBody) SetCode(v string) *EndConferenceResponseBody {
  s.Code = &v
  return s
}

func (s *EndConferenceResponseBody) SetData(v *EndConferenceResponseBodyData) *EndConferenceResponseBody {
  s.Data = v
  return s
}

func (s *EndConferenceResponseBody) SetHttpStatusCode(v int32) *EndConferenceResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EndConferenceResponseBody) SetMessage(v string) *EndConferenceResponseBody {
  s.Message = &v
  return s
}

func (s *EndConferenceResponseBody) SetParams(v []*string) *EndConferenceResponseBody {
  s.Params = v
  return s
}

func (s *EndConferenceResponseBody) SetRequestId(v string) *EndConferenceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndConferenceResponseBody) Validate() error {
  return dara.Validate(s)
}

type EndConferenceResponseBodyData struct {
  CallContext *EndConferenceResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
  // example:
  // 
  // 103655
  ContextId *int64 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
  UserContext *EndConferenceResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s EndConferenceResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceResponseBodyData) GoString() string {
  return s.String()
}

func (s *EndConferenceResponseBodyData) GetCallContext() *EndConferenceResponseBodyDataCallContext  {
  return s.CallContext
}

func (s *EndConferenceResponseBodyData) GetContextId() *int64  {
  return s.ContextId
}

func (s *EndConferenceResponseBodyData) GetUserContext() *EndConferenceResponseBodyDataUserContext  {
  return s.UserContext
}

func (s *EndConferenceResponseBodyData) SetCallContext(v *EndConferenceResponseBodyDataCallContext) *EndConferenceResponseBodyData {
  s.CallContext = v
  return s
}

func (s *EndConferenceResponseBodyData) SetContextId(v int64) *EndConferenceResponseBodyData {
  s.ContextId = &v
  return s
}

func (s *EndConferenceResponseBodyData) SetUserContext(v *EndConferenceResponseBodyDataUserContext) *EndConferenceResponseBodyData {
  s.UserContext = v
  return s
}

func (s *EndConferenceResponseBodyData) Validate() error {
  return dara.Validate(s)
}

type EndConferenceResponseBodyDataCallContext struct {
  ChannelContexts []*EndConferenceResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
  // example:
  // 
  // ccc-test
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // job-6538214103685****
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s EndConferenceResponseBodyDataCallContext) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceResponseBodyDataCallContext) GoString() string {
  return s.String()
}

func (s *EndConferenceResponseBodyDataCallContext) GetChannelContexts() []*EndConferenceResponseBodyDataCallContextChannelContexts  {
  return s.ChannelContexts
}

func (s *EndConferenceResponseBodyDataCallContext) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EndConferenceResponseBodyDataCallContext) GetJobId() *string  {
  return s.JobId
}

func (s *EndConferenceResponseBodyDataCallContext) SetChannelContexts(v []*EndConferenceResponseBodyDataCallContextChannelContexts) *EndConferenceResponseBodyDataCallContext {
  s.ChannelContexts = v
  return s
}

func (s *EndConferenceResponseBodyDataCallContext) SetInstanceId(v string) *EndConferenceResponseBodyDataCallContext {
  s.InstanceId = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContext) SetJobId(v string) *EndConferenceResponseBodyDataCallContext {
  s.JobId = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContext) Validate() error {
  return dara.Validate(s)
}

type EndConferenceResponseBodyDataCallContextChannelContexts struct {
  // example:
  // 
  // OUTBOUND
  CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
  // example:
  // 
  // ch:user:131888****->8001****:1609225718294:job-65700074013925376
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

func (s EndConferenceResponseBodyDataCallContextChannelContexts) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceResponseBodyDataCallContextChannelContexts) GoString() string {
  return s.String()
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetCallType() *string  {
  return s.CallType
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetChannelId() *string  {
  return s.ChannelId
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetChannelState() *string  {
  return s.ChannelState
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetDestination() *string  {
  return s.Destination
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetJobId() *string  {
  return s.JobId
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetOriginator() *string  {
  return s.Originator
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string  {
  return s.ReleaseInitiator
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string  {
  return s.ReleaseReason
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64  {
  return s.Timestamp
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetUserExtension() *string  {
  return s.UserExtension
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) GetUserId() *string  {
  return s.UserId
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetCallType(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.CallType = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.ChannelId = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.ChannelState = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetDestination(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.Destination = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetJobId(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.JobId = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.Originator = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.ReleaseInitiator = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.ReleaseReason = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.Timestamp = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.UserExtension = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) SetUserId(v string) *EndConferenceResponseBodyDataCallContextChannelContexts {
  s.UserId = &v
  return s
}

func (s *EndConferenceResponseBodyDataCallContextChannelContexts) Validate() error {
  return dara.Validate(s)
}

type EndConferenceResponseBodyDataUserContext struct {
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
  // ccc-test
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // job-6538214103685****
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // example:
  // 
  // false
  OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
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

func (s EndConferenceResponseBodyDataUserContext) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceResponseBodyDataUserContext) GoString() string {
  return s.String()
}

func (s *EndConferenceResponseBodyDataUserContext) GetBreakCode() *string  {
  return s.BreakCode
}

func (s *EndConferenceResponseBodyDataUserContext) GetDeviceId() *string  {
  return s.DeviceId
}

func (s *EndConferenceResponseBodyDataUserContext) GetDeviceState() *string  {
  return s.DeviceState
}

func (s *EndConferenceResponseBodyDataUserContext) GetExtension() *string  {
  return s.Extension
}

func (s *EndConferenceResponseBodyDataUserContext) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EndConferenceResponseBodyDataUserContext) GetJobId() *string  {
  return s.JobId
}

func (s *EndConferenceResponseBodyDataUserContext) GetOutboundScenario() *bool  {
  return s.OutboundScenario
}

func (s *EndConferenceResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string  {
  return s.SignedSkillGroupIdList
}

func (s *EndConferenceResponseBodyDataUserContext) GetUserId() *string  {
  return s.UserId
}

func (s *EndConferenceResponseBodyDataUserContext) GetUserState() *string  {
  return s.UserState
}

func (s *EndConferenceResponseBodyDataUserContext) GetWorkMode() *string  {
  return s.WorkMode
}

func (s *EndConferenceResponseBodyDataUserContext) SetBreakCode(v string) *EndConferenceResponseBodyDataUserContext {
  s.BreakCode = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetDeviceId(v string) *EndConferenceResponseBodyDataUserContext {
  s.DeviceId = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetDeviceState(v string) *EndConferenceResponseBodyDataUserContext {
  s.DeviceState = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetExtension(v string) *EndConferenceResponseBodyDataUserContext {
  s.Extension = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetInstanceId(v string) *EndConferenceResponseBodyDataUserContext {
  s.InstanceId = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetJobId(v string) *EndConferenceResponseBodyDataUserContext {
  s.JobId = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetOutboundScenario(v bool) *EndConferenceResponseBodyDataUserContext {
  s.OutboundScenario = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *EndConferenceResponseBodyDataUserContext {
  s.SignedSkillGroupIdList = v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetUserId(v string) *EndConferenceResponseBodyDataUserContext {
  s.UserId = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetUserState(v string) *EndConferenceResponseBodyDataUserContext {
  s.UserState = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) SetWorkMode(v string) *EndConferenceResponseBodyDataUserContext {
  s.WorkMode = &v
  return s
}

func (s *EndConferenceResponseBodyDataUserContext) Validate() error {
  return dara.Validate(s)
}

