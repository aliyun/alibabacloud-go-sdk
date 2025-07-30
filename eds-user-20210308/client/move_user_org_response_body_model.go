// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveUserOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveUserOrgResponseBody
	GetRequestId() *string
}

type MoveUserOrgResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveUserOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveUserOrgResponseBody) GoString() string {
	return s.String()
}

func (s *MoveUserOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveUserOrgResponseBody) SetRequestId(v string) *MoveUserOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveUserOrgResponseBody) Validate() error {
	return dara.Validate(s)
}
