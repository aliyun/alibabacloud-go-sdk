// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateQueueResponseBody
	GetRequestId() *string
}

type CreateQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 59B52E2C-0B8E-44EC-A314-D0314A50***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
