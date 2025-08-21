// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAsyncTaskResponseBody
	GetRequestId() *string
}

type CreateAsyncTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAsyncTaskResponseBody) SetRequestId(v string) *CreateAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
