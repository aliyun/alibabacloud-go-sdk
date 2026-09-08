// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadyForServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadyForServiceResponseBody
	GetCode() *string
	SetData(v *ReadyForServiceResponseBodyData) *ReadyForServiceResponseBody
	GetData() *ReadyForServiceResponseBodyData
	SetHttpStatusCode(v int32) *ReadyForServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReadyForServiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *ReadyForServiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *ReadyForServiceResponseBody
	GetRequestId() *string
}

type ReadyForServiceResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ReadyForServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of response parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// CC49060B-87ED-489A-AD3D-00E57775DBFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReadyForServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadyForServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ReadyForServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadyForServiceResponseBody) GetData() *ReadyForServiceResponseBodyData {
	return s.Data
}

func (s *ReadyForServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReadyForServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadyForServiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ReadyForServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadyForServiceResponseBody) SetCode(v string) *ReadyForServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetData(v *ReadyForServiceResponseBodyData) *ReadyForServiceResponseBody {
	s.Data = v
	return s
}

func (s *ReadyForServiceResponseBody) SetHttpStatusCode(v int32) *ReadyForServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetMessage(v string) *ReadyForServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetParams(v []*string) *ReadyForServiceResponseBody {
	s.Params = v
	return s
}

func (s *ReadyForServiceResponseBody) SetRequestId(v string) *ReadyForServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadyForServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadyForServiceResponseBodyData struct {
	// Break status code, which can be either System-defined or Custom-defined. System-defined break codes include: Warm-up (temporary break state after the agent is published and before becoming idle), RingingTimeout (break caused by agent ringing timeout), and RejectCall (break caused by agent rejecting a call). There are no restrictions on Custom-defined status codes; customers can define them according to their business needs.
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Device ID, which is the identity ID of a browser-based Web Real-Time Communication (WebRTC) softphone or a physical phone device. Only one type of device can be registered at a time.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Agent extension number.
	//
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the agent is in outbound-only mode.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// List of skill group IDs that the agent has signed into.
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Agent status.
	//
	// example:
	//
	// READY
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ReadyForServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadyForServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadyForServiceResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ReadyForServiceResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ReadyForServiceResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ReadyForServiceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReadyForServiceResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ReadyForServiceResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ReadyForServiceResponseBodyData) GetSignedSkillGroupIdList() []*string {
	return s.SignedSkillGroupIdList
}

func (s *ReadyForServiceResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ReadyForServiceResponseBodyData) GetUserState() *string {
	return s.UserState
}

func (s *ReadyForServiceResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ReadyForServiceResponseBodyData) SetBreakCode(v string) *ReadyForServiceResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetDeviceId(v string) *ReadyForServiceResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetExtension(v string) *ReadyForServiceResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetInstanceId(v string) *ReadyForServiceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetJobId(v string) *ReadyForServiceResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetOutboundScenario(v bool) *ReadyForServiceResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ReadyForServiceResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetUserId(v string) *ReadyForServiceResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetUserState(v string) *ReadyForServiceResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetWorkMode(v string) *ReadyForServiceResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
