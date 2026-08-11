// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceResponseBody
	GetCode() *string
	SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody
	GetData() *GetInstanceResponseBodyData
	SetHttpStatusCode(v int32) *GetInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetInstanceResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
}

type GetInstanceResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance details.
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not permitted to operate this instance. User=206770505484719639, Instance=null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceResponseBody) GetData() *GetInstanceResponseBodyData {
	return s.Data
}

func (s *GetInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetParams(v []*string) *GetInstanceResponseBody {
	s.Params = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyData struct {
	// The number of concurrent connections.
	//
	// example:
	//
	// 50
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1769653616000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// This is a large language model chatbot
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// e49ad122-15a1-443a-a102-84a78a93be79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Questionnaire
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Xiaomi business space information.
	//
	// example:
	//
	// {"agentId":"1380415","agentInstanceId":"outbound_d1d1a8e6-a14e-46c8-b580-e50d94cb2d7e","agentKey":"329cb6fdb880431d82400a8365380100_p_outbound_public"}
	NluProfile *string `json:"NluProfile,omitempty" xml:"NluProfile,omitempty"`
	// The service mode.
	//
	// example:
	//
	// STANDARD
	ServiceMode *string `json:"ServiceMode,omitempty" xml:"ServiceMode,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1308144684576765
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 1769653616000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetInstanceResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetInstanceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetInstanceResponseBodyData) GetNluProfile() *string {
	return s.NluProfile
}

func (s *GetInstanceResponseBodyData) GetServiceMode() *string {
	return s.ServiceMode
}

func (s *GetInstanceResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetInstanceResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetInstanceResponseBodyData) SetConcurrency(v int32) *GetInstanceResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCreatedTime(v int64) *GetInstanceResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDescription(v string) *GetInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetName(v string) *GetInstanceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetNluProfile(v string) *GetInstanceResponseBodyData {
	s.NluProfile = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetServiceMode(v string) *GetInstanceResponseBodyData {
	s.ServiceMode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTenantId(v string) *GetInstanceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUpdatedTime(v int64) *GetInstanceResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
