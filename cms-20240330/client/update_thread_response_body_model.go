// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateThreadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateThreadResponseBody
	GetRequestId() *string
	SetThreadId(v string) *UpdateThreadResponseBody
	GetThreadId() *string
	SetVersion(v int64) *UpdateThreadResponseBody
	GetVersion() *int64
}

type UpdateThreadResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// thread_id01
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// 0.1.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateThreadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateThreadResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateThreadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateThreadResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *UpdateThreadResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateThreadResponseBody) SetRequestId(v string) *UpdateThreadResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateThreadResponseBody) SetThreadId(v string) *UpdateThreadResponseBody {
	s.ThreadId = &v
	return s
}

func (s *UpdateThreadResponseBody) SetVersion(v int64) *UpdateThreadResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateThreadResponseBody) Validate() error {
	return dara.Validate(s)
}
