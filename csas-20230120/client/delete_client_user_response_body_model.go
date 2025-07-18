// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClientUserResponseBody
	GetRequestId() *string
}

type DeleteClientUserResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientUserResponseBody) SetRequestId(v string) *DeleteClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
