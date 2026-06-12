// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestTaskLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceTestTaskLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTestTaskLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceTestTaskLogsResponseBody
	GetRequestId() *string
	SetTaskLogs(v []*ListServiceTestTaskLogsResponseBodyTaskLogs) *ListServiceTestTaskLogsResponseBody
	GetTaskLogs() []*ListServiceTestTaskLogsResponseBodyTaskLogs
}

type ListServiceTestTaskLogsResponseBody struct {
	// The number of entries returned per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDEE055B-D5F4-5B92-8F21-999D408F1214
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The log entries.
	TaskLogs []*ListServiceTestTaskLogsResponseBodyTaskLogs `json:"TaskLogs,omitempty" xml:"TaskLogs,omitempty" type:"Repeated"`
}

func (s ListServiceTestTaskLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTaskLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTestTaskLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTestTaskLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceTestTaskLogsResponseBody) GetTaskLogs() []*ListServiceTestTaskLogsResponseBodyTaskLogs {
	return s.TaskLogs
}

func (s *ListServiceTestTaskLogsResponseBody) SetMaxResults(v int32) *ListServiceTestTaskLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) SetNextToken(v string) *ListServiceTestTaskLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) SetRequestId(v string) *ListServiceTestTaskLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) SetTaskLogs(v []*ListServiceTestTaskLogsResponseBodyTaskLogs) *ListServiceTestTaskLogsResponseBody {
	s.TaskLogs = v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) Validate() error {
	if s.TaskLogs != nil {
		for _, item := range s.TaskLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceTestTaskLogsResponseBodyTaskLogs struct {
	// The log content.
	//
	// example:
	//
	// log content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The timestamp.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-12-12T20:00:09Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceTestTaskLogsResponseBodyTaskLogs) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTaskLogsResponseBodyTaskLogs) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) GetContent() *string {
	return s.Content
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) SetContent(v string) *ListServiceTestTaskLogsResponseBodyTaskLogs {
	s.Content = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) SetTimestamp(v string) *ListServiceTestTaskLogsResponseBodyTaskLogs {
	s.Timestamp = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) Validate() error {
	return dara.Validate(s)
}
