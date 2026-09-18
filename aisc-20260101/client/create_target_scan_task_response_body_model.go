// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTargetScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateTargetScanTaskResponseBodyData) *CreateTargetScanTaskResponseBody
	GetData() *CreateTargetScanTaskResponseBodyData
	SetRequestId(v string) *CreateTargetScanTaskResponseBody
	GetRequestId() *string
}

type CreateTargetScanTaskResponseBody struct {
	// The creation result, which contains the TaskId of the new scan task.
	Data *CreateTargetScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier of the request, which is used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTargetScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTargetScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTargetScanTaskResponseBody) GetData() *CreateTargetScanTaskResponseBodyData {
	return s.Data
}

func (s *CreateTargetScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTargetScanTaskResponseBody) SetData(v *CreateTargetScanTaskResponseBodyData) *CreateTargetScanTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTargetScanTaskResponseBody) SetRequestId(v string) *CreateTargetScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTargetScanTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTargetScanTaskResponseBodyData struct {
	// The unique identifier of the scan task. The initial task status is PREPARING (asynchronous preparation in progress). You can call ListScanTasksByTarget to query the task status and progress.
	//
	// example:
	//
	// task-abc123def4567
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTargetScanTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTargetScanTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTargetScanTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTargetScanTaskResponseBodyData) SetTaskId(v string) *CreateTargetScanTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateTargetScanTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
