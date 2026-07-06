// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *UpdateNamespaceRequest
	GetAutoCreateRepo() *bool
	SetDefaultRepoConfiguration(v *RepoConfiguration) *UpdateNamespaceRequest
	GetDefaultRepoConfiguration() *RepoConfiguration
	SetDefaultRepoType(v string) *UpdateNamespaceRequest
	GetDefaultRepoType() *string
	SetInstanceId(v string) *UpdateNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *UpdateNamespaceRequest
	GetNamespaceName() *string
}

type UpdateNamespaceRequest struct {
	// Whether to automatically create a repository when an image is pushed.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default configuration for automatically created repositories.
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type for automatically created repositories. This parameter applies only if `AutoCreateRepo` is set to `true`. Valid values:
	//
	// - `PUBLIC`: a public repository
	//
	// - `PRIVATE`: a private repository
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
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *UpdateNamespaceRequest) GetDefaultRepoConfiguration() *RepoConfiguration {
	return s.DefaultRepoConfiguration
}

func (s *UpdateNamespaceRequest) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *UpdateNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateNamespaceRequest) SetAutoCreateRepo(v bool) *UpdateNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *UpdateNamespaceRequest) SetDefaultRepoConfiguration(v *RepoConfiguration) *UpdateNamespaceRequest {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *UpdateNamespaceRequest) SetDefaultRepoType(v string) *UpdateNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *UpdateNamespaceRequest) SetInstanceId(v string) *UpdateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceName(v string) *UpdateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateNamespaceRequest) Validate() error {
	if s.DefaultRepoConfiguration != nil {
		if err := s.DefaultRepoConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
