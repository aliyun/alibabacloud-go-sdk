// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceAutoScalerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBehavior(v *UpdateServiceAutoScalerRequestBehavior) *UpdateServiceAutoScalerRequest
	GetBehavior() *UpdateServiceAutoScalerRequestBehavior
	SetMax(v int32) *UpdateServiceAutoScalerRequest
	GetMax() *int32
	SetMin(v int32) *UpdateServiceAutoScalerRequest
	GetMin() *int32
	SetScaleStrategies(v []*UpdateServiceAutoScalerRequestScaleStrategies) *UpdateServiceAutoScalerRequest
	GetScaleStrategies() []*UpdateServiceAutoScalerRequestScaleStrategies
}

type UpdateServiceAutoScalerRequest struct {
	// The Autoscaler operation.
	Behavior *UpdateServiceAutoScalerRequestBehavior `json:"behavior,omitempty" xml:"behavior,omitempty" type:"Struct"`
	// The maximum number of instances. The value must be greater than that of the min parameter.
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
	// The auto scaling policies.
	//
	// This parameter is required.
	ScaleStrategies []*UpdateServiceAutoScalerRequestScaleStrategies `json:"scaleStrategies,omitempty" xml:"scaleStrategies,omitempty" type:"Repeated"`
}

func (s UpdateServiceAutoScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequest) GetBehavior() *UpdateServiceAutoScalerRequestBehavior {
	return s.Behavior
}

func (s *UpdateServiceAutoScalerRequest) GetMax() *int32 {
	return s.Max
}

func (s *UpdateServiceAutoScalerRequest) GetMin() *int32 {
	return s.Min
}

func (s *UpdateServiceAutoScalerRequest) GetScaleStrategies() []*UpdateServiceAutoScalerRequestScaleStrategies {
	return s.ScaleStrategies
}

func (s *UpdateServiceAutoScalerRequest) SetBehavior(v *UpdateServiceAutoScalerRequestBehavior) *UpdateServiceAutoScalerRequest {
	s.Behavior = v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetMax(v int32) *UpdateServiceAutoScalerRequest {
	s.Max = &v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetMin(v int32) *UpdateServiceAutoScalerRequest {
	s.Min = &v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetScaleStrategies(v []*UpdateServiceAutoScalerRequestScaleStrategies) *UpdateServiceAutoScalerRequest {
	s.ScaleStrategies = v
	return s
}

func (s *UpdateServiceAutoScalerRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceAutoScalerRequestBehavior struct {
	// The operation that reduces the number of instances to 0.
	OnZero *UpdateServiceAutoScalerRequestBehaviorOnZero `json:"onZero,omitempty" xml:"onZero,omitempty" type:"Struct"`
	// The scale-in operation.
	ScaleDown *UpdateServiceAutoScalerRequestBehaviorScaleDown `json:"scaleDown,omitempty" xml:"scaleDown,omitempty" type:"Struct"`
	// The scale-out operation.
	ScaleUp *UpdateServiceAutoScalerRequestBehaviorScaleUp `json:"scaleUp,omitempty" xml:"scaleUp,omitempty" type:"Struct"`
}

func (s UpdateServiceAutoScalerRequestBehavior) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehavior) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehavior) GetOnZero() *UpdateServiceAutoScalerRequestBehaviorOnZero {
	return s.OnZero
}

func (s *UpdateServiceAutoScalerRequestBehavior) GetScaleDown() *UpdateServiceAutoScalerRequestBehaviorScaleDown {
	return s.ScaleDown
}

func (s *UpdateServiceAutoScalerRequestBehavior) GetScaleUp() *UpdateServiceAutoScalerRequestBehaviorScaleUp {
	return s.ScaleUp
}

func (s *UpdateServiceAutoScalerRequestBehavior) SetOnZero(v *UpdateServiceAutoScalerRequestBehaviorOnZero) *UpdateServiceAutoScalerRequestBehavior {
	s.OnZero = v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehavior) SetScaleDown(v *UpdateServiceAutoScalerRequestBehaviorScaleDown) *UpdateServiceAutoScalerRequestBehavior {
	s.ScaleDown = v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehavior) SetScaleUp(v *UpdateServiceAutoScalerRequestBehaviorScaleUp) *UpdateServiceAutoScalerRequestBehavior {
	s.ScaleUp = v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehavior) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceAutoScalerRequestBehaviorOnZero struct {
	// The time window that is required before the number of instances is reduced to 0. Default value: 600. The number of instances can be reduced to 0 only if no request is available or no traffic exists in the specified time window.
	//
	// example:
	//
	// 600
	ScaleDownGracePeriodSeconds *int32 `json:"scaleDownGracePeriodSeconds,omitempty" xml:"scaleDownGracePeriodSeconds,omitempty"`
	// The number of instances that you want to create at a time if the number of instances is scaled out from 0. Default value: 1.
	//
	// example:
	//
	// 1
	ScaleUpActivationReplicas *int32 `json:"scaleUpActivationReplicas,omitempty" xml:"scaleUpActivationReplicas,omitempty"`
}

func (s UpdateServiceAutoScalerRequestBehaviorOnZero) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehaviorOnZero) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) GetScaleDownGracePeriodSeconds() *int32 {
	return s.ScaleDownGracePeriodSeconds
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) GetScaleUpActivationReplicas() *int32 {
	return s.ScaleUpActivationReplicas
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) SetScaleDownGracePeriodSeconds(v int32) *UpdateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleDownGracePeriodSeconds = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) SetScaleUpActivationReplicas(v int32) *UpdateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleUpActivationReplicas = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceAutoScalerRequestBehaviorScaleDown struct {
	// The time window that is required before the scale-in operation is performed. Default value: 300. The scale-in operation can be performed only if the specified metric drops below the threshold in the specified time window.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleDown) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleDown) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleDown) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleDown) SetStabilizationWindowSeconds(v int32) *UpdateServiceAutoScalerRequestBehaviorScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleDown) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceAutoScalerRequestBehaviorScaleUp struct {
	// The time window that is required before the scale-out operation is performed. Default value: 0. The scale-out operation can be performed only if the specified metric exceeds the specified threshold in the specified time window.
	//
	// example:
	//
	// 0
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleUp) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleUp) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleUp) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleUp) SetStabilizationWindowSeconds(v int32) *UpdateServiceAutoScalerRequestBehaviorScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleUp) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceAutoScalerRequestScaleStrategies struct {
	// The name of the metric for triggering auto scaling. Valid values:
	//
	// 	- qps: the queries per second (QPS) for an individual instance.
	//
	// 	- cpu: the CPU utilization.
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
	// 	- If you set metricName to QPS, scale-out is triggered when the average QPS for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to CPU, scale-out is triggered when the average CPU utilization for a single instance is greater than this threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s UpdateServiceAutoScalerRequestScaleStrategies) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestScaleStrategies) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) GetMetricName() *string {
	return s.MetricName
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) GetService() *string {
	return s.Service
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) SetMetricName(v string) *UpdateServiceAutoScalerRequestScaleStrategies {
	s.MetricName = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) SetService(v string) *UpdateServiceAutoScalerRequestScaleStrategies {
	s.Service = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) SetThreshold(v float32) *UpdateServiceAutoScalerRequestScaleStrategies {
	s.Threshold = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) Validate() error {
	return dara.Validate(s)
}
