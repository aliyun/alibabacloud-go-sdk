// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneImportTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInterveneImportTaskInfoResponseBody
	GetCode() *string
	SetData(v *GetInterveneImportTaskInfoResponseBodyData) *GetInterveneImportTaskInfoResponseBody
	GetData() *GetInterveneImportTaskInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetInterveneImportTaskInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInterveneImportTaskInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInterveneImportTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInterveneImportTaskInfoResponseBody
	GetSuccess() *bool
}

type GetInterveneImportTaskInfoResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneImportTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneImportTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInterveneImportTaskInfoResponseBody) GetData() *GetInterveneImportTaskInfoResponseBodyData {
	return s.Data
}

func (s *GetInterveneImportTaskInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInterveneImportTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInterveneImportTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterveneImportTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInterveneImportTaskInfoResponseBody) SetCode(v string) *GetInterveneImportTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetData(v *GetInterveneImportTaskInfoResponseBodyData) *GetInterveneImportTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetHttpStatusCode(v int32) *GetInterveneImportTaskInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetMessage(v string) *GetInterveneImportTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetRequestId(v string) *GetInterveneImportTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetSuccess(v bool) *GetInterveneImportTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneImportTaskInfoResponseBodyData struct {
	Code   *int32                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Status *GetInterveneImportTaskInfoResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s GetInterveneImportTaskInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *GetInterveneImportTaskInfoResponseBodyData) GetStatus() *GetInterveneImportTaskInfoResponseBodyDataStatus {
	return s.Status
}

func (s *GetInterveneImportTaskInfoResponseBodyData) SetCode(v int32) *GetInterveneImportTaskInfoResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyData) SetStatus(v *GetInterveneImportTaskInfoResponseBodyDataStatus) *GetInterveneImportTaskInfoResponseBodyData {
	s.Status = v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyData) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneImportTaskInfoResponseBodyDataStatus struct {
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 80
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 41405255
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ft-task-20190101m8rnK
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetInterveneImportTaskInfoResponseBodyDataStatus) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) GetMsg() *string {
	return s.Msg
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) GetPercentage() *int32 {
	return s.Percentage
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) GetStatus() *int32 {
	return s.Status
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) GetTaskName() *string {
	return s.TaskName
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetMsg(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Msg = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetPercentage(v int32) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Percentage = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetStatus(v int32) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Status = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetTaskId(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.TaskId = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetTaskName(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.TaskName = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) Validate() error {
	return dara.Validate(s)
}
