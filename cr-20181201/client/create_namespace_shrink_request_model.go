// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *CreateNamespaceShrinkRequest
	GetAutoCreateRepo() *bool
	SetDefaultRepoConfigurationShrink(v string) *CreateNamespaceShrinkRequest
	GetDefaultRepoConfigurationShrink() *string
	SetDefaultRepoType(v string) *CreateNamespaceShrinkRequest
	GetDefaultRepoType() *string
	SetInstanceId(v string) *CreateNamespaceShrinkRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateNamespaceShrinkRequest
	GetNamespaceName() *string
}

type CreateNamespaceShrinkRequest struct {
	// Specifies whether to automatically create an image repository in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo                 *bool   `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfigurationShrink *string `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of the repositories that are automatically created in the namespace. Valid values:
	//
	// 	- `PUBLIC`: public repositories
	//
	// 	- `PRIVATE`: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace. The name must be 2 to 120 characters in length, and can contain lowercase letters, digits, and the following delimiters: underscores (_), hyphens (-), and periods (.). The name cannot start or end with a delimiter.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateNamespaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceShrinkRequest) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *CreateNamespaceShrinkRequest) GetDefaultRepoConfigurationShrink() *string {
	return s.DefaultRepoConfigurationShrink
}

func (s *CreateNamespaceShrinkRequest) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *CreateNamespaceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNamespaceShrinkRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateNamespaceShrinkRequest) SetAutoCreateRepo(v bool) *CreateNamespaceShrinkRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetDefaultRepoConfigurationShrink(v string) *CreateNamespaceShrinkRequest {
	s.DefaultRepoConfigurationShrink = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetDefaultRepoType(v string) *CreateNamespaceShrinkRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetInstanceId(v string) *CreateNamespaceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetNamespaceName(v string) *CreateNamespaceShrinkRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
