// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceInstancesResponseBody
	GetRequestId() *string
}

type DeleteServiceInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DB140E67-D75F-5585-946E-41D8DC8F4E00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
