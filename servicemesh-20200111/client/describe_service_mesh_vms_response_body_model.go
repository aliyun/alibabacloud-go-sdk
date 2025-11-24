// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshVMsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceMeshVMsResponseBody
	GetRequestId() *string
	SetVMs(v []*DescribeServiceMeshVMsResponseBodyVMs) *DescribeServiceMeshVMsResponseBody
	GetVMs() []*DescribeServiceMeshVMsResponseBodyVMs
}

type DescribeServiceMeshVMsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4b2c0fe0-6705-4614-8521-6b9d289163c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ECS instances that reside in the same VPC as the ASM instance.
	VMs []*DescribeServiceMeshVMsResponseBodyVMs `json:"VMs,omitempty" xml:"VMs,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshVMsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshVMsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshVMsResponseBody) GetVMs() []*DescribeServiceMeshVMsResponseBodyVMs {
	return s.VMs
}

func (s *DescribeServiceMeshVMsResponseBody) SetRequestId(v string) *DescribeServiceMeshVMsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBody) SetVMs(v []*DescribeServiceMeshVMsResponseBodyVMs) *DescribeServiceMeshVMsResponseBody {
	s.VMs = v
	return s
}

func (s *DescribeServiceMeshVMsResponseBody) Validate() error {
	if s.VMs != nil {
		for _, item := range s.VMs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceMeshVMsResponseBodyVMs struct {
	// Indicates whether the ECS instance has labels.
	//
	// example:
	//
	// false
	HasTag *bool `json:"HasTag,omitempty" xml:"HasTag,omitempty"`
	// The host name.
	//
	// example:
	//
	// iZ2ze45cgxkx4q12eh****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-2ze0kub9scdguom****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the ECS instance.
	//
	// example:
	//
	// 192.168.2.241
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The security group to which the ECS instance belongs.
	//
	// example:
	//
	// sg-2zeaqy08amco9osl****
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// ccb37ff104caf419fbf48fb38e6f3****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The state of the ECS instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServiceMeshVMsResponseBodyVMs) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshVMsResponseBodyVMs) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetHasTag() *bool {
	return s.HasTag
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetHostName() *string {
	return s.HostName
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetRegion() *string {
	return s.Region
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) GetStatus() *string {
	return s.Status
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetHasTag(v bool) *DescribeServiceMeshVMsResponseBodyVMs {
	s.HasTag = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetHostName(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.HostName = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetInstanceId(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.InstanceId = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetIpAddress(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.IpAddress = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetRegion(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.Region = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetSecurityGroupIds(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.SecurityGroupIds = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetServiceMeshId(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetStatus(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.Status = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) Validate() error {
	return dara.Validate(s)
}
