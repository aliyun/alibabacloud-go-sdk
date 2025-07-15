// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalAccelerationInstanceAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGlobalAccelerationInstanceAttributesResponseBody
	GetRequestId() *string
}

type ModifyGlobalAccelerationInstanceAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD5BCEE8-F62C-40C2-9AC3-89XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalAccelerationInstanceAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalAccelerationInstanceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponseBody) SetRequestId(v string) *ModifyGlobalAccelerationInstanceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
