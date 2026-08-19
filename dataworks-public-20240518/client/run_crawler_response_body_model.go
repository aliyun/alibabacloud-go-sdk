// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RunCrawlerResponseBody
	GetId() *int64
	SetRequestId(v string) *RunCrawlerResponseBody
	GetRequestId() *string
	SetRunAccepted(v bool) *RunCrawlerResponseBody
	GetRunAccepted() *bool
	SetRunStatus(v string) *RunCrawlerResponseBody
	GetRunStatus() *string
	SetSuccess(v bool) *RunCrawlerResponseBody
	GetSuccess() *bool
	SetTaskInstanceId(v int64) *RunCrawlerResponseBody
	GetTaskInstanceId() *int64
}

type RunCrawlerResponseBody struct {
	// The ID of the metadata crawler.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the run request was accepted. A value of true indicates that the request was accepted, but does not indicate that the collection task is complete.
	RunAccepted *bool `json:"RunAccepted,omitempty" xml:"RunAccepted,omitempty"`
	// The initial run status after submission. The value is WAITING when the run request is successfully accepted. To query the final status, call ListCrawlerRuns.
	//
	// example:
	//
	// WAITING
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The associated DataWorks task instance ID. This field may be empty. To query the final run record, call ListCrawlerRuns.
	//
	// example:
	//
	// 1234
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s RunCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *RunCrawlerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *RunCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCrawlerResponseBody) GetRunAccepted() *bool {
	return s.RunAccepted
}

func (s *RunCrawlerResponseBody) GetRunStatus() *string {
	return s.RunStatus
}

func (s *RunCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunCrawlerResponseBody) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *RunCrawlerResponseBody) SetId(v int64) *RunCrawlerResponseBody {
	s.Id = &v
	return s
}

func (s *RunCrawlerResponseBody) SetRequestId(v string) *RunCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCrawlerResponseBody) SetRunAccepted(v bool) *RunCrawlerResponseBody {
	s.RunAccepted = &v
	return s
}

func (s *RunCrawlerResponseBody) SetRunStatus(v string) *RunCrawlerResponseBody {
	s.RunStatus = &v
	return s
}

func (s *RunCrawlerResponseBody) SetSuccess(v bool) *RunCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *RunCrawlerResponseBody) SetTaskInstanceId(v int64) *RunCrawlerResponseBody {
	s.TaskInstanceId = &v
	return s
}

func (s *RunCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
