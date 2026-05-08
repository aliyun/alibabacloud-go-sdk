// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateProductImageResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateProductImageResponseBody
	GetTaskId() *string
}

type CreateProductImageResponseBody struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateProductImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateProductImageResponseBody) SetRequestId(v string) *CreateProductImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductImageResponseBody) SetTaskId(v string) *CreateProductImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateProductImageResponseBody) Validate() error {
	return dara.Validate(s)
}
