// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVmcoreDiagnosisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVmcoreDiagnosisTaskResponseBody
	GetCode() *string
	SetData(v []*ListVmcoreDiagnosisTaskResponseBodyData) *ListVmcoreDiagnosisTaskResponseBody
	GetData() []*ListVmcoreDiagnosisTaskResponseBodyData
	SetMessage(v string) *ListVmcoreDiagnosisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVmcoreDiagnosisTaskResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListVmcoreDiagnosisTaskResponseBody
	GetTotal() *int64
}

type ListVmcoreDiagnosisTaskResponseBody struct {
	// example:
	//
	// Success
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListVmcoreDiagnosisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListVmcoreDiagnosisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVmcoreDiagnosisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListVmcoreDiagnosisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVmcoreDiagnosisTaskResponseBody) GetData() []*ListVmcoreDiagnosisTaskResponseBodyData {
	return s.Data
}

func (s *ListVmcoreDiagnosisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVmcoreDiagnosisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVmcoreDiagnosisTaskResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListVmcoreDiagnosisTaskResponseBody) SetCode(v string) *ListVmcoreDiagnosisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBody) SetData(v []*ListVmcoreDiagnosisTaskResponseBodyData) *ListVmcoreDiagnosisTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBody) SetMessage(v string) *ListVmcoreDiagnosisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBody) SetRequestId(v string) *ListVmcoreDiagnosisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBody) SetTotal(v int64) *ListVmcoreDiagnosisTaskResponseBody {
	s.Total = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBody) Validate() error {
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

type ListVmcoreDiagnosisTaskResponseBodyData struct {
	// example:
	//
	// 2025-12-02T17:36:12
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// error message
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// running
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	// example:
	//
	// vmcore
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s ListVmcoreDiagnosisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVmcoreDiagnosisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) SetCreatedAt(v string) *ListVmcoreDiagnosisTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) SetErrorMsg(v string) *ListVmcoreDiagnosisTaskResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) SetTaskId(v string) *ListVmcoreDiagnosisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) SetTaskStatus(v string) *ListVmcoreDiagnosisTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) SetTaskType(v string) *ListVmcoreDiagnosisTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
