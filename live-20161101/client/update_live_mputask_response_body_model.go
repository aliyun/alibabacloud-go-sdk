// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveMPUTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveMPUTaskResponseBody
	GetRequestId() *string
}

type UpdateLiveMPUTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0F72851F-5DC1-1979-9B2C-450040316C3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveMPUTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveMPUTaskResponseBody) SetRequestId(v string) *UpdateLiveMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveMPUTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
