// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeDiagnosisResponseBody
	GetCode() *string
	SetData(v *InvokeDiagnosisResponseBodyData) *InvokeDiagnosisResponseBody
	GetData() *InvokeDiagnosisResponseBodyData
	SetMessage(v string) *InvokeDiagnosisResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeDiagnosisResponseBody
	GetRequestId() *string
}

type InvokeDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *InvokeDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s InvokeDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeDiagnosisResponseBody) GetData() *InvokeDiagnosisResponseBodyData {
	return s.Data
}

func (s *InvokeDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeDiagnosisResponseBody) SetCode(v string) *InvokeDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetData(v *InvokeDiagnosisResponseBodyData) *InvokeDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetMessage(v string) *InvokeDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetRequestId(v string) *InvokeDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}

type InvokeDiagnosisResponseBodyData struct {
	// example:
	//
	// ihqhAcrt
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InvokeDiagnosisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InvokeDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *InvokeDiagnosisResponseBodyData) SetTaskId(v string) *InvokeDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *InvokeDiagnosisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
