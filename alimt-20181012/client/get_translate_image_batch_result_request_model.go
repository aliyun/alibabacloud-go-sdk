// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslateImageBatchResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTranslateImageBatchResultRequest
	GetTaskId() *string
}

type GetTranslateImageBatchResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// EEA28E6D-4828-5031-BD8C-8FF1B3216842
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTranslateImageBatchResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateImageBatchResultRequest) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTranslateImageBatchResultRequest) SetTaskId(v string) *GetTranslateImageBatchResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetTranslateImageBatchResultRequest) Validate() error {
	return dara.Validate(s)
}
