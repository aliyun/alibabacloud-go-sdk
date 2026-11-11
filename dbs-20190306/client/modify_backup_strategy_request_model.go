// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupLogIntervalSeconds(v int32) *ModifyBackupStrategyRequest
	GetBackupLogIntervalSeconds() *int32
	SetBackupPeriod(v string) *ModifyBackupStrategyRequest
	GetBackupPeriod() *string
	SetBackupPlanId(v string) *ModifyBackupStrategyRequest
	GetBackupPlanId() *string
	SetBackupStartTime(v string) *ModifyBackupStrategyRequest
	GetBackupStartTime() *string
	SetBackupStrategyType(v string) *ModifyBackupStrategyRequest
	GetBackupStrategyType() *string
	SetClientToken(v string) *ModifyBackupStrategyRequest
	GetClientToken() *string
	SetOwnerId(v string) *ModifyBackupStrategyRequest
	GetOwnerId() *string
}

type ModifyBackupStrategyRequest struct {
	// The incremental backup interval in seconds.
	//
	// > Only physical backups are supported.
	//
	// example:
	//
	// 1000
	BackupLogIntervalSeconds *int32 `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	// The full backup cycle. Valid values:
	//
	// - Monday: Monday
	//
	// - Tuesday: Tuesday
	//
	// - Wednesday: Wednesday
	//
	// - Thursday: Thursday
	//
	// - Friday: Friday
	//
	// - Saturday: Saturday
	//
	// - Sunday: Sunday
	//
	// > You must specify at least two days of the week for BackupPeriod and separate the days with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday,Tuesday,Wednesday
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01XXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The start time of the full backup. Format: HH:mm.
	//
	// > This parameter is required.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The full backup epoch. Valid values:
	//
	// - **simple**: periodic backup. This value is used together with BackupPeriod and BackupStartTime.
	//
	// - **manual**: manual backup.
	//
	// > Default value: **simple**.
	//
	// example:
	//
	// simple
	BackupStrategyType *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyBackupStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupStrategyRequest) GetBackupLogIntervalSeconds() *int32 {
	return s.BackupLogIntervalSeconds
}

func (s *ModifyBackupStrategyRequest) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *ModifyBackupStrategyRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupStrategyRequest) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *ModifyBackupStrategyRequest) GetBackupStrategyType() *string {
	return s.BackupStrategyType
}

func (s *ModifyBackupStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBackupStrategyRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyBackupStrategyRequest) SetBackupLogIntervalSeconds(v int32) *ModifyBackupStrategyRequest {
	s.BackupLogIntervalSeconds = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupPeriod(v string) *ModifyBackupStrategyRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupPlanId(v string) *ModifyBackupStrategyRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupStartTime(v string) *ModifyBackupStrategyRequest {
	s.BackupStartTime = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupStrategyType(v string) *ModifyBackupStrategyRequest {
	s.BackupStrategyType = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetClientToken(v string) *ModifyBackupStrategyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetOwnerId(v string) *ModifyBackupStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupStrategyRequest) Validate() error {
	return dara.Validate(s)
}
