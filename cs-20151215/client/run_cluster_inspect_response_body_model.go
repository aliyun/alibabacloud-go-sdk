// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterInspectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *RunClusterInspectResponseBody
	GetReportId() *string
	SetRequestId(v string) *RunClusterInspectResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RunClusterInspectResponseBody
	GetTaskId() *string
}

type RunClusterInspectResponseBody struct {
	// The inspection report ID.
	//
	// example:
	//
	// 5d6557c983064c45bed62ab2a2119cc7
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AB4D067-4DD7-5471-B90A-FCC564BC3337
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The inspection task ID.
	//
	// example:
	//
	// T-67d7ec016ce37c0106000***
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunClusterInspectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunClusterInspectResponseBody) GoString() string {
	return s.String()
}

func (s *RunClusterInspectResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *RunClusterInspectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunClusterInspectResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RunClusterInspectResponseBody) SetReportId(v string) *RunClusterInspectResponseBody {
	s.ReportId = &v
	return s
}

func (s *RunClusterInspectResponseBody) SetRequestId(v string) *RunClusterInspectResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunClusterInspectResponseBody) SetTaskId(v string) *RunClusterInspectResponseBody {
	s.TaskId = &v
	return s
}

func (s *RunClusterInspectResponseBody) Validate() error {
	return dara.Validate(s)
}
