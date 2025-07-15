// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateInstanceConfigRequest
	GetConfig() *string
	SetInstanceId(v string) *UpdateInstanceConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *UpdateInstanceConfigRequest
	GetRegionId() *string
}

type UpdateInstanceConfigRequest struct {
	// The configurations that you want to update for the ApsaraMQ for Kafka instance. The value must be a valid JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"kafka.log.retention.hours":"33"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceConfigRequest) SetConfig(v string) *UpdateInstanceConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateInstanceConfigRequest) SetInstanceId(v string) *UpdateInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceConfigRequest) SetRegionId(v string) *UpdateInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
