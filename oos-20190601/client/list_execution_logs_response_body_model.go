// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionLogs(v []*ListExecutionLogsResponseBodyExecutionLogs) *ListExecutionLogsResponseBody
	GetExecutionLogs() []*ListExecutionLogsResponseBodyExecutionLogs
	SetIsTruncated(v bool) *ListExecutionLogsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListExecutionLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExecutionLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExecutionLogsResponseBody
	GetRequestId() *string
}

type ListExecutionLogsResponseBody struct {
	// The logs of the execution.
	ExecutionLogs []*ListExecutionLogsResponseBodyExecutionLogs `json:"ExecutionLogs,omitempty" xml:"ExecutionLogs,omitempty" type:"Repeated"`
	// Indicates whether the log is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABdpsGWjX8dJ-a6dl_pvoS7AFxNHSNJKHLCAJJ0ylgA53nWW5V4HTEZKCYTaEPNOrxFir4z43UTOjE150cFr8AGTifA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExecutionLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponseBody) GetExecutionLogs() []*ListExecutionLogsResponseBodyExecutionLogs {
	return s.ExecutionLogs
}

func (s *ListExecutionLogsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListExecutionLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutionLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutionLogsResponseBody) SetExecutionLogs(v []*ListExecutionLogsResponseBodyExecutionLogs) *ListExecutionLogsResponseBody {
	s.ExecutionLogs = v
	return s
}

func (s *ListExecutionLogsResponseBody) SetIsTruncated(v bool) *ListExecutionLogsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetMaxResults(v int32) *ListExecutionLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetNextToken(v string) *ListExecutionLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetRequestId(v string) *ListExecutionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionLogsResponseBody) Validate() error {
	if s.ExecutionLogs != nil {
		for _, item := range s.ExecutionLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutionLogsResponseBodyExecutionLogs struct {
	// The log type.
	//
	// example:
	//
	// System
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The details of the task execution.
	//
	// example:
	//
	// The task CheckDiskCategory completed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task execution ID.
	//
	// example:
	//
	// exec-1234567zxcvb.t0010
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The timestamp when the task was run.
	//
	// example:
	//
	// 2019-05-24T:02:29:07Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListExecutionLogsResponseBodyExecutionLogs) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionLogsResponseBodyExecutionLogs) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) GetLogType() *string {
	return s.LogType
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) GetMessage() *string {
	return s.Message
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetLogType(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.LogType = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetMessage(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.Message = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetTaskExecutionId(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetTimestamp(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.Timestamp = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) Validate() error {
	return dara.Validate(s)
}
