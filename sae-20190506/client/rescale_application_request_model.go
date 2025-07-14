// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RescaleApplicationRequest
	GetAppId() *string
	SetAutoEnableApplicationScalingRule(v bool) *RescaleApplicationRequest
	GetAutoEnableApplicationScalingRule() *bool
	SetMinReadyInstanceRatio(v int32) *RescaleApplicationRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *RescaleApplicationRequest
	GetMinReadyInstances() *int32
	SetReplicas(v int32) *RescaleApplicationRequest
	GetReplicas() *int32
}

type RescaleApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to automatically enable an auto scaling policy for the application. Take note of the following rules:
	//
	// 	- **true**: turns on Logon-free Sharing
	//
	// 	- **false**: turns off Logon-free Sharing
	//
	// example:
	//
	// true
	AutoEnableApplicationScalingRule *bool `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	// The percentage of the minimum number of available instances. Take note of the following rules:
	//
	// 	- If you set the value to **-1**, the minimum number of available instances is not determined based on this parameter. Default value: -1.
	//
	// 	- If you set the value to a number **from 0 to 100**, the minimum number of available instances is calculated by using the following formula: Current number of instances × (Value of MinReadyInstanceRatio × 100%). The value is the nearest integer rounded up from the calculated result. For example, if the percentage is set to **50**% and five instances are available, the minimum number of available instances is 3.
	//
	// > When **MinReadyInstance*	- and **MinReadyInstanceRatio*	- are specified and **MinReadyInstanceRatio*	- is set to a number from 0 to 100, the value of MinReadyInstanceRatio*	- takes precedence.***	- For example, if **MinReadyInstances*	- is set to **5, and **MinReadyInstanceRatio*	- is set to **50**, the minimum number of available instances is set to the nearest integer rounded up from the calculated result of the following formula: Current number of instances × **50%**.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Special values:
	//
	// 	- If you set the value to **0**, business interruptions occur when the application is updated.
	//
	// 	- If you set the value to \\*\\*-1\\*\\*, the minimum number of available instances is automatically set to a system-recommended value. The value is the nearest integer to which the calculated result of the following formula is rounded up: Current number of instances × 25%. For example, if five instances are available, the minimum number of available instances is calculated by using the following formula: 5 × 25% = 1.25. In this case, the minimum number of available instances is 2.
	//
	// > Make sure that at least one instance is available during application deployment and rollback to prevent business interruptions.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The expected number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s RescaleApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RescaleApplicationRequest) GetAutoEnableApplicationScalingRule() *bool {
	return s.AutoEnableApplicationScalingRule
}

func (s *RescaleApplicationRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *RescaleApplicationRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *RescaleApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *RescaleApplicationRequest) SetAppId(v string) *RescaleApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationRequest) SetAutoEnableApplicationScalingRule(v bool) *RescaleApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *RescaleApplicationRequest) SetMinReadyInstanceRatio(v int32) *RescaleApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RescaleApplicationRequest) SetMinReadyInstances(v int32) *RescaleApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *RescaleApplicationRequest) SetReplicas(v int32) *RescaleApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *RescaleApplicationRequest) Validate() error {
	return dara.Validate(s)
}
