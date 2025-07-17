// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvDropMetricsRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateEnvDropMetricsRuleRequest
	GetAliyunLang() *string
	SetDropMetrics(v string) *UpdateEnvDropMetricsRuleRequest
	GetDropMetrics() *string
	SetEnvironmentId(v string) *UpdateEnvDropMetricsRuleRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *UpdateEnvDropMetricsRuleRequest
	GetRegionId() *string
}

type UpdateEnvDropMetricsRuleRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// metric_1
	//
	// metric_2
	//
	// metric_3
	DropMetrics *string `json:"DropMetrics,omitempty" xml:"DropMetrics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEnvDropMetricsRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvDropMetricsRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvDropMetricsRuleRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateEnvDropMetricsRuleRequest) GetDropMetrics() *string {
	return s.DropMetrics
}

func (s *UpdateEnvDropMetricsRuleRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateEnvDropMetricsRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnvDropMetricsRuleRequest) SetAliyunLang(v string) *UpdateEnvDropMetricsRuleRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleRequest) SetDropMetrics(v string) *UpdateEnvDropMetricsRuleRequest {
	s.DropMetrics = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleRequest) SetEnvironmentId(v string) *UpdateEnvDropMetricsRuleRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleRequest) SetRegionId(v string) *UpdateEnvDropMetricsRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleRequest) Validate() error {
	return dara.Validate(s)
}
