// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerIdeInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateServerIdeInstanceShrinkRequest
	GetCredentialConfigShrink() *string
	SetCu(v int32) *CreateServerIdeInstanceShrinkRequest
	GetCu() *int32
	SetDatasetsShrink(v string) *CreateServerIdeInstanceShrinkRequest
	GetDatasetsShrink() *string
	SetImageId(v string) *CreateServerIdeInstanceShrinkRequest
	GetImageId() *string
	SetImageUrl(v string) *CreateServerIdeInstanceShrinkRequest
	GetImageUrl() *string
	SetInstanceName(v string) *CreateServerIdeInstanceShrinkRequest
	GetInstanceName() *string
	SetOwner(v string) *CreateServerIdeInstanceShrinkRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateServerIdeInstanceShrinkRequest
	GetProjectId() *int64
	SetResourceGroupId(v string) *CreateServerIdeInstanceShrinkRequest
	GetResourceGroupId() *string
	SetUserCommandShrink(v string) *CreateServerIdeInstanceShrinkRequest
	GetUserCommandShrink() *string
	SetUserVpcShrink(v string) *CreateServerIdeInstanceShrinkRequest
	GetUserVpcShrink() *string
}

type CreateServerIdeInstanceShrinkRequest struct {
	// The credential injection configuration for the instance. After this feature is enabled, you can use the default RAM role chain or specify a custom RAM role.
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The number of CUs used by the instance.
	//
	// example:
	//
	// 10
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The list of datasets mounted to the instance.
	DatasetsShrink *string `json:"Datasets,omitempty" xml:"Datasets,omitempty"`
	// The image ID. You can call ListServerIdeImages to obtain the image ID.
	//
	// example:
	//
	// System_serveride_notebook_20240822
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image URL. This parameter is required when you use a non-official DataWorks image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/serveride:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The name of the personal development environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// notebook_dev
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The Alibaba Cloud account ID of the user who owns the instance. If this parameter is not specified, the current caller is used by default.
	//
	// example:
	//
	// 20933221576142****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The DataWorks resource group identifier. You can specify the numeric ID of the resource group or the full identifier in the Serverless_res_group_{tenantId}_{resgId} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_123456789012345_9876543210****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The user command configuration to be executed when the instance starts.
	UserCommandShrink *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// The Virtual Private Cloud (VPC) configuration used by the instance.
	UserVpcShrink *string `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s CreateServerIdeInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateServerIdeInstanceShrinkRequest) GetCu() *int32 {
	return s.Cu
}

func (s *CreateServerIdeInstanceShrinkRequest) GetDatasetsShrink() *string {
	return s.DatasetsShrink
}

func (s *CreateServerIdeInstanceShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateServerIdeInstanceShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateServerIdeInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateServerIdeInstanceShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateServerIdeInstanceShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateServerIdeInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServerIdeInstanceShrinkRequest) GetUserCommandShrink() *string {
	return s.UserCommandShrink
}

func (s *CreateServerIdeInstanceShrinkRequest) GetUserVpcShrink() *string {
	return s.UserVpcShrink
}

func (s *CreateServerIdeInstanceShrinkRequest) SetCredentialConfigShrink(v string) *CreateServerIdeInstanceShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetCu(v int32) *CreateServerIdeInstanceShrinkRequest {
	s.Cu = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetDatasetsShrink(v string) *CreateServerIdeInstanceShrinkRequest {
	s.DatasetsShrink = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetImageId(v string) *CreateServerIdeInstanceShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetImageUrl(v string) *CreateServerIdeInstanceShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetInstanceName(v string) *CreateServerIdeInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetOwner(v string) *CreateServerIdeInstanceShrinkRequest {
	s.Owner = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetProjectId(v int64) *CreateServerIdeInstanceShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServerIdeInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetUserCommandShrink(v string) *CreateServerIdeInstanceShrinkRequest {
	s.UserCommandShrink = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) SetUserVpcShrink(v string) *CreateServerIdeInstanceShrinkRequest {
	s.UserVpcShrink = &v
	return s
}

func (s *CreateServerIdeInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
