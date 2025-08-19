// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveAccountResponseBody
	GetRequestId() *string
}

type MoveAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveAccountResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveAccountResponseBody) SetRequestId(v string) *MoveAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
