// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PushResourceRuleShrinkRequest
	GetInstanceId() *string
	SetMetricInfoShrink(v string) *PushResourceRuleShrinkRequest
	GetMetricInfoShrink() *string
}

type PushResourceRuleShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfoShrink *string `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
}

func (s PushResourceRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushResourceRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushResourceRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PushResourceRuleShrinkRequest) GetMetricInfoShrink() *string {
	return s.MetricInfoShrink
}

func (s *PushResourceRuleShrinkRequest) SetInstanceId(v string) *PushResourceRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *PushResourceRuleShrinkRequest) SetMetricInfoShrink(v string) *PushResourceRuleShrinkRequest {
	s.MetricInfoShrink = &v
	return s
}

func (s *PushResourceRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
