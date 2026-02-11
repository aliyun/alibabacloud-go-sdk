// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrometheusAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePrometheusAlertRuleResponseBody
	GetSuccess() *bool
}

type DeletePrometheusAlertRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePrometheusAlertRuleResponseBody) SetRequestId(v string) *DeletePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) SetSuccess(v bool) *DeletePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
