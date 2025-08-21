// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*DescribeInstanceIpAddressResponseBodyInstanceList) *DescribeInstanceIpAddressResponseBody
	GetInstanceList() []*DescribeInstanceIpAddressResponseBodyInstanceList
	SetRequestId(v string) *DescribeInstanceIpAddressResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeInstanceIpAddressResponseBody
	GetTotal() *int32
}

type DescribeInstanceIpAddressResponseBody struct {
	// An array that consists of details of the instance.
	InstanceList []*DescribeInstanceIpAddressResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0907F8-A9F3-5E11-977B-D59CD98C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBody) GetInstanceList() []*DescribeInstanceIpAddressResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeInstanceIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceIpAddressResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstanceIpAddressResponseBody) SetInstanceList(v []*DescribeInstanceIpAddressResponseBodyInstanceList) *DescribeInstanceIpAddressResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) SetRequestId(v string) *DescribeInstanceIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) SetTotal(v int32) *DescribeInstanceIpAddressResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceIpAddressResponseBodyInstanceList struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DDoS mitigation status of the instance. Valid values:
	//
	// 	- **normal**: not under DDoS attacks.
	//
	// 	- **abnormal**: under DDoS attacks.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **ecs**
	//
	// 	- **slb**
	//
	// 	- **eip**
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// An array that consists of the details of the asset.
	IpAddressConfig []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig `json:"IpAddressConfig,omitempty" xml:"IpAddressConfig,omitempty" type:"Repeated"`
}

func (s DescribeInstanceIpAddressResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) GetIpAddressConfig() []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	return s.IpAddressConfig
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceName(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceStatus(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetIpAddressConfig(v []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.IpAddressConfig = v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig struct {
	// The basic protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 5200
	BlackholeThreshold *int32 `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	// The traffic scrubbing threshold for the asset measured in Mbit/s. Unit: Mbit/s.
	//
	// example:
	//
	// 300
	DefenseBpsThreshold *int32 `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	// The traffic scrubbing threshold for the asset measured in packets per second (PPS). Unit: PPS.
	//
	// example:
	//
	// 70000
	DefensePpsThreshold *int32 `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	// The burstable protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 12310
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **mitigating**: indicates that traffic scrubbing is in progress.
	//
	// 	- **blackholed**: indicates that blackhole filtering is triggered for the asset.
	//
	// 	- **normal**: indicates that no DDoS attacks are launched against the asset.
	//
	// example:
	//
	// normal
	IpStatus *string `json:"IpStatus,omitempty" xml:"IpStatus,omitempty"`
	// The IP version of the IP address. Valid values:
	//
	// 	- **v4**: IPv4.
	//
	// 	- **v6**: IPv6.
	//
	// example:
	//
	// v4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the asset is added to the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsBgppack *bool `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
	// Indicates whether best-effort protection is enabled for the asset. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 0
	IsFullProtection *int32 `json:"IsFullProtection,omitempty" xml:"IsFullProtection,omitempty"`
	// The region code of the asset.
	//
	// example:
	//
	// cn-hangzhou-dg-a01
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetBlackholeThreshold() *int32 {
	return s.BlackholeThreshold
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetDefenseBpsThreshold() *int32 {
	return s.DefenseBpsThreshold
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetDefensePpsThreshold() *int32 {
	return s.DefensePpsThreshold
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetElasticThreshold() *int32 {
	return s.ElasticThreshold
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetIpStatus() *string {
	return s.IpStatus
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetIsBgppack() *bool {
	return s.IsBgppack
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetIsFullProtection() *int32 {
	return s.IsFullProtection
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetBlackholeThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetDefenseBpsThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetDefensePpsThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetElasticThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetInstanceIp(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIpStatus(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IpStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIpVersion(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIsBgppack(v bool) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IsBgppack = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIsFullProtection(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IsFullProtection = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetRegionId(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) Validate() error {
	return dara.Validate(s)
}
