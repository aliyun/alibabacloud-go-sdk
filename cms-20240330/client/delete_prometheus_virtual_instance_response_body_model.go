// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusVirtualInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrometheusVirtualInstanceResponseBody
	GetRequestId() *string
}

type DeletePrometheusVirtualInstanceResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePrometheusVirtualInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusVirtualInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusVirtualInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusVirtualInstanceResponseBody) SetRequestId(v string) *DeletePrometheusVirtualInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusVirtualInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
