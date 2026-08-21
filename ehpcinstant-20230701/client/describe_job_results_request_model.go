// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrayIndex(v int32) *DescribeJobResultsRequest
	GetArrayIndex() *int32
	SetContentEncoding(v string) *DescribeJobResultsRequest
	GetContentEncoding() *string
	SetJobId(v string) *DescribeJobResultsRequest
	GetJobId() *string
	SetLimitBytes(v string) *DescribeJobResultsRequest
	GetLimitBytes() *string
	SetStartTime(v string) *DescribeJobResultsRequest
	GetStartTime() *string
	SetTaskName(v string) *DescribeJobResultsRequest
	GetTaskName() *string
}

type DescribeJobResultsRequest struct {
	// The executor index.
	//
	// example:
	//
	// 0_1
	ArrayIndex *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// The encoding method for the CommandContent and Output fields in the response. Valid values:
	//
	// - PlainText: returns the original script content and output.
	//
	// - Base64: returns Base64-encoded script content and output.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The job ID.
	//
	// example:
	//
	// job-xxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum size of the log. Unit: bytes. Valid values: 1 to 1048576 (1 MB).
	//
	// example:
	//
	// 1048576
	LimitBytes *string `json:"LimitBytes,omitempty" xml:"LimitBytes,omitempty"`
	// The start time in UTC, in RFC 3339 format.
	//
	// example:
	//
	// 2024-09-02T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobResultsRequest) GetArrayIndex() *int32 {
	return s.ArrayIndex
}

func (s *DescribeJobResultsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeJobResultsRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResultsRequest) GetLimitBytes() *string {
	return s.LimitBytes
}

func (s *DescribeJobResultsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeJobResultsRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeJobResultsRequest) SetArrayIndex(v int32) *DescribeJobResultsRequest {
	s.ArrayIndex = &v
	return s
}

func (s *DescribeJobResultsRequest) SetContentEncoding(v string) *DescribeJobResultsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeJobResultsRequest) SetJobId(v string) *DescribeJobResultsRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobResultsRequest) SetLimitBytes(v string) *DescribeJobResultsRequest {
	s.LimitBytes = &v
	return s
}

func (s *DescribeJobResultsRequest) SetStartTime(v string) *DescribeJobResultsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeJobResultsRequest) SetTaskName(v string) *DescribeJobResultsRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeJobResultsRequest) Validate() error {
	return dara.Validate(s)
}
