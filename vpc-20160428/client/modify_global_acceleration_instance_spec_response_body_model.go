// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalAccelerationInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGlobalAccelerationInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyGlobalAccelerationInstanceSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD5BCEE8-F62C-40C2-9AC3-89XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalAccelerationInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalAccelerationInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalAccelerationInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalAccelerationInstanceSpecResponseBody) SetRequestId(v string) *ModifyGlobalAccelerationInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
