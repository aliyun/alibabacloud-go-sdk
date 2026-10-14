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
	// The IDs, editions, remarks, and IP types of the Anti-DDoS Pro or Anti-DDoS Premium instances.
	InstanceIds []*DescribeInstanceIdsResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
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
	if s.InstanceIds != nil {
		for _, item := range s.InstanceIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceIdsResponseBodyInstanceIds struct {
	// The edition of the instance. Valid values:
	//
	// - **0**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Insurance Edition.
	//
	// - **1**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Unlimited Edition.
	//
	// - **2**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Acceleration Line.
	//
	// - **3**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Secure Acceleration Line.
	//
	// - **9**: Anti-DDoS Pro or Anti-DDoS Premium (the Chinese mainland) Professional Edition.
	//
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ddoscoo-cn-zvp2eibz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP forwarding mode of the instance. Valid values:
	//
	// - **fnat**: The IP versions of the origin server and the client are the same. IPv4 client requests are forwarded to the IPv4 origin server, and IPv6 client requests are forwarded to the IPv6 origin server.
	//
	// - **v6tov4**: Both IPv6 and IPv4 client requests are forwarded to the IPv4 origin server.
	//
	// example:
	//
	// fnat
	IpMode *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	// The IP protocol version of the instance. Valid values:
	//
	// - **Ipv4**: IPv4.
	//
	// - **Ipv6**: IPv6.
	//
	// example:
	//
	// Ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The remark of the instance.
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
