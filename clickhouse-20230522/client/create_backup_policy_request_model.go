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
	// The number of days for which you can retain the backup data.
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
	// The backup cycle, which indicates the day of the week when the system regularly backs up data. Separate multiple dates with commas (`,`).
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup time window within which the backup task is performed. The time is displayed in `UTC`. For example, `12:00Z-13:00Z` indicates that the backup time window ranges from `12:00` (UTC) to `13:00` `(UTC)`.
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
