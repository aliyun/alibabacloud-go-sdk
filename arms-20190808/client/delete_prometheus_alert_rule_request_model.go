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
	SetClusterId(v string) *DeletePrometheusAlertRuleRequest
	GetClusterId() *string
}

type DeletePrometheusAlertRuleRequest struct {
	// The ID of the alert rule. You can call the ListPrometheusAlertRules operation to query the ID of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3888704
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The cluster ID of the Prometheus monitoring alarm rule.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s *DeletePrometheusAlertRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeletePrometheusAlertRuleRequest) SetAlertId(v int64) *DeletePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *DeletePrometheusAlertRuleRequest) SetClusterId(v string) *DeletePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *DeletePrometheusAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
