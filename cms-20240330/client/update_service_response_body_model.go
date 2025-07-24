// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *UpdateServiceResponseBody
	GetServiceId() *string
}

type UpdateServiceResponseBody struct {
	// example:
	//
	// 123-0F43-23423-AC43-34234
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// em87vd@c2e25bcfe0e21ce0***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceResponseBody) SetServiceId(v string) *UpdateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
