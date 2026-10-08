// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRCInstanceList(v []*DescribeRCInstanceIpAddressResponseBodyRCInstanceList) *DescribeRCInstanceIpAddressResponseBody
	GetRCInstanceList() []*DescribeRCInstanceIpAddressResponseBodyRCInstanceList
	SetRequestId(v string) *DescribeRCInstanceIpAddressResponseBody
	GetRequestId() *string
	SetTotal(v string) *DescribeRCInstanceIpAddressResponseBody
	GetTotal() *string
}

type DescribeRCInstanceIpAddressResponseBody struct {
	// The details of instances to which the assets that are assigned public IP addresses belong.
	RCInstanceList []*DescribeRCInstanceIpAddressResponseBodyRCInstanceList `json:"RCInstanceList,omitempty" xml:"RCInstanceList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C048E440-EA84-5E97-8C81-2A7060D0****_th**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of assets that are assigned public IP addresses returned.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRCInstanceIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceIpAddressResponseBody) GetRCInstanceList() []*DescribeRCInstanceIpAddressResponseBodyRCInstanceList {
	return s.RCInstanceList
}

func (s *DescribeRCInstanceIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceIpAddressResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeRCInstanceIpAddressResponseBody) SetRCInstanceList(v []*DescribeRCInstanceIpAddressResponseBodyRCInstanceList) *DescribeRCInstanceIpAddressResponseBody {
	s.RCInstanceList = v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBody) SetRequestId(v string) *DescribeRCInstanceIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBody) SetTotal(v string) *DescribeRCInstanceIpAddressResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBody) Validate() error {
	if s.RCInstanceList != nil {
		for _, item := range s.RCInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceIpAddressResponseBodyRCInstanceList struct {
	// The Custom instance ID.
	//
	// example:
	//
	// rc-kti8hw44yy0x53******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Custom instance name.
	//
	// example:
	//
	// rc-kti8hw44yy0x53******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DDoS mitigation status of the instance. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **abnormal**: Under attack.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the assets that are assigned public IP addresses. The value is fixed as **ecs**.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The details of the assets that are assigned public IP addresses.
	IpAddressConfig []*DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig `json:"IpAddressConfig,omitempty" xml:"IpAddressConfig,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceIpAddressResponseBodyRCInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceIpAddressResponseBodyRCInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) GetIpAddressConfig() []*DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	return s.IpAddressConfig
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) SetInstanceId(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) SetInstanceName(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceList {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) SetInstanceStatus(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) SetInstanceType(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) SetIpAddressConfig(v []*DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) *DescribeRCInstanceIpAddressResponseBodyRCInstanceList {
	s.IpAddressConfig = v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceList) Validate() error {
	if s.IpAddressConfig != nil {
		for _, item := range s.IpAddressConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig struct {
	// The basic DDoS Mitigation Threshold of the assets that are assigned public IP addresses. Unit: Mbit/s.
	//
	// example:
	//
	// 5200
	BlackholeThreshold *int32 `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	// The traffic scrubbing threshold of the assets that are assigned public IP addresses. Unit: Mbit/s.
	//
	// example:
	//
	// 300
	DefenseBpsThreshold *int32 `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	// The message rate scrubbing threshold of the assets that are assigned public IP addresses. Unit: pps.
	//
	// example:
	//
	// 70000
	DefensePpsThreshold *int32 `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	// The DDoS burstable Mitigation Threshold of the assets that are assigned public IP addresses. Unit: Mbit/s.
	//
	// example:
	//
	// 12310
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	// The IP address of the assets that are assigned public IP addresses.
	//
	// example:
	//
	// 39.105.XXX.XXX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The DDoS mitigation status of the assets that are assigned public IP addresses. Valid values:
	//
	// - **mitigating**: Cleaning.
	//
	// - **blackholed**: Black Hole Activated.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// normal
	IpStatus *string `json:"IpStatus,omitempty" xml:"IpStatus,omitempty"`
	// The IP protocol version of the instance. Valid values:
	//
	// - **v4**
	//
	// - **v6**
	//
	// example:
	//
	// v4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the assets that are assigned public IP addresses is attached to Anti-DDoS Origin. Valid values:
	//
	// - **true**: Attached.
	//
	// - **false**: Not attached.
	//
	// example:
	//
	// true
	IsBgppack *bool `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
	// Indicates whether best-effort protection is enabled for the assets that are assigned public IP addresses in Anti-DDoS Origin. Valid values:
	//
	// - **0**: Best-effort protection is not enabled.
	//
	// - **1**: Best-effort protection is enabled.
	//
	// example:
	//
	// 0
	IsFullProtection *int32 `json:"IsFullProtection,omitempty" xml:"IsFullProtection,omitempty"`
	// The region encoding of the assets that are assigned public IP addresses.
	//
	// example:
	//
	// cn-beijing-wt97-a01
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetBlackholeThreshold() *int32 {
	return s.BlackholeThreshold
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetDefenseBpsThreshold() *int32 {
	return s.DefenseBpsThreshold
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetDefensePpsThreshold() *int32 {
	return s.DefensePpsThreshold
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetElasticThreshold() *int32 {
	return s.ElasticThreshold
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetIpStatus() *string {
	return s.IpStatus
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetIsBgppack() *bool {
	return s.IsBgppack
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetIsFullProtection() *int32 {
	return s.IsFullProtection
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetBlackholeThreshold(v int32) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetDefenseBpsThreshold(v int32) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetDefensePpsThreshold(v int32) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetElasticThreshold(v int32) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetInstanceIp(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.InstanceIp = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetIpStatus(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.IpStatus = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetIpVersion(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.IpVersion = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetIsBgppack(v bool) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.IsBgppack = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetIsFullProtection(v int32) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.IsFullProtection = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) SetRegionId(v string) *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponseBodyRCInstanceListIpAddressConfig) Validate() error {
	return dara.Validate(s)
}
