// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListExecutorLogsResponseBodyLogs) *ListExecutorLogsResponseBody
	GetLogs() []*ListExecutorLogsResponseBodyLogs
	SetMaxResults(v int32) *ListExecutorLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExecutorLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExecutorLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListExecutorLogsResponseBody
	GetTotalCount() *int32
}

type ListExecutorLogsResponseBody struct {
	Logs []*ListExecutorLogsResponseBodyLogs `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListExecutorLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorLogsResponseBody) GetLogs() []*ListExecutorLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListExecutorLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutorLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutorLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutorLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListExecutorLogsResponseBody) SetLogs(v []*ListExecutorLogsResponseBodyLogs) *ListExecutorLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListExecutorLogsResponseBody) SetMaxResults(v int32) *ListExecutorLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutorLogsResponseBody) SetNextToken(v string) *ListExecutorLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutorLogsResponseBody) SetRequestId(v string) *ListExecutorLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorLogsResponseBody) SetTotalCount(v int32) *ListExecutorLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExecutorLogsResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutorLogsResponseBodyLogs struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 6383327
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// log4j.log
	LogName *string `json:"logName,omitempty" xml:"logName,omitempty"`
	// example:
	//
	// log4j
	LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
	// example:
	//
	// 1745390462
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListExecutorLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListExecutorLogsResponseBodyLogs) GetFileName() *string {
	return s.FileName
}

func (s *ListExecutorLogsResponseBodyLogs) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListExecutorLogsResponseBodyLogs) GetLogName() *string {
	return s.LogName
}

func (s *ListExecutorLogsResponseBodyLogs) GetLogType() *string {
	return s.LogType
}

func (s *ListExecutorLogsResponseBodyLogs) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListExecutorLogsResponseBodyLogs) SetFileName(v string) *ListExecutorLogsResponseBodyLogs {
	s.FileName = &v
	return s
}

func (s *ListExecutorLogsResponseBodyLogs) SetFileSize(v int64) *ListExecutorLogsResponseBodyLogs {
	s.FileSize = &v
	return s
}

func (s *ListExecutorLogsResponseBodyLogs) SetLogName(v string) *ListExecutorLogsResponseBodyLogs {
	s.LogName = &v
	return s
}

func (s *ListExecutorLogsResponseBodyLogs) SetLogType(v string) *ListExecutorLogsResponseBodyLogs {
	s.LogType = &v
	return s
}

func (s *ListExecutorLogsResponseBodyLogs) SetUpdateTime(v int64) *ListExecutorLogsResponseBodyLogs {
	s.UpdateTime = &v
	return s
}

func (s *ListExecutorLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
