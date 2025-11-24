// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSwimLaneGroupResponseBody
	GetRequestId() *string
}

type DeleteSwimLaneGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSwimLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSwimLaneGroupResponseBody) SetRequestId(v string) *DeleteSwimLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSwimLaneGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
