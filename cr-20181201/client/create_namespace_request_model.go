// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *CreateNamespaceRequest
	GetAutoCreateRepo() *bool
	SetDefaultRepoConfiguration(v *RepoConfiguration) *CreateNamespaceRequest
	GetDefaultRepoConfiguration() *RepoConfiguration
	SetDefaultRepoType(v string) *CreateNamespaceRequest
	GetDefaultRepoType() *string
	SetInstanceId(v string) *CreateNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateNamespaceRequest
	GetNamespaceName() *string
}

type CreateNamespaceRequest struct {
	// Specifies whether to automatically create an image repository in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo           *bool              `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
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

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *CreateNamespaceRequest) GetDefaultRepoConfiguration() *RepoConfiguration {
	return s.DefaultRepoConfiguration
}

func (s *CreateNamespaceRequest) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *CreateNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateNamespaceRequest) SetAutoCreateRepo(v bool) *CreateNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *CreateNamespaceRequest) SetDefaultRepoConfiguration(v *RepoConfiguration) *CreateNamespaceRequest {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *CreateNamespaceRequest) SetDefaultRepoType(v string) *CreateNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *CreateNamespaceRequest) SetInstanceId(v string) *CreateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceName(v string) *CreateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
