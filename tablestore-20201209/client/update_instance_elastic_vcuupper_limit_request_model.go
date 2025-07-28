// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceElasticVCUUpperLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticVCUUpperLimit(v float32) *UpdateInstanceElasticVCUUpperLimitRequest
	GetElasticVCUUpperLimit() *float32
	SetInstanceName(v string) *UpdateInstanceElasticVCUUpperLimitRequest
	GetInstanceName() *string
}

type UpdateInstanceElasticVCUUpperLimitRequest struct {
	// The upper limit for the VCUs of the instance.
	//
	// >  Valid values of the upper limit for the VCUs of an instance: **Number of reserved VCUs+0.1 to 2000**. You can upgrade or downgrade configurations to modify the number of reserved VCUs by increments or decrements of 1. You can dynamically modify the upper limit for the VCUs of an instance by increments or decrements of 0.1
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	ElasticVCUUpperLimit *float32 `json:"ElasticVCUUpperLimit,omitempty" xml:"ElasticVCUUpperLimit,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// workshop-bj-ots1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceElasticVCUUpperLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceElasticVCUUpperLimitRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) GetElasticVCUUpperLimit() *float32 {
	return s.ElasticVCUUpperLimit
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) SetElasticVCUUpperLimit(v float32) *UpdateInstanceElasticVCUUpperLimitRequest {
	s.ElasticVCUUpperLimit = &v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) SetInstanceName(v string) *UpdateInstanceElasticVCUUpperLimitRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) Validate() error {
	return dara.Validate(s)
}
