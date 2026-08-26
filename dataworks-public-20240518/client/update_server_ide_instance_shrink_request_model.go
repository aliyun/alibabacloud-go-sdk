// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerIdeInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *UpdateServerIdeInstanceShrinkRequest
	GetCredentialConfigShrink() *string
	SetCu(v int32) *UpdateServerIdeInstanceShrinkRequest
	GetCu() *int32
	SetDatasetsShrink(v string) *UpdateServerIdeInstanceShrinkRequest
	GetDatasetsShrink() *string
	SetImageId(v string) *UpdateServerIdeInstanceShrinkRequest
	GetImageId() *string
	SetImageUrl(v string) *UpdateServerIdeInstanceShrinkRequest
	GetImageUrl() *string
	SetInstanceId(v string) *UpdateServerIdeInstanceShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateServerIdeInstanceShrinkRequest
	GetInstanceName() *string
	SetProjectId(v int64) *UpdateServerIdeInstanceShrinkRequest
	GetProjectId() *int64
	SetUserVpcShrink(v string) *UpdateServerIdeInstanceShrinkRequest
	GetUserVpcShrink() *string
}

type UpdateServerIdeInstanceShrinkRequest struct {
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
	// The image ID. You can call ListServerIdeImages to obtain the ID.
	//
	// example:
	//
	// System_serveride_notebook_20240822
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image URL. This parameter is required when you use a non-DataWorks official image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/serveride:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The personal development environment instance ID. You can call ListServerIdeInstances to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the personal development environment instance.
	//
	// example:
	//
	// notebook_dev
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The VPC configuration used by the instance.
	UserVpcShrink *string `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s UpdateServerIdeInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetCu() *int32 {
	return s.Cu
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetDatasetsShrink() *string {
	return s.DatasetsShrink
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateServerIdeInstanceShrinkRequest) GetUserVpcShrink() *string {
	return s.UserVpcShrink
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetCredentialConfigShrink(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetCu(v int32) *UpdateServerIdeInstanceShrinkRequest {
	s.Cu = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetDatasetsShrink(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.DatasetsShrink = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetImageId(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetImageUrl(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetInstanceId(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetInstanceName(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetProjectId(v int64) *UpdateServerIdeInstanceShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) SetUserVpcShrink(v string) *UpdateServerIdeInstanceShrinkRequest {
	s.UserVpcShrink = &v
	return s
}

func (s *UpdateServerIdeInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
