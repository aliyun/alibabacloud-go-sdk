// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*DescribeInstanceIdsResponseBodyInstanceIds) *DescribeInstanceIdsResponseBody
	GetInstanceIds() []*DescribeInstanceIdsResponseBodyInstanceIds
	SetRequestId(v string) *DescribeInstanceIdsResponseBody
	GetRequestId() *string
}

type DescribeInstanceIdsResponseBody struct {
	// The ID, type, description, and IP version of the instance.
	InstanceIds []*DescribeInstanceIdsResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 310A41FD-0990-5610-92E0-A6A55D7C6444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsResponseBody) GetInstanceIds() []*DescribeInstanceIdsResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *DescribeInstanceIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceIdsResponseBody) SetInstanceIds(v []*DescribeInstanceIdsResponseBodyInstanceIds) *DescribeInstanceIdsResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstanceIdsResponseBody) SetRequestId(v string) *DescribeInstanceIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceIdsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceIdsResponseBodyInstanceIds struct {
	// The type of the instance. Valid values:
	//
	// 	- **0**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Insurance mitigation plan
	//
	// 	- **1**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Unlimited mitigation plan
	//
	// 	- **2**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the CMA mitigation plan
	//
	// 	- **3**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Sec-CMA mitigation plan
	//
	// 	- **9**: Anti-DDoS Proxy (Chinese Mainland) instance of the Profession mitigation plan
	//
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-zvp2eibz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The description of the instance.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeInstanceIdsResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIdsResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) GetEdition() *int32 {
	return s.Edition
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) GetIpMode() *string {
	return s.IpMode
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetEdition(v int32) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetInstanceId(v string) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetIpMode(v string) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.IpMode = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetIpVersion(v string) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetRemark(v string) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}
