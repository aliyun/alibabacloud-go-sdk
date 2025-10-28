// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateApplicationScalingRuleRequest
	GetAppId() *string
	SetScalingBehaviour(v string) *CreateApplicationScalingRuleRequest
	GetScalingBehaviour() *string
	SetScalingRuleEnable(v bool) *CreateApplicationScalingRuleRequest
	GetScalingRuleEnable() *bool
	SetScalingRuleMetric(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleMetric() *string
	SetScalingRuleName(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleName() *string
	SetScalingRuleTimer(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleTimer() *string
	SetScalingRuleTrigger(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleTrigger() *string
	SetScalingRuleType(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleType() *string
}

type CreateApplicationScalingRuleRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplicationlink](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Configure custom elastic behavior, refer to the example for specific data structure.
	//
	// example:
	//
	// {
	//
	//       "scaleUp": {
	//
	//             "stabilizationWindowSeconds": "0",
	//
	//             "selectPolicy": "Max",
	//
	//             "policies": [
	//
	//                   {
	//
	//                         "type": "Pods",
	//
	//                         "value": 5,
	//
	//                         "periodSeconds": 15
	//
	//                   }
	//
	//             ]
	//
	//       },
	//
	//       "scaleDown": {
	//
	//             "stabilizationWindowSeconds": "300",
	//
	//             "selectPolicy": "Max",
	//
	//             "policies": [
	//
	//                   {
	//
	//                         "type": "Percent",
	//
	//                         "value": 200,
	//
	//                         "periodSeconds": 15
	//
	//                   }
	//
	//             ]
	//
	//       }
	//
	// }
	ScalingBehaviour *string `json:"ScalingBehaviour,omitempty" xml:"ScalingBehaviour,omitempty"`
	// Specifies whether to enable the auto scaling policy. Valid values:
	//
	// 	- **true**: enables the auto scaling policy.
	//
	// 	- **false**: disables the auto scaling policy.
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
	// The name of the auto scaling policy. The name must start with a lowercase letter, and can contain lowercase letters, digits, and hyphens (-). The name must be 1 to 32 characters in length.
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
	// The trigger policy for the auto scaling policy. Set the value in the JSON format by using the ScalingRuleTriggerDTO class. For more information, see Additional information about request parameters.
	//
	// example:
	//
	// ScalingRuleTriggerDTO{......}
	ScalingRuleTrigger *string `json:"ScalingRuleTrigger,omitempty" xml:"ScalingRuleTrigger,omitempty"`
	// The type of the auto scaling policy. Set the value to **trigger**.
	//
	// example:
	//
	// trigger
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s CreateApplicationScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationScalingRuleRequest) GetScalingBehaviour() *string {
	return s.ScalingBehaviour
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleEnable() *bool {
	return s.ScalingRuleEnable
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleMetric() *string {
	return s.ScalingRuleMetric
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleTimer() *string {
	return s.ScalingRuleTimer
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleTrigger() *string {
	return s.ScalingRuleTrigger
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *CreateApplicationScalingRuleRequest) SetAppId(v string) *CreateApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingBehaviour(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingBehaviour = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleEnable(v bool) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleEnable = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleMetric(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleMetric = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleName(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleTimer(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleTimer = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleTrigger(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleTrigger = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleType(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
