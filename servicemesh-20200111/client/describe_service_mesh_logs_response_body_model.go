// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*DescribeServiceMeshLogsResponseBodyLogs) *DescribeServiceMeshLogsResponseBody
	GetLogs() []*DescribeServiceMeshLogsResponseBodyLogs
	SetRequestId(v string) *DescribeServiceMeshLogsResponseBody
	GetRequestId() *string
}

type DescribeServiceMeshLogsResponseBody struct {
	// The details of the logs.
	Logs []*DescribeServiceMeshLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsResponseBody) GetLogs() []*DescribeServiceMeshLogsResponseBodyLogs {
	return s.Logs
}

func (s *DescribeServiceMeshLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshLogsResponseBody) SetLogs(v []*DescribeServiceMeshLogsResponseBodyLogs) *DescribeServiceMeshLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeServiceMeshLogsResponseBody) SetRequestId(v string) *DescribeServiceMeshLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshLogsResponseBody) Validate() error {
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

type DescribeServiceMeshLogsResponseBodyLogs struct {
	// The point in time when the logs were generated.
	//
	// example:
	//
	// 2021-11-19T15:21:53+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The content of the logs.
	//
	// example:
	//
	// [RequestID: 31d3a0f0-07ed-4f6e-9004-1804498c****, UID-110982038403****] c096d641835af4658827a4c66c234***	- | Start to add cluster c186a6d9641a24098b5499d4d8313****
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// ca04bc38979214bf2882be79d39b4****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) GetLog() *string {
	return s.Log
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) SetCreationTime(v string) *DescribeServiceMeshLogsResponseBodyLogs {
	s.CreationTime = &v
	return s
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) SetLog(v string) *DescribeServiceMeshLogsResponseBodyLogs {
	s.Log = &v
	return s
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) SetServiceMeshId(v string) *DescribeServiceMeshLogsResponseBodyLogs {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
