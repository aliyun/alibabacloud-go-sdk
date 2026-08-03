// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CopyImageRequest
	GetClientToken() *string
	SetDestinationDescription(v string) *CopyImageRequest
	GetDestinationDescription() *string
	SetDestinationImageName(v string) *CopyImageRequest
	GetDestinationImageName() *string
	SetDestinationRegionId(v string) *CopyImageRequest
	GetDestinationRegionId() *string
	SetDryRun(v bool) *CopyImageRequest
	GetDryRun() *bool
	SetEncryptAlgorithm(v string) *CopyImageRequest
	GetEncryptAlgorithm() *string
	SetEncrypted(v bool) *CopyImageRequest
	GetEncrypted() *bool
	SetImageId(v string) *CopyImageRequest
	GetImageId() *string
	SetKMSKeyId(v string) *CopyImageRequest
	GetKMSKeyId() *string
	SetOwnerAccount(v string) *CopyImageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CopyImageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CopyImageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CopyImageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CopyImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopyImageRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CopyImageRequestTag) *CopyImageRequest
	GetTag() []*CopyImageRequestTag
}

type CopyImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the copied image. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is a description example.
	DestinationDescription *string `json:"DestinationDescription,omitempty" xml:"DestinationDescription,omitempty"`
	// The name of the copied image. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `aliyun` or `acs:`. The name cannot contain `http://` or `https://`. The name can contain digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// YourImageName
	DestinationImageName *string `json:"DestinationImageName,omitempty" xml:"DestinationImageName,omitempty"`
	// The ID of the destination region to which the image is copied.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The system checks whether your AccessKey pair is valid, whether Resource Access Management (RAM) user authorization is complete, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the error code `DryRunOperation` is returned.
	//
	// - false: performs a dry run and sends the request. If the request passes the dry run, a 2XX HTTP status code is returned and the operation is Normal.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// hide
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the copied image.
	//
	// - true: encrypts the copied image.
	//
	// - false: does not encrypt the copied image.
	//
	// - Not specified: determined by the backend. For more information, see the supplementary description below.
	//
	// Default value: not specified.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the source custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp1h46wfpjsjastc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the key used to encrypt the image.
	//
	// example:
	//
	// e522b26d-abf6-4e0d-b5da-04b7******3c
	KMSKeyId     *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the source custom image. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the copied image. If you do not set this parameter, the copied image belongs to the default resource group.
	//
	// > If you are a Resource Access Management (RAM) user and invoke this operation with the `ResourceGroupId` parameter left empty, take note of the following: if the RAM user does not have permissions on the default resource group, the error message `Forbidden: User not authorized to operate on the specified resource` is returned. Set the `ResourceGroupId` parameter to a resource group ID that the RAM user has permissions on, or grant the RAM user permissions on the default resource group by using the corresponding Alibaba Cloud account before you invoke this operation again. The `ResourceGroupId` parameter determines which resource group the replicated image belongs to.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CopyImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CopyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyImageRequest) GoString() string {
	return s.String()
}

func (s *CopyImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CopyImageRequest) GetDestinationDescription() *string {
	return s.DestinationDescription
}

func (s *CopyImageRequest) GetDestinationImageName() *string {
	return s.DestinationImageName
}

func (s *CopyImageRequest) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CopyImageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CopyImageRequest) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CopyImageRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CopyImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CopyImageRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CopyImageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CopyImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CopyImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopyImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopyImageRequest) GetTag() []*CopyImageRequestTag {
	return s.Tag
}

func (s *CopyImageRequest) SetClientToken(v string) *CopyImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CopyImageRequest) SetDestinationDescription(v string) *CopyImageRequest {
	s.DestinationDescription = &v
	return s
}

func (s *CopyImageRequest) SetDestinationImageName(v string) *CopyImageRequest {
	s.DestinationImageName = &v
	return s
}

func (s *CopyImageRequest) SetDestinationRegionId(v string) *CopyImageRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CopyImageRequest) SetDryRun(v bool) *CopyImageRequest {
	s.DryRun = &v
	return s
}

func (s *CopyImageRequest) SetEncryptAlgorithm(v string) *CopyImageRequest {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CopyImageRequest) SetEncrypted(v bool) *CopyImageRequest {
	s.Encrypted = &v
	return s
}

func (s *CopyImageRequest) SetImageId(v string) *CopyImageRequest {
	s.ImageId = &v
	return s
}

func (s *CopyImageRequest) SetKMSKeyId(v string) *CopyImageRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CopyImageRequest) SetOwnerAccount(v string) *CopyImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CopyImageRequest) SetOwnerId(v int64) *CopyImageRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyImageRequest) SetRegionId(v string) *CopyImageRequest {
	s.RegionId = &v
	return s
}

func (s *CopyImageRequest) SetResourceGroupId(v string) *CopyImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CopyImageRequest) SetResourceOwnerAccount(v string) *CopyImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyImageRequest) SetResourceOwnerId(v int64) *CopyImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyImageRequest) SetTag(v []*CopyImageRequestTag) *CopyImageRequest {
	s.Tag = v
	return s
}

func (s *CopyImageRequest) Validate() error {
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

type CopyImageRequestTag struct {
	// The tag key of the copied image. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the copied image. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CopyImageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CopyImageRequestTag) GoString() string {
	return s.String()
}

func (s *CopyImageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CopyImageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CopyImageRequestTag) SetKey(v string) *CopyImageRequestTag {
	s.Key = &v
	return s
}

func (s *CopyImageRequestTag) SetValue(v string) *CopyImageRequestTag {
	s.Value = &v
	return s
}

func (s *CopyImageRequestTag) Validate() error {
	return dara.Validate(s)
}
