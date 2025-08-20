// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoSourceCodeRepoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v string) *UpdateRepoSourceCodeRepoRequest
	GetAutoBuild() *string
	SetCodeRepoId(v string) *UpdateRepoSourceCodeRepoRequest
	GetCodeRepoId() *string
	SetCodeRepoName(v string) *UpdateRepoSourceCodeRepoRequest
	GetCodeRepoName() *string
	SetCodeRepoNamespaceName(v string) *UpdateRepoSourceCodeRepoRequest
	GetCodeRepoNamespaceName() *string
	SetCodeRepoType(v string) *UpdateRepoSourceCodeRepoRequest
	GetCodeRepoType() *string
	SetDisableCacheBuild(v string) *UpdateRepoSourceCodeRepoRequest
	GetDisableCacheBuild() *string
	SetInstanceId(v string) *UpdateRepoSourceCodeRepoRequest
	GetInstanceId() *string
	SetOverseaBuild(v string) *UpdateRepoSourceCodeRepoRequest
	GetOverseaBuild() *string
	SetRepoId(v string) *UpdateRepoSourceCodeRepoRequest
	GetRepoId() *string
}

type UpdateRepoSourceCodeRepoRequest struct {
	// Specifies whether to enable automatic image building when code is committed. Valid values:
	//
	// 	- `true`: enables automatic image building when code is committed.
	//
	// 	- `false`: disables automatic image building when code is committed.
	//
	// example:
	//
	// true
	AutoBuild *string `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The ID of the source code repository.
	//
	// example:
	//
	// crr-cp7d6sget5r****
	CodeRepoId *string `json:"CodeRepoId,omitempty" xml:"CodeRepoId,omitempty"`
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
	// The type of the source code hosting platform. Valid values: GITHUB, GITLAB, GITEE, CODE, and CODEUP.
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
	DisableCacheBuild *string `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
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
	// 	- `false`: disables Build With Servers Deployed Outside Chinese Mainland.
	//
	// example:
	//
	// false
	OverseaBuild *string `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-gzsrlevmvoa****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s UpdateRepoSourceCodeRepoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoSourceCodeRepoRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepoSourceCodeRepoRequest) GetAutoBuild() *string {
	return s.AutoBuild
}

func (s *UpdateRepoSourceCodeRepoRequest) GetCodeRepoId() *string {
	return s.CodeRepoId
}

func (s *UpdateRepoSourceCodeRepoRequest) GetCodeRepoName() *string {
	return s.CodeRepoName
}

func (s *UpdateRepoSourceCodeRepoRequest) GetCodeRepoNamespaceName() *string {
	return s.CodeRepoNamespaceName
}

func (s *UpdateRepoSourceCodeRepoRequest) GetCodeRepoType() *string {
	return s.CodeRepoType
}

func (s *UpdateRepoSourceCodeRepoRequest) GetDisableCacheBuild() *string {
	return s.DisableCacheBuild
}

func (s *UpdateRepoSourceCodeRepoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRepoSourceCodeRepoRequest) GetOverseaBuild() *string {
	return s.OverseaBuild
}

func (s *UpdateRepoSourceCodeRepoRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *UpdateRepoSourceCodeRepoRequest) SetAutoBuild(v string) *UpdateRepoSourceCodeRepoRequest {
	s.AutoBuild = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoId(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoId = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoName(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoName = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoNamespaceName(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoNamespaceName = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoType(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoType = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetDisableCacheBuild(v string) *UpdateRepoSourceCodeRepoRequest {
	s.DisableCacheBuild = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetInstanceId(v string) *UpdateRepoSourceCodeRepoRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetOverseaBuild(v string) *UpdateRepoSourceCodeRepoRequest {
	s.OverseaBuild = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetRepoId(v string) *UpdateRepoSourceCodeRepoRequest {
	s.RepoId = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) Validate() error {
	return dara.Validate(s)
}
