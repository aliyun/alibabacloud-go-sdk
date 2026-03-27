// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ClaimCallResponseBody
	GetCode() *string
	SetData(v *ClaimCallResponseBodyData) *ClaimCallResponseBody
	GetData() *ClaimCallResponseBodyData
	SetHttpStatusCode(v int32) *ClaimCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ClaimCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *ClaimCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *ClaimCallResponseBody
	GetRequestId() *string
}

type ClaimCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClaimCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 93CDC17E-3E8A-48F2-99E5-FA2E238DE8B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClaimCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallResponseBody) GoString() string {
	return s.String()
}

func (s *ClaimCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClaimCallResponseBody) GetData() *ClaimCallResponseBodyData {
	return s.Data
}

func (s *ClaimCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ClaimCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClaimCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ClaimCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClaimCallResponseBody) SetCode(v string) *ClaimCallResponseBody {
	s.Code = &v
	return s
}

func (s *ClaimCallResponseBody) SetData(v *ClaimCallResponseBodyData) *ClaimCallResponseBody {
	s.Data = v
	return s
}

func (s *ClaimCallResponseBody) SetHttpStatusCode(v int32) *ClaimCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ClaimCallResponseBody) SetMessage(v string) *ClaimCallResponseBody {
	s.Message = &v
	return s
}

func (s *ClaimCallResponseBody) SetParams(v []*string) *ClaimCallResponseBody {
	s.Params = v
	return s
}

func (s *ClaimCallResponseBody) SetRequestId(v string) *ClaimCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClaimCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClaimCallResponseBodyData struct {
	CallContext *ClaimCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	ContextId   *int64                                `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	UserContext *ClaimCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s ClaimCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClaimCallResponseBodyData) GetCallContext() *ClaimCallResponseBodyDataCallContext {
	return s.CallContext
}

func (s *ClaimCallResponseBodyData) GetContextId() *int64 {
	return s.ContextId
}

func (s *ClaimCallResponseBodyData) GetUserContext() *ClaimCallResponseBodyDataUserContext {
	return s.UserContext
}

func (s *ClaimCallResponseBodyData) SetCallContext(v *ClaimCallResponseBodyDataCallContext) *ClaimCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *ClaimCallResponseBodyData) SetContextId(v int64) *ClaimCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *ClaimCallResponseBodyData) SetUserContext(v *ClaimCallResponseBodyDataUserContext) *ClaimCallResponseBodyData {
	s.UserContext = v
	return s
}

func (s *ClaimCallResponseBodyData) Validate() error {
	if s.CallContext != nil {
		if err := s.CallContext.Validate(); err != nil {
			return err
		}
	}
	if s.UserContext != nil {
		if err := s.UserContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClaimCallResponseBodyDataCallContext struct {
	ChannelContexts []*ClaimCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ClaimCallResponseBodyDataCallContext) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *ClaimCallResponseBodyDataCallContext) GetChannelContexts() []*ClaimCallResponseBodyDataCallContextChannelContexts {
	return s.ChannelContexts
}

func (s *ClaimCallResponseBodyDataCallContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClaimCallResponseBodyDataCallContext) GetJobId() *string {
	return s.JobId
}

func (s *ClaimCallResponseBodyDataCallContext) SetChannelContexts(v []*ClaimCallResponseBodyDataCallContextChannelContexts) *ClaimCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

func (s *ClaimCallResponseBodyDataCallContext) SetInstanceId(v string) *ClaimCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContext) SetJobId(v string) *ClaimCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContext) Validate() error {
	if s.ChannelContexts != nil {
		for _, item := range s.ChannelContexts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClaimCallResponseBodyDataCallContextChannelContexts struct {
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
	// ANSWERED
	ChannelState *string `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	// example:
	//
	// a=b;c=d;
	ChannelVariables *string `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	// example:
	//
	// 1390501****
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// job-6573574060089****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0830019****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// example:
	//
	// 1390501****
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ReleaseReason    *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// 1609225718295
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 8059****
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// example:
	//
	// invoker@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ClaimCallResponseBodyDataCallContextChannelContexts) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetCallType() *string {
	return s.CallType
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetChannelId() *string {
	return s.ChannelId
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetChannelState() *string {
	return s.ChannelState
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetChannelVariables() *string {
	return s.ChannelVariables
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetDestination() *string {
	return s.Destination
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetJobId() *string {
	return s.JobId
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetOriginator() *string {
	return s.Originator
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetUserExtension() *string {
	return s.UserExtension
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) GetUserId() *string {
	return s.UserId
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetChannelVariables(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.ChannelVariables = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *ClaimCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *ClaimCallResponseBodyDataCallContextChannelContexts) Validate() error {
	return dara.Validate(s)
}

type ClaimCallResponseBodyDataUserContext struct {
	// example:
	//
	// Customized
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// example:
	//
	// device-xxxx
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 0830019****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6573574060089****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// false
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// example:
	//
	// user@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// Dialing
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ClaimCallResponseBodyDataUserContext) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *ClaimCallResponseBodyDataUserContext) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ClaimCallResponseBodyDataUserContext) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ClaimCallResponseBodyDataUserContext) GetExtension() *string {
	return s.Extension
}

func (s *ClaimCallResponseBodyDataUserContext) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClaimCallResponseBodyDataUserContext) GetJobId() *string {
	return s.JobId
}

func (s *ClaimCallResponseBodyDataUserContext) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ClaimCallResponseBodyDataUserContext) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ClaimCallResponseBodyDataUserContext) GetUserId() *string {
	return s.UserId
}

func (s *ClaimCallResponseBodyDataUserContext) GetUserState() *string {
	return s.UserState
}

func (s *ClaimCallResponseBodyDataUserContext) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ClaimCallResponseBodyDataUserContext) SetBreakCode(v string) *ClaimCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetDeviceId(v string) *ClaimCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetExtension(v string) *ClaimCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetInstanceId(v string) *ClaimCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetJobId(v string) *ClaimCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *ClaimCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *ClaimCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetUserId(v string) *ClaimCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetUserState(v string) *ClaimCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) SetWorkMode(v string) *ClaimCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *ClaimCallResponseBodyDataUserContext) Validate() error {
	return dara.Validate(s)
}
