// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceTaskResponseBody
	GetRequestId() *string
}

type DeleteServiceTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceTaskResponseBody) SetRequestId(v string) *DeleteServiceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
