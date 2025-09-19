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
	// The UUID of the Hybrid Backup Recovery (HBR) agent that is used to restore the data of the database on your server.
	//
	// >  You can call the [DescribeUniBackupDatabase](~~DescribeUniBackupDatabase~~) operation to query the UUID.
	//
	// example:
	//
	// ac457b30598d11ed800000163e02****
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The ID of the anti-ransomware policy.
	//
	// >  You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The identifier of the point in time for restoration in the backup version that you want to use. The database is an Oracle database.****
	//
	// >  You can call the [DescribeUniRecoverableList](~~DescribeUniRecoverableList~~) operation to query the value.
	//
	// example:
	//
	// 925702.0
	ResetScn *string `json:"ResetScn,omitempty" xml:"ResetScn,omitempty"`
	// The point in time for restoration in the backup version that you want to use. The database is an Oracle database.****
	//
	// >  You can call the [DescribeUniRecoverableList](~~DescribeUniRecoverableList~~) operation to query the value.
	//
	// example:
	//
	// 2022-10-29 01:06:24
	ResetTime *string `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	// The information about the database. This parameter is available when the database is a Microsoft SQL Server (MSSQL) database. The value is a JSON string. Valid values:
	//
	// 	- **name**: the name of the database
	//
	// 	- **files**: the path to the database files
	//
	// >  You can call the [DescribeUniRecoverableList](~~DescribeUniRecoverableList~~) operation to query the information.
	//
	// example:
	//
	// {"files": {"qtc":"F:\\\\database\\\\qtc.mdf","qtc_log":"F:\\\\database\\\\qtc_0.ldf"},
	//
	// "name":"qtc"}
	RestoreInfo *string `json:"RestoreInfo,omitempty" xml:"RestoreInfo,omitempty"`
	// The point in time to which you want to restore data.
	//
	// >  You can call the [DescribeRestorePlans](~~DescribeRestorePlans~~) operation to query the point in time.
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
