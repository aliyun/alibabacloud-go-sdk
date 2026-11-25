// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *UpdateNamespaceShrinkRequest
	GetAutoCreateRepo() *bool
	SetDefaultRepoConfigurationShrink(v string) *UpdateNamespaceShrinkRequest
	GetDefaultRepoConfigurationShrink() *string
	SetDefaultRepoType(v string) *UpdateNamespaceShrinkRequest
	GetDefaultRepoType() *string
	SetInstanceId(v string) *UpdateNamespaceShrinkRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *UpdateNamespaceShrinkRequest
	GetNamespaceName() *string
}

type UpdateNamespaceShrinkRequest struct {
	// Specifies whether to allow automatic creation of a repository upon push.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository configuration.
	DefaultRepoConfigurationShrink *string `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default repository type. Valid values:
	//
	// - `PUBLIC`: public repository
	//
	// - `PRIVATE`: private repository
	//
	// example:
	//
	// PRIVATE
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceShrinkRequest) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *UpdateNamespaceShrinkRequest) GetDefaultRepoConfigurationShrink() *string {
	return s.DefaultRepoConfigurationShrink
}

func (s *UpdateNamespaceShrinkRequest) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *UpdateNamespaceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNamespaceShrinkRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateNamespaceShrinkRequest) SetAutoCreateRepo(v bool) *UpdateNamespaceShrinkRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetDefaultRepoConfigurationShrink(v string) *UpdateNamespaceShrinkRequest {
	s.DefaultRepoConfigurationShrink = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetDefaultRepoType(v string) *UpdateNamespaceShrinkRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetInstanceId(v string) *UpdateNamespaceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetNamespaceName(v string) *UpdateNamespaceShrinkRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
