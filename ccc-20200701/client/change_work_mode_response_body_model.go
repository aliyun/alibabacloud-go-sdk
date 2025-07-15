// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeWorkModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeWorkModeResponseBody
	GetCode() *string
	SetData(v *ChangeWorkModeResponseBodyData) *ChangeWorkModeResponseBody
	GetData() *ChangeWorkModeResponseBodyData
	SetHttpStatusCode(v int32) *ChangeWorkModeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ChangeWorkModeResponseBody
	GetMessage() *string
	SetParams(v []*string) *ChangeWorkModeResponseBody
	GetParams() []*string
	SetRequestId(v string) *ChangeWorkModeResponseBody
	GetRequestId() *string
}

type ChangeWorkModeResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ChangeWorkModeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 87731ED1-6224-48A5-99E3-6237FF9B1C00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeWorkModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeWorkModeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeWorkModeResponseBody) GetData() *ChangeWorkModeResponseBodyData {
	return s.Data
}

func (s *ChangeWorkModeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChangeWorkModeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeWorkModeResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ChangeWorkModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeWorkModeResponseBody) SetCode(v string) *ChangeWorkModeResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetData(v *ChangeWorkModeResponseBodyData) *ChangeWorkModeResponseBody {
	s.Data = v
	return s
}

func (s *ChangeWorkModeResponseBody) SetHttpStatusCode(v int32) *ChangeWorkModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetMessage(v string) *ChangeWorkModeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetParams(v []*string) *ChangeWorkModeResponseBody {
	s.Params = v
	return s
}

func (s *ChangeWorkModeResponseBody) SetRequestId(v string) *ChangeWorkModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeWorkModeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeWorkModeResponseBodyData struct {
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
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// false
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// OFFLINE
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ChangeWorkModeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeWorkModeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ChangeWorkModeResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ChangeWorkModeResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ChangeWorkModeResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeWorkModeResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ChangeWorkModeResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ChangeWorkModeResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ChangeWorkModeResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ChangeWorkModeResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *ChangeWorkModeResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ChangeWorkModeResponseBodyData) SetBreakCode(v string) *ChangeWorkModeResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetDeviceId(v string) *ChangeWorkModeResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetExtension(v string) *ChangeWorkModeResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetInstanceId(v string) *ChangeWorkModeResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetJobId(v string) *ChangeWorkModeResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetOutboundScenario(v bool) *ChangeWorkModeResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ChangeWorkModeResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetUserId(v string) *ChangeWorkModeResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetUserState(v string) *ChangeWorkModeResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetWorkMode(v string) *ChangeWorkModeResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
