// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() []*DescribeInstancesResponseBodyInstances
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeInstancesResponseBody struct {
	// The details about the instances.
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A0AF40CC-814A-5A86-AEAA-6F19E88B8A39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() []*DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	// The time when the instance was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637751953000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The overdue status of the instance. The value is fixed as **0**, which indicates that your Alibaba Cloud account does not have overdue payments. The instance supports only the subscription billing method.
	//
	// example:
	//
	// 0
	DebtStatus *int32 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// The mitigation plan of the instance. Valid values:
	//
	// 	- **0**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Insurance mitigation plan
	//
	// 	- **1**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Unlimited mitigation plan
	//
	// 	- **2**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Chinese Mainland Acceleration (CMA) mitigation plan
	//
	// 	- **9**: Anti-DDoS Proxy (Chinese Mainland) instance of the Profession mitigation plan
	//
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The traffic forwarding status of the instance. Valid values:
	//
	// 	- **0**: The instance no longer forwards service traffic.
	//
	// 	- **1**: The instance forwards service traffic as expected.
	//
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when the instance expires. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640361600000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-7pp2g9ed****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.199.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP address-based forwarding mode of the instance. Valid values:
	//
	// 	- **fnat**: Requests from IPv4 addresses are forwarded to origin servers that use IPv4 addresses and requests from IPv6 addresses are forwarded to origin servers that use IPv6 addresses.
	//
	// 	- **v6tov4**: All requests are forwarded to origin servers that use IPv4 addresses.
	//
	// example:
	//
	// fnat
	IpMode *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	// The IP version of the instance. Valid values:
	//
	// 	- **Ipv4**
	//
	// 	- **Ipv6**
	//
	// example:
	//
	// Ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the metering method of the 95th percentile burstable clean bandwidth is enabled for the instance. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	IsFirstOpenBw *int64 `json:"IsFirstOpenBw,omitempty" xml:"IsFirstOpenBw,omitempty"`
	// Indicates whether the metering method of the 95th percentile burstable QPS is enabled for the instance. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	IsFirstOpenQps *int64 `json:"IsFirstOpenQps,omitempty" xml:"IsFirstOpenQps,omitempty"`
	// The remarks of the instance.
	//
	// example:
	//
	// doc-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **1**: normal
	//
	// 	- **2**: expired
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeInstancesResponseBodyInstances) GetDebtStatus() *int32 {
	return s.DebtStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetEdition() *int32 {
	return s.Edition
}

func (s *DescribeInstancesResponseBodyInstances) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeInstancesResponseBodyInstances) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesResponseBodyInstances) GetIpMode() *string {
	return s.IpMode
}

func (s *DescribeInstancesResponseBodyInstances) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstancesResponseBodyInstances) GetIsFirstOpenBw() *int64 {
	return s.IsFirstOpenBw
}

func (s *DescribeInstancesResponseBodyInstances) GetIsFirstOpenQps() *int64 {
	return s.IsFirstOpenQps
}

func (s *DescribeInstancesResponseBodyInstances) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstancesResponseBodyInstances) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstances) SetCreateTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDebtStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEdition(v int32) *DescribeInstancesResponseBodyInstances {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnabled(v int32) *DescribeInstancesResponseBodyInstances {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIp(v string) *DescribeInstancesResponseBodyInstances {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIpMode(v string) *DescribeInstancesResponseBodyInstances {
	s.IpMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIpVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIsFirstOpenBw(v int64) *DescribeInstancesResponseBodyInstances {
	s.IsFirstOpenBw = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIsFirstOpenQps(v int64) *DescribeInstancesResponseBodyInstances {
	s.IsFirstOpenQps = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRemark(v string) *DescribeInstancesResponseBodyInstances {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
