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
}

type DescribePrometheusAlertRuleRequest struct {
	// This parameter is required.
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
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

func (s *DescribePrometheusAlertRuleRequest) SetAlertId(v int64) *DescribePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *DescribePrometheusAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
