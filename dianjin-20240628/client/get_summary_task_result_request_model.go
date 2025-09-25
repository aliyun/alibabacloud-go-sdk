// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetSummaryTaskResultRequest
	GetTaskId() *string
}

type GetSummaryTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetSummaryTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSummaryTaskResultRequest) SetTaskId(v string) *GetSummaryTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetSummaryTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
