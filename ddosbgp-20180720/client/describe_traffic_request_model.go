// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int32) *DescribeTrafficRequest
	GetEndTime() *int32
	SetFlowType(v string) *DescribeTrafficRequest
	GetFlowType() *string
	SetInstanceId(v string) *DescribeTrafficRequest
	GetInstanceId() *string
	SetInterval(v int32) *DescribeTrafficRequest
	GetInterval() *int32
	SetIp(v string) *DescribeTrafficRequest
	GetIp() *string
	SetIpnet(v string) *DescribeTrafficRequest
	GetIpnet() *string
	SetRegionId(v string) *DescribeTrafficRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTrafficRequest
	GetResourceGroupId() *string
	SetStartTime(v int32) *DescribeTrafficRequest
	GetStartTime() *int32
}

type DescribeTrafficRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// If you do not specify this parameter, the current system time is used as the end time.
	//
	// example:
	//
	// 1563445054
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the traffic statistics to query. Valid values:
	//
	// 	- **max**: the peak traffic within the specified interval.
	//
	// 	- **avg**: the average traffic within the specified interval.
	//
	// Enumerated values:
	//
	// 	- all
	//
	// 	- avg
	//
	// 	- max
	//
	// example:
	//
	// max
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// If you specify an on-demand instance, you must configure the **Interval*	- parameter.
	//
	// example:
	//
	// ddosbgp-cn-n6w203qg****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The interval at which the traffic statistics are calculated. Unit: seconds. Default value: **5**.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The public IP address of the asset to query. If you do not specify this parameter, the traffic statistics of all public IP addresses that are protected by the Anti-DDoS Origin instance are queried.
	//
	// >  The public IP address must be a protected object of the Anti-DDoS Origin instance. You can call the [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) operation to query all protected objects of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// 39.XX.XX.96
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The Classless Inter-Domain Routing (CIDR) block of the on-demand instance that you want to query.
	//
	// example:
	//
	// 111.XX.XX.0/24
	Ipnet *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
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
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619798400
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DescribeTrafficRequest) GetFlowType() *string {
	return s.FlowType
}

func (s *DescribeTrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTrafficRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeTrafficRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeTrafficRequest) GetIpnet() *string {
	return s.Ipnet
}

func (s *DescribeTrafficRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTrafficRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTrafficRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *DescribeTrafficRequest) SetEndTime(v int32) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetFlowType(v string) *DescribeTrafficRequest {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficRequest) SetInstanceId(v string) *DescribeTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int32) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetIp(v string) *DescribeTrafficRequest {
	s.Ip = &v
	return s
}

func (s *DescribeTrafficRequest) SetIpnet(v string) *DescribeTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTrafficRequest) SetRegionId(v string) *DescribeTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTrafficRequest) SetResourceGroupId(v string) *DescribeTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int32) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTrafficRequest) Validate() error {
	return dara.Validate(s)
}
