// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusViewId(v string) *UpdatePrometheusViewResponseBody
	GetPrometheusViewId() *string
	SetRequestId(v string) *UpdatePrometheusViewResponseBody
	GetRequestId() *string
}

type UpdatePrometheusViewResponseBody struct {
	// example:
	//
	// rw-xxxxxx
	PrometheusViewId *string `json:"prometheusViewId,omitempty" xml:"prometheusViewId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 64D964F5-76C7-19A2-9399-457744AB3619
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePrometheusViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusViewResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusViewResponseBody) GetPrometheusViewId() *string {
	return s.PrometheusViewId
}

func (s *UpdatePrometheusViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusViewResponseBody) SetPrometheusViewId(v string) *UpdatePrometheusViewResponseBody {
	s.PrometheusViewId = &v
	return s
}

func (s *UpdatePrometheusViewResponseBody) SetRequestId(v string) *UpdatePrometheusViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusViewResponseBody) Validate() error {
	return dara.Validate(s)
}
