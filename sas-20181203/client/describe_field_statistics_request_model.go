// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v string) *DescribeFieldStatisticsRequest
	GetMachineTypes() *string
	SetRegionId(v string) *DescribeFieldStatisticsRequest
	GetRegionId() *string
	SetResourceDirectoryAccountId(v int64) *DescribeFieldStatisticsRequest
	GetResourceDirectoryAccountId() *int64
}

type DescribeFieldStatisticsRequest struct {
	// The type of asset to query. If you do not specify an asset type, the statistics information of all asset types is queried. Valid values:
	//
	// - **ecs**: server
	//
	// - **cloud_product**: cloud product.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The ID of the region where the server resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account of the member accounts in the resource directory.
	//
	// >You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DescribeFieldStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeFieldStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFieldStatisticsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeFieldStatisticsRequest) SetMachineTypes(v string) *DescribeFieldStatisticsRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeFieldStatisticsRequest) SetRegionId(v string) *DescribeFieldStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFieldStatisticsRequest) SetResourceDirectoryAccountId(v int64) *DescribeFieldStatisticsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeFieldStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
