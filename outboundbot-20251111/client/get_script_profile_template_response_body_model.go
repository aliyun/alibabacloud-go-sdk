// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptProfileTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScriptProfileTemplateResponseBody
	GetCode() *string
	SetData(v *GetScriptProfileTemplateResponseBodyData) *GetScriptProfileTemplateResponseBody
	GetData() *GetScriptProfileTemplateResponseBodyData
	SetHttpStatusCode(v int32) *GetScriptProfileTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetScriptProfileTemplateResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetScriptProfileTemplateResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetScriptProfileTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScriptProfileTemplateResponseBody
	GetSuccess() *bool
}

type GetScriptProfileTemplateResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetScriptProfileTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Instance does not exist. Instance=392db13c-8901-4a25-b566-91d0d8114cec
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScriptProfileTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScriptProfileTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptProfileTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScriptProfileTemplateResponseBody) GetData() *GetScriptProfileTemplateResponseBodyData {
	return s.Data
}

func (s *GetScriptProfileTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetScriptProfileTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScriptProfileTemplateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetScriptProfileTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScriptProfileTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScriptProfileTemplateResponseBody) SetCode(v string) *GetScriptProfileTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetData(v *GetScriptProfileTemplateResponseBodyData) *GetScriptProfileTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetHttpStatusCode(v int32) *GetScriptProfileTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetMessage(v string) *GetScriptProfileTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetParams(v []*string) *GetScriptProfileTemplateResponseBody {
	s.Params = v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetRequestId(v string) *GetScriptProfileTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetSuccess(v bool) *GetScriptProfileTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptProfileTemplateResponseBodyData struct {
	// The creation time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description.
	//
	// example:
	//
	// As a survey specialist, sequentially ask about overall service satisfaction, service highlights, improvement suggestions, service efficiency, employee attitude, and willingness to choose again, and collect information
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label definition.
	//
	// example:
	//
	// [{\\"name\\":\\"Overall Satisfaction\\",\\"description\\":\\"Collect evaluation information on overall satisfaction with the service\\",\\"candidateValues\\":[\\"Very Satisfied\\",\\"Satisfied\\",\\"Average\\",\\"Dissatisfied\\",\\"Very Dissatisfied\\"]},{\\"name\\":\\"Service Highlights\\",\\"description\\":\\"Customer description of service highlights or satisfactory aspects\\",\\"candidateValues\\":[]},{\\"name\\":\\"Improvement Suggestions\\",\\"description\\":\\"Customer description of improvement suggestions\\",\\"candidateValues\\":[]},{\\"name\\":\\"Service Efficiency\\",\\"description\\":\\"Customer feedback on service response speed and timeliness of service completion\\",\\"candidateValues\\":[]},{\\"name\\":\\"Employee Attitude\\",\\"description\\":\\"Customer evaluation of the professionalism and attitude of service personnel\\",\\"candidateValues\\":[]},{\\"name\\":\\"Willingness to Choose Again\\",\\"description\\":\\"Whether the customer is willing to choose again\\",\\"candidateValues\\":[\\"Yes\\",\\"No\\"]}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The name.
	//
	// example:
	//
	// Service Satisfaction Survey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template details.
	//
	// example:
	//
	// {\\"name\\":\\"Li Ming\\",\\"gender\\":\\"Male\\"}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b59
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The update time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The variable definition.
	//
	// example:
	//
	// [{\\"name\\":\\"name\\",\\"description\\":\\"Customer name\\"},{\\"name\\":\\"gender\\",\\"description\\":\\"Customer gender\\"}]
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s GetScriptProfileTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScriptProfileTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScriptProfileTemplateResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetScriptProfileTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetScriptProfileTemplateResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *GetScriptProfileTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetScriptProfileTemplateResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetScriptProfileTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetScriptProfileTemplateResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetScriptProfileTemplateResponseBodyData) GetVariables() *string {
	return s.Variables
}

func (s *GetScriptProfileTemplateResponseBodyData) SetCreatedTime(v int64) *GetScriptProfileTemplateResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetDescription(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetLabels(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Labels = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetName(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetSchema(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetTemplateId(v string) *GetScriptProfileTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetUpdatedTime(v int64) *GetScriptProfileTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetVariables(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Variables = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
