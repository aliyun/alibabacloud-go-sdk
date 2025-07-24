// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionWithScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateProcessDefinitionWithScheduleResponseBody
	GetCode() *int32
	SetData(v *CreateProcessDefinitionWithScheduleResponseBodyData) *CreateProcessDefinitionWithScheduleResponseBody
	GetData() *CreateProcessDefinitionWithScheduleResponseBodyData
	SetFailed(v string) *CreateProcessDefinitionWithScheduleResponseBody
	GetFailed() *string
	SetHttpStatusCode(v int32) *CreateProcessDefinitionWithScheduleResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *CreateProcessDefinitionWithScheduleResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateProcessDefinitionWithScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateProcessDefinitionWithScheduleResponseBody
	GetSuccess() *string
}

type CreateProcessDefinitionWithScheduleResponseBody struct {
	// The code that is returned by the backend server.
	//
	// example:
	//
	// 1400009
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreateProcessDefinitionWithScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request failed.
	//
	// example:
	//
	// false
	Failed *string `json:"failed,omitempty" xml:"failed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the returned code.
	//
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetData() *CreateProcessDefinitionWithScheduleResponseBodyData {
	return s.Data
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetFailed() *string {
	return s.Failed
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetCode(v int32) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetData(v *CreateProcessDefinitionWithScheduleResponseBodyData) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Data = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetFailed(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Failed = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetHttpStatusCode(v int32) *CreateProcessDefinitionWithScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetMsg(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetRequestId(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetSuccess(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionWithScheduleResponseBodyData struct {
	// The workflow ID.
	//
	// example:
	//
	// 160************
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The serial number of the workflow.
	//
	// example:
	//
	// 12342
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) SetCode(v int64) *CreateProcessDefinitionWithScheduleResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) SetId(v int32) *CreateProcessDefinitionWithScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
