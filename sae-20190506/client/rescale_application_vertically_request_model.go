// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationVerticallyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RescaleApplicationVerticallyRequest
	GetAppId() *string
	SetCpu(v string) *RescaleApplicationVerticallyRequest
	GetCpu() *string
	SetDeploy(v bool) *RescaleApplicationVerticallyRequest
	GetDeploy() *bool
	SetDiskSize(v string) *RescaleApplicationVerticallyRequest
	GetDiskSize() *string
	SetMemory(v string) *RescaleApplicationVerticallyRequest
	GetMemory() *string
	SetResourceType(v string) *RescaleApplicationVerticallyRequest
	GetResourceType() *string
	SetVSwitchId(v string) *RescaleApplicationVerticallyRequest
	GetVSwitchId() *string
	SetAutoEnableApplicationScalingRule(v bool) *RescaleApplicationVerticallyRequest
	GetAutoEnableApplicationScalingRule() *bool
	SetMinReadyInstanceRatio(v int32) *RescaleApplicationVerticallyRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *RescaleApplicationVerticallyRequest
	GetMinReadyInstances() *int32
}

type RescaleApplicationVerticallyRequest struct {
	// The app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Target CPU specification. Unit: millicore.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Cpu    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Deploy *bool   `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 20
	DiskSize *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// Target memory specification. Unit: MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2048
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Enable application scale rules automatically.
	//
	// example:
	//
	// true
	AutoEnableApplicationScalingRule *bool `json:"autoEnableApplicationScalingRule,omitempty" xml:"autoEnableApplicationScalingRule,omitempty"`
	// The ratio of minimum ready instances.
	//
	// example:
	//
	// 50
	MinReadyInstanceRatio *int32 `json:"minReadyInstanceRatio,omitempty" xml:"minReadyInstanceRatio,omitempty"`
	// Minimum ready instances.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"minReadyInstances,omitempty" xml:"minReadyInstances,omitempty"`
}

func (s RescaleApplicationVerticallyRequest) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationVerticallyRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyRequest) GetAppId() *string {
	return s.AppId
}

func (s *RescaleApplicationVerticallyRequest) GetCpu() *string {
	return s.Cpu
}

func (s *RescaleApplicationVerticallyRequest) GetDeploy() *bool {
	return s.Deploy
}

func (s *RescaleApplicationVerticallyRequest) GetDiskSize() *string {
	return s.DiskSize
}

func (s *RescaleApplicationVerticallyRequest) GetMemory() *string {
	return s.Memory
}

func (s *RescaleApplicationVerticallyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RescaleApplicationVerticallyRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RescaleApplicationVerticallyRequest) GetAutoEnableApplicationScalingRule() *bool {
	return s.AutoEnableApplicationScalingRule
}

func (s *RescaleApplicationVerticallyRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *RescaleApplicationVerticallyRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *RescaleApplicationVerticallyRequest) SetAppId(v string) *RescaleApplicationVerticallyRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetCpu(v string) *RescaleApplicationVerticallyRequest {
	s.Cpu = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetDeploy(v bool) *RescaleApplicationVerticallyRequest {
	s.Deploy = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetDiskSize(v string) *RescaleApplicationVerticallyRequest {
	s.DiskSize = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetMemory(v string) *RescaleApplicationVerticallyRequest {
	s.Memory = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetResourceType(v string) *RescaleApplicationVerticallyRequest {
	s.ResourceType = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetVSwitchId(v string) *RescaleApplicationVerticallyRequest {
	s.VSwitchId = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetAutoEnableApplicationScalingRule(v bool) *RescaleApplicationVerticallyRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetMinReadyInstanceRatio(v int32) *RescaleApplicationVerticallyRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetMinReadyInstances(v int32) *RescaleApplicationVerticallyRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) Validate() error {
	return dara.Validate(s)
}
