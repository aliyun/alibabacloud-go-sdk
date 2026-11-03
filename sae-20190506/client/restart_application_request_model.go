// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RestartApplicationRequest
	GetAppId() *string
	SetAutoEnableApplicationScalingRule(v bool) *RestartApplicationRequest
	GetAutoEnableApplicationScalingRule() *bool
	SetMinReadyInstanceRatio(v int32) *RestartApplicationRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *RestartApplicationRequest
	GetMinReadyInstances() *int32
}

type RestartApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether the elastic scaling policy for the application is automatically enabled. Valid values:
	//
	// - **true**: enables the policy.
	//
	// - **false**: disables the policy.
	//
	// example:
	//
	// false
	AutoEnableApplicationScalingRule *bool `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	// The percentage of the minimum number of ready instances. Valid values:
	//
	//  - **-1**: the initial value, which indicates that the percentage is not used.
	//
	//  - **0 to 100**: the percentage, which is rounded up. For example, if you set this parameter to **50*	- and you have 5 instances, the minimum number of ready instances is 3.
	//
	// > If both **MinReadyInstances*	- and **MinReadyInstanceRatio*	- are specified and the value of **MinReadyInstanceRatio*	- is not **-1**, the **MinReadyInstanceRatio*	- parameter takes precedence. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the system uses **50*	- to calculate the minimum number of ready instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of ready instances. Valid values:
	//
	// - If you set this parameter to **0**, services are interrupted during the application upgrade.
	//
	// - If you set this parameter to **-1**, the system uses the recommended value, which is 25% of the current number of instances. For example, if you have 5 instances, 5 × 25% = 1.25. The value is rounded up to 2.
	//
	// > For each rolling deployment, set the minimum number of ready instances to 1 or greater to ensure service continuity.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
}

func (s RestartApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartApplicationRequest) GoString() string {
	return s.String()
}

func (s *RestartApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RestartApplicationRequest) GetAutoEnableApplicationScalingRule() *bool {
	return s.AutoEnableApplicationScalingRule
}

func (s *RestartApplicationRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *RestartApplicationRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *RestartApplicationRequest) SetAppId(v string) *RestartApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RestartApplicationRequest) SetAutoEnableApplicationScalingRule(v bool) *RestartApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *RestartApplicationRequest) SetMinReadyInstanceRatio(v int32) *RestartApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RestartApplicationRequest) SetMinReadyInstances(v int32) *RestartApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *RestartApplicationRequest) Validate() error {
	return dara.Validate(s)
}
