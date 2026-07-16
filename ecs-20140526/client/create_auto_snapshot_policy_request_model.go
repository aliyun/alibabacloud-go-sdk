// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCopiedSnapshotsRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest
	GetCopiedSnapshotsRetentionDays() *int32
	SetCopyEncryptionConfiguration(v *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) *CreateAutoSnapshotPolicyRequest
	GetCopyEncryptionConfiguration() *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration
	SetEnableCrossRegionCopy(v bool) *CreateAutoSnapshotPolicyRequest
	GetEnableCrossRegionCopy() *bool
	SetOwnerId(v int64) *CreateAutoSnapshotPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateAutoSnapshotPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateAutoSnapshotPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoSnapshotPolicyRequest
	GetResourceOwnerId() *int64
	SetStorageLocationArn(v string) *CreateAutoSnapshotPolicyRequest
	GetStorageLocationArn() *string
	SetTag(v []*CreateAutoSnapshotPolicyRequestTag) *CreateAutoSnapshotPolicyRequest
	GetTag() []*CreateAutoSnapshotPolicyRequestTag
	SetTargetCopyRegions(v string) *CreateAutoSnapshotPolicyRequest
	GetTargetCopyRegions() *string
	SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyName() *string
	SetRegionId(v string) *CreateAutoSnapshotPolicyRequest
	GetRegionId() *string
	SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest
	GetRepeatWeekdays() *string
	SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
	SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest
	GetTimePoints() *string
}

type CreateAutoSnapshotPolicyRequest struct {
	// The retention period of cross-region snapshot replicas. Unit: days. Valid values:
	//
	// - -1: Snapshot replicas are permanently retained.
	//
	// - 1 to 65535: Snapshot replicas are retained for the specified number of days.
	//
	// Default value: -1.
	//
	// example:
	//
	// 30
	CopiedSnapshotsRetentionDays *int32 `json:"CopiedSnapshotsRetentionDays,omitempty" xml:"CopiedSnapshotsRetentionDays,omitempty"`
	// The backup encryption parameter object for snapshot geo-redundancy.
	CopyEncryptionConfiguration *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration `json:"CopyEncryptionConfiguration,omitempty" xml:"CopyEncryptionConfiguration,omitempty" type:"Struct"`
	// Specifies whether to enable automatic cross-region replication.
	//
	// - true: enables automatic cross-region replication.
	//
	// - false: disables automatic cross-region replication.
	//
	// example:
	//
	// false
	EnableCrossRegionCopy *bool  `json:"EnableCrossRegionCopy,omitempty" xml:"EnableCrossRegionCopy,omitempty"`
	OwnerId               *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2kkmhmhs****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The tags of the automatic snapshot policy.
	Tag []*CreateAutoSnapshotPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The destination region to which snapshots are replicated. You can specify only one destination region.
	//
	// example:
	//
	// ["cn-hangzhou"]
	TargetCopyRegions *string `json:"TargetCopyRegions,omitempty" xml:"TargetCopyRegions,omitempty"`
	// The name of the automatic snapshot policy. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// Default value: null.
	//
	// example:
	//
	// TestName
	AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty" xml:"autoSnapshotPolicyName,omitempty"`
	// The region to which the automatic snapshot policy belongs. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The days of the week on which automatic snapshots are created. Unit: days. The cycle is weekly. Valid values: 1 to 7. For example, 1 indicates Monday. Format description:
	//
	// - The parameter value must be a JSON array. For example, ["1"\\] indicates that automatic snapshots are created every Monday.
	//
	// - To create multiple automatic snapshots within a week, specify multiple days separated by commas (,). You can specify a maximum of 7 days. For example, ["1","3","5"\\] indicates that automatic snapshots are created every Monday, Wednesday, and Friday.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2"]
	RepeatWeekdays *string `json:"repeatWeekdays,omitempty" xml:"repeatWeekdays,omitempty"`
	// The retention period of automatic snapshots. Unit: days. Valid values:
	//
	// - -1: Automatic snapshots are permanently retained.
	//
	// - 1 to 65535: Automatic snapshots are retained for the specified number of days.
	//
	// Default value: -1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The points in time at which automatic snapshots are created. The time is displayed in UTC+8. Unit: hours. Valid values: 0 to 23, which represent the 24 points in time from 00:00 to 23:00. For example, 1 indicates 01:00. Format description:
	//
	// - The parameter value must be a JSON array. For example, ["1"\\] indicates that automatic snapshots are created at 01:00.
	//
	// - To create multiple automatic snapshots within a day, specify multiple points in time separated by commas (,). You can specify a maximum of 24 points in time. For example, ["1","3","5"\\] indicates that automatic snapshots are created at 01:00, 03:00, and 05:00.
	//
	// > If a disk contains a large amount of data and the time required to create an automatic snapshot exceeds the interval between two consecutive points in time, the next point in time is skipped. For example, you set 09:00, 10:00, 11:00, and 12:00 as the points in time for automatic snapshot creation. The snapshot creation starts at 09:00 and is completed at 10:20, which takes 80 minutes. The system skips the 10:00 point in time and creates the next automatic snapshot at 11:00.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["0", "1", … "23"]
	TimePoints *string `json:"timePoints,omitempty" xml:"timePoints,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) GetCopiedSnapshotsRetentionDays() *int32 {
	return s.CopiedSnapshotsRetentionDays
}

func (s *CreateAutoSnapshotPolicyRequest) GetCopyEncryptionConfiguration() *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	return s.CopyEncryptionConfiguration
}

func (s *CreateAutoSnapshotPolicyRequest) GetEnableCrossRegionCopy() *bool {
	return s.EnableCrossRegionCopy
}

func (s *CreateAutoSnapshotPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoSnapshotPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAutoSnapshotPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoSnapshotPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoSnapshotPolicyRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *CreateAutoSnapshotPolicyRequest) GetTag() []*CreateAutoSnapshotPolicyRequestTag {
	return s.Tag
}

func (s *CreateAutoSnapshotPolicyRequest) GetTargetCopyRegions() *string {
	return s.TargetCopyRegions
}

func (s *CreateAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *CreateAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAutoSnapshotPolicyRequest) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *CreateAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateAutoSnapshotPolicyRequest) GetTimePoints() *string {
	return s.TimePoints
}

func (s *CreateAutoSnapshotPolicyRequest) SetCopiedSnapshotsRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.CopiedSnapshotsRetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetCopyEncryptionConfiguration(v *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) *CreateAutoSnapshotPolicyRequest {
	s.CopyEncryptionConfiguration = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetEnableCrossRegionCopy(v bool) *CreateAutoSnapshotPolicyRequest {
	s.EnableCrossRegionCopy = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetOwnerId(v int64) *CreateAutoSnapshotPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetResourceGroupId(v string) *CreateAutoSnapshotPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetResourceOwnerAccount(v string) *CreateAutoSnapshotPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetResourceOwnerId(v int64) *CreateAutoSnapshotPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetStorageLocationArn(v string) *CreateAutoSnapshotPolicyRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTag(v []*CreateAutoSnapshotPolicyRequestTag) *CreateAutoSnapshotPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTargetCopyRegions(v string) *CreateAutoSnapshotPolicyRequest {
	s.TargetCopyRegions = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRegionId(v string) *CreateAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) Validate() error {
	if s.CopyEncryptionConfiguration != nil {
		if err := s.CopyEncryptionConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration struct {
	// > This parameter is not publicly available.
	Arn []*CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable encryption for cross-region snapshot backup. Valid values:
	//
	// - true: enables encryption.
	//
	// - false: disables encryption.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The KMS key ID used for encrypted cross-region snapshot backup.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GetArn() []*CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	return s.Arn
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) SetArn(v []*CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	s.Arn = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) SetEncrypted(v bool) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	s.Encrypted = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) SetKMSKeyId(v string) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	s.KMSKeyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn struct {
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 1000000000
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// hide
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// hide
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) SetAssumeRoleFor(v int64) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) SetRoleType(v string) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	s.RoleType = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) SetRolearn(v string) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	s.Rolearn = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) Validate() error {
	return dara.Validate(s)
}

type CreateAutoSnapshotPolicyRequestTag struct {
	// The tag key of the automatic snapshot policy. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the automatic snapshot policy. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with acs:. The tag value cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAutoSnapshotPolicyRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAutoSnapshotPolicyRequestTag) SetKey(v string) *CreateAutoSnapshotPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestTag) SetValue(v string) *CreateAutoSnapshotPolicyRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestTag) Validate() error {
	return dara.Validate(s)
}
