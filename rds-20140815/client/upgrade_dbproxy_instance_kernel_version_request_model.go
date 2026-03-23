// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBProxyInstanceKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetDBProxyEngineType() *string
	SetOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetTargetMinorVersion() *string
	SetUpgradeTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetUpgradeTime() *string
}

type UpgradeDBProxyInstanceKernelVersionRequest struct {
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBProxyEngineType    *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetMinorVersion   *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	UpgradeTime          *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
}

func (s UpgradeDBProxyInstanceKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBProxyInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetDBInstanceId(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetDBProxyEngineType(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetSwitchTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetTargetMinorVersion(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetUpgradeTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
