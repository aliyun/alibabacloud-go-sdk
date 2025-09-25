// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityCheckTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetQualityCheckTaskResultRequest
	GetTaskId() *string
}

type GetQualityCheckTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetQualityCheckTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetQualityCheckTaskResultRequest) SetTaskId(v string) *GetQualityCheckTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetQualityCheckTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
