// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrometheusAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *DescribePrometheusAlertRuleRequest
	GetAlertId() *int64
	SetClusterId(v string) *DescribePrometheusAlertRuleRequest
	GetClusterId() *string
}

type DescribePrometheusAlertRuleRequest struct {
	// The ID of the alert rule. You can call the ListPrometheusAlertRules operation to query the ID of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3888704
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribePrometheusAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *DescribePrometheusAlertRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePrometheusAlertRuleRequest) SetAlertId(v int64) *DescribePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *DescribePrometheusAlertRuleRequest) SetClusterId(v string) *DescribePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribePrometheusAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
