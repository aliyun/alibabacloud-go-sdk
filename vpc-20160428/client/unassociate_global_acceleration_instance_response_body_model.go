// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateGlobalAccelerationInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateGlobalAccelerationInstanceResponseBody
	GetRequestId() *string
}

type UnassociateGlobalAccelerationInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD5BCEE8-F62C-40C2-9AC3-89XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateGlobalAccelerationInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateGlobalAccelerationInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateGlobalAccelerationInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateGlobalAccelerationInstanceResponseBody) SetRequestId(v string) *UnassociateGlobalAccelerationInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
