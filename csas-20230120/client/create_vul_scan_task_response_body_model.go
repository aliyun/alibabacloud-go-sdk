// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVulScanTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateVulScanTaskResponseBody
	GetTaskId() *string
}

type CreateVulScanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created vulnerability scanning task.
	//
	// example:
	//
	// vul-scan-task-4d7b1e9a6c38****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVulScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVulScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVulScanTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVulScanTaskResponseBody) SetRequestId(v string) *CreateVulScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVulScanTaskResponseBody) SetTaskId(v string) *CreateVulScanTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVulScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
