// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpiredDays(v string) *ModifyDBInstanceNetworkTypeRequest
	GetClassicExpiredDays() *string
	SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetDBInstanceId() *string
	SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest
	GetInstanceNetworkType() *string
	SetOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest
	GetPrivateIpAddress() *string
	SetReadWriteSplittingClassicExpiredDays(v int32) *ModifyDBInstanceNetworkTypeRequest
	GetReadWriteSplittingClassicExpiredDays() *int32
	SetReadWriteSplittingPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest
	GetReadWriteSplittingPrivateIpAddress() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest
	GetResourceOwnerId() *int64
	SetRetainClassic(v string) *ModifyDBInstanceNetworkTypeRequest
	GetRetainClassic() *string
	SetVPCId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetVSwitchId() *string
}

type ModifyDBInstanceNetworkTypeRequest struct {
	// The number of days for which you want to retain the classic network endpoint. Valid values: **1 to 120**. Default value: **7**.
	//
	// > If you set the **RetainClassic*	- parameter to **True**, you must also specify this parameter.
	//
	// example:
	//
	// 7
	ClassicExpiredDays *string `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type after the modification. Set the value to **VPC**.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The internal IP address of the instance. The internal IP address must be within the CIDR block supported by the specified vSwitch. The system automatically assigns a private IP address to an instance based on the values of **VPCId*	- and **VSwitchId**.
	//
	// example:
	//
	// 172.10.40.25
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of days for which you want to retain the read/write splitting endpoint of the classic network type. Valid values: **1 to 120**. Default value: **7**.
	//
	// >  This parameter takes effect only when a read/write splitting endpoint of the classic network type exists and the **RetainClassic*	- parameter is set to **True**.
	//
	// example:
	//
	// 7
	ReadWriteSplittingClassicExpiredDays *int32 `json:"ReadWriteSplittingClassicExpiredDays,omitempty" xml:"ReadWriteSplittingClassicExpiredDays,omitempty"`
	// The internal IP address that corresponds to the read/write splitting endpoint of the instance. The internal IP address must be within the CIDR block supported by the specified vSwitch. The system automatically assigns a private IP address to an instance based on the values of **VPCId*	- and **VSwitchId**.
	//
	// >  This parameter is valid when a read/write splitting endpoint of the classic network type exists.
	//
	// example:
	//
	// 192.168.0.22
	ReadWriteSplittingPrivateIpAddress *string `json:"ReadWriteSplittingPrivateIpAddress,omitempty" xml:"ReadWriteSplittingPrivateIpAddress,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to retain the classic network endpoint. Valid values:
	//
	// 	- **True**: retains the classic network endpoint.
	//
	// 	- **False*	- (default): does not retain the classic network endpoint.
	//
	// example:
	//
	// True
	RetainClassic *string `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6f7l4fg90xxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch. This parameter is required if the **VPCId*	- parameter is specified.
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetClassicExpiredDays() *string {
	return s.ClassicExpiredDays
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetReadWriteSplittingClassicExpiredDays() *int32 {
	return s.ReadWriteSplittingClassicExpiredDays
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetReadWriteSplittingPrivateIpAddress() *string {
	return s.ReadWriteSplittingPrivateIpAddress
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetRetainClassic() *string {
	return s.RetainClassic
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetClassicExpiredDays(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetReadWriteSplittingClassicExpiredDays(v int32) *ModifyDBInstanceNetworkTypeRequest {
	s.ReadWriteSplittingClassicExpiredDays = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetReadWriteSplittingPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.ReadWriteSplittingPrivateIpAddress = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetRetainClassic(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.RetainClassic = &v
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
