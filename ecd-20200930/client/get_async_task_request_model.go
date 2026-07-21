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
	// The asynchronous task ID. When you call the [CopyCdsFile](https://help.aliyun.com/document_detail/2247626.html) operation to copy a folder, this field is returned because the copy is performed asynchronously in the background. Call this operation and pass in the asynchronous task ID to retrieve the task details.
	//
	// This parameter is required.
	//
	// example:
	//
	// 81a8a07a-aec4-4dd5-80da-ae69e482****
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// The enterprise network disk ID.
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
