// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v *DescribeInstanceResponseBodyInstanceList) *DescribeInstanceResponseBody
	GetInstanceList() *DescribeInstanceResponseBodyInstanceList
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeInstanceResponseBody
	GetTotal() *int32
}

type DescribeInstanceResponseBody struct {
	// The details of the assets.
	InstanceList *DescribeInstanceResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetInstanceList() *DescribeInstanceResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstanceResponseBody) SetInstanceList(v *DescribeInstanceResponseBodyInstanceList) *DescribeInstanceResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTotal(v int32) *DescribeInstanceResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyInstanceList struct {
	Instance []*DescribeInstanceResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceList) GetInstance() []*DescribeInstanceResponseBodyInstanceListInstance {
	return s.Instance
}

func (s *DescribeInstanceResponseBodyInstanceList) SetInstance(v []*DescribeInstanceResponseBodyInstanceListInstance) *DescribeInstanceResponseBodyInstanceList {
	s.Instance = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyInstanceListInstance struct {
	// The basic protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 5200
	BlackholeThreshold *int32 `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	// The traffic scrubbing threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 300
	DefenseBpsThreshold *int32 `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	// The packet scrubbing threshold for the asset. Unit: packets per second (pps).
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
	// The ID of the asset.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 121.199.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **mitigating**: indicates that traffic scrubbing is triggered for the asset.
	//
	// 	- **blackholed**: indicates that blackhole filtering is triggered for the asset.
	//
	// 	- **normal**: indicates that the instance is normal.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the asset.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP protocol that is supported by the asset. Valid values:
	//
	// 	- **v4**: IPv4
	//
	// 	- **v6**: IPv6
	//
	// example:
	//
	// v4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the asset is associated with an Anti-DDoS Origin Basic instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	IsBgppack *bool `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceListInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetBlackholeThreshold() *int32 {
	return s.BlackholeThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetDefenseBpsThreshold() *int32 {
	return s.DefenseBpsThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetDefensePpsThreshold() *int32 {
	return s.DefensePpsThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetElasticThreshold() *int32 {
	return s.ElasticThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetIsBgppack() *bool {
	return s.IsBgppack
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetBlackholeThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefenseBpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefensePpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetElasticThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceIp(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceStatus(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceType(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIpVersion(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIsBgppack(v bool) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IsBgppack = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) Validate() error {
	return dara.Validate(s)
}
