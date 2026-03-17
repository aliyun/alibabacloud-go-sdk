// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHealthCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHealthCheckResponseBody
	GetRequestId() *string
}

type DeleteHealthCheckResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7F0B079C-2D0E-4ABF-A970-C079F785A09C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHealthCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHealthCheckResponseBody) SetRequestId(v string) *DeleteHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHealthCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
