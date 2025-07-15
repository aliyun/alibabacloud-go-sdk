// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ClaimChatResponseBody
	GetCode() *string
	SetData(v *ClaimChatResponseBodyData) *ClaimChatResponseBody
	GetData() *ClaimChatResponseBodyData
	SetHttpStatusCode(v int32) *ClaimChatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ClaimChatResponseBody
	GetMessage() *string
	SetParams(v []*string) *ClaimChatResponseBody
	GetParams() []*string
	SetRequestId(v string) *ClaimChatResponseBody
	GetRequestId() *string
}

type ClaimChatResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClaimChatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClaimChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClaimChatResponseBody) GoString() string {
	return s.String()
}

func (s *ClaimChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClaimChatResponseBody) GetData() *ClaimChatResponseBodyData {
	return s.Data
}

func (s *ClaimChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ClaimChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClaimChatResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ClaimChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClaimChatResponseBody) SetCode(v string) *ClaimChatResponseBody {
	s.Code = &v
	return s
}

func (s *ClaimChatResponseBody) SetData(v *ClaimChatResponseBodyData) *ClaimChatResponseBody {
	s.Data = v
	return s
}

func (s *ClaimChatResponseBody) SetHttpStatusCode(v int32) *ClaimChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ClaimChatResponseBody) SetMessage(v string) *ClaimChatResponseBody {
	s.Message = &v
	return s
}

func (s *ClaimChatResponseBody) SetParams(v []*string) *ClaimChatResponseBody {
	s.Params = v
	return s
}

func (s *ClaimChatResponseBody) SetRequestId(v string) *ClaimChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClaimChatResponseBody) Validate() error {
	return dara.Validate(s)
}

type ClaimChatResponseBodyData struct {
	ChatContexts []*ClaimChatResponseBodyDataChatContexts `json:"ChatContexts,omitempty" xml:"ChatContexts,omitempty" type:"Repeated"`
	// example:
	//
	// 123456789
	ContextId   *int64                                `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	UserContext *ClaimChatResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s ClaimChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClaimChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClaimChatResponseBodyData) GetChatContexts() []*ClaimChatResponseBodyDataChatContexts {
	return s.ChatContexts
}

func (s *ClaimChatResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *ClaimChatResponseBodyData) GetUserContext() *ClaimChatResponseBodyDataUserContext {
	return s.UserContext
}

func (s *ClaimChatResponseBodyData) SetChatContexts(v []*ClaimChatResponseBodyDataChatContexts) *ClaimChatResponseBodyData {
	s.ChatContexts = v
	return s
}

func (s *ClaimChatResponseBodyData) SetContextId(v int64) *ClaimChatResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *ClaimChatResponseBodyData) SetUserContext(v *ClaimChatResponseBodyDataUserContext) *ClaimChatResponseBodyData {
	s.UserContext = v
	return s
}

func (s *ClaimChatResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ClaimChatResponseBodyDataChatContexts struct {
	// example:
	//
	// 226****-cbb6-****-8fea-1e71baf7bfa7
	AccessChannelId   *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	AccessChannelName *string `json:"AccessChannelName,omitempty" xml:"AccessChannelName,omitempty"`
	// example:
	//
	// Web
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// example:
	//
	// true
	BeingAssigned *bool   `json:"BeingAssigned,omitempty" xml:"BeingAssigned,omitempty"`
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// example:
	//
	// INBOUND
	ChatType *string `json:"ChatType,omitempty" xml:"ChatType,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ClaimChatResponseBodyDataChatContexts) String() string {
	return dara.Prettify(s)
}

func (s ClaimChatResponseBodyDataChatContexts) GoString() string {
	return s.String()
}

func (s *ClaimChatResponseBodyDataChatContexts) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ClaimChatResponseBodyDataChatContexts) GetAccessChannelName() *string {
	return s.AccessChannelName
}

func (s *ClaimChatResponseBodyDataChatContexts) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ClaimChatResponseBodyDataChatContexts) GetBeingAssigned() *bool {
	return s.BeingAssigned
}

func (s *ClaimChatResponseBodyDataChatContexts) GetCallVariables() *string {
	return s.CallVariables
}

func (s *ClaimChatResponseBodyDataChatContexts) GetChatType() *string {
	return s.ChatType
}

func (s *ClaimChatResponseBodyDataChatContexts) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClaimChatResponseBodyDataChatContexts) GetJobId() *string {
	return s.JobId
}

func (s *ClaimChatResponseBodyDataChatContexts) SetAccessChannelId(v string) *ClaimChatResponseBodyDataChatContexts {
	s.AccessChannelId = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetAccessChannelName(v string) *ClaimChatResponseBodyDataChatContexts {
	s.AccessChannelName = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetAccessChannelType(v string) *ClaimChatResponseBodyDataChatContexts {
	s.AccessChannelType = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetBeingAssigned(v bool) *ClaimChatResponseBodyDataChatContexts {
	s.BeingAssigned = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetCallVariables(v string) *ClaimChatResponseBodyDataChatContexts {
	s.CallVariables = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetChatType(v string) *ClaimChatResponseBodyDataChatContexts {
	s.ChatType = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetInstanceId(v string) *ClaimChatResponseBodyDataChatContexts {
	s.InstanceId = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) SetJobId(v string) *ClaimChatResponseBodyDataChatContexts {
	s.JobId = &v
	return s
}

func (s *ClaimChatResponseBodyDataChatContexts) Validate() error {
	return dara.Validate(s)
}

type ClaimChatResponseBodyDataUserContext struct {
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// example:
	//
	// CCC-169.254.165.2-browser125.0.0-bs48b41903450e6c8
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// ONLINE
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
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 18******102
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// false
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// example:
	//
	// userId@ccc-test
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

func (s ClaimChatResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s ClaimChatResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *ClaimChatResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ClaimChatResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ClaimChatResponseBodyDataUserContext) GetDeviceState() *string {
	return s.DeviceState
}

func (s *ClaimChatResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *ClaimChatResponseBodyDataUserContext) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *ClaimChatResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClaimChatResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *ClaimChatResponseBodyDataUserContext) GetMobile() *string {
	return s.Mobile
}

func (s *ClaimChatResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ClaimChatResponseBodyDataUserContext) GetReserved() *int64 {
	return s.Reserved
}

func (s *ClaimChatResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ClaimChatResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *ClaimChatResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *ClaimChatResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ClaimChatResponseBodyDataUserContext) SetBreakCode(v string) *ClaimChatResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetDeviceId(v string) *ClaimChatResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetDeviceState(v string) *ClaimChatResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetExtension(v string) *ClaimChatResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetHeartbeat(v int64) *ClaimChatResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetInstanceId(v string) *ClaimChatResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetJobId(v string) *ClaimChatResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetMobile(v string) *ClaimChatResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetOutboundScenario(v bool) *ClaimChatResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetReserved(v int64) *ClaimChatResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *ClaimChatResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetUserId(v string) *ClaimChatResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetUserState(v string) *ClaimChatResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) SetWorkMode(v string) *ClaimChatResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *ClaimChatResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
