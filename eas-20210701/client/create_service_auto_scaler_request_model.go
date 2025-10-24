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
	// The Autoscaler operation.
	Behavior *CreateServiceAutoScalerRequestBehavior `json:"behavior,omitempty" xml:"behavior,omitempty" type:"Struct"`
	// The maximum number of instances in the service. The value of max must be greater than the value of min.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// The minimum number of instances in the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
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
	// The operation that reduces the number of instances to 0.
	OnZero *CreateServiceAutoScalerRequestBehaviorOnZero `json:"onZero,omitempty" xml:"onZero,omitempty" type:"Struct"`
	// The scale-in operation.
	ScaleDown *CreateServiceAutoScalerRequestBehaviorScaleDown `json:"scaleDown,omitempty" xml:"scaleDown,omitempty" type:"Struct"`
	// The scale-out operation.
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
	// The time window that is required before the number of instances is reduced to 0. The number of instances can be reduced to 0 only if no request is available or no traffic exists in the specified time window. Default value: 600.
	//
	// example:
	//
	// 600
	ScaleDownGracePeriodSeconds *int32 `json:"scaleDownGracePeriodSeconds,omitempty" xml:"scaleDownGracePeriodSeconds,omitempty"`
	// The number of instances that you want to create at a time if the number of instances is 0. Default value: 1.
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
	// The time window that is required before the scale-in operation is performed. The scale-in operation can be performed only if the specified metric drops below the specified threshold in the specified time window. Default value: 300.
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
	// The time window that is required before the scale-out operation is performed. The scale-out operation can be performed only if the specified metric exceeds the specified threshold in the specified time window. Default value: 0.
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
	// The name of the metric for triggering auto scaling. Valid values:
	//
	// 	- qps: the queries per second (qps) for an individual instance.
	//
	// 	- cpu: the cpu utilization.
	//
	// 	- gpu[util]: gpu utilization.
	//
	// This parameter is required.
	//
	// example:
	//
	// qps
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The threshold of the metric that triggers auto scaling.
	//
	// 	- If you set metricName to qps, scale-out is triggered when the average qps for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to cpu, scale-out is triggered when the average cpu utilization for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to gpu, scale-out is triggered when the average gpu utilization for a single instance is greater than this threshold.
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
