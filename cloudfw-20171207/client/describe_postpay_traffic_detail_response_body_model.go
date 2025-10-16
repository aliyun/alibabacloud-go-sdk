// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayTrafficDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePostpayTrafficDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePostpayTrafficDetailResponseBody
	GetTotalCount() *int32
	SetTrafficList(v []*DescribePostpayTrafficDetailResponseBodyTrafficList) *DescribePostpayTrafficDetailResponseBody
	GetTrafficList() []*DescribePostpayTrafficDetailResponseBodyTrafficList
}

type DescribePostpayTrafficDetailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0525EADE-C112-5702-A5BC-0E2F6F94DB23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The statistics on traffic.
	TrafficList []*DescribePostpayTrafficDetailResponseBodyTrafficList `json:"TrafficList,omitempty" xml:"TrafficList,omitempty" type:"Repeated"`
}

func (s DescribePostpayTrafficDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayTrafficDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePostpayTrafficDetailResponseBody) GetTrafficList() []*DescribePostpayTrafficDetailResponseBodyTrafficList {
	return s.TrafficList
}

func (s *DescribePostpayTrafficDetailResponseBody) SetRequestId(v string) *DescribePostpayTrafficDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBody) SetTotalCount(v int32) *DescribePostpayTrafficDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBody) SetTrafficList(v []*DescribePostpayTrafficDetailResponseBodyTrafficList) *DescribePostpayTrafficDetailResponseBody {
	s.TrafficList = v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBody) Validate() error {
	if s.TrafficList != nil {
		for _, item := range s.TrafficList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePostpayTrafficDetailResponseBodyTrafficList struct {
	// The inbound network throughput, which indicates the total number of bytes that are received Unit: bytes.
	//
	// example:
	//
	// 1115096939
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-8vb2d7c9mtn0bo9qcraq
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The asset type. This value takes effect only for the Internet firewall.
	//
	// example:
	//
	// EcsPublicIP
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The outbound network throughput, which indicates the total number of bytes that are sent. Unit: bytes.
	//
	// example:
	//
	// 100000000
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// Protection duration. Unit: hours.
	//
	// example:
	//
	// 20
	ProtectionDuration *int64 `json:"ProtectionDuration,omitempty" xml:"ProtectionDuration,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The resource ID. The resource ID for the Internet firewall is the public IP address that is protected the Internet firewall, and the resource ID for a NAT firewall is the instance ID of the NAT firewall.
	//
	// example:
	//
	// 39.106.146.214
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The total inbound and outbound network throughput, which indicates the total number of bytes that are received and sent. Unit: bytes.
	//
	// example:
	//
	// 1215096939
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The date on which the statistics are collected.
	//
	// example:
	//
	// 20231001
	TrafficDay *string `json:"TrafficDay,omitempty" xml:"TrafficDay,omitempty"`
	// The traffic type. Valid values:
	//
	// 	- **EIP_TRAFFIC**: traffic for the Internet firewall
	//
	// 	- **NatGateway_TRAFFIC**: traffic for NAT firewalls
	//
	// 	- **VPC_TRAFFIC**: traffic for VPC firewalls
	//
	// example:
	//
	// EIP_TRAFFIC
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribePostpayTrafficDetailResponseBodyTrafficList) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficDetailResponseBodyTrafficList) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetProtectionDuration() *int64 {
	return s.ProtectionDuration
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetTrafficDay() *string {
	return s.TrafficDay
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) GetTrafficType() *string {
	return s.TrafficType
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetInBytes(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.InBytes = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetInstanceId(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.InstanceId = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetInstanceType(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.InstanceType = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetOutBytes(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.OutBytes = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetProtectionDuration(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.ProtectionDuration = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetRegionNo(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.RegionNo = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetResourceId(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.ResourceId = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetTotalBytes(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.TotalBytes = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetTrafficDay(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.TrafficDay = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetTrafficType(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.TrafficType = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) Validate() error {
	return dara.Validate(s)
}
