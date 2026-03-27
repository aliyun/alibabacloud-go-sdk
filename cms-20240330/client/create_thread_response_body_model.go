// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateThreadResponseBody
	GetRequestId() *string
	SetThreadId(v string) *CreateThreadResponseBody
	GetThreadId() *string
}

type CreateThreadResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// thread_id01
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateThreadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThreadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateThreadResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateThreadResponseBody) SetRequestId(v string) *CreateThreadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThreadResponseBody) SetThreadId(v string) *CreateThreadResponseBody {
	s.ThreadId = &v
	return s
}

func (s *CreateThreadResponseBody) Validate() error {
	return dara.Validate(s)
}
