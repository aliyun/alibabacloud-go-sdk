// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeregisterResourceTypeResponseBody
	GetRequestId() *string
}

type DeregisterResourceTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeregisterResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeregisterResourceTypeResponseBody) SetRequestId(v string) *DeregisterResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeregisterResourceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
