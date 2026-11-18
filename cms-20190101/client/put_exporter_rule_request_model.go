// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutExporterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescribe(v string) *PutExporterRuleRequest
	GetDescribe() *string
	SetDstNames(v []*string) *PutExporterRuleRequest
	GetDstNames() []*string
	SetMetricName(v string) *PutExporterRuleRequest
	GetMetricName() *string
	SetNamespace(v string) *PutExporterRuleRequest
	GetNamespace() *string
	SetRegionId(v string) *PutExporterRuleRequest
	GetRegionId() *string
	SetRuleName(v string) *PutExporterRuleRequest
	GetRuleName() *string
	SetTargetWindows(v string) *PutExporterRuleRequest
	GetTargetWindows() *string
}

type PutExporterRuleRequest struct {
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// This parameter is required.
	DstNames      []*string `json:"DstNames,omitempty" xml:"DstNames,omitempty" type:"Repeated"`
	MetricName    *string   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace     *string   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName      *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	TargetWindows *string   `json:"TargetWindows,omitempty" xml:"TargetWindows,omitempty"`
}

func (s PutExporterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutExporterRuleRequest) GoString() string {
	return s.String()
}

func (s *PutExporterRuleRequest) GetDescribe() *string {
	return s.Describe
}

func (s *PutExporterRuleRequest) GetDstNames() []*string {
	return s.DstNames
}

func (s *PutExporterRuleRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *PutExporterRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutExporterRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutExporterRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutExporterRuleRequest) GetTargetWindows() *string {
	return s.TargetWindows
}

func (s *PutExporterRuleRequest) SetDescribe(v string) *PutExporterRuleRequest {
	s.Describe = &v
	return s
}

func (s *PutExporterRuleRequest) SetDstNames(v []*string) *PutExporterRuleRequest {
	s.DstNames = v
	return s
}

func (s *PutExporterRuleRequest) SetMetricName(v string) *PutExporterRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutExporterRuleRequest) SetNamespace(v string) *PutExporterRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutExporterRuleRequest) SetRegionId(v string) *PutExporterRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PutExporterRuleRequest) SetRuleName(v string) *PutExporterRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutExporterRuleRequest) SetTargetWindows(v string) *PutExporterRuleRequest {
	s.TargetWindows = &v
	return s
}

func (s *PutExporterRuleRequest) Validate() error {
	return dara.Validate(s)
}
