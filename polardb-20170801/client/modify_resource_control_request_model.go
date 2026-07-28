// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuCount(v int32) *ModifyResourceControlRequest
	GetCpuCount() *int32
	SetDBClusterId(v string) *ModifyResourceControlRequest
	GetDBClusterId() *string
	SetMaxCpu(v int32) *ModifyResourceControlRequest
	GetMaxCpu() *int32
	SetRegionId(v string) *ModifyResourceControlRequest
	GetRegionId() *string
	SetResourceControlName(v string) *ModifyResourceControlRequest
	GetResourceControlName() *string
}

type ModifyResourceControlRequest struct {
	// The total number of CPU cores.
	//
	// example:
	//
	// 4
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
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
	// cn-hangzhou
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

func (s ModifyResourceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceControlRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceControlRequest) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *ModifyResourceControlRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyResourceControlRequest) GetMaxCpu() *int32 {
	return s.MaxCpu
}

func (s *ModifyResourceControlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceControlRequest) GetResourceControlName() *string {
	return s.ResourceControlName
}

func (s *ModifyResourceControlRequest) SetCpuCount(v int32) *ModifyResourceControlRequest {
	s.CpuCount = &v
	return s
}

func (s *ModifyResourceControlRequest) SetDBClusterId(v string) *ModifyResourceControlRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyResourceControlRequest) SetMaxCpu(v int32) *ModifyResourceControlRequest {
	s.MaxCpu = &v
	return s
}

func (s *ModifyResourceControlRequest) SetRegionId(v string) *ModifyResourceControlRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceControlRequest) SetResourceControlName(v string) *ModifyResourceControlRequest {
	s.ResourceControlName = &v
	return s
}

func (s *ModifyResourceControlRequest) Validate() error {
	return dara.Validate(s)
}
