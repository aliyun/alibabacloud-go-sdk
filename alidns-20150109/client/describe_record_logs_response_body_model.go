// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeRecordLogsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRecordLogsResponseBody
	GetPageSize() *int64
	SetRecordLogs(v *DescribeRecordLogsResponseBodyRecordLogs) *DescribeRecordLogsResponseBody
	GetRecordLogs() *DescribeRecordLogsResponseBodyRecordLogs
	SetRequestId(v string) *DescribeRecordLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeRecordLogsResponseBody
	GetTotalCount() *int64
}

type DescribeRecordLogsResponseBody struct {
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
	// The operation logs.
	RecordLogs *DescribeRecordLogsResponseBodyRecordLogs `json:"RecordLogs,omitempty" xml:"RecordLogs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRecordLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRecordLogsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRecordLogsResponseBody) GetRecordLogs() *DescribeRecordLogsResponseBodyRecordLogs {
	return s.RecordLogs
}

func (s *DescribeRecordLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRecordLogsResponseBody) SetPageNumber(v int64) *DescribeRecordLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetPageSize(v int64) *DescribeRecordLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetRecordLogs(v *DescribeRecordLogsResponseBodyRecordLogs) *DescribeRecordLogsResponseBody {
	s.RecordLogs = v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetRequestId(v string) *DescribeRecordLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetTotalCount(v int64) *DescribeRecordLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordLogsResponseBodyRecordLogs struct {
	RecordLog []*DescribeRecordLogsResponseBodyRecordLogsRecordLog `json:"RecordLog,omitempty" xml:"RecordLog,omitempty" type:"Repeated"`
}

func (s DescribeRecordLogsResponseBodyRecordLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordLogsResponseBodyRecordLogs) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) GetRecordLog() []*DescribeRecordLogsResponseBodyRecordLogsRecordLog {
	return s.RecordLog
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) SetRecordLog(v []*DescribeRecordLogsResponseBodyRecordLogsRecordLog) *DescribeRecordLogsResponseBodyRecordLogs {
	s.RecordLog = v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordLogsResponseBodyRecordLogsRecordLog struct {
	// The operation that you performed.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when you performed the operation.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	ActionTime *string `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	// The time when you performed the operation. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 134514540000
	ActionTimestamp *int64 `json:"ActionTimestamp,omitempty" xml:"ActionTimestamp,omitempty"`
	// The IP address of the operator.
	//
	// example:
	//
	// 182.92.253.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The operation message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeRecordLogsResponseBodyRecordLogsRecordLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordLogsResponseBodyRecordLogsRecordLog) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) GetAction() *string {
	return s.Action
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) GetActionTime() *string {
	return s.ActionTime
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) GetActionTimestamp() *int64 {
	return s.ActionTimestamp
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) GetMessage() *string {
	return s.Message
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) SetAction(v string) *DescribeRecordLogsResponseBodyRecordLogsRecordLog {
	s.Action = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) SetActionTime(v string) *DescribeRecordLogsResponseBodyRecordLogsRecordLog {
	s.ActionTime = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) SetActionTimestamp(v int64) *DescribeRecordLogsResponseBodyRecordLogsRecordLog {
	s.ActionTimestamp = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) SetClientIp(v string) *DescribeRecordLogsResponseBodyRecordLogsRecordLog {
	s.ClientIp = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) SetMessage(v string) *DescribeRecordLogsResponseBodyRecordLogsRecordLog {
	s.Message = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogsRecordLog) Validate() error {
	return dara.Validate(s)
}
