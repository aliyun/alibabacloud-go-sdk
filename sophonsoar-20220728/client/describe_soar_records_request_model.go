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
	// The start time when the task was completed. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1755676363777
	CompletedBeginTime *int64 `json:"CompletedBeginTime,omitempty" xml:"CompletedBeginTime,omitempty"`
	// The end time when the task was completed. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1683526284584
	CompletedEndTime *int64 `json:"CompletedEndTime,omitempty" xml:"CompletedEndTime,omitempty"`
	// The end time of the task run. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1683772744953
	EndMillis *int64 `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 10 entries are returned on each page.
	//
	// > Specify a value for PageSize.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the playbook.
	//
	// > For more information, see [DescribePlaybooks](~~DescribePlaybooks~~).
	//
	// example:
	//
	// 8f55e76d-b5d5-4720-9cd7-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The input parameter of the playbook.
	//
	// example:
	//
	// input
	QueryValue *string `json:"QueryValue,omitempty" xml:"QueryValue,omitempty"`
	// The UUID of the playbook task execution.
	//
	// > For more information, see [DescribeSoarRecords](https://help.aliyun.com/document_detail/2627455.html).
	//
	// example:
	//
	// 6d412cfa-0905-4567-8a83-xxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The start time of the task run. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1683526284584
	StartMillis *int64 `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
	// The status of the task run. Valid values:
	//
	// - **success**: The task is successful.
	//
	// - **failed**: The task failed.
	//
	// - **inprogress**: The task is in progress.
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
	// The trigger type of the task. Valid values:
	//
	// - **stream**: stream
	//
	// - **debug**: test
	//
	// - **manual**: manual
	//
	// - **timer**: scheduled
	//
	// - **SubInvoke**: child flow
	//
	// - **siem**: triggered by a SIEM product
	//
	// example:
	//
	// debug
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The ID of the Alibaba Cloud account that runs the playbook task.
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
