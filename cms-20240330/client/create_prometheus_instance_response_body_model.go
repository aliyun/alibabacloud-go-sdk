// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusInstanceId(v string) *CreatePrometheusInstanceResponseBody
	GetPrometheusInstanceId() *string
	SetRequestId(v string) *CreatePrometheusInstanceResponseBody
	GetRequestId() *string
}

type CreatePrometheusInstanceResponseBody struct {
	// Instance ID.
	//
	// example:
	//
	// rw-abc123
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// ID of the request.
	//
	// example:
	//
	// 264C3E89-BE6E-5F82-A484-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceResponseBody) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *CreatePrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusInstanceResponseBody) SetPrometheusInstanceId(v string) *CreatePrometheusInstanceResponseBody {
	s.PrometheusInstanceId = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) SetRequestId(v string) *CreatePrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
