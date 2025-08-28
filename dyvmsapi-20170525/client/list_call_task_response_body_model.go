// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallTaskResponseBody
	GetCode() *string
	SetData(v []*ListCallTaskResponseBodyData) *ListCallTaskResponseBody
	GetData() []*ListCallTaskResponseBodyData
	SetPageNumber(v int64) *ListCallTaskResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCallTaskResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCallTaskResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCallTaskResponseBody
	GetTotal() *int64
}

type ListCallTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task information.
	Data []*ListCallTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5B0F201F-DCDA-45C2-AA92-1AE177F94991
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallTaskResponseBody) GetData() []*ListCallTaskResponseBodyData {
	return s.Data
}

func (s *ListCallTaskResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCallTaskResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallTaskResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCallTaskResponseBody) SetCode(v string) *ListCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallTaskResponseBody) SetData(v []*ListCallTaskResponseBodyData) *ListCallTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListCallTaskResponseBody) SetPageNumber(v int64) *ListCallTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskResponseBody) SetPageSize(v int64) *ListCallTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskResponseBody) SetRequestId(v string) *ListCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallTaskResponseBody) SetTotal(v int64) *ListCallTaskResponseBody {
	s.Total = &v
	return s
}

func (s *ListCallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCallTaskResponseBodyData struct {
	// The type of the task template. Valid values:
	//
	// 	- **VMS_VOICE_TTS**: the TTS notification template.
	//
	// 	- **VMS_VOICE_CODE**: the voice notification template.
	//
	// 	- **VMS_TTS**: the voice verification code template.
	//
	// example:
	//
	// VMS_VOICE_TTS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The time when the task was completed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1614330986000
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The number of tasks that were complete.
	//
	// example:
	//
	// 2
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// The task progress.
	//
	// example:
	//
	// 26%
	CompletedRate *int32 `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The type of the called number.
	//
	// example:
	//
	// LIST
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The time when the scheduled task was started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1614330978000
	FireTime *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 123879546214
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 0571000****
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The task state. Valid values:
	//
	// 	- **INIT**: The task was in the initial state.
	//
	// 	- **RELEASE**: The task was being parsed.
	//
	// 	- **RUNNING**: The task was running.
	//
	// 	- **STOP**: The task was manually suspended.
	//
	// 	- **SYSTEM_STOP**: The task was suspended by the system.
	//
	// 	- **CANCEL**: The task was manually canceled.
	//
	// 	- **SYSTEM_CANCEL**: The task was canceled by the system.
	//
	// 	- **DONE**: The task was complete.
	//
	// example:
	//
	// DONE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// Aliyun
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The ID of the voice template.
	//
	// example:
	//
	// TTS_2100****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// Test Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The total number of called numbers.
	//
	// example:
	//
	// 600
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCallTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallTaskResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *ListCallTaskResponseBodyData) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListCallTaskResponseBodyData) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *ListCallTaskResponseBodyData) GetCompletedRate() *int32 {
	return s.CompletedRate
}

func (s *ListCallTaskResponseBodyData) GetData() *string {
	return s.Data
}

func (s *ListCallTaskResponseBodyData) GetDataType() *string {
	return s.DataType
}

func (s *ListCallTaskResponseBodyData) GetFireTime() *string {
	return s.FireTime
}

func (s *ListCallTaskResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListCallTaskResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *ListCallTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListCallTaskResponseBodyData) GetStopTime() *string {
	return s.StopTime
}

func (s *ListCallTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListCallTaskResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListCallTaskResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListCallTaskResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCallTaskResponseBodyData) SetBizType(v string) *ListCallTaskResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetCompleteTime(v string) *ListCallTaskResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetCompletedCount(v int64) *ListCallTaskResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetCompletedRate(v int32) *ListCallTaskResponseBodyData {
	s.CompletedRate = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetData(v string) *ListCallTaskResponseBodyData {
	s.Data = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetDataType(v string) *ListCallTaskResponseBodyData {
	s.DataType = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetFireTime(v string) *ListCallTaskResponseBodyData {
	s.FireTime = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetId(v int64) *ListCallTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetResource(v string) *ListCallTaskResponseBodyData {
	s.Resource = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetStatus(v string) *ListCallTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetStopTime(v string) *ListCallTaskResponseBodyData {
	s.StopTime = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTaskName(v string) *ListCallTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTemplateCode(v string) *ListCallTaskResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTemplateName(v string) *ListCallTaskResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTotalCount(v int64) *ListCallTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
