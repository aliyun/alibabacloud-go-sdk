// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskOperationLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListTaskOperationLogsResponseBodyPagingInfo) *ListTaskOperationLogsResponseBody
	GetPagingInfo() *ListTaskOperationLogsResponseBodyPagingInfo
	SetRequestId(v string) *ListTaskOperationLogsResponseBody
	GetRequestId() *string
}

type ListTaskOperationLogsResponseBody struct {
	// The pagination information.
	PagingInfo *ListTaskOperationLogsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTaskOperationLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskOperationLogsResponseBody) GetPagingInfo() *ListTaskOperationLogsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListTaskOperationLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskOperationLogsResponseBody) SetPagingInfo(v *ListTaskOperationLogsResponseBodyPagingInfo) *ListTaskOperationLogsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListTaskOperationLogsResponseBody) SetRequestId(v string) *ListTaskOperationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskOperationLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTaskOperationLogsResponseBodyPagingInfo struct {
	// The operation logs.
	OperationLogs []*ListTaskOperationLogsResponseBodyPagingInfoOperationLogs `json:"OperationLogs,omitempty" xml:"OperationLogs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskOperationLogsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTaskOperationLogsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) GetOperationLogs() []*ListTaskOperationLogsResponseBodyPagingInfoOperationLogs {
	return s.OperationLogs
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) SetOperationLogs(v []*ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) *ListTaskOperationLogsResponseBodyPagingInfo {
	s.OperationLogs = v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) SetPageNumber(v int32) *ListTaskOperationLogsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) SetPageSize(v int32) *ListTaskOperationLogsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) SetTotalCount(v int32) *ListTaskOperationLogsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListTaskOperationLogsResponseBodyPagingInfoOperationLogs struct {
	// The time when the operation log was generated.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The operation content.
	//
	// example:
	//
	// Freeze tasks
	OperationContent *string `json:"OperationContent,omitempty" xml:"OperationContent,omitempty"`
	// The serial number of the operation.
	//
	// example:
	//
	// 1111
	OperationSeq *int64 `json:"OperationSeq,omitempty" xml:"OperationSeq,omitempty"`
	// The ID of the task on which the operation was performed.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 1000
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) String() string {
	return dara.Prettify(s)
}

func (s ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) GoString() string {
	return s.String()
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) GetOperationContent() *string {
	return s.OperationContent
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) GetOperationSeq() *int64 {
	return s.OperationSeq
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) GetUser() *string {
	return s.User
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) SetCreateTime(v int64) *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs {
	s.CreateTime = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) SetOperationContent(v string) *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs {
	s.OperationContent = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) SetOperationSeq(v int64) *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs {
	s.OperationSeq = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) SetTaskId(v int64) *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs {
	s.TaskId = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) SetUser(v string) *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs {
	s.User = &v
	return s
}

func (s *ListTaskOperationLogsResponseBodyPagingInfoOperationLogs) Validate() error {
	return dara.Validate(s)
}
