// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVMsInServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVMsInServiceMeshResponseBody
	GetRequestId() *string
	SetVMs(v []*DescribeVMsInServiceMeshResponseBodyVMs) *DescribeVMsInServiceMeshResponseBody
	GetVMs() []*DescribeVMsInServiceMeshResponseBodyVMs
}

type DescribeVMsInServiceMeshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4b2c0fe0-6705-4614-8521-6b9d289163c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VMs that are added to the ASM instance.
	VMs []*DescribeVMsInServiceMeshResponseBodyVMs `json:"VMs,omitempty" xml:"VMs,omitempty" type:"Repeated"`
}

func (s DescribeVMsInServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVMsInServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVMsInServiceMeshResponseBody) GetVMs() []*DescribeVMsInServiceMeshResponseBodyVMs {
	return s.VMs
}

func (s *DescribeVMsInServiceMeshResponseBody) SetRequestId(v string) *DescribeVMsInServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBody) SetVMs(v []*DescribeVMsInServiceMeshResponseBodyVMs) *DescribeVMsInServiceMeshResponseBody {
	s.VMs = v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBody) Validate() error {
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

type DescribeVMsInServiceMeshResponseBodyVMs struct {
	// Indicates whether the ECS instance has labels.
	//
	// example:
	//
	// true
	HasTag *bool `json:"HasTag,omitempty" xml:"HasTag,omitempty"`
	// The host name.
	//
	// example:
	//
	// iZ2ze45cgxkx4q12eh9l****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-2ze45cgxkx4q12e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the ECS instance.
	//
	// example:
	//
	// 10.0.*,***
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
	// The state of the ECS instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVMsInServiceMeshResponseBodyVMs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVMsInServiceMeshResponseBodyVMs) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetHasTag() *bool {
	return s.HasTag
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetHostName() *string {
	return s.HostName
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetRegion() *string {
	return s.Region
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetHasTag(v bool) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.HasTag = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetHostName(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.HostName = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetInstanceId(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.InstanceId = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetIpAddress(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.IpAddress = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetRegion(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.Region = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetSecurityGroupIds(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.SecurityGroupIds = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetStatus(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.Status = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) Validate() error {
	return dara.Validate(s)
}
