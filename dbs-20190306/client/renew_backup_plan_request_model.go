// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *RenewBackupPlanRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *RenewBackupPlanRequest
	GetClientToken() *string
	SetOwnerId(v string) *RenewBackupPlanRequest
	GetOwnerId() *string
	SetPeriod(v string) *RenewBackupPlanRequest
	GetPeriod() *string
	SetUsedTime(v int32) *RenewBackupPlanRequest
	GetUsedTime() *int32
}

type RenewBackupPlanRequest struct {
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
	// HKAJHFIUEQWBFIJSNFO****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to use yearly subscription or monthly subscription for the instance. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If the Period parameter is set to Year, the value of the UsedTime parameter ranges from 1 to 9.
	//
	// 	- If the Period parameter is set to Month, the value of the UsedTime parameter ranges from 1 to 11.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s RenewBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *RenewBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *RenewBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *RenewBackupPlanRequest) GetPeriod() *string {
	return s.Period
}

func (s *RenewBackupPlanRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *RenewBackupPlanRequest) SetBackupPlanId(v string) *RenewBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *RenewBackupPlanRequest) SetClientToken(v string) *RenewBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewBackupPlanRequest) SetOwnerId(v string) *RenewBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewBackupPlanRequest) SetPeriod(v string) *RenewBackupPlanRequest {
	s.Period = &v
	return s
}

func (s *RenewBackupPlanRequest) SetUsedTime(v int32) *RenewBackupPlanRequest {
	s.UsedTime = &v
	return s
}

func (s *RenewBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
