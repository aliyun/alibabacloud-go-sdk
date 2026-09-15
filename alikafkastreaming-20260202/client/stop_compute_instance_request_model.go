// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopComputeInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *StopComputeInstanceRequest
	GetRegionId() *string
}

type StopComputeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alikafka_streaming-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopComputeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopComputeInstanceRequest) SetInstanceId(v string) *StopComputeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopComputeInstanceRequest) SetRegionId(v string) *StopComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
