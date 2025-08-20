// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChartNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *CreateChartNamespaceRequest
	GetAutoCreateRepo() *bool
	SetDefaultRepoType(v string) *CreateChartNamespaceRequest
	GetDefaultRepoType() *string
	SetInstanceId(v string) *CreateChartNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateChartNamespaceRequest
	GetNamespaceName() *string
}

type CreateChartNamespaceRequest struct {
	// Specifies whether to automatically create repositories in the namespace. Valid values:
	//
	// \\-`  true `: automatically creates repositories in the namespace.
	//
	// \\-`  false `: does not automatically create repositories in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
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
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateChartNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateChartNamespaceRequest) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *CreateChartNamespaceRequest) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *CreateChartNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateChartNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateChartNamespaceRequest) SetAutoCreateRepo(v bool) *CreateChartNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *CreateChartNamespaceRequest) SetDefaultRepoType(v string) *CreateChartNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *CreateChartNamespaceRequest) SetInstanceId(v string) *CreateChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChartNamespaceRequest) SetNamespaceName(v string) *CreateChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateChartNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
