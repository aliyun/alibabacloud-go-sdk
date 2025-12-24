// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedBeginTime(v int64) *DescribeSoarRecordsRequest
	GetCompletedBeginTime() *int64
	SetCompletedEndTime(v int64) *DescribeSoarRecordsRequest
	GetCompletedEndTime() *int64
	SetEndMillis(v int64) *DescribeSoarRecordsRequest
	GetEndMillis() *int64
	SetLang(v string) *DescribeSoarRecordsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeSoarRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarRecordsRequest
	GetPageSize() *int32
	SetPlaybookUuid(v string) *DescribeSoarRecordsRequest
	GetPlaybookUuid() *string
	SetQueryValue(v string) *DescribeSoarRecordsRequest
	GetQueryValue() *string
	SetRequestUuid(v string) *DescribeSoarRecordsRequest
	GetRequestUuid() *string
	SetStartMillis(v int64) *DescribeSoarRecordsRequest
	GetStartMillis() *int64
	SetTaskStatus(v string) *DescribeSoarRecordsRequest
	GetTaskStatus() *string
	SetTaskflowMd5(v string) *DescribeSoarRecordsRequest
	GetTaskflowMd5() *string
	SetTriggerType(v string) *DescribeSoarRecordsRequest
	GetTriggerType() *string
	SetTriggerUser(v string) *DescribeSoarRecordsRequest
	GetTriggerUser() *string
}

type DescribeSoarRecordsRequest struct {
	CompletedBeginTime *int64 `json:"CompletedBeginTime,omitempty" xml:"CompletedBeginTime,omitempty"`
	CompletedEndTime   *int64 `json:"CompletedEndTime,omitempty" xml:"CompletedEndTime,omitempty"`
	// The end time of the task execution, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683772744953
	EndMillis *int64 `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	// Set the language type for requests and received messages. The default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Set which page to start displaying the query results from. The default value is 1, indicating the first page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specify the maximum number of data entries per page when performing a paginated query. The default number of entries per page is 20. If the PageSize parameter is empty, it will return 10 entries by default.
	//
	// > It is recommended not to leave the PageSize value empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the playbook.
	//
	// > You can obtain this parameter by calling the [DescribePlaybooks](~~DescribePlaybooks~~) interface.
	//
	// example:
	//
	// 8f55e76d-b5d5-4720-9cd7-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	QueryValue   *string `json:"QueryValue,omitempty" xml:"QueryValue,omitempty"`
	// UUID of the playbook task execution.
	//
	// > You can obtain this parameter by calling the [DescribeSoarRecords](https://help.aliyun.com/document_detail/2627455.html) interface.
	//
	// example:
	//
	// 6d412cfa-0905-4567-8a83-xxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The start time of the task execution, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683526284584
	StartMillis *int64 `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
	// The status of the task execution. Values:
	//
	// - **success**: Successful task.
	//
	// - **failed**: Failed task.
	//
	// - **inprogress**: Task in progress
	//
	// example:
	//
	// inprogress
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The MD5 value of the playbook configuration.
	//
	// example:
	//
	// be0a4ef084dd174abe478df52xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The Alibaba Cloud account ID that executed the playbook task.
	//
	// example:
	//
	// 127xxxx4392
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsRequest) GetCompletedBeginTime() *int64 {
	return s.CompletedBeginTime
}

func (s *DescribeSoarRecordsRequest) GetCompletedEndTime() *int64 {
	return s.CompletedEndTime
}

func (s *DescribeSoarRecordsRequest) GetEndMillis() *int64 {
	return s.EndMillis
}

func (s *DescribeSoarRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSoarRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarRecordsRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeSoarRecordsRequest) GetQueryValue() *string {
	return s.QueryValue
}

func (s *DescribeSoarRecordsRequest) GetRequestUuid() *string {
	return s.RequestUuid
}

func (s *DescribeSoarRecordsRequest) GetStartMillis() *int64 {
	return s.StartMillis
}

func (s *DescribeSoarRecordsRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeSoarRecordsRequest) GetTaskflowMd5() *string {
	return s.TaskflowMd5
}

func (s *DescribeSoarRecordsRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *DescribeSoarRecordsRequest) GetTriggerUser() *string {
	return s.TriggerUser
}

func (s *DescribeSoarRecordsRequest) SetCompletedBeginTime(v int64) *DescribeSoarRecordsRequest {
	s.CompletedBeginTime = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetCompletedEndTime(v int64) *DescribeSoarRecordsRequest {
	s.CompletedEndTime = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetEndMillis(v int64) *DescribeSoarRecordsRequest {
	s.EndMillis = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetLang(v string) *DescribeSoarRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPageNumber(v int32) *DescribeSoarRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPageSize(v int32) *DescribeSoarRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPlaybookUuid(v string) *DescribeSoarRecordsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetQueryValue(v string) *DescribeSoarRecordsRequest {
	s.QueryValue = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetRequestUuid(v string) *DescribeSoarRecordsRequest {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetStartMillis(v int64) *DescribeSoarRecordsRequest {
	s.StartMillis = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTaskStatus(v string) *DescribeSoarRecordsRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTaskflowMd5(v string) *DescribeSoarRecordsRequest {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTriggerType(v string) *DescribeSoarRecordsRequest {
	s.TriggerType = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTriggerUser(v string) *DescribeSoarRecordsRequest {
	s.TriggerUser = &v
	return s
}

func (s *DescribeSoarRecordsRequest) Validate() error {
	return dara.Validate(s)
}
