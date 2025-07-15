// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitAuditTaskResponseBody
	GetCode() *string
	SetData(v *SubmitAuditTaskResponseBodyData) *SubmitAuditTaskResponseBody
	GetData() *SubmitAuditTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitAuditTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitAuditTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAuditTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAuditTaskResponseBody
	GetSuccess() *bool
}

type SubmitAuditTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitAuditTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAuditTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAuditTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAuditTaskResponseBody) GetData() *SubmitAuditTaskResponseBodyData {
	return s.Data
}

func (s *SubmitAuditTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitAuditTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAuditTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAuditTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAuditTaskResponseBody) SetCode(v string) *SubmitAuditTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAuditTaskResponseBody) SetData(v *SubmitAuditTaskResponseBodyData) *SubmitAuditTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAuditTaskResponseBody) SetHttpStatusCode(v int32) *SubmitAuditTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAuditTaskResponseBody) SetMessage(v string) *SubmitAuditTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAuditTaskResponseBody) SetRequestId(v string) *SubmitAuditTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAuditTaskResponseBody) SetSuccess(v bool) *SubmitAuditTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAuditTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAuditTaskResponseBodyData struct {
	// example:
	//
	// xx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitAuditTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAuditTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitAuditTaskResponseBodyData) SetTaskId(v string) *SubmitAuditTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitAuditTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
