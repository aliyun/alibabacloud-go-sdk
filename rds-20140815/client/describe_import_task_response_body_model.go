// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *DescribeImportTaskResponseBody
	GetAccount() *string
	SetDbVersion(v string) *DescribeImportTaskResponseBody
	GetDbVersion() *string
	SetDetail(v string) *DescribeImportTaskResponseBody
	GetDetail() *string
	SetRequestId(v string) *DescribeImportTaskResponseBody
	GetRequestId() *string
	SetSourceCategory(v string) *DescribeImportTaskResponseBody
	GetSourceCategory() *string
	SetSourceIp(v string) *DescribeImportTaskResponseBody
	GetSourceIp() *string
	SetSourcePort(v string) *DescribeImportTaskResponseBody
	GetSourcePort() *string
	SetStatus(v string) *DescribeImportTaskResponseBody
	GetStatus() *string
	SetTargetInstanceName(v string) *DescribeImportTaskResponseBody
	GetTargetInstanceName() *string
	SetTaskId(v int64) *DescribeImportTaskResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *DescribeImportTaskResponseBody
	GetTaskName() *string
	SetTaskType(v string) *DescribeImportTaskResponseBody
	GetTaskType() *string
}

type DescribeImportTaskResponseBody struct {
	// example:
	//
	// myadmin
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// 5.7
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// Error Message
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// A103039D-B1B2-4C57-B989-7D7C0DA95426
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// aliyunRDS
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// example:
	//
	// 59.172.25.122
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 3306
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// example:
	//
	// Importing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// rm-t4neh0q12v1******
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// example:
	//
	// 416980000
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// test01
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// import
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportTaskResponseBody) GetAccount() *string {
	return s.Account
}

func (s *DescribeImportTaskResponseBody) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeImportTaskResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *DescribeImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImportTaskResponseBody) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *DescribeImportTaskResponseBody) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeImportTaskResponseBody) GetSourcePort() *string {
	return s.SourcePort
}

func (s *DescribeImportTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeImportTaskResponseBody) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *DescribeImportTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeImportTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeImportTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeImportTaskResponseBody) SetAccount(v string) *DescribeImportTaskResponseBody {
	s.Account = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetDbVersion(v string) *DescribeImportTaskResponseBody {
	s.DbVersion = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetDetail(v string) *DescribeImportTaskResponseBody {
	s.Detail = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetRequestId(v string) *DescribeImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetSourceCategory(v string) *DescribeImportTaskResponseBody {
	s.SourceCategory = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetSourceIp(v string) *DescribeImportTaskResponseBody {
	s.SourceIp = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetSourcePort(v string) *DescribeImportTaskResponseBody {
	s.SourcePort = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetStatus(v string) *DescribeImportTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetTargetInstanceName(v string) *DescribeImportTaskResponseBody {
	s.TargetInstanceName = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetTaskId(v int64) *DescribeImportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetTaskName(v string) *DescribeImportTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeImportTaskResponseBody) SetTaskType(v string) *DescribeImportTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
