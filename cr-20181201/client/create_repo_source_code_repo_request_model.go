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
	// Specifies whether to trigger image building when source code is committed. Valid values:
	//
	// 	- `true`: triggers image building when source code is committed.
	//
	// 	- `false`: does not trigger image building when source code is committed.
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
	// The namespace to which the source code repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the source code hosting platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	//
	// This parameter is required.
	//
	// example:
	//
	// GITHUB
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Specifies whether to disable building caches. Valid values:
	//
	// 	- `true`: disables building caches.
	//
	// 	- `false`: enables building caches.
	//
	// example:
	//
	// false
	DisableCacheBuild *bool `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable Build With Servers Deployed Outside Chinese Mainland. Valid values:
	//
	// 	- `true`: enables Build With Servers Deployed Outside Chinese Mainland.
	//
	// 	- `false`: does not enable Build With Servers Deployed Outside Chinese Mainland.
	//
	// example:
	//
	// false
	OverseaBuild *bool `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the image repository.
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
