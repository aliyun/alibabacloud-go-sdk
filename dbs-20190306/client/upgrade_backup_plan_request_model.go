// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *UpgradeBackupPlanRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *UpgradeBackupPlanRequest
	GetClientToken() *string
	SetInstanceClass(v string) *UpgradeBackupPlanRequest
	GetInstanceClass() *string
	SetOwnerId(v string) *UpgradeBackupPlanRequest
	GetOwnerId() *string
}

type UpgradeBackupPlanRequest struct {
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01e****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The specifications of the instance. Valid values:
	//
	// 	- micro
	//
	// 	- small
	//
	// 	- medium
	//
	// 	- large
	//
	// 	- xlarge
	//
	// This parameter is required.
	//
	// example:
	//
	// micro
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UpgradeBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *UpgradeBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeBackupPlanRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *UpgradeBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpgradeBackupPlanRequest) SetBackupPlanId(v string) *UpgradeBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *UpgradeBackupPlanRequest) SetClientToken(v string) *UpgradeBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeBackupPlanRequest) SetInstanceClass(v string) *UpgradeBackupPlanRequest {
	s.InstanceClass = &v
	return s
}

func (s *UpgradeBackupPlanRequest) SetOwnerId(v string) *UpgradeBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
