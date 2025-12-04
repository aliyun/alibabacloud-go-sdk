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
	SetDBClusterId(v string) *CreateBackupPolicyRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *CreateBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *CreateBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetRegionId(v string) *CreateBackupPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type CreateBackupPolicyRequest struct {
	// The backup retention period. The default retention period is seven days. Valid values: 7 to 730. Unit: day.
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
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The day of a week when the system regularly backs up data. If you specify multiple days of a week, separate them with commas (,). Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. Specify the time in the ISO 8601 standard in the HH:mmZ-HH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// For example, if you set the backup window to 00:00Z-01:00Z, the data of the cluster can be backed up from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00Z-11:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *CreateBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *CreateBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBackupPolicyRequest) SetBackupRetentionPeriod(v string) *CreateBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetDBClusterId(v string) *CreateBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetOwnerAccount(v string) *CreateBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetOwnerId(v int64) *CreateBackupPolicyRequest {
	s.OwnerId = &v
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

func (s *CreateBackupPolicyRequest) SetResourceOwnerAccount(v string) *CreateBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetResourceOwnerId(v int64) *CreateBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
