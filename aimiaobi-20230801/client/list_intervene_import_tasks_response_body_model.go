// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneImportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInterveneImportTasksResponseBody
	GetCode() *string
	SetData(v *ListInterveneImportTasksResponseBodyData) *ListInterveneImportTasksResponseBody
	GetData() *ListInterveneImportTasksResponseBodyData
	SetHttpStatusCode(v int32) *ListInterveneImportTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInterveneImportTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInterveneImportTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInterveneImportTasksResponseBody
	GetSuccess() *bool
}

type ListInterveneImportTasksResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInterveneImportTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneImportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneImportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInterveneImportTasksResponseBody) GetData() *ListInterveneImportTasksResponseBodyData {
	return s.Data
}

func (s *ListInterveneImportTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInterveneImportTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInterveneImportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterveneImportTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInterveneImportTasksResponseBody) SetCode(v string) *ListInterveneImportTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetData(v *ListInterveneImportTasksResponseBodyData) *ListInterveneImportTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetHttpStatusCode(v int32) *ListInterveneImportTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetMessage(v string) *ListInterveneImportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetRequestId(v string) *ListInterveneImportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetSuccess(v bool) *ListInterveneImportTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInterveneImportTasksResponseBodyData struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*ListInterveneImportTasksResponseBodyDataStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListInterveneImportTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneImportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ListInterveneImportTasksResponseBodyData) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListInterveneImportTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterveneImportTasksResponseBodyData) GetStatusList() []*ListInterveneImportTasksResponseBodyDataStatusList {
	return s.StatusList
}

func (s *ListInterveneImportTasksResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListInterveneImportTasksResponseBodyData) SetCode(v int32) *ListInterveneImportTasksResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetPageIndex(v int32) *ListInterveneImportTasksResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetPageSize(v int32) *ListInterveneImportTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetStatusList(v []*ListInterveneImportTasksResponseBodyDataStatusList) *ListInterveneImportTasksResponseBodyData {
	s.StatusList = v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetTotalSize(v int32) *ListInterveneImportTasksResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) Validate() error {
	if s.StatusList != nil {
		for _, item := range s.StatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInterveneImportTasksResponseBodyDataStatusList struct {
	// example:
	//
	// Success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 5
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// Success
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 4854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 12344454
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListInterveneImportTasksResponseBodyDataStatusList) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneImportTasksResponseBodyDataStatusList) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) GetMsg() *string {
	return s.Msg
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) GetStatus() *int32 {
	return s.Status
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) GetTaskName() *string {
	return s.TaskName
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetMsg(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Msg = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetPercentage(v int32) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Percentage = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetStatus(v int32) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Status = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetTaskId(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.TaskId = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetTaskName(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.TaskName = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) Validate() error {
	return dara.Validate(s)
}
