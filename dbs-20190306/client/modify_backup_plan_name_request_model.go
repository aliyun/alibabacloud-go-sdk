// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPlanNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyBackupPlanNameRequest
	GetBackupPlanId() *string
	SetBackupPlanName(v string) *ModifyBackupPlanNameRequest
	GetBackupPlanName() *string
	SetClientToken(v string) *ModifyBackupPlanNameRequest
	GetClientToken() *string
	SetOwnerId(v string) *ModifyBackupPlanNameRequest
	GetOwnerId() *string
}

type ModifyBackupPlanNameRequest struct {
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0XXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The name of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0XXXX
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyBackupPlanNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPlanNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanNameRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupPlanNameRequest) GetBackupPlanName() *string {
	return s.BackupPlanName
}

func (s *ModifyBackupPlanNameRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBackupPlanNameRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyBackupPlanNameRequest) SetBackupPlanId(v string) *ModifyBackupPlanNameRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) SetBackupPlanName(v string) *ModifyBackupPlanNameRequest {
	s.BackupPlanName = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) SetClientToken(v string) *ModifyBackupPlanNameRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) SetOwnerId(v string) *ModifyBackupPlanNameRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) Validate() error {
	return dara.Validate(s)
}
