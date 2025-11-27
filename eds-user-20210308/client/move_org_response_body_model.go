// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveOrgResponseBody
	GetRequestId() *string
}

type MoveOrgResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveOrgResponseBody) GoString() string {
	return s.String()
}

func (s *MoveOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveOrgResponseBody) SetRequestId(v string) *MoveOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveOrgResponseBody) Validate() error {
	return dara.Validate(s)
}
