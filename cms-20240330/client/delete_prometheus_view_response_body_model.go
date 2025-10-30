// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrometheusViewResponseBody
	GetRequestId() *string
}

type DeletePrometheusViewResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePrometheusViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusViewResponseBody) SetRequestId(v string) *DeletePrometheusViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusViewResponseBody) Validate() error {
	return dara.Validate(s)
}
