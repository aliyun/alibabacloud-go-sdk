// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskCode(v string) *BatchQueryTaskStatusRequest
	GetTaskCode() *string
	SetTaskIds(v []*string) *BatchQueryTaskStatusRequest
	GetTaskIds() []*string
}

type BatchQueryTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// EssayCorrectionTask
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	// example:
	//
	// ["xxxx1","xxxx2"]
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s BatchQueryTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryTaskStatusRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *BatchQueryTaskStatusRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *BatchQueryTaskStatusRequest) SetTaskCode(v string) *BatchQueryTaskStatusRequest {
	s.TaskCode = &v
	return s
}

func (s *BatchQueryTaskStatusRequest) SetTaskIds(v []*string) *BatchQueryTaskStatusRequest {
	s.TaskIds = v
	return s
}

func (s *BatchQueryTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
