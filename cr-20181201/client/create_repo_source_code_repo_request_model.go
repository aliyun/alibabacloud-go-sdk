// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSourceCodeRepoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v bool) *CreateRepoSourceCodeRepoRequest
	GetAutoBuild() *bool
	SetCodeRepoName(v string) *CreateRepoSourceCodeRepoRequest
	GetCodeRepoName() *string
	SetCodeRepoNamespaceName(v string) *CreateRepoSourceCodeRepoRequest
	GetCodeRepoNamespaceName() *string
	SetCodeRepoType(v string) *CreateRepoSourceCodeRepoRequest
	GetCodeRepoType() *string
	SetDisableCacheBuild(v bool) *CreateRepoSourceCodeRepoRequest
	GetDisableCacheBuild() *bool
	SetInstanceId(v string) *CreateRepoSourceCodeRepoRequest
	GetInstanceId() *string
	SetOverseaBuild(v bool) *CreateRepoSourceCodeRepoRequest
	GetOverseaBuild() *bool
	SetRepoId(v string) *CreateRepoSourceCodeRepoRequest
	GetRepoId() *string
}

type CreateRepoSourceCodeRepoRequest struct {
	// Specifies whether to enable automatic building. Valid values:
	//
	// - `true`: enables automatic building.
	//
	// - `false`: disables automatic building.
	//
	// example:
	//
	// true
	AutoBuild *bool `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The name of the source code repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace name of the source code repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the source code platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	//
	// This parameter is required.
	//
	// example:
	//
	// GITHUB
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Specifies whether to disable the build cache. Valid values:
	//
	// - `true`: disables the build cache.
	//
	// - `false`: enables the build cache.
	//
	// example:
	//
	// false
	DisableCacheBuild *bool `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable intelligent build acceleration for overseas sources. Valid values:
	//
	// - `true`: enables intelligent build acceleration for overseas sources.
	//
	// - `false`: disables intelligent build acceleration for overseas sources.
	//
	// example:
	//
	// false
	OverseaBuild *bool `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The image repository ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateRepoSourceCodeRepoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSourceCodeRepoRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSourceCodeRepoRequest) GetAutoBuild() *bool {
	return s.AutoBuild
}

func (s *CreateRepoSourceCodeRepoRequest) GetCodeRepoName() *string {
	return s.CodeRepoName
}

func (s *CreateRepoSourceCodeRepoRequest) GetCodeRepoNamespaceName() *string {
	return s.CodeRepoNamespaceName
}

func (s *CreateRepoSourceCodeRepoRequest) GetCodeRepoType() *string {
	return s.CodeRepoType
}

func (s *CreateRepoSourceCodeRepoRequest) GetDisableCacheBuild() *bool {
	return s.DisableCacheBuild
}

func (s *CreateRepoSourceCodeRepoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoSourceCodeRepoRequest) GetOverseaBuild() *bool {
	return s.OverseaBuild
}

func (s *CreateRepoSourceCodeRepoRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepoSourceCodeRepoRequest) SetAutoBuild(v bool) *CreateRepoSourceCodeRepoRequest {
	s.AutoBuild = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetCodeRepoName(v string) *CreateRepoSourceCodeRepoRequest {
	s.CodeRepoName = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetCodeRepoNamespaceName(v string) *CreateRepoSourceCodeRepoRequest {
	s.CodeRepoNamespaceName = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetCodeRepoType(v string) *CreateRepoSourceCodeRepoRequest {
	s.CodeRepoType = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetDisableCacheBuild(v bool) *CreateRepoSourceCodeRepoRequest {
	s.DisableCacheBuild = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetInstanceId(v string) *CreateRepoSourceCodeRepoRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetOverseaBuild(v bool) *CreateRepoSourceCodeRepoRequest {
	s.OverseaBuild = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetRepoId(v string) *CreateRepoSourceCodeRepoRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) Validate() error {
	return dara.Validate(s)
}
