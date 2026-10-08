// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateApplicationScalingRuleRequest
	GetAppId() *string
	SetScalingBehaviour(v string) *UpdateApplicationScalingRuleRequest
	GetScalingBehaviour() *string
	SetScalingRuleEnable(v bool) *UpdateApplicationScalingRuleRequest
	GetScalingRuleEnable() *bool
	SetScalingRuleMetric(v string) *UpdateApplicationScalingRuleRequest
	GetScalingRuleMetric() *string
	SetScalingRuleName(v string) *UpdateApplicationScalingRuleRequest
	GetScalingRuleName() *string
	SetScalingRuleTimer(v string) *UpdateApplicationScalingRuleRequest
	GetScalingRuleTimer() *string
	SetScalingRuleTrigger(v string) *UpdateApplicationScalingRuleRequest
	GetScalingRuleTrigger() *string
	SetScalingRuleType(v string) *UpdateApplicationScalingRuleRequest
	GetScalingRuleType() *string
}

type UpdateApplicationScalingRuleRequest struct {
	// The ID of the application. Call the [ListApplication](https://help.aliyun.com/document_detail/149390.html) operation to obtain this ID.
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configuration of custom scaling behaviors. For more information about the data structure, see the example.
	//
	// example:
	//
	// {"scaleUp":{"stabilizationWindowSeconds":"0","selectPolicy":"Max","policies":[{"type":"Pods","value":5,"periodSeconds":15}]},"scaleDown":{"stabilizationWindowSeconds":"300","selectPolicy":"Max","policies":[{"type":"Percent","value":200,"periodSeconds":15}]}}
	ScalingBehaviour *string `json:"ScalingBehaviour,omitempty" xml:"ScalingBehaviour,omitempty"`
	// The status of the Auto Scaling policy.
	//
	// - **true**: enabled
	//
	// - **false**: disabled
	//
	// example:
	//
	// true
	ScalingRuleEnable *bool `json:"ScalingRuleEnable,omitempty" xml:"ScalingRuleEnable,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	ScalingRuleMetric *string `json:"ScalingRuleMetric,omitempty" xml:"ScalingRuleMetric,omitempty"`
	// The name of the Auto Scaling policy.
	//
	// example:
	//
	// cpu-trigger
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	ScalingRuleTimer *string `json:"ScalingRuleTimer,omitempty" xml:"ScalingRuleTimer,omitempty"`
	// The trigger policy, which is a JSON string of a ScalingRuleTriggerDTO object. For more information about the format, see the Additional information about request parameters section.
	//
	// example:
	//
	// ScalingRuleTriggerDTO{......}
	ScalingRuleTrigger *string `json:"ScalingRuleTrigger,omitempty" xml:"ScalingRuleTrigger,omitempty"`
	// The type of the Auto Scaling policy. Only the following type is supported:
	//
	// - trigger: a trigger-based policy.
	//
	// example:
	//
	// trigger
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s UpdateApplicationScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingBehaviour() *string {
	return s.ScalingBehaviour
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingRuleEnable() *bool {
	return s.ScalingRuleEnable
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingRuleMetric() *string {
	return s.ScalingRuleMetric
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingRuleTimer() *string {
	return s.ScalingRuleTimer
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingRuleTrigger() *string {
	return s.ScalingRuleTrigger
}

func (s *UpdateApplicationScalingRuleRequest) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *UpdateApplicationScalingRuleRequest) SetAppId(v string) *UpdateApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingBehaviour(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingBehaviour = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleEnable(v bool) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleEnable = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleMetric(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleMetric = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleName(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleTimer(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleTimer = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleTrigger(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleTrigger = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleType(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
