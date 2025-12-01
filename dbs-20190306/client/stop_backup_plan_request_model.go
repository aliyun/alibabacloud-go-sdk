// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *StopBackupPlanRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *StopBackupPlanRequest
	GetClientToken() *string
	SetOwnerId(v string) *StopBackupPlanRequest
	GetOwnerId() *string
	SetStopMethod(v string) *StopBackupPlanRequest
	GetStopMethod() *string
}

type StopBackupPlanRequest struct {
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The method that is used to stop the backup schedule. Valid values:
	//
	// 	- ALL: stops the backup schedule, full data backup tasks, incremental log backup tasks, and restore tasks
	//
	// 	- PLAN: stops only the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	StopMethod *string `json:"StopMethod,omitempty" xml:"StopMethod,omitempty"`
}

func (s StopBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s StopBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *StopBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *StopBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StopBackupPlanRequest) GetStopMethod() *string {
	return s.StopMethod
}

func (s *StopBackupPlanRequest) SetBackupPlanId(v string) *StopBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *StopBackupPlanRequest) SetClientToken(v string) *StopBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *StopBackupPlanRequest) SetOwnerId(v string) *StopBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *StopBackupPlanRequest) SetStopMethod(v string) *StopBackupPlanRequest {
	s.StopMethod = &v
	return s
}

func (s *StopBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
