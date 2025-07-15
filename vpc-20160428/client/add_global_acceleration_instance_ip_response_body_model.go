// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGlobalAccelerationInstanceIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddGlobalAccelerationInstanceIpResponseBody
	GetRequestId() *string
}

type AddGlobalAccelerationInstanceIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 01FDDD49-C4B7-4D2A-A8E5-A93915C450A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGlobalAccelerationInstanceIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGlobalAccelerationInstanceIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddGlobalAccelerationInstanceIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGlobalAccelerationInstanceIpResponseBody) SetRequestId(v string) *AddGlobalAccelerationInstanceIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpResponseBody) Validate() error {
	return dara.Validate(s)
}
