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
	SetOrderId(v string) *GetComputeInstanceRequest
	GetOrderId() *string
	SetRegionId(v string) *GetComputeInstanceRequest
	GetRegionId() *string
}

type GetComputeInstanceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// This parameter is required.
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

func (s *GetComputeInstanceRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *GetComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeInstanceRequest) SetInstanceId(v string) *GetComputeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetComputeInstanceRequest) SetOrderId(v string) *GetComputeInstanceRequest {
	s.OrderId = &v
	return s
}

func (s *GetComputeInstanceRequest) SetRegionId(v string) *GetComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
