// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuCount(v int32) *CreateResourceControlRequest
	GetCpuCount() *int32
	SetDBClusterId(v string) *CreateResourceControlRequest
	GetDBClusterId() *string
	SetMaxCpu(v int32) *CreateResourceControlRequest
	GetMaxCpu() *int32
	SetRegionId(v string) *CreateResourceControlRequest
	GetRegionId() *string
	SetResourceControlName(v string) *CreateResourceControlRequest
	GetResourceControlName() *string
}

type CreateResourceControlRequest struct {
	// The maximum number of CPU cores that the resource control rule can use. The minimum value is 1. The maximum value is determined by the cluster kernel parameter resource_control_cpu_count_limit. You must specify one and only one of this parameter and MaxCpu.
	//
	// example:
	//
	// 4
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The PolarDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maximum CPU quota percentage that the resource control rule can use. Valid values: 1 to 100. You must specify one and only one of this parameter and CpuCount.
	//
	// example:
	//
	// 20
	MaxCpu *int32 `json:"MaxCpu,omitempty" xml:"MaxCpu,omitempty"`
	// The region ID of the PolarDB cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource control rule. The name must be 1 to 63 ASCII bytes in length, start with a letter, and can contain only letters, digits, and underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_rc
	ResourceControlName *string `json:"ResourceControlName,omitempty" xml:"ResourceControlName,omitempty"`
}

func (s CreateResourceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceControlRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceControlRequest) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *CreateResourceControlRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateResourceControlRequest) GetMaxCpu() *int32 {
	return s.MaxCpu
}

func (s *CreateResourceControlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateResourceControlRequest) GetResourceControlName() *string {
	return s.ResourceControlName
}

func (s *CreateResourceControlRequest) SetCpuCount(v int32) *CreateResourceControlRequest {
	s.CpuCount = &v
	return s
}

func (s *CreateResourceControlRequest) SetDBClusterId(v string) *CreateResourceControlRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateResourceControlRequest) SetMaxCpu(v int32) *CreateResourceControlRequest {
	s.MaxCpu = &v
	return s
}

func (s *CreateResourceControlRequest) SetRegionId(v string) *CreateResourceControlRequest {
	s.RegionId = &v
	return s
}

func (s *CreateResourceControlRequest) SetResourceControlName(v string) *CreateResourceControlRequest {
	s.ResourceControlName = &v
	return s
}

func (s *CreateResourceControlRequest) Validate() error {
	return dara.Validate(s)
}
