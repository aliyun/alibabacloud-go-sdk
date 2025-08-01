// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDiagnosisResultRequest
	GetTaskId() *string
}

type GetDiagnosisResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// quzuYl23
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDiagnosisResultRequest) SetTaskId(v string) *GetDiagnosisResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDiagnosisResultRequest) Validate() error {
	return dara.Validate(s)
}
