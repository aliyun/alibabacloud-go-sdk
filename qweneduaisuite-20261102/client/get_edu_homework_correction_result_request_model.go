// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEduHomeworkCorrectionResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetEduHomeworkCorrectionResultRequest
	GetTaskId() *string
}

type GetEduHomeworkCorrectionResultRequest struct {
	// The task ID.
	//
	// example:
	//
	// b18eabc4-5d84-40b0-ae20-e0eb2f585781
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetEduHomeworkCorrectionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultRequest) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEduHomeworkCorrectionResultRequest) SetTaskId(v string) *GetEduHomeworkCorrectionResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultRequest) Validate() error {
	return dara.Validate(s)
}
