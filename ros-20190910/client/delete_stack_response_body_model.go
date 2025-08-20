// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStackResponseBody
	GetRequestId() *string
}

type DeleteStackResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStackResponseBody) SetRequestId(v string) *DeleteStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStackResponseBody) Validate() error {
	return dara.Validate(s)
}
