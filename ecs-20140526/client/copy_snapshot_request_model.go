// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v []*CopySnapshotRequestArn) *CopySnapshotRequest
	GetArn() []*CopySnapshotRequestArn
	SetClientToken(v string) *CopySnapshotRequest
	GetClientToken() *string
	SetDestinationRegionId(v string) *CopySnapshotRequest
	GetDestinationRegionId() *string
	SetDestinationSnapshotDescription(v string) *CopySnapshotRequest
	GetDestinationSnapshotDescription() *string
	SetDestinationSnapshotName(v string) *CopySnapshotRequest
	GetDestinationSnapshotName() *string
	SetDestinationStorageLocationArn(v string) *CopySnapshotRequest
	GetDestinationStorageLocationArn() *string
	SetEncrypted(v bool) *CopySnapshotRequest
	GetEncrypted() *bool
	SetKMSKeyId(v string) *CopySnapshotRequest
	GetKMSKeyId() *string
	SetOwnerId(v int64) *CopySnapshotRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CopySnapshotRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CopySnapshotRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CopySnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopySnapshotRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *CopySnapshotRequest
	GetRetentionDays() *int32
	SetSnapshotId(v string) *CopySnapshotRequest
	GetSnapshotId() *string
	SetTag(v []*CopySnapshotRequestTag) *CopySnapshotRequest
	GetTag() []*CopySnapshotRequestTag
}

type CopySnapshotRequest struct {
	// > This parameter is in invitational preview and is not publicly available.
	Arn []*CopySnapshotRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Ensures the idempotence of the request. The value is generated from your client and must be unique among different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the destination region to which to copy the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// us-east-1
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The description of the new snapshot. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// Default value: null.
	//
	// This parameter is required.
	//
	// example:
	//
	// CopySnapshotDemo
	DestinationSnapshotDescription *string `json:"DestinationSnapshotDescription,omitempty" xml:"DestinationSnapshotDescription,omitempty"`
	// The name of the new snapshot. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It cannot start with http:// or https://. The name can contain letters, digits, and Unicode characters that are categorized under the letter classification. It can also contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Default value: null.
	//
	// This parameter is required.
	//
	// example:
	//
	// CopySnapshotDemo
	DestinationSnapshotName *string `json:"DestinationSnapshotName,omitempty" xml:"DestinationSnapshotName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	DestinationStorageLocationArn *string `json:"DestinationStorageLocationArn,omitempty" xml:"DestinationStorageLocationArn,omitempty"`
	// Specifies whether to encrypt the cloud disk. Valid values:
	//
	// - true: encrypts the cloud disk.
	//
	// - false: does not encrypt the cloud disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The customer master key (CMK) in the destination region.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the source snapshot. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the new snapshot, in days. The snapshot undergoes automatic release when the retention period expires. Valid values: 1 to 65536.
	//
	// Default value: null, which indicates that the snapshot does not undergo automatic release.
	//
	// example:
	//
	// 60
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the source snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tag information of the new snapshot.
	Tag []*CopySnapshotRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CopySnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotRequest) GoString() string {
	return s.String()
}

func (s *CopySnapshotRequest) GetArn() []*CopySnapshotRequestArn {
	return s.Arn
}

func (s *CopySnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CopySnapshotRequest) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CopySnapshotRequest) GetDestinationSnapshotDescription() *string {
	return s.DestinationSnapshotDescription
}

func (s *CopySnapshotRequest) GetDestinationSnapshotName() *string {
	return s.DestinationSnapshotName
}

func (s *CopySnapshotRequest) GetDestinationStorageLocationArn() *string {
	return s.DestinationStorageLocationArn
}

func (s *CopySnapshotRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CopySnapshotRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CopySnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopySnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopySnapshotRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CopySnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopySnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopySnapshotRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CopySnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CopySnapshotRequest) GetTag() []*CopySnapshotRequestTag {
	return s.Tag
}

func (s *CopySnapshotRequest) SetArn(v []*CopySnapshotRequestArn) *CopySnapshotRequest {
	s.Arn = v
	return s
}

func (s *CopySnapshotRequest) SetClientToken(v string) *CopySnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CopySnapshotRequest) SetDestinationRegionId(v string) *CopySnapshotRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CopySnapshotRequest) SetDestinationSnapshotDescription(v string) *CopySnapshotRequest {
	s.DestinationSnapshotDescription = &v
	return s
}

func (s *CopySnapshotRequest) SetDestinationSnapshotName(v string) *CopySnapshotRequest {
	s.DestinationSnapshotName = &v
	return s
}

func (s *CopySnapshotRequest) SetDestinationStorageLocationArn(v string) *CopySnapshotRequest {
	s.DestinationStorageLocationArn = &v
	return s
}

func (s *CopySnapshotRequest) SetEncrypted(v bool) *CopySnapshotRequest {
	s.Encrypted = &v
	return s
}

func (s *CopySnapshotRequest) SetKMSKeyId(v string) *CopySnapshotRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CopySnapshotRequest) SetOwnerId(v int64) *CopySnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *CopySnapshotRequest) SetRegionId(v string) *CopySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CopySnapshotRequest) SetResourceGroupId(v string) *CopySnapshotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CopySnapshotRequest) SetResourceOwnerAccount(v string) *CopySnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopySnapshotRequest) SetResourceOwnerId(v int64) *CopySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopySnapshotRequest) SetRetentionDays(v int32) *CopySnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CopySnapshotRequest) SetSnapshotId(v string) *CopySnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *CopySnapshotRequest) SetTag(v []*CopySnapshotRequestTag) *CopySnapshotRequest {
	s.Tag = v
	return s
}

func (s *CopySnapshotRequest) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CopySnapshotRequestArn struct {
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CopySnapshotRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotRequestArn) GoString() string {
	return s.String()
}

func (s *CopySnapshotRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CopySnapshotRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CopySnapshotRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CopySnapshotRequestArn) SetAssumeRoleFor(v int64) *CopySnapshotRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CopySnapshotRequestArn) SetRoleType(v string) *CopySnapshotRequestArn {
	s.RoleType = &v
	return s
}

func (s *CopySnapshotRequestArn) SetRolearn(v string) *CopySnapshotRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CopySnapshotRequestArn) Validate() error {
	return dara.Validate(s)
}

type CopySnapshotRequestTag struct {
	// The tag key of the new snapshot. Once specified, the tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the new snapshot. Once specified, the tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CopySnapshotRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotRequestTag) GoString() string {
	return s.String()
}

func (s *CopySnapshotRequestTag) GetKey() *string {
	return s.Key
}

func (s *CopySnapshotRequestTag) GetValue() *string {
	return s.Value
}

func (s *CopySnapshotRequestTag) SetKey(v string) *CopySnapshotRequestTag {
	s.Key = &v
	return s
}

func (s *CopySnapshotRequestTag) SetValue(v string) *CopySnapshotRequestTag {
	s.Value = &v
	return s
}

func (s *CopySnapshotRequestTag) Validate() error {
	return dara.Validate(s)
}
