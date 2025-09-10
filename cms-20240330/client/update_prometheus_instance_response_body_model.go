// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusInstanceId(v string) *UpdatePrometheusInstanceResponseBody
	GetPrometheusInstanceId() *string
	SetRequestId(v string) *UpdatePrometheusInstanceResponseBody
	GetRequestId() *string
}

type UpdatePrometheusInstanceResponseBody struct {
	// example:
	//
	// rw-abc123
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1E92F783-E057-58F1-BD5C-92DED088E7A5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusInstanceResponseBody) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *UpdatePrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusInstanceResponseBody) SetPrometheusInstanceId(v string) *UpdatePrometheusInstanceResponseBody {
	s.PrometheusInstanceId = &v
	return s
}

func (s *UpdatePrometheusInstanceResponseBody) SetRequestId(v string) *UpdatePrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
