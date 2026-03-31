// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoginProfileResponseBody
	GetRequestId() *string
}

type UpdateLoginProfileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoginProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoginProfileResponseBody) SetRequestId(v string) *UpdateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoginProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
