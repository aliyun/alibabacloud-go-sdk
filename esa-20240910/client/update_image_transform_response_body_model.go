// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTransformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateImageTransformResponseBody
	GetRequestId() *string
}

type UpdateImageTransformResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateImageTransformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageTransformResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageTransformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageTransformResponseBody) SetRequestId(v string) *UpdateImageTransformResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageTransformResponseBody) Validate() error {
	return dara.Validate(s)
}
