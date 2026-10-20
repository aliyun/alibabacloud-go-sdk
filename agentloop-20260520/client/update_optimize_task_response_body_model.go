// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOptimizeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOptimizeTaskResponseBody
	GetRequestId() *string
}

type UpdateOptimizeTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateOptimizeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOptimizeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOptimizeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOptimizeTaskResponseBody) SetRequestId(v string) *UpdateOptimizeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOptimizeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
