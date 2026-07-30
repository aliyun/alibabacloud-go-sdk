// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBridgeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVirtualBridgeStatusResponseBody
	GetRequestId() *string
}

type ModifyVirtualBridgeStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVirtualBridgeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBridgeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBridgeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVirtualBridgeStatusResponseBody) SetRequestId(v string) *ModifyVirtualBridgeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVirtualBridgeStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
