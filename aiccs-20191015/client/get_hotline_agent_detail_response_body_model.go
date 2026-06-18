// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineAgentDetailResponseBody
	GetCode() *string
	SetData(v *GetHotlineAgentDetailResponseBodyData) *GetHotlineAgentDetailResponseBody
	GetData() *GetHotlineAgentDetailResponseBodyData
	SetHttpStatusCode(v int64) *GetHotlineAgentDetailResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetHotlineAgentDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineAgentDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotlineAgentDetailResponseBody
	GetSuccess() *bool
}

type GetHotlineAgentDetailResponseBody struct {
	// Status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Agent service data.
	Data *GetHotlineAgentDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineAgentDetailResponseBody) GetData() *GetHotlineAgentDetailResponseBodyData {
	return s.Data
}

func (s *GetHotlineAgentDetailResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetHotlineAgentDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineAgentDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineAgentDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotlineAgentDetailResponseBody) SetCode(v string) *GetHotlineAgentDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetData(v *GetHotlineAgentDetailResponseBodyData) *GetHotlineAgentDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetMessage(v string) *GetHotlineAgentDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetRequestId(v string) *GetHotlineAgentDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetSuccess(v bool) *GetHotlineAgentDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotlineAgentDetailResponseBodyData struct {
	// Agent ID.
	//
	// example:
	//
	// 2235****
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Agent status. Valid values: **1~6**.
	//
	// example:
	//
	// 1
	AgentStatus *int32 `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// Agent status code. Valid values:
	//
	// - **AgentCheckout**: Agent logged off.
	//
	// - **AgentReady**: Agent idle.
	//
	// - **AgentBreak**: Agent on break.
	//
	// - **AgentAcw**: Post-processing after a call.
	//
	// - **AgentBusyForCall**: In a call.
	//
	// example:
	//
	// AgentCheckout
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	// Indicates whether the agent is assigned. Valid values:
	//
	// **false**: Not assigned (no call).
	//
	// **true**: Assigned (in a call).
	//
	// example:
	//
	// false
	Assigned *bool `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// Break type. Valid values:
	//
	// - **1**: Short break.
	//
	// - **2**: Meal break.
	//
	// - **3**: Meeting.
	//
	// - **4**: Coaching.
	//
	// - **5**: Training.
	//
	// example:
	//
	// 1
	RestType *int32 `json:"RestType,omitempty" xml:"RestType,omitempty"`
	// Tenant ID to which the agent belongs, corresponding to the instance ID in the input parameter.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Heartbeat signature.
	//
	// example:
	//
	// dnthF_oinHg7JMJDmKqex3Ux****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetHotlineAgentDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBodyData) GetAgentId() *int64 {
	return s.AgentId
}

func (s *GetHotlineAgentDetailResponseBodyData) GetAgentStatus() *int32 {
	return s.AgentStatus
}

func (s *GetHotlineAgentDetailResponseBodyData) GetAgentStatusCode() *string {
	return s.AgentStatusCode
}

func (s *GetHotlineAgentDetailResponseBodyData) GetAssigned() *bool {
	return s.Assigned
}

func (s *GetHotlineAgentDetailResponseBodyData) GetRestType() *int32 {
	return s.RestType
}

func (s *GetHotlineAgentDetailResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetHotlineAgentDetailResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatus(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatusCode(v string) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAssigned(v bool) *GetHotlineAgentDetailResponseBodyData {
	s.Assigned = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetRestType(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.RestType = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetTenantId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetToken(v string) *GetHotlineAgentDetailResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
