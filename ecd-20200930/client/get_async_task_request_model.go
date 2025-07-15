// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *GetAsyncTaskRequest
	GetAsyncTaskId() *string
	SetCdsId(v string) *GetAsyncTaskRequest
	GetCdsId() *string
}

type GetAsyncTaskRequest struct {
	// The asynchronous task ID. This parameter is not returned if you copy files. This parameter is returned if you copy folders in the backend in an asynchronous manner. You can call the GetAsyncTask operation to obtain the ID and information about an asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 81a8a07a-aec4-4dd5-80da-ae69e482****
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
}

func (s GetAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskRequest) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *GetAsyncTaskRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *GetAsyncTaskRequest) SetAsyncTaskId(v string) *GetAsyncTaskRequest {
	s.AsyncTaskId = &v
	return s
}

func (s *GetAsyncTaskRequest) SetCdsId(v string) *GetAsyncTaskRequest {
	s.CdsId = &v
	return s
}

func (s *GetAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
