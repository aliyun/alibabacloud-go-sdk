// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomGroupResponseBody
	GetRequestId() *string
}

type DeleteCustomGroupResponseBody struct {
	// The ID of the request. This parameter is unique.
	//
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomGroupResponseBody) SetRequestId(v string) *DeleteCustomGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
