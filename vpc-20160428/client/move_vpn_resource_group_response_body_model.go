// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveVpnResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveVpnResourceGroupResponseBody
	GetRequestId() *string
}

type MoveVpnResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89ED47AF-3319-566E-A5F5-94E3F47F54A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveVpnResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveVpnResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveVpnResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveVpnResourceGroupResponseBody) SetRequestId(v string) *MoveVpnResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveVpnResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
