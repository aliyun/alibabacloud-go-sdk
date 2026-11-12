// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationBackupPolicyRequest
	GetApplicationId() *string
	SetApplicationType(v string) *ModifyApplicationBackupPolicyRequest
	GetApplicationType() *string
	SetPreferredBackupPeriod(v string) *ModifyApplicationBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupRetentionPeriod(v string) *ModifyApplicationBackupPolicyRequest
	GetPreferredBackupRetentionPeriod() *string
	SetPreferredBackupTime(v string) *ModifyApplicationBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetRegionId(v string) *ModifyApplicationBackupPolicyRequest
	GetRegionId() *string
}

type ModifyApplicationBackupPolicyRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type, such as polarclaw. This parameter is required. If it is not provided, the service returns InvalidApplicationType.Malformed. Valid values: polarclaw, supabase, and polarclawsaas. Note: The ApplicationType field in the DescribeApplications response contains the polardb_ prefix (such as polardb_polarclaw). When you call this operation, remove the prefix and use the short format.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The days of the week when automatic backups are performed. Separate multiple values with commas. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	//
	// example:
	//
	// Monday,Tuesday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The retention period of level-1 backups. Unit: days.
	//
	// example:
	//
	// 7
	PreferredBackupRetentionPeriod *string `json:"PreferredBackupRetentionPeriod,omitempty" xml:"PreferredBackupRetentionPeriod,omitempty"`
	// The time window when automatic backups are performed. Specify the time in the HH:mmZ-HH:mmZ format in UTC. The start time and end time must be on the hour and exactly one hour apart.
	//
	// example:
	//
	// 07:00Z-08:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The region ID of the application.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyApplicationBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationBackupPolicyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationBackupPolicyRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ModifyApplicationBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyApplicationBackupPolicyRequest) GetPreferredBackupRetentionPeriod() *string {
	return s.PreferredBackupRetentionPeriod
}

func (s *ModifyApplicationBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyApplicationBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApplicationBackupPolicyRequest) SetApplicationId(v string) *ModifyApplicationBackupPolicyRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationBackupPolicyRequest) SetApplicationType(v string) *ModifyApplicationBackupPolicyRequest {
	s.ApplicationType = &v
	return s
}

func (s *ModifyApplicationBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyApplicationBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyApplicationBackupPolicyRequest) SetPreferredBackupRetentionPeriod(v string) *ModifyApplicationBackupPolicyRequest {
	s.PreferredBackupRetentionPeriod = &v
	return s
}

func (s *ModifyApplicationBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyApplicationBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyApplicationBackupPolicyRequest) SetRegionId(v string) *ModifyApplicationBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApplicationBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
