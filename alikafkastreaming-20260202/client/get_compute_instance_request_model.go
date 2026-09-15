// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetComputeInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetComputeInstanceRequest
	GetRegionId() *string
}

type GetComputeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alikafka_streaming-cn-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetComputeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeInstanceRequest) SetInstanceId(v string) *GetComputeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetComputeInstanceRequest) SetRegionId(v string) *GetComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
