// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *DescribeIpLocationServiceResponseBodyInstance) *DescribeIpLocationServiceResponseBody
	GetInstance() *DescribeIpLocationServiceResponseBodyInstance
	SetRequestId(v string) *DescribeIpLocationServiceResponseBody
	GetRequestId() *string
}

type DescribeIpLocationServiceResponseBody struct {
	// The details of the asset.
	Instance *DescribeIpLocationServiceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponseBody) GetInstance() *DescribeIpLocationServiceResponseBodyInstance {
	return s.Instance
}

func (s *DescribeIpLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpLocationServiceResponseBody) SetInstance(v *DescribeIpLocationServiceResponseBodyInstance) *DescribeIpLocationServiceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeIpLocationServiceResponseBody) SetRequestId(v string) *DescribeIpLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIpLocationServiceResponseBodyInstance struct {
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
	// The type of the asset. Valid values:
	//
	// 	- **ecs**: an ECS instance.
	//
	// 	- **slb**: an SLB instance.
	//
	// 	- **eip**: an EIP.
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The region to which the public IP address of the asset belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeIpLocationServiceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpLocationServiceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpLocationServiceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeIpLocationServiceResponseBodyInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeIpLocationServiceResponseBodyInstance) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeIpLocationServiceResponseBodyInstance) GetRegion() *string {
	return s.Region
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceId(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceName(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceType(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInternetIp(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetRegion(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.Region = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
