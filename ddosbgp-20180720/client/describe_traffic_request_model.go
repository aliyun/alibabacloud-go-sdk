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
	// The end time of the query. Specify a UNIX timestamp. Unit: seconds.
	//
	// If you do not set this parameter, the current time is used as the end time.
	//
	// example:
	//
	// 1563445054
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of traffic statistics to query. Valid values:
	//
	// - **max**: the peak traffic within the statistical interval.
	//
	// - **avg**: the average traffic within the statistical interval.
	//
	// example:
	//
	// max
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// > You can call [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) to query the IDs of all Anti-DDoS Origin instances.
	//
	//
	// If the instance specified here is used for traffic diversion, you must set the **Interval*	- request parameter.
	//
	// example:
	//
	// ddosbgp-cn-n6w203qg****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval for traffic statistics. Unit: seconds. This parameter specifies the length of each interval for which traffic data is aggregated. Default value: **5**.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The assets that are assigned public IP addresses to query. If you do not specify this parameter in Settings, the traffic statistics of all assets that are assigned public IP addresses protected by the Anti-DDoS Origin instance are queried.
	//
	// > The assets that are assigned public IP addresses must have been added as protected objects of the Anti-DDoS Origin instance. You can invoke [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) to query all protected objects of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// 39.XX.XX.96
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The CIDR block used for traffic diversion to query.
	//
	// example:
	//
	// 111.XX.XX.0/24
	Ipnet *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// > You can call [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) to query all region IDs supported by Anti-DDoS Origin.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not set this parameter, the default resource group is used.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the query. Specify a UNIX timestamp. Unit: seconds.
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
