// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *UpgradeDBInstanceKernelVersionRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest
	GetTargetMinorVersion() *string
	SetUpgradeTime(v string) *UpgradeDBInstanceKernelVersionRequest
	GetUpgradeTime() *string
}

type UpgradeDBInstanceKernelVersionRequest struct {
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetMinorVersion   *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	UpgradeTime          *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSwitchTime(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetUpgradeTime(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
