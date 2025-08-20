// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChartNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *UpdateChartNamespaceRequest
	GetAutoCreateRepo() *bool
	SetDefaultRepoType(v string) *UpdateChartNamespaceRequest
	GetDefaultRepoType() *string
	SetInstanceId(v string) *UpdateChartNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *UpdateChartNamespaceRequest
	GetNamespaceName() *string
}

type UpdateChartNamespaceRequest struct {
	// Specifies whether to automatically create repositories in the namespace. Valid values:
	//
	// 	- `true`: automatically creates repositories in the namespace.
	//
	// 	- `false`: does not automatically create repositories in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default type of the repository. Valid values:
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
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateChartNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateChartNamespaceRequest) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *UpdateChartNamespaceRequest) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *UpdateChartNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateChartNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateChartNamespaceRequest) SetAutoCreateRepo(v bool) *UpdateChartNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *UpdateChartNamespaceRequest) SetDefaultRepoType(v string) *UpdateChartNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *UpdateChartNamespaceRequest) SetInstanceId(v string) *UpdateChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChartNamespaceRequest) SetNamespaceName(v string) *UpdateChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateChartNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
