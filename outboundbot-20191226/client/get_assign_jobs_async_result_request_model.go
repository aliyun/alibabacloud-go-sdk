// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssignJobsAsyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *GetAssignJobsAsyncResultRequest
	GetAsyncTaskId() *string
}

type GetAssignJobsAsyncResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dc79b0f9-a781-4305-85e2-d5d56679ae69
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
}

func (s GetAssignJobsAsyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssignJobsAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetAssignJobsAsyncResultRequest) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *GetAssignJobsAsyncResultRequest) SetAsyncTaskId(v string) *GetAssignJobsAsyncResultRequest {
	s.AsyncTaskId = &v
	return s
}

func (s *GetAssignJobsAsyncResultRequest) Validate() error {
	return dara.Validate(s)
}
