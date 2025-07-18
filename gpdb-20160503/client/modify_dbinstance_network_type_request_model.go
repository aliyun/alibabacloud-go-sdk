// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetDBInstanceId() *string
	SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest
	GetInstanceNetworkType() *string
	SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest
	GetPrivateIpAddress() *string
	SetVPCId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetVSwitchId() *string
}

type ModifyDBInstanceNetworkTypeRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new network type of the instance. Valid values:
	//
	// 	- VPC
	//
	// 	- Classic
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The internal IP address of the instance.
	//
	// example:
	//
	// 10.10.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The virtual private cloud (VPC) ID of the instance.
	//
	// example:
	//
	// vpc-bp19ame5m1r3oejns****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance. This parameter must be specified when VPCId is specified.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVPCId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) Validate() error {
	return dara.Validate(s)
}
