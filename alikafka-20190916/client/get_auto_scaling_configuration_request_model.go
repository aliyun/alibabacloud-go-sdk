// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAutoScalingConfigurationRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetAutoScalingConfigurationRequest
	GetRegionId() *string
}

type GetAutoScalingConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAutoScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoScalingConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAutoScalingConfigurationRequest) SetInstanceId(v string) *GetAutoScalingConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigurationRequest) SetRegionId(v string) *GetAutoScalingConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *GetAutoScalingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
