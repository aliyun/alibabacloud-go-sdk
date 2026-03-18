// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int32) *DescribeDdosEventRequest
	GetEndTime() *int32
	SetInstanceId(v string) *DescribeDdosEventRequest
	GetInstanceId() *string
	SetIp(v string) *DescribeDdosEventRequest
	GetIp() *string
	SetPageNo(v int32) *DescribeDdosEventRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeDdosEventRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDdosEventRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDdosEventRequest
	GetResourceGroupId() *string
	SetStartTime(v int32) *DescribeDdosEventRequest
	GetStartTime() *int32
}

type DescribeDdosEventRequest struct {
	// The end time of the DDoS attack events to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638288000
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The attacked IP address to query.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the DDoS attack events to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1633017600
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDdosEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DescribeDdosEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDdosEventRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDdosEventRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeDdosEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDdosEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDdosEventRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDdosEventRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *DescribeDdosEventRequest) SetEndTime(v int32) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetInstanceId(v string) *DescribeDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetIp(v string) *DescribeDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int32) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int32) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetRegionId(v string) *DescribeDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetResourceGroupId(v string) *DescribeDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int32) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventRequest) Validate() error {
	return dara.Validate(s)
}
