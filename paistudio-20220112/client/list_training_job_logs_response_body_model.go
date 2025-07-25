// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*string) *ListTrainingJobLogsResponseBody
	GetLogs() []*string
	SetRequestId(v string) *ListTrainingJobLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrainingJobLogsResponseBody
	GetTotalCount() *string
}

type ListTrainingJobLogsResponseBody struct {
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// CBF05F13-B24C-5129-9048-4FA684DCD579
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobLogsResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *ListTrainingJobLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrainingJobLogsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrainingJobLogsResponseBody) SetLogs(v []*string) *ListTrainingJobLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListTrainingJobLogsResponseBody) SetRequestId(v string) *ListTrainingJobLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobLogsResponseBody) SetTotalCount(v string) *ListTrainingJobLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
