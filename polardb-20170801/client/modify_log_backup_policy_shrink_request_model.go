// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedLogPoliciesShrink(v string) *ModifyLogBackupPolicyShrinkRequest
	GetAdvancedLogPoliciesShrink() *string
	SetDBClusterId(v string) *ModifyLogBackupPolicyShrinkRequest
	GetDBClusterId() *string
	SetLogBackupAnotherRegionRegion(v string) *ModifyLogBackupPolicyShrinkRequest
	GetLogBackupAnotherRegionRegion() *string
	SetLogBackupAnotherRegionRetentionPeriod(v string) *ModifyLogBackupPolicyShrinkRequest
	GetLogBackupAnotherRegionRetentionPeriod() *string
	SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyShrinkRequest
	GetLogBackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyLogBackupPolicyShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLogBackupPolicyShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLogBackupPolicyShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyLogBackupPolicyShrinkRequest struct {
	AdvancedLogPoliciesShrink *string `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information of all clusters that are deployed in a specific region, such as the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region in which you want to store cross-region log backups. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// cn-hangzhou
	LogBackupAnotherRegionRegion *string `json:"LogBackupAnotherRegionRegion,omitempty" xml:"LogBackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region log backups. Valid values:
	//
	// 	- **0**: The cross-region backup feature is disabled.
	//
	// 	- **30 to 7300**: Cross-region log backups are retained for 30 to 7,300 days.
	//
	// 	- **-1**: The log backups are permanently retained.
	//
	// >  When you create a cluster, the default value of this parameter is **0**.
	//
	// example:
	//
	// 30
	LogBackupAnotherRegionRetentionPeriod *string `json:"LogBackupAnotherRegionRetentionPeriod,omitempty" xml:"LogBackupAnotherRegionRetentionPeriod,omitempty"`
	// The retention period of the log backups. Valid values:
	//
	// 	- 3 to 7300: The log backups are retained for 3 to 7,300 days.
	//
	// 	- \\-1: The log backups are permanently retained.
	//
	// example:
	//
	// 3
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLogBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetAdvancedLogPoliciesShrink() *string {
	return s.AdvancedLogPoliciesShrink
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetLogBackupAnotherRegionRegion() *string {
	return s.LogBackupAnotherRegionRegion
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetLogBackupAnotherRegionRetentionPeriod() *string {
	return s.LogBackupAnotherRegionRetentionPeriod
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLogBackupPolicyShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetAdvancedLogPoliciesShrink(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.AdvancedLogPoliciesShrink = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetLogBackupAnotherRegionRegion(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.LogBackupAnotherRegionRegion = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetLogBackupAnotherRegionRetentionPeriod(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.LogBackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
