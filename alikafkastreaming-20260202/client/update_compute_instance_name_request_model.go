// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateComputeInstanceNameRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateComputeInstanceNameRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateComputeInstanceNameRequest
	GetInstanceName() *string
	SetRegionId(v string) *UpdateComputeInstanceNameRequest
	GetRegionId() *string
}

type UpdateComputeInstanceNameRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateComputeInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeInstanceNameRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateComputeInstanceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateComputeInstanceNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateComputeInstanceNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateComputeInstanceNameRequest) SetClientToken(v string) *UpdateComputeInstanceNameRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateComputeInstanceNameRequest) SetInstanceId(v string) *UpdateComputeInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateComputeInstanceNameRequest) SetInstanceName(v string) *UpdateComputeInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateComputeInstanceNameRequest) SetRegionId(v string) *UpdateComputeInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateComputeInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
