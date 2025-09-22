// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocExtractionResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDocExtractionResultRequest
	GetTaskId() *string
}

type GetDocExtractionResultRequest struct {
	// - Task ID.
	//
	// - taskId is obtained from the SubmitDocExtractionTaskAdvance and SubmitDocExtractionTask interfaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97693b4c-17a8-4198-aa28-798d3c855577mhrv
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDocExtractionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocExtractionResultRequest) SetTaskId(v string) *GetDocExtractionResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDocExtractionResultRequest) Validate() error {
	return dara.Validate(s)
}
