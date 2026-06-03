// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobGroupsAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *ListJobGroupsAsyncRequest
	GetAsyncTaskId() *string
}

type ListJobGroupsAsyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6243d904-939d-42ce-a8e4-886a139e77a3
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
}

func (s ListJobGroupsAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncRequest) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncRequest) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ListJobGroupsAsyncRequest) SetAsyncTaskId(v string) *ListJobGroupsAsyncRequest {
	s.AsyncTaskId = &v
	return s
}

func (s *ListJobGroupsAsyncRequest) Validate() error {
	return dara.Validate(s)
}
