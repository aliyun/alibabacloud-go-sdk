// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugResourceRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DebugResourceRuleShrinkRequest
	GetInstanceId() *string
	SetMetricInfoShrink(v string) *DebugResourceRuleShrinkRequest
	GetMetricInfoShrink() *string
	SetRegionId(v string) *DebugResourceRuleShrinkRequest
	GetRegionId() *string
}

type DebugResourceRuleShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfoShrink *string `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DebugResourceRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugResourceRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DebugResourceRuleShrinkRequest) GetMetricInfoShrink() *string {
	return s.MetricInfoShrink
}

func (s *DebugResourceRuleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DebugResourceRuleShrinkRequest) SetInstanceId(v string) *DebugResourceRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugResourceRuleShrinkRequest) SetMetricInfoShrink(v string) *DebugResourceRuleShrinkRequest {
	s.MetricInfoShrink = &v
	return s
}

func (s *DebugResourceRuleShrinkRequest) SetRegionId(v string) *DebugResourceRuleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DebugResourceRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
