// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v string) *CreateBackupPolicyRequest
	GetBackupRetentionPeriod() *string
	SetDBInstanceId(v string) *CreateBackupPolicyRequest
	GetDBInstanceId() *string
	SetPreferredBackupPeriod(v string) *CreateBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *CreateBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetRegionId(v string) *CreateBackupPolicyRequest
	GetRegionId() *string
}

type CreateBackupPolicyRequest struct {
	// The number of days to retain backups.
	//
	// example:
	//
	// 8
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-2ze0eb0w182xh8549
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The backup cycle. Specify the days of the week for backups. Separate multiple days with commas (`,`).
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The UTC time range to perform the backup. For example, `12:00Z-13:00Z` means that the backup starts between 12:00 and 13:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00Z-11:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyRequest) GetBackupRetentionPeriod() *string {
	return s.BackupRetentionPeriod
}

func (s *CreateBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *CreateBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *CreateBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBackupPolicyRequest) SetBackupRetentionPeriod(v string) *CreateBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetDBInstanceId(v string) *CreateBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetPreferredBackupPeriod(v string) *CreateBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetPreferredBackupTime(v string) *CreateBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetRegionId(v string) *CreateBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
