// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *ListTasksResponseBody
	GetResult() interface{}
}

type ListTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetResult(v interface{}) *ListTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
