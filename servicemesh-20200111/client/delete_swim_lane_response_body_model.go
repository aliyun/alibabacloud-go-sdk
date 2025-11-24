// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSwimLaneResponseBody
	GetRequestId() *string
}

type DeleteSwimLaneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSwimLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimLaneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSwimLaneResponseBody) SetRequestId(v string) *DeleteSwimLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSwimLaneResponseBody) Validate() error {
	return dara.Validate(s)
}
