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
	// The interval at which you want to perform incremental log backups. Unit: seconds.
	//
	// > This parameter takes effect only when physical backups are performed.
	//
	// example:
	//
	// 1000
	BackupLogIntervalSeconds *int32 `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	// The day of each week when the full backup task runs. Valid values:
	//
	// 	- Monday
	//
	// 	- Tuesday
	//
	// 	- Wednesday
	//
	// 	- Thursday
	//
	// 	- Friday
	//
	// 	- Saturday
	//
	// 	- Sunday
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01XXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The start time of the full backup task. Specify the time in the HH:mm format.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup method that you want to use for full backups. Valid values:
	//
	// 	- **simple**: scheduled backup. If you specify this value for the BackupStrategyType parameter, you must also specify the BackupPeriod and BackupStartTime parameters.
	//
	// 	- **Manual**: manual backup.
	//
	// > Default value: **simple**.
	//
	// example:
	//
	// simple
	BackupStrategyType *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
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
