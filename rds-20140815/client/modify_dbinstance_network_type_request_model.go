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
	ClassicExpiredDays *string `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	InstanceNetworkType                  *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount                         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PrivateIpAddress                     *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ReadWriteSplittingClassicExpiredDays *int32  `json:"ReadWriteSplittingClassicExpiredDays,omitempty" xml:"ReadWriteSplittingClassicExpiredDays,omitempty"`
	ReadWriteSplittingPrivateIpAddress   *string `json:"ReadWriteSplittingPrivateIpAddress,omitempty" xml:"ReadWriteSplittingPrivateIpAddress,omitempty"`
	ResourceOwnerAccount                 *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RetainClassic                        *string `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	VPCId                                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId                            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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
