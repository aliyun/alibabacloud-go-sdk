// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SwitchDBInstanceVpcRequest
	GetDBInstanceId() *string
	SetPrivateIpAddress(v string) *SwitchDBInstanceVpcRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *SwitchDBInstanceVpcRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *SwitchDBInstanceVpcRequest
	GetVPCId() *string
	SetVSwitchId(v string) *SwitchDBInstanceVpcRequest
	GetVSwitchId() *string
}

type SwitchDBInstanceVpcRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The private IP address of the instance. The private IP address must be within the CIDR block of the vSwitch that is specified by the **VSwitchId*	- parameter.
	//
	// >  You can call the DescribeVSwitches operation to query the CIDR block of the vSwitch.
	//
	// example:
	//
	// 10.23.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The VPC ID.
	//
	// > The VPC must reside in the same region as the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6f7l4fg90*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance.
	//
	// > The vSwitch must belong to the same zone as the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6adz52c2p*****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s SwitchDBInstanceVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceVpcRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceVpcRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchDBInstanceVpcRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *SwitchDBInstanceVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchDBInstanceVpcRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *SwitchDBInstanceVpcRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *SwitchDBInstanceVpcRequest) SetDBInstanceId(v string) *SwitchDBInstanceVpcRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetPrivateIpAddress(v string) *SwitchDBInstanceVpcRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetResourceOwnerId(v int64) *SwitchDBInstanceVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetVPCId(v string) *SwitchDBInstanceVpcRequest {
	s.VPCId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetVSwitchId(v string) *SwitchDBInstanceVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) Validate() error {
	return dara.Validate(s)
}
