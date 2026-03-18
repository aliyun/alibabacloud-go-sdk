// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody
	GetInstanceSpecs() []*DescribeInstanceSpecsResponseBodyInstanceSpecs
	SetRequestId(v string) *DescribeInstanceSpecsResponseBody
	GetRequestId() *string
}

type DescribeInstanceSpecsResponseBody struct {
	// The specifications of the Anti-DDoS Origin instance, including whether best-effort protection is enabled, the number of available best-effort protection sessions, and the number of used best-effort protection sessions.
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 5840AB9F-1419-4620-807D-5EA476090247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) GetInstanceSpecs() []*DescribeInstanceSpecsResponseBodyInstanceSpecs {
	return s.InstanceSpecs
}

func (s *DescribeInstanceSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) Validate() error {
	if s.InstanceSpecs != nil {
		for _, item := range s.InstanceSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	// The available best-effort protection sessions.
	//
	// example:
	//
	// 2
	AvailableDefenseTimes *int32 `json:"AvailableDefenseTimes,omitempty" xml:"AvailableDefenseTimes,omitempty"`
	// The number of times that blackhole filtering can be deactivated.
	//
	// example:
	//
	// 100
	AvailableDeleteBlackholeCount *string `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	// The percentage of the used best-effort protection sessions. Unit: %.
	//
	// example:
	//
	// 30
	DefenseTimesPercent *int32 `json:"DefenseTimesPercent,omitempty" xml:"DefenseTimesPercent,omitempty"`
	// Indicates whether the instance is downgraded. Valid value:
	//
	// 	- **8**: The instance is downgraded due to excessive bandwidth usage.
	//
	// example:
	//
	// 8
	DowngradeStatus *int32 `json:"DowngradeStatus,omitempty" xml:"DowngradeStatus,omitempty"`
	// The ID of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether best-effort protection is enabled. Valid values:
	//
	// 	- **0**: Best-effort protection is disabled.
	//
	// 	- **1**: Best-effort protection is enabled.
	//
	// example:
	//
	// 1
	IsFullDefenseMode *int32 `json:"IsFullDefenseMode,omitempty" xml:"IsFullDefenseMode,omitempty"`
	// The configurations of the Anti-DDoS Origin instance, including the number of protected IP addresses and protection bandwidth.
	PackConfig *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the name of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The total best-effort protection sessions.
	//
	// example:
	//
	// 2
	TotalDefenseTimes *int32 `json:"TotalDefenseTimes,omitempty" xml:"TotalDefenseTimes,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetAvailableDefenseTimes() *int32 {
	return s.AvailableDefenseTimes
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetAvailableDeleteBlackholeCount() *string {
	return s.AvailableDeleteBlackholeCount
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetDefenseTimesPercent() *int32 {
	return s.DefenseTimesPercent
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetDowngradeStatus() *int32 {
	return s.DowngradeStatus
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetIsFullDefenseMode() *int32 {
	return s.IsFullDefenseMode
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetPackConfig() *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	return s.PackConfig
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetTotalDefenseTimes() *int32 {
	return s.TotalDefenseTimes
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDeleteBlackholeCount(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseTimesPercent(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseTimesPercent = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDowngradeStatus(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DowngradeStatus = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetIsFullDefenseMode(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.IsFullDefenseMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPackConfig(v *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PackConfig = v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetRegion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.Region = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetTotalDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.TotalDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) Validate() error {
	if s.PackConfig != nil {
		if err := s.PackConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig struct {
	// The bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of IP addresses that are protected by the Anti-DDoS Origin Enterprise instance.
	//
	// example:
	//
	// 0
	BindIpCount *int32 `json:"BindIpCount,omitempty" xml:"BindIpCount,omitempty"`
	// The burstable clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	ElasticBwMbps *int32 `json:"ElasticBwMbps,omitempty" xml:"ElasticBwMbps,omitempty"`
	// The metering method of burstable clean bandwidth. Valid values:
	//
	// 	- **month**: the monthly 95th percentile metering method.
	//
	// 	- **day**: the daily 95th percentile metering method.
	//
	// example:
	//
	// day
	ElasticBwMode *string `json:"ElasticBwMode,omitempty" xml:"ElasticBwMode,omitempty"`
	// The burstable protection bandwidth of each protected IP address. Unit: Gbit/s.
	//
	// example:
	//
	// 300
	IpAdvanceThre *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	// The basic protection bandwidth of each protected IP address. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	IpBasicThre *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	// The number of IP addresses that can be protected by the Anti-DDoS Origin Enterprise instance.
	//
	// example:
	//
	// 100
	IpSpec *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
	// The clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	NormalBandwidth *int32 `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	// The burstable protection bandwidth of the Anti-DDoS Origin instance. Unit: Gbit/s.
	//
	// example:
	//
	// 300
	PackAdvThre *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	// The basic protection bandwidth of the Anti-DDoS Origin instance. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	PackBasicThre *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetBindIpCount() *int32 {
	return s.BindIpCount
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetElasticBwMbps() *int32 {
	return s.ElasticBwMbps
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetElasticBwMode() *string {
	return s.ElasticBwMode
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetIpAdvanceThre() *int32 {
	return s.IpAdvanceThre
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetIpBasicThre() *int32 {
	return s.IpBasicThre
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetIpSpec() *int32 {
	return s.IpSpec
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetNormalBandwidth() *int32 {
	return s.NormalBandwidth
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetPackAdvThre() *int32 {
	return s.PackAdvThre
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GetPackBasicThre() *int32 {
	return s.PackBasicThre
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBandwidth(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBindIpCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.BindIpCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetElasticBwMbps(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.ElasticBwMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetElasticBwMode(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.ElasticBwMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpAdvanceThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpSpec(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetNormalBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.NormalBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackAdvThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) Validate() error {
	return dara.Validate(s)
}
