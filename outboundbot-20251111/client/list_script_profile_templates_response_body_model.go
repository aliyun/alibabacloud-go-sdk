// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptProfileTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptProfileTemplatesResponseBody
	GetCode() *string
	SetData(v []*ListScriptProfileTemplatesResponseBodyData) *ListScriptProfileTemplatesResponseBody
	GetData() []*ListScriptProfileTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListScriptProfileTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptProfileTemplatesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListScriptProfileTemplatesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListScriptProfileTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScriptProfileTemplatesResponseBody
	GetSuccess() *bool
}

type ListScriptProfileTemplatesResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*ListScriptProfileTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// Instance does not exist. Instance=ob-9876543210.
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

func (s ListScriptProfileTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptProfileTemplatesResponseBody) GetData() []*ListScriptProfileTemplatesResponseBodyData {
	return s.Data
}

func (s *ListScriptProfileTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptProfileTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptProfileTemplatesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListScriptProfileTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptProfileTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptProfileTemplatesResponseBody) SetCode(v string) *ListScriptProfileTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetData(v []*ListScriptProfileTemplatesResponseBodyData) *ListScriptProfileTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetHttpStatusCode(v int32) *ListScriptProfileTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetMessage(v string) *ListScriptProfileTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetParams(v []*string) *ListScriptProfileTemplatesResponseBody {
	s.Params = v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetRequestId(v string) *ListScriptProfileTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetSuccess(v bool) *ListScriptProfileTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptProfileTemplatesResponseBodyData struct {
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
	// [{\\"name\\":\\"Overall Satisfaction\\",\\"description\\":\\"Collect evaluation information on overall service satisfaction\\",\\"candidateValues\\":[\\"Very Satisfied\\",\\"Satisfied\\",\\"Average\\",\\"Dissatisfied\\",\\"Very Dissatisfied\\"]},{\\"name\\":\\"Service Highlights\\",\\"description\\":\\"Customer description of service highlights or satisfactory aspects\\",\\"candidateValues\\":[]},{\\"name\\":\\"Improvement Suggestions\\",\\"description\\":\\"Customer description of improvement suggestions\\",\\"candidateValues\\":[]},{\\"name\\":\\"Service Efficiency\\",\\"description\\":\\"Customer feedback on service response speed and timeliness of service completion\\",\\"candidateValues\\":[]},{\\"name\\":\\"Employee Attitude\\",\\"description\\":\\"Customer evaluation of service personnel professionalism and attitude\\",\\"candidateValues\\":[]},{\\"name\\":\\"Willingness to Choose Again\\",\\"description\\":\\"Whether the customer is willing to choose again\\",\\"candidateValues\\":[\\"Yes\\",\\"No\\"]}]
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
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
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

func (s ListScriptProfileTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetVariables() *string {
	return s.Variables
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetCreatedTime(v int64) *ListScriptProfileTemplatesResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetDescription(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetLabels(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Labels = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetName(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetSchema(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Schema = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetTemplateId(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetUpdatedTime(v int64) *ListScriptProfileTemplatesResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetVariables(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Variables = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
