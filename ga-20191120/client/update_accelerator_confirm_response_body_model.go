// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorConfirmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAcceleratorConfirmResponseBody
	GetRequestId() *string
}

type UpdateAcceleratorConfirmResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorConfirmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorConfirmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAcceleratorConfirmResponseBody) SetRequestId(v string) *UpdateAcceleratorConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAcceleratorConfirmResponseBody) Validate() error {
	return dara.Validate(s)
}
