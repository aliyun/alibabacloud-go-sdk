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
	// The total number of CPU cores.
	//
	// example:
	//
	// 16
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maximum number of CPUs. Unit: 0.001 CPU. A value of 1000 indicates one CPU. If you specify this parameter, instances whose CPU count is less than the specified value are returned.
	//
	// example:
	//
	// 1000000
	MaxCpu *int32 `json:"MaxCpu,omitempty" xml:"MaxCpu,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource control name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-rc
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
