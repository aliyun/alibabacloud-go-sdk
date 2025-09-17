// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *CreateResourceResponseBody
	GetResourceId() *string
	SetResourcePath(v string) *CreateResourceResponseBody
	GetResourcePath() *string
	SetTaskId(v string) *CreateResourceResponseBody
	GetTaskId() *string
}

type CreateResourceResponseBody struct {
	// The ID of a request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cctest
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The path of the resources. The relative resource ID. The resource path contains the complete resource location (parent resource/child resource).
	//
	// example:
	//
	// Instance/r-8vbf5abe31c9c4d4/Account/cctest
	ResourcePath *string `json:"resourcePath,omitempty" xml:"resourcePath,omitempty"`
	// The ID of the asynchronous task. If the operation is asynchronous, this field is returned. In this case, the HTTP status code 202 is returned.
	//
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateResourceResponseBody) GetResourcePath() *string {
	return s.ResourcePath
}

func (s *CreateResourceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceId(v string) *CreateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourcePath(v string) *CreateResourceResponseBody {
	s.ResourcePath = &v
	return s
}

func (s *CreateResourceResponseBody) SetTaskId(v string) *CreateResourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
