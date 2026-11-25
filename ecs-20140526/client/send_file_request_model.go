// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SendFileRequest
	GetClientToken() *string
	SetContent(v string) *SendFileRequest
	GetContent() *string
	SetContentType(v string) *SendFileRequest
	GetContentType() *string
	SetDescription(v string) *SendFileRequest
	GetDescription() *string
	SetFileGroup(v string) *SendFileRequest
	GetFileGroup() *string
	SetFileMode(v string) *SendFileRequest
	GetFileMode() *string
	SetFileOwner(v string) *SendFileRequest
	GetFileOwner() *string
	SetInstanceId(v []*string) *SendFileRequest
	GetInstanceId() []*string
	SetName(v string) *SendFileRequest
	GetName() *string
	SetOverwrite(v bool) *SendFileRequest
	GetOverwrite() *bool
	SetOwnerAccount(v string) *SendFileRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SendFileRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SendFileRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SendFileRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *SendFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendFileRequest
	GetResourceOwnerId() *int64
	SetTag(v []*SendFileRequestTag) *SendFileRequest
	GetTag() []*SendFileRequestTag
	SetTargetDir(v string) *SendFileRequest
	GetTargetDir() *string
	SetTimeout(v int64) *SendFileRequest
	GetTimeout() *int64
}

type SendFileRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The content of the file. The file size cannot exceed 32 KB after Base64 encoding.
	//
	// - When the `ContentType` parameter is set to `PlainText`, this field is plain text in plaintext format.
	//
	// - When the `ContentType` parameter is set to `Base64`, this field is Base64-encoded text.
	//
	// This parameter is required.
	//
	// example:
	//
	// #!/bin/bash  echo "Current User is :"  echo $(ps | grep "$$" | awk \\"{print $2}\\")  --------  oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file. Valid values:
	//
	// - PlainText: plain text.
	//
	// - Base64: Base64-encoded text.
	//
	// Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the file. All character sets are supported. The description cannot exceed 512 characters in length.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file. This parameter takes effect only on Linux instances. Default value: root. The value cannot exceed 64 characters in length.
	//
	// > When using another user group, make sure that the user group exists on the instance.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions of the file. This parameter takes effect only on Linux instances. The setting method is the same as that of the chmod command.
	//
	// Default value: 0644, which indicates that the owner has read and write permissions, and the user group and other users have read-only permissions.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. This parameter takes effect only on Linux instances. Default value: root. The value cannot exceed 64 characters in length.
	//
	// > When using another user, make sure that the user exists on the instance.
	//
	// example:
	//
	// test
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The list of ECS instances to which the file is sent. You can specify up to 50 ECS instance IDs. Valid values of N: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp185dy2o3o6n****"]
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The name of the file. All character sets are supported. The name cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to overwrite the file if a file with the same name already exists in the target directory. Valid values:
	//
	// - true: Overwrite the file.
	//
	// - false: Do not overwrite the file.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Overwrite    *bool   `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the target ECS instances. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the file sending task belongs. When this parameter is specified:
	//
	// - The ECS instances specified by InstanceId must belong to this resource group.
	//
	// - You can filter the corresponding file sending results by specifying this parameter (by calling [DescribeSendFileResults](https://help.aliyun.com/document_detail/184117.html)).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tag []*SendFileRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The folder on the target ECS instances where the file is sent. If the folder does not exist, automatic creation is performed. The folder path cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout period for sending the file. Unit: seconds.
	//
	// - A timeout occurs when the file cannot be sent due to process issues, missing modules, or a missing Cloud Assistant Agent.
	//
	// - If the specified timeout period is less than 10 seconds, the system automatically sets it to 10 seconds to ensure successful sending.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SendFileRequest) GetContent() *string {
	return s.Content
}

func (s *SendFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendFileRequest) GetDescription() *string {
	return s.Description
}

func (s *SendFileRequest) GetFileGroup() *string {
	return s.FileGroup
}

func (s *SendFileRequest) GetFileMode() *string {
	return s.FileMode
}

func (s *SendFileRequest) GetFileOwner() *string {
	return s.FileOwner
}

func (s *SendFileRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *SendFileRequest) GetName() *string {
	return s.Name
}

func (s *SendFileRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *SendFileRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SendFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SendFileRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SendFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendFileRequest) GetTag() []*SendFileRequestTag {
	return s.Tag
}

func (s *SendFileRequest) GetTargetDir() *string {
	return s.TargetDir
}

func (s *SendFileRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *SendFileRequest) SetClientToken(v string) *SendFileRequest {
	s.ClientToken = &v
	return s
}

func (s *SendFileRequest) SetContent(v string) *SendFileRequest {
	s.Content = &v
	return s
}

func (s *SendFileRequest) SetContentType(v string) *SendFileRequest {
	s.ContentType = &v
	return s
}

func (s *SendFileRequest) SetDescription(v string) *SendFileRequest {
	s.Description = &v
	return s
}

func (s *SendFileRequest) SetFileGroup(v string) *SendFileRequest {
	s.FileGroup = &v
	return s
}

func (s *SendFileRequest) SetFileMode(v string) *SendFileRequest {
	s.FileMode = &v
	return s
}

func (s *SendFileRequest) SetFileOwner(v string) *SendFileRequest {
	s.FileOwner = &v
	return s
}

func (s *SendFileRequest) SetInstanceId(v []*string) *SendFileRequest {
	s.InstanceId = v
	return s
}

func (s *SendFileRequest) SetName(v string) *SendFileRequest {
	s.Name = &v
	return s
}

func (s *SendFileRequest) SetOverwrite(v bool) *SendFileRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileRequest) SetOwnerAccount(v string) *SendFileRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SendFileRequest) SetOwnerId(v int64) *SendFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SendFileRequest) SetRegionId(v string) *SendFileRequest {
	s.RegionId = &v
	return s
}

func (s *SendFileRequest) SetResourceGroupId(v string) *SendFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SendFileRequest) SetResourceOwnerAccount(v string) *SendFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendFileRequest) SetResourceOwnerId(v int64) *SendFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendFileRequest) SetTag(v []*SendFileRequestTag) *SendFileRequest {
	s.Tag = v
	return s
}

func (s *SendFileRequest) SetTargetDir(v string) *SendFileRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileRequest) SetTimeout(v int64) *SendFileRequest {
	s.Timeout = &v
	return s
}

func (s *SendFileRequest) Validate() error {
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

type SendFileRequestTag struct {
	// The tag key of the file sending task. Valid values of N: 1 to 20. The tag key cannot be an empty string once specified.
	//
	// If you use a single tag to filter resources, the returned resource count cannot exceed 1,000. If you use multiple tags to filter resources, the returned resource count for resources attached to all specified tags cannot exceed 1,000. If the resource count exceeds 1,000, you must call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the file sending task. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SendFileRequestTag) String() string {
	return dara.Prettify(s)
}

func (s SendFileRequestTag) GoString() string {
	return s.String()
}

func (s *SendFileRequestTag) GetKey() *string {
	return s.Key
}

func (s *SendFileRequestTag) GetValue() *string {
	return s.Value
}

func (s *SendFileRequestTag) SetKey(v string) *SendFileRequestTag {
	s.Key = &v
	return s
}

func (s *SendFileRequestTag) SetValue(v string) *SendFileRequestTag {
	s.Value = &v
	return s
}

func (s *SendFileRequestTag) Validate() error {
	return dara.Validate(s)
}
