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
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
