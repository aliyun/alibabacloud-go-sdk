// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateQueueResponseBody
	GetName() *string
	SetRequestId(v string) *CreateQueueResponseBody
	GetRequestId() *string
}

type CreateQueueResponseBody struct {
	// The name of the created queue.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueueResponseBody) SetName(v string) *CreateQueueResponseBody {
	s.Name = &v
	return s
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
