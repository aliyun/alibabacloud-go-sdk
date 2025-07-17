// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRumFileStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRumFileStatusResponseBody
	GetRequestId() *string
}

type UpdateRumFileStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 626037F5-FDEB-45B0-804C-B3C92797****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRumFileStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRumFileStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRumFileStatusResponseBody) SetRequestId(v string) *UpdateRumFileStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRumFileStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
