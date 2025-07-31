// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpiredDays(v int32) *ModifyDBInstanceNetworkTypeRequest
	GetClassicExpiredDays() *int32
	SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetDBInstanceId() *string
	SetNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest
	GetResourceOwnerId() *int64
	SetRetainClassic(v string) *ModifyDBInstanceNetworkTypeRequest
	GetRetainClassic() *string
	SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetVpcId() *string
	SetZoneId(v string) *ModifyDBInstanceNetworkTypeRequest
	GetZoneId() *string
}

type ModifyDBInstanceNetworkTypeRequest struct {
	// The retention period of the original classic network address when you change the network type to VPC. Valid values: **14**, **30**, **60**, and **120**. Unit: days.
	//
	// > This parameter is required when the **NetworkType*	- parameter is set to **VPC*	- and the **RetainClassic*	- parameter is set to **True**.
	//
	// example:
	//
	// 30
	ClassicExpiredDays *int32 `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp11483712c1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type to switch to. Valid value:
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to retain the original classic network address when you change the network type to VPC. Valid values:
	//
	// 	- **True**: retains the original classic network address.
	//
	// 	- **False**: does not retain the original classic network address.
	//
	// >
	//
	// 	- This parameter is required when the **NetworkType*	- parameter is set to **VPC**.
	//
	// 	- If you set this parameter to **True**, you must also specify the **ClassicExpiredDays*	- parameter.
	//
	// example:
	//
	// False
	RetainClassic *string `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// > This parameter is required when the **NetworkType*	- parameter is set to **VPC**.
	//
	// example:
	//
	// vsw-bp1vj604nj5a9zz74****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// > This parameter is required when the **NetworkType*	- parameter is set to **VPC**.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 可用区ID，您可以通过调用[DescribeRegions](https://help.aliyun.com/document_detail/61933.html)接口查询可用区ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetClassicExpiredDays() *int32 {
	return s.ClassicExpiredDays
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *ModifyDBInstanceNetworkTypeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyDBInstanceNetworkTypeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetClassicExpiredDays(v int32) *ModifyDBInstanceNetworkTypeRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.NetworkType = &v
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

func (s *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVpcId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetZoneId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) Validate() error {
	return dara.Validate(s)
}
