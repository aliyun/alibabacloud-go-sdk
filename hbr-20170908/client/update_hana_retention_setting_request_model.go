// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaRetentionSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateHanaRetentionSettingRequest
	GetClusterId() *string
	SetDatabaseName(v string) *UpdateHanaRetentionSettingRequest
	GetDatabaseName() *string
	SetDisabled(v bool) *UpdateHanaRetentionSettingRequest
	GetDisabled() *bool
	SetRetentionDays(v int64) *UpdateHanaRetentionSettingRequest
	GetRetentionDays() *int64
	SetSchedule(v string) *UpdateHanaRetentionSettingRequest
	GetSchedule() *string
	SetVaultId(v string) *UpdateHanaRetentionSettingRequest
	GetVaultId() *string
}

type UpdateHanaRetentionSettingRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-00024vyjj******srrq
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Specifies whether to permanently retain the backup. Valid values:
	//
	// 	- true: The backup is permanently retained.
	//
	// 	- false: The backup is retained for the specified number of days.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The retention period of the backup data. Unit: days. If you set the Disabled parameter to false, the backup is retained for the number of days specified by this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	RetentionDays *int64 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The policy to update the retention period. Format: `I|{startTime}|{interval}`. The retention period is updated at an interval of {interval} starting from {startTime}.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour, and P1D indicates an interval of one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// I|0|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-000fb0v2ly******k6
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaRetentionSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaRetentionSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaRetentionSettingRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateHanaRetentionSettingRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *UpdateHanaRetentionSettingRequest) GetDisabled() *bool {
	return s.Disabled
}

func (s *UpdateHanaRetentionSettingRequest) GetRetentionDays() *int64 {
	return s.RetentionDays
}

func (s *UpdateHanaRetentionSettingRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateHanaRetentionSettingRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateHanaRetentionSettingRequest) SetClusterId(v string) *UpdateHanaRetentionSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetDatabaseName(v string) *UpdateHanaRetentionSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetDisabled(v bool) *UpdateHanaRetentionSettingRequest {
	s.Disabled = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetRetentionDays(v int64) *UpdateHanaRetentionSettingRequest {
	s.RetentionDays = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetSchedule(v string) *UpdateHanaRetentionSettingRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetVaultId(v string) *UpdateHanaRetentionSettingRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) Validate() error {
	return dara.Validate(s)
}
