// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusViewId(v string) *CreatePrometheusViewResponseBody
	GetPrometheusViewId() *string
	SetRequestId(v string) *CreatePrometheusViewResponseBody
	GetRequestId() *string
}

type CreatePrometheusViewResponseBody struct {
	// example:
	//
	// cd5237f7dbd574cf9bbd648ff9efb16cd
	PrometheusViewId *string `json:"prometheusViewId,omitempty" xml:"prometheusViewId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePrometheusViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusViewResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusViewResponseBody) GetPrometheusViewId() *string {
	return s.PrometheusViewId
}

func (s *CreatePrometheusViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusViewResponseBody) SetPrometheusViewId(v string) *CreatePrometheusViewResponseBody {
	s.PrometheusViewId = &v
	return s
}

func (s *CreatePrometheusViewResponseBody) SetRequestId(v string) *CreatePrometheusViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusViewResponseBody) Validate() error {
	return dara.Validate(s)
}
