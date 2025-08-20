// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoSourceCodeRepoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v string) *GetRepoSourceCodeRepoResponseBody
	GetAutoBuild() *string
	SetCode(v string) *GetRepoSourceCodeRepoResponseBody
	GetCode() *string
	SetCodeRepoDomain(v string) *GetRepoSourceCodeRepoResponseBody
	GetCodeRepoDomain() *string
	SetCodeRepoName(v string) *GetRepoSourceCodeRepoResponseBody
	GetCodeRepoName() *string
	SetCodeRepoNamespaceName(v string) *GetRepoSourceCodeRepoResponseBody
	GetCodeRepoNamespaceName() *string
	SetCodeRepoType(v string) *GetRepoSourceCodeRepoResponseBody
	GetCodeRepoType() *string
	SetDisableCacheBuild(v string) *GetRepoSourceCodeRepoResponseBody
	GetDisableCacheBuild() *string
	SetIsSuccess(v bool) *GetRepoSourceCodeRepoResponseBody
	GetIsSuccess() *bool
	SetOverseaBuild(v string) *GetRepoSourceCodeRepoResponseBody
	GetOverseaBuild() *string
	SetRepoId(v string) *GetRepoSourceCodeRepoResponseBody
	GetRepoId() *string
	SetRequestId(v string) *GetRepoSourceCodeRepoResponseBody
	GetRequestId() *string
}

type GetRepoSourceCodeRepoResponseBody struct {
	// Indicates whether image building is automatically triggered when source code is committed. Valid values:
	//
	// 	- `true`: Image building is automatically triggered when source code is committed.
	//
	// 	- `false`: Image building is not triggered when source code is committed.
	//
	// example:
	//
	// true
	AutoBuild *string `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The address of the source code repository.
	//
	// example:
	//
	// https://github.com
	CodeRepoDomain *string `json:"CodeRepoDomain,omitempty" xml:"CodeRepoDomain,omitempty"`
	// The name of the source code repository.
	//
	// example:
	//
	// repo
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	//
	// example:
	//
	// namespace
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the code hosting platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	//
	// example:
	//
	// GITHUB
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Indicates whether build cache is disabled. Valid values:
	//
	// 	- `true`: Build cache is disabled.
	//
	// 	- `false`: Build cache is enabled.
	//
	// example:
	//
	// false
	DisableCacheBuild *string `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Indicates whether image building is accelerated for servers outside the Chinese mainland. Valid values:
	//
	// 	- `true`: Image building is accelerated for servers outside the Chinese mainland.
	//
	// 	- `false`: Image building is not accelerated for servers outside the Chinese mainland.
	//
	// example:
	//
	// false
	OverseaBuild *string `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRepoSourceCodeRepoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSourceCodeRepoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoSourceCodeRepoResponseBody) GetAutoBuild() *string {
	return s.AutoBuild
}

func (s *GetRepoSourceCodeRepoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoSourceCodeRepoResponseBody) GetCodeRepoDomain() *string {
	return s.CodeRepoDomain
}

func (s *GetRepoSourceCodeRepoResponseBody) GetCodeRepoName() *string {
	return s.CodeRepoName
}

func (s *GetRepoSourceCodeRepoResponseBody) GetCodeRepoNamespaceName() *string {
	return s.CodeRepoNamespaceName
}

func (s *GetRepoSourceCodeRepoResponseBody) GetCodeRepoType() *string {
	return s.CodeRepoType
}

func (s *GetRepoSourceCodeRepoResponseBody) GetDisableCacheBuild() *string {
	return s.DisableCacheBuild
}

func (s *GetRepoSourceCodeRepoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoSourceCodeRepoResponseBody) GetOverseaBuild() *string {
	return s.OverseaBuild
}

func (s *GetRepoSourceCodeRepoResponseBody) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepoSourceCodeRepoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoSourceCodeRepoResponseBody) SetAutoBuild(v string) *GetRepoSourceCodeRepoResponseBody {
	s.AutoBuild = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCode(v string) *GetRepoSourceCodeRepoResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoDomain(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoDomain = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoName(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoName = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoNamespaceName(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoNamespaceName = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoType(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoType = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetDisableCacheBuild(v string) *GetRepoSourceCodeRepoResponseBody {
	s.DisableCacheBuild = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetIsSuccess(v bool) *GetRepoSourceCodeRepoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetOverseaBuild(v string) *GetRepoSourceCodeRepoResponseBody {
	s.OverseaBuild = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetRepoId(v string) *GetRepoSourceCodeRepoResponseBody {
	s.RepoId = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetRequestId(v string) *GetRepoSourceCodeRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) Validate() error {
	return dara.Validate(s)
}
