// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanClusterVulsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ScanClusterVulsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ScanClusterVulsResponseBody
	GetTaskId() *string
}

type ScanClusterVulsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314A1E1
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Task ID.
	//
	// example:
	//
	// T-xascadasd*****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ScanClusterVulsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanClusterVulsResponseBody) GoString() string {
	return s.String()
}

func (s *ScanClusterVulsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanClusterVulsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ScanClusterVulsResponseBody) SetRequestId(v string) *ScanClusterVulsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanClusterVulsResponseBody) SetTaskId(v string) *ScanClusterVulsResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScanClusterVulsResponseBody) Validate() error {
	return dara.Validate(s)
}
