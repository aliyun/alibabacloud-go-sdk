// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalyzeTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDocumentAnalyzeTaskStatusRequest
	GetTaskId() *string
}

type GetDocumentAnalyzeTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocumentAnalyzeTaskStatusRequest) SetTaskId(v string) *GetDocumentAnalyzeTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
