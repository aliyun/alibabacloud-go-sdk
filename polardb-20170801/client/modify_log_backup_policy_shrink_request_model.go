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
	// The advanced backup policies.
	//
	// > - - This parameter is not supported for PolarDB for PostgreSQL (Oracle Compatible) or PolarDB for PostgreSQL.
	//
	// >
	//
	// > - - This parameter is supported only for clusters for which the BackupPolicyLevel parameter is set to Advanced.
	AdvancedLogPoliciesShrink *string `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty"`
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view information about all clusters in a specific region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The destination region for cross-region log backups. For information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > - - After you enable the advanced backup feature, this parameter is no longer valid. Use the AdvancedLogPolicies parameter instead.
	//
	// example:
	//
	// cn-hangzhou
	LogBackupAnotherRegionRegion *string `json:"LogBackupAnotherRegionRegion,omitempty" xml:"LogBackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region log backups. Valid values:
	//
	// - **0**: Disables the cross-region log backup feature.
	//
	// - **30 to 7300**: The retention period in days.
	//
	// - **-1**: long-term retention.
	//
	// > 	- 	- When you create a cluster, the default value of this parameter is **0**. This value disables the cross-region log backup feature.
	//
	// >
	//
	// > 	- - After you enable the advanced backup feature, this parameter is no longer valid. Use the AdvancedLogPolicies parameter instead.
	//
	// example:
	//
	// 30
	LogBackupAnotherRegionRetentionPeriod *string `json:"LogBackupAnotherRegionRetentionPeriod,omitempty" xml:"LogBackupAnotherRegionRetentionPeriod,omitempty"`
	// The retention period of log backups. Valid values:
	//
	// - 3 to 7300: The retention period in days.
	//
	// - -1: long-term retention.
	//
	// > 	- 	- After you enable the advanced backup feature, this parameter is no longer valid. Use the AdvancedLogPolicies parameter instead.
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
