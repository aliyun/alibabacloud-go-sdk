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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The content of the file. The file content cannot exceed 32 KB after Base64 encoding.
	//
	// - If `ContentType` is set to `PlainText`, this parameter specifies the plain text content.
	//
	// - If `ContentType` is set to `Base64`, this parameter specifies the Base64-encoded content.
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
	// - Base64: Base64-encoded.
	//
	// Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the file. The full character set is supported. The description cannot exceed 512 characters in length.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group of the file. This parameter takes effect only on Linux instances. Default value: root. The value cannot exceed 64 characters in length.
	//
	// > If you specify a different user group, make sure that the user group exists on the instance.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file. This parameter takes effect only on Linux instances. You can configure this parameter in the same way as you run the chmod command.
	//
	// Default value: 0644, which indicates that the owner has read and write permissions, and the group and other users have read-only permissions.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. This parameter takes effect only on Linux instances. Default value: root. The value cannot exceed 64 characters in length.
	//
	// > If you specify a different user, make sure that the user exists on the instance.
	//
	// example:
	//
	// test
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The IDs of the ECS instances to which you want to send the file. You can specify up to 50 instance IDs. Valid values of N: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp185dy2o3o6n****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The name of the file. The full character set is supported. The name cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to overwrite a file with the same name in the destination directory. Valid values:
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
	// The region ID of the target ECS instances. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group for file sending. If you specify this parameter:
	//
	// - The ECS instances specified by InstanceId must belong to this resource group.
	//
	// - You can filter file sending results by specifying this parameter when you call [DescribeSendFileResults](https://help.aliyun.com/document_detail/184117.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*SendFileRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The destination directory on the target ECS instances where the file is sent. If the directory does not exist, it is automatically created. The directory path cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout period for sending the file. Unit: seconds.
	//
	// - A timeout may occur when the file cannot be sent due to a process issue, a missing module, or a missing Cloud Assistant Agent.
	//
	// - If the specified timeout period is less than 10 seconds, the system automatically sets the timeout period to 10 seconds to ensure successful delivery.
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
	// The key of the tag for file sending. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count with this tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count with all the specified tags attached cannot exceed 1,000. If the resource count exceeds 1,000, call [ListTagResources](https://help.aliyun.com/document_detail/110425.html) to query the resources.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag for file sending. Valid values of N: 1 to 20. The tag value can be an empty string.
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
