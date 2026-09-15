// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReopenComputeInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *ReopenComputeInstanceRequest
	GetRegionId() *string
}

type ReopenComputeInstanceRequest struct {
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

func (s ReopenComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReopenComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReopenComputeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReopenComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReopenComputeInstanceRequest) SetInstanceId(v string) *ReopenComputeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReopenComputeInstanceRequest) SetRegionId(v string) *ReopenComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ReopenComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
