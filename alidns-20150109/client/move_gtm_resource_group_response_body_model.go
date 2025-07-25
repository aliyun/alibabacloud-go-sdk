// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGtmResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveGtmResourceGroupResponseBody
	GetRequestId() *string
}

type MoveGtmResourceGroupResponseBody struct {
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveGtmResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveGtmResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveGtmResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveGtmResourceGroupResponseBody) SetRequestId(v string) *MoveGtmResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveGtmResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
