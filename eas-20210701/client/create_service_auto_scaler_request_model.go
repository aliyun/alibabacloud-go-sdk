// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAutoScalerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBehavior(v *CreateServiceAutoScalerRequestBehavior) *CreateServiceAutoScalerRequest
	GetBehavior() *CreateServiceAutoScalerRequestBehavior
	SetMax(v int32) *CreateServiceAutoScalerRequest
	GetMax() *int32
	SetMin(v int32) *CreateServiceAutoScalerRequest
	GetMin() *int32
	SetScaleStrategies(v []*CreateServiceAutoScalerRequestScaleStrategies) *CreateServiceAutoScalerRequest
	GetScaleStrategies() []*CreateServiceAutoScalerRequestScaleStrategies
}

type CreateServiceAutoScalerRequest struct {
	// The behavior during auto-scaling.
	Behavior *CreateServiceAutoScalerRequestBehavior `json:"behavior,omitempty" xml:"behavior,omitempty" type:"Struct"`
	// The maximum number of instances. The value must be greater than the minimum number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// The minimum number of instances. The value must be greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// The list of scaling policies.
	//
	// This parameter is required.
	ScaleStrategies []*CreateServiceAutoScalerRequestScaleStrategies `json:"scaleStrategies,omitempty" xml:"scaleStrategies,omitempty" type:"Repeated"`
}

func (s CreateServiceAutoScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequest) GetBehavior() *CreateServiceAutoScalerRequestBehavior {
	return s.Behavior
}

func (s *CreateServiceAutoScalerRequest) GetMax() *int32 {
	return s.Max
}

func (s *CreateServiceAutoScalerRequest) GetMin() *int32 {
	return s.Min
}

func (s *CreateServiceAutoScalerRequest) GetScaleStrategies() []*CreateServiceAutoScalerRequestScaleStrategies {
	return s.ScaleStrategies
}

func (s *CreateServiceAutoScalerRequest) SetBehavior(v *CreateServiceAutoScalerRequestBehavior) *CreateServiceAutoScalerRequest {
	s.Behavior = v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetMax(v int32) *CreateServiceAutoScalerRequest {
	s.Max = &v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetMin(v int32) *CreateServiceAutoScalerRequest {
	s.Min = &v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetScaleStrategies(v []*CreateServiceAutoScalerRequestScaleStrategies) *CreateServiceAutoScalerRequest {
	s.ScaleStrategies = v
	return s
}

func (s *CreateServiceAutoScalerRequest) Validate() error {
	if s.Behavior != nil {
		if err := s.Behavior.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleStrategies != nil {
		for _, item := range s.ScaleStrategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServiceAutoScalerRequestBehavior struct {
	// The behavior when scaling in to zero.
	OnZero *CreateServiceAutoScalerRequestBehaviorOnZero `json:"onZero,omitempty" xml:"onZero,omitempty" type:"Struct"`
	// The behavior during scale-in.
	ScaleDown *CreateServiceAutoScalerRequestBehaviorScaleDown `json:"scaleDown,omitempty" xml:"scaleDown,omitempty" type:"Struct"`
	// The behavior during scale-out.
	ScaleUp *CreateServiceAutoScalerRequestBehaviorScaleUp `json:"scaleUp,omitempty" xml:"scaleUp,omitempty" type:"Struct"`
}

func (s CreateServiceAutoScalerRequestBehavior) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehavior) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehavior) GetOnZero() *CreateServiceAutoScalerRequestBehaviorOnZero {
	return s.OnZero
}

func (s *CreateServiceAutoScalerRequestBehavior) GetScaleDown() *CreateServiceAutoScalerRequestBehaviorScaleDown {
	return s.ScaleDown
}

func (s *CreateServiceAutoScalerRequestBehavior) GetScaleUp() *CreateServiceAutoScalerRequestBehaviorScaleUp {
	return s.ScaleUp
}

func (s *CreateServiceAutoScalerRequestBehavior) SetOnZero(v *CreateServiceAutoScalerRequestBehaviorOnZero) *CreateServiceAutoScalerRequestBehavior {
	s.OnZero = v
	return s
}

func (s *CreateServiceAutoScalerRequestBehavior) SetScaleDown(v *CreateServiceAutoScalerRequestBehaviorScaleDown) *CreateServiceAutoScalerRequestBehavior {
	s.ScaleDown = v
	return s
}

func (s *CreateServiceAutoScalerRequestBehavior) SetScaleUp(v *CreateServiceAutoScalerRequestBehaviorScaleUp) *CreateServiceAutoScalerRequestBehavior {
	s.ScaleUp = v
	return s
}

func (s *CreateServiceAutoScalerRequestBehavior) Validate() error {
	if s.OnZero != nil {
		if err := s.OnZero.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleDown != nil {
		if err := s.ScaleDown.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleUp != nil {
		if err := s.ScaleUp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceAutoScalerRequestBehaviorOnZero struct {
	// The grace period before scaling in to zero. When the metrics meet the conditions for scaling in to zero, the scale-in to zero is executed after the grace period elapses. Default value: 600.
	//
	// example:
	//
	// 600
	ScaleDownGracePeriodSeconds *int32 `json:"scaleDownGracePeriodSeconds,omitempty" xml:"scaleDownGracePeriodSeconds,omitempty"`
	// The number of instances created at a time when scaling out from zero. Default value: 1.
	//
	// example:
	//
	// 1
	ScaleUpActivationReplicas *int32 `json:"scaleUpActivationReplicas,omitempty" xml:"scaleUpActivationReplicas,omitempty"`
}

func (s CreateServiceAutoScalerRequestBehaviorOnZero) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehaviorOnZero) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) GetScaleDownGracePeriodSeconds() *int32 {
	return s.ScaleDownGracePeriodSeconds
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) GetScaleUpActivationReplicas() *int32 {
	return s.ScaleUpActivationReplicas
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) SetScaleDownGracePeriodSeconds(v int32) *CreateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleDownGracePeriodSeconds = &v
	return s
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) SetScaleUpActivationReplicas(v int32) *CreateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleUpActivationReplicas = &v
	return s
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) Validate() error {
	return dara.Validate(s)
}

type CreateServiceAutoScalerRequestBehaviorScaleDown struct {
	// The stabilization window period for scale-in. Default value: 300. Scale-in is executed only when the metrics do not meet the specified conditions over the past period of time.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s CreateServiceAutoScalerRequestBehaviorScaleDown) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehaviorScaleDown) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleDown) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleDown) SetStabilizationWindowSeconds(v int32) *CreateServiceAutoScalerRequestBehaviorScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleDown) Validate() error {
	return dara.Validate(s)
}

type CreateServiceAutoScalerRequestBehaviorScaleUp struct {
	// The stabilization window period for scale-out. Default value: 0. Scale-out is executed only when the metrics meet the specified conditions over the past period of time.
	//
	// example:
	//
	// 0
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s CreateServiceAutoScalerRequestBehaviorScaleUp) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehaviorScaleUp) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleUp) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleUp) SetStabilizationWindowSeconds(v int32) *CreateServiceAutoScalerRequestBehaviorScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleUp) Validate() error {
	return dara.Validate(s)
}

type CreateServiceAutoScalerRequestScaleStrategies struct {
	// The name of the metric. Valid values:
	//
	// - qps: the queries per second (QPS) threshold per instance
	//
	// - cpu: the CPU utilization threshold
	//
	// - gpu[util]: the GPU utilization threshold
	//
	// This parameter is required.
	//
	// example:
	//
	// qps
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service to which the metric belongs. If this parameter is not specified, the current service is used by default.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The threshold that triggers auto-scaling.
	//
	// - If the metric name is set to qps, scale-out is triggered when the average QPS per instance exceeds this value.
	//
	// - If the metric name is set to cpu, scale-out is triggered when the average CPU utilization per instance exceeds this value.
	//
	// - If the metric name is set to gpu[util], scale-out is triggered when the average GPU utilization per instance exceeds this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s CreateServiceAutoScalerRequestScaleStrategies) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerRequestScaleStrategies) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) GetService() *string {
	return s.Service
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) SetMetricName(v string) *CreateServiceAutoScalerRequestScaleStrategies {
	s.MetricName = &v
	return s
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) SetService(v string) *CreateServiceAutoScalerRequestScaleStrategies {
	s.Service = &v
	return s
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) SetThreshold(v float32) *CreateServiceAutoScalerRequestScaleStrategies {
	s.Threshold = &v
	return s
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) Validate() error {
	return dara.Validate(s)
}
