// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHealthCheckTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHealthCheckTemplatesResponseBody
	GetRequestId() *string
}

type DeleteHealthCheckTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHealthCheckTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHealthCheckTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHealthCheckTemplatesResponseBody) SetRequestId(v string) *DeleteHealthCheckTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHealthCheckTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
