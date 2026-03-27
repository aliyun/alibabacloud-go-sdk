// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrometheusInstanceResponseBody
	GetRequestId() *string
}

type DeletePrometheusInstanceResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusInstanceResponseBody) SetRequestId(v string) *DeletePrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
