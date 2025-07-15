// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTakeBreakResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TakeBreakResponseBody
	GetCode() *string
	SetData(v *TakeBreakResponseBodyData) *TakeBreakResponseBody
	GetData() *TakeBreakResponseBodyData
	SetHttpStatusCode(v int32) *TakeBreakResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TakeBreakResponseBody
	GetMessage() *string
	SetParams(v []*string) *TakeBreakResponseBody
	GetParams() []*string
	SetRequestId(v string) *TakeBreakResponseBody
	GetRequestId() *string
}

type TakeBreakResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TakeBreakResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// B59382D2-5755-4C6D-861F-FA2AAD8F89F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TakeBreakResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TakeBreakResponseBody) GoString() string {
	return s.String()
}

func (s *TakeBreakResponseBody) GetCode() *string {
	return s.Code
}

func (s *TakeBreakResponseBody) GetData() *TakeBreakResponseBodyData {
	return s.Data
}

func (s *TakeBreakResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TakeBreakResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TakeBreakResponseBody) GetParams() []*string {
	return s.Params
}

func (s *TakeBreakResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TakeBreakResponseBody) SetCode(v string) *TakeBreakResponseBody {
	s.Code = &v
	return s
}

func (s *TakeBreakResponseBody) SetData(v *TakeBreakResponseBodyData) *TakeBreakResponseBody {
	s.Data = v
	return s
}

func (s *TakeBreakResponseBody) SetHttpStatusCode(v int32) *TakeBreakResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TakeBreakResponseBody) SetMessage(v string) *TakeBreakResponseBody {
	s.Message = &v
	return s
}

func (s *TakeBreakResponseBody) SetParams(v []*string) *TakeBreakResponseBody {
	s.Params = v
	return s
}

func (s *TakeBreakResponseBody) SetRequestId(v string) *TakeBreakResponseBody {
	s.RequestId = &v
	return s
}

func (s *TakeBreakResponseBody) Validate() error {
	return dara.Validate(s)
}

type TakeBreakResponseBodyData struct {
	// example:
	//
	// lunchtime
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 1609249563836
	Heartbeat *int64 `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// 1609234221864
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

func (s TakeBreakResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TakeBreakResponseBodyData) GoString() string {
	return s.String()
}

func (s *TakeBreakResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *TakeBreakResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *TakeBreakResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *TakeBreakResponseBodyData) GetHeartbeat() *int64 {
	return s.Heartbeat
}

func (s *TakeBreakResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TakeBreakResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *TakeBreakResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *TakeBreakResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *TakeBreakResponseBodyData) GetReserved() *int64 {
	return s.Reserved
}

func (s *TakeBreakResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *TakeBreakResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *TakeBreakResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *TakeBreakResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *TakeBreakResponseBodyData) SetBreakCode(v string) *TakeBreakResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetDeviceId(v string) *TakeBreakResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetExtension(v string) *TakeBreakResponseBodyData {
	s.Extension = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetHeartbeat(v int64) *TakeBreakResponseBodyData {
	s.Heartbeat = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetInstanceId(v string) *TakeBreakResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetJobId(v string) *TakeBreakResponseBodyData {
	s.JobId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetMobile(v string) *TakeBreakResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetOutboundScenario(v bool) *TakeBreakResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetReserved(v int64) *TakeBreakResponseBodyData {
	s.Reserved = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetSignedSkillGroupIdList(v []*string) *TakeBreakResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *TakeBreakResponseBodyData) SetUserId(v string) *TakeBreakResponseBodyData {
	s.UserId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetUserState(v string) *TakeBreakResponseBodyData {
	s.UserState = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetWorkMode(v string) *TakeBreakResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *TakeBreakResponseBodyData) Validate() error {
	return dara.Validate(s)
}
