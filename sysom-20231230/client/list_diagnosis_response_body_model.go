// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDiagnosisResponseBody
	GetRequestId() *string
	SetCode(v string) *ListDiagnosisResponseBody
	GetCode() *string
	SetData(v []*ListDiagnosisResponseBodyData) *ListDiagnosisResponseBody
	GetData() []*ListDiagnosisResponseBodyData
	SetMessage(v string) *ListDiagnosisResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListDiagnosisResponseBody
	GetTotal() *int64
}

type ListDiagnosisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Total   *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDiagnosisResponseBody) GetData() []*ListDiagnosisResponseBodyData {
	return s.Data
}

func (s *ListDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDiagnosisResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListDiagnosisResponseBody) SetRequestId(v string) *ListDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetCode(v string) *ListDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetData(v []*ListDiagnosisResponseBodyData) *ListDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *ListDiagnosisResponseBody) SetMessage(v string) *ListDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetTotal(v int64) *ListDiagnosisResponseBody {
	s.Total = &v
	return s
}

func (s *ListDiagnosisResponseBody) Validate() error {
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

type ListDiagnosisResponseBodyData struct {
	Code        *int32      `json:"code,omitempty" xml:"code,omitempty"`
	Command     interface{} `json:"command,omitempty" xml:"command,omitempty"`
	CreatedAt   *string     `json:"created_at,omitempty" xml:"created_at,omitempty"`
	ErrMsg      *string     `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	Params      interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Result      interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ServiceName *string     `json:"service_name,omitempty" xml:"service_name,omitempty"`
	Status      *string     `json:"status,omitempty" xml:"status,omitempty"`
	TaskId      *string     `json:"task_id,omitempty" xml:"task_id,omitempty"`
	UpdatedAt   *string     `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Url         *string     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListDiagnosisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ListDiagnosisResponseBodyData) GetCommand() interface{} {
	return s.Command
}

func (s *ListDiagnosisResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListDiagnosisResponseBodyData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ListDiagnosisResponseBodyData) GetParams() interface{} {
	return s.Params
}

func (s *ListDiagnosisResponseBodyData) GetResult() interface{} {
	return s.Result
}

func (s *ListDiagnosisResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListDiagnosisResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDiagnosisResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListDiagnosisResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListDiagnosisResponseBodyData) SetCode(v int32) *ListDiagnosisResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetCommand(v interface{}) *ListDiagnosisResponseBodyData {
	s.Command = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetCreatedAt(v string) *ListDiagnosisResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetErrMsg(v string) *ListDiagnosisResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetParams(v interface{}) *ListDiagnosisResponseBodyData {
	s.Params = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetResult(v interface{}) *ListDiagnosisResponseBodyData {
	s.Result = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetServiceName(v string) *ListDiagnosisResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetStatus(v string) *ListDiagnosisResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetTaskId(v string) *ListDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetUpdatedAt(v string) *ListDiagnosisResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetUrl(v string) *ListDiagnosisResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
