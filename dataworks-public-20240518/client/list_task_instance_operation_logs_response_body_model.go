// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstanceOperationLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListTaskInstanceOperationLogsResponseBodyPagingInfo) *ListTaskInstanceOperationLogsResponseBody
	GetPagingInfo() *ListTaskInstanceOperationLogsResponseBodyPagingInfo
	SetRequestId(v string) *ListTaskInstanceOperationLogsResponseBody
	GetRequestId() *string
}

type ListTaskInstanceOperationLogsResponseBody struct {
	// The pagination information.
	PagingInfo *ListTaskInstanceOperationLogsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTaskInstanceOperationLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstanceOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskInstanceOperationLogsResponseBody) GetPagingInfo() *ListTaskInstanceOperationLogsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListTaskInstanceOperationLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskInstanceOperationLogsResponseBody) SetPagingInfo(v *ListTaskInstanceOperationLogsResponseBodyPagingInfo) *ListTaskInstanceOperationLogsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBody) SetRequestId(v string) *ListTaskInstanceOperationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskInstanceOperationLogsResponseBodyPagingInfo struct {
	// The operation logs.
	OperationLogs []*ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs `json:"OperationLogs,omitempty" xml:"OperationLogs,omitempty" type:"Repeated"`
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

func (s ListTaskInstanceOperationLogsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstanceOperationLogsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) GetOperationLogs() []*ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs {
	return s.OperationLogs
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) SetOperationLogs(v []*ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) *ListTaskInstanceOperationLogsResponseBodyPagingInfo {
	s.OperationLogs = v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) SetPageNumber(v int32) *ListTaskInstanceOperationLogsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) SetPageSize(v int32) *ListTaskInstanceOperationLogsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) SetTotalCount(v int32) *ListTaskInstanceOperationLogsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfo) Validate() error {
	if s.OperationLogs != nil {
		for _, item := range s.OperationLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs struct {
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
	// The ID of the instance on which the operation was performed.
	//
	// example:
	//
	// 1234
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 1000
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) GoString() string {
	return s.String()
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) GetOperationContent() *string {
	return s.OperationContent
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) GetOperationSeq() *int64 {
	return s.OperationSeq
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) GetUser() *string {
	return s.User
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) SetCreateTime(v int64) *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs {
	s.CreateTime = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) SetOperationContent(v string) *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs {
	s.OperationContent = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) SetOperationSeq(v int64) *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs {
	s.OperationSeq = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) SetTaskInstanceId(v int64) *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs {
	s.TaskInstanceId = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) SetUser(v string) *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs {
	s.User = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponseBodyPagingInfoOperationLogs) Validate() error {
	return dara.Validate(s)
}
