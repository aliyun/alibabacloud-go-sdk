// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobCuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateComputeJobCuRequest
	GetClientToken() *string
	SetCuLimit(v float64) *UpdateComputeJobCuRequest
	GetCuLimit() *float64
	SetCuReserved(v float64) *UpdateComputeJobCuRequest
	GetCuReserved() *float64
	SetInstanceId(v string) *UpdateComputeJobCuRequest
	GetInstanceId() *string
	SetJobName(v string) *UpdateComputeJobCuRequest
	GetJobName() *string
	SetRegionId(v string) *UpdateComputeJobCuRequest
	GetRegionId() *string
}

type UpdateComputeJobCuRequest struct {
	ClientToken *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CuLimit     *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	CuReserved  *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateComputeJobCuRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobCuRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobCuRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateComputeJobCuRequest) GetCuLimit() *float64 {
	return s.CuLimit
}

func (s *UpdateComputeJobCuRequest) GetCuReserved() *float64 {
	return s.CuReserved
}

func (s *UpdateComputeJobCuRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateComputeJobCuRequest) GetJobName() *string {
	return s.JobName
}

func (s *UpdateComputeJobCuRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateComputeJobCuRequest) SetClientToken(v string) *UpdateComputeJobCuRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateComputeJobCuRequest) SetCuLimit(v float64) *UpdateComputeJobCuRequest {
	s.CuLimit = &v
	return s
}

func (s *UpdateComputeJobCuRequest) SetCuReserved(v float64) *UpdateComputeJobCuRequest {
	s.CuReserved = &v
	return s
}

func (s *UpdateComputeJobCuRequest) SetInstanceId(v string) *UpdateComputeJobCuRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateComputeJobCuRequest) SetJobName(v string) *UpdateComputeJobCuRequest {
	s.JobName = &v
	return s
}

func (s *UpdateComputeJobCuRequest) SetRegionId(v string) *UpdateComputeJobCuRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateComputeJobCuRequest) Validate() error {
	return dara.Validate(s)
}
