// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVLExtractionResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVLExtractionResultRequest
	GetTaskId() *string
}

type GetVLExtractionResultRequest struct {
	// - taskID.
	//
	// - The taskId is obtained from the interfaces SubmitVLExtractionTaskAdvance and SubmitVLExtractionTask.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1436b6f5-ddea-4308-9d1c-60939e5d5ea8
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVLExtractionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultRequest) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVLExtractionResultRequest) SetTaskId(v string) *GetVLExtractionResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetVLExtractionResultRequest) Validate() error {
	return dara.Validate(s)
}
