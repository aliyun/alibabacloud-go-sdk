// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniRestorePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v string) *CreateUniRestorePlanRequest
	GetDatabase() *string
	SetInstanceUuid(v string) *CreateUniRestorePlanRequest
	GetInstanceUuid() *string
	SetPolicyId(v int64) *CreateUniRestorePlanRequest
	GetPolicyId() *int64
	SetResetScn(v string) *CreateUniRestorePlanRequest
	GetResetScn() *string
	SetResetTime(v string) *CreateUniRestorePlanRequest
	GetResetTime() *string
	SetRestoreInfo(v string) *CreateUniRestorePlanRequest
	GetRestoreInfo() *string
	SetTimePoint(v int64) *CreateUniRestorePlanRequest
	GetTimePoint() *int64
}

type CreateUniRestorePlanRequest struct {
	// The name of the database.
	//
	// example:
	//
	// qtc
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The unique identifier of the database backup client on the destination server for restoration.
	//
	// >Call the [DescribeUniBackupDatabase](~~DescribeUniBackupDatabase~~) operation to obtain this parameter.
	//
	// example:
	//
	// ac457b30598d11ed800000163e02****
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The ID of the database anti-ransomware backup policy.
	//
	// >Call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The **reset_scn*	- value of the selected record from the recoverable points in time when you query backups for an Oracle database.
	//
	// >Call the [DescribeUniRecoverableList](~~DescribeUniRecoverableList~~) operation to obtain this parameter.
	//
	// example:
	//
	// 925702.0
	ResetScn *string `json:"ResetScn,omitempty" xml:"ResetScn,omitempty"`
	// The **reset_time*	- value of the selected record from the recoverable points in time when you query backups for an Oracle database.
	//
	// >Call the [DescribeUniRecoverableList](~~DescribeUniRecoverableList~~) operation to obtain this parameter.
	//
	// example:
	//
	// 2022-10-29 01:06:24
	ResetTime *string `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	// The database restoration information when the database type is MSSQL. The value is a JSON string. Valid values:
	//
	// - **name**: the name of the database.
	//
	// - **files**: the file path of the database.
	//
	// >Call the [DescribeUniRecoverableList](~~DescribeUniRecoverableList~~) operation to obtain this parameter.
	//
	// example:
	//
	// {"files": {"qtc":"F:\\\\database\\\\qtc.mdf","qtc_log":"F:\\\\database\\\\qtc_0.ldf"},
	//
	// "name":"qtc"}
	RestoreInfo *string `json:"RestoreInfo,omitempty" xml:"RestoreInfo,omitempty"`
	// The point in time to which you want to restore the database.
	//
	// >Call the [DescribeRestorePlans](~~DescribeRestorePlans~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656957664000
	TimePoint *int64 `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s CreateUniRestorePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUniRestorePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateUniRestorePlanRequest) GetDatabase() *string {
	return s.Database
}

func (s *CreateUniRestorePlanRequest) GetInstanceUuid() *string {
	return s.InstanceUuid
}

func (s *CreateUniRestorePlanRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *CreateUniRestorePlanRequest) GetResetScn() *string {
	return s.ResetScn
}

func (s *CreateUniRestorePlanRequest) GetResetTime() *string {
	return s.ResetTime
}

func (s *CreateUniRestorePlanRequest) GetRestoreInfo() *string {
	return s.RestoreInfo
}

func (s *CreateUniRestorePlanRequest) GetTimePoint() *int64 {
	return s.TimePoint
}

func (s *CreateUniRestorePlanRequest) SetDatabase(v string) *CreateUniRestorePlanRequest {
	s.Database = &v
	return s
}

func (s *CreateUniRestorePlanRequest) SetInstanceUuid(v string) *CreateUniRestorePlanRequest {
	s.InstanceUuid = &v
	return s
}

func (s *CreateUniRestorePlanRequest) SetPolicyId(v int64) *CreateUniRestorePlanRequest {
	s.PolicyId = &v
	return s
}

func (s *CreateUniRestorePlanRequest) SetResetScn(v string) *CreateUniRestorePlanRequest {
	s.ResetScn = &v
	return s
}

func (s *CreateUniRestorePlanRequest) SetResetTime(v string) *CreateUniRestorePlanRequest {
	s.ResetTime = &v
	return s
}

func (s *CreateUniRestorePlanRequest) SetRestoreInfo(v string) *CreateUniRestorePlanRequest {
	s.RestoreInfo = &v
	return s
}

func (s *CreateUniRestorePlanRequest) SetTimePoint(v int64) *CreateUniRestorePlanRequest {
	s.TimePoint = &v
	return s
}

func (s *CreateUniRestorePlanRequest) Validate() error {
	return dara.Validate(s)
}
