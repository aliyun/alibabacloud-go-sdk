// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *DeletePrometheusAlertRuleRequest
	GetAlertId() *int64
}

type DeletePrometheusAlertRuleRequest struct {
	// This parameter is required.
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DeletePrometheusAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *DeletePrometheusAlertRuleRequest) SetAlertId(v int64) *DeletePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *DeletePrometheusAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
