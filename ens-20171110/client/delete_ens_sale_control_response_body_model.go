// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnsSaleControlResponseBody
	GetRequestId() *string
}

type DeleteEnsSaleControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnsSaleControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnsSaleControlResponseBody) SetRequestId(v string) *DeleteEnsSaleControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnsSaleControlResponseBody) Validate() error {
	return dara.Validate(s)
}
