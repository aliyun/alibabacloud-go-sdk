// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAICInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentVar(v string) *ModifyAICInstanceTypeRequest
	GetEnvironmentVar() *string
	SetFrequency(v int64) *ModifyAICInstanceTypeRequest
	GetFrequency() *int64
	SetImageId(v string) *ModifyAICInstanceTypeRequest
	GetImageId() *string
	SetInstanceId(v string) *ModifyAICInstanceTypeRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyAICInstanceTypeRequest
	GetInstanceType() *string
	SetResolution(v string) *ModifyAICInstanceTypeRequest
	GetResolution() *string
}

type ModifyAICInstanceTypeRequest struct {
	// example:
	//
	// [object Object]
	EnvironmentVar *string `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty"`
	// example:
	//
	// 30
	Frequency *int64 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourImage ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aic-xxxx-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aic.cf52r.c1.np
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 720*1280
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s ModifyAICInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAICInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAICInstanceTypeRequest) GetEnvironmentVar() *string {
	return s.EnvironmentVar
}

func (s *ModifyAICInstanceTypeRequest) GetFrequency() *int64 {
	return s.Frequency
}

func (s *ModifyAICInstanceTypeRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyAICInstanceTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAICInstanceTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyAICInstanceTypeRequest) GetResolution() *string {
	return s.Resolution
}

func (s *ModifyAICInstanceTypeRequest) SetEnvironmentVar(v string) *ModifyAICInstanceTypeRequest {
	s.EnvironmentVar = &v
	return s
}

func (s *ModifyAICInstanceTypeRequest) SetFrequency(v int64) *ModifyAICInstanceTypeRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyAICInstanceTypeRequest) SetImageId(v string) *ModifyAICInstanceTypeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyAICInstanceTypeRequest) SetInstanceId(v string) *ModifyAICInstanceTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAICInstanceTypeRequest) SetInstanceType(v string) *ModifyAICInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyAICInstanceTypeRequest) SetResolution(v string) *ModifyAICInstanceTypeRequest {
	s.Resolution = &v
	return s
}

func (s *ModifyAICInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
