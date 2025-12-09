// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveResourceGroupResponseBody
	GetRequestId() *string
}

type MoveResourceGroupResponseBody struct {
	// example:
	//
	// 5BCD2F6C-7A9D-47C1-8588-2CC6A4E0BE5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
