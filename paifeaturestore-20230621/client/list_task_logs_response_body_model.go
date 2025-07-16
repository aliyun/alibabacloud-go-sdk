// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*string) *ListTaskLogsResponseBody
	GetLogs() []*string
	SetRequestId(v string) *ListTaskLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTaskLogsResponseBody
	GetTotalCount() *int32
}

type ListTaskLogsResponseBody struct {
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 72F15A8A-5A28-5B18-A0DE-0EABD7D3245A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskLogsResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *ListTaskLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskLogsResponseBody) SetLogs(v []*string) *ListTaskLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListTaskLogsResponseBody) SetRequestId(v string) *ListTaskLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskLogsResponseBody) SetTotalCount(v int32) *ListTaskLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTaskLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
