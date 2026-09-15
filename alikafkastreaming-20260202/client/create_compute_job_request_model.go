// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCuLimit(v float64) *CreateComputeJobRequest
	GetCuLimit() *float64
	SetCuReserved(v float64) *CreateComputeJobRequest
	GetCuReserved() *float64
	SetInstanceId(v string) *CreateComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *CreateComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *CreateComputeJobRequest
	GetRegionId() *string
	SetRemark(v string) *CreateComputeJobRequest
	GetRemark() *string
}

type CreateComputeJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	CuLimit *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	CuReserved *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
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
	// order_enrichment
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 订单流实时清洗
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeJobRequest) GetCuLimit() *float64 {
	return s.CuLimit
}

func (s *CreateComputeJobRequest) GetCuReserved() *float64 {
	return s.CuReserved
}

func (s *CreateComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateComputeJobRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateComputeJobRequest) SetCuLimit(v float64) *CreateComputeJobRequest {
	s.CuLimit = &v
	return s
}

func (s *CreateComputeJobRequest) SetCuReserved(v float64) *CreateComputeJobRequest {
	s.CuReserved = &v
	return s
}

func (s *CreateComputeJobRequest) SetInstanceId(v string) *CreateComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateComputeJobRequest) SetJobName(v string) *CreateComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateComputeJobRequest) SetRegionId(v string) *CreateComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateComputeJobRequest) SetRemark(v string) *CreateComputeJobRequest {
	s.Remark = &v
	return s
}

func (s *CreateComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
