// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRepositoryResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetRepositoryResponseBody
	GetCreateTime() *int64
	SetDetail(v string) *GetRepositoryResponseBody
	GetDetail() *string
	SetInstanceId(v string) *GetRepositoryResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetRepositoryResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetRepositoryResponseBody
	GetModifiedTime() *int64
	SetRepoBuildType(v string) *GetRepositoryResponseBody
	GetRepoBuildType() *string
	SetRepoId(v string) *GetRepositoryResponseBody
	GetRepoId() *string
	SetRepoName(v string) *GetRepositoryResponseBody
	GetRepoName() *string
	SetRepoNamespaceName(v string) *GetRepositoryResponseBody
	GetRepoNamespaceName() *string
	SetRepoStatus(v string) *GetRepositoryResponseBody
	GetRepoStatus() *string
	SetRepoType(v string) *GetRepositoryResponseBody
	GetRepoType() *string
	SetRequestId(v string) *GetRepositoryResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetRepositoryResponseBody
	GetResourceGroupId() *string
	SetSummary(v string) *GetRepositoryResponseBody
	GetSummary() *string
	SetTagImmutability(v bool) *GetRepositoryResponseBody
	GetTagImmutability() *bool
}

type GetRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the repository was created.
	//
	// example:
	//
	// 1570759546000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The details of the repository.
	//
	// example:
	//
	// test
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the repository was last modified.
	//
	// example:
	//
	// 1570759546100
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates how the repository was created. Valid values:
	//
	// 	- `MANUAL`: The repository was manually created.
	//
	// 	- `AUTO`: The repository was automatically created.
	//
	// example:
	//
	// MANUAL
	RepoBuildType *string `json:"RepoBuildType,omitempty" xml:"RepoBuildType,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-l5eoubonp0l****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`: public repository.
	//
	// 	- `PRIVATE`: private repository.
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 915E6734-3E50-4640-8DBA-126D2D94DE29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// Automatically created repository
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Indicates whether the feature of image tag immutability is enabled. Valid values:
	//
	// 	- `true`: The feature of image tag immutability is enabled.
	//
	// 	- `false`: The feature of image tag immutability is disabled.
	//
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s GetRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepositoryResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRepositoryResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *GetRepositoryResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepositoryResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetRepositoryResponseBody) GetRepoBuildType() *string {
	return s.RepoBuildType
}

func (s *GetRepositoryResponseBody) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepositoryResponseBody) GetRepoName() *string {
	return s.RepoName
}

func (s *GetRepositoryResponseBody) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetRepositoryResponseBody) GetRepoStatus() *string {
	return s.RepoStatus
}

func (s *GetRepositoryResponseBody) GetRepoType() *string {
	return s.RepoType
}

func (s *GetRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepositoryResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRepositoryResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *GetRepositoryResponseBody) GetTagImmutability() *bool {
	return s.TagImmutability
}

func (s *GetRepositoryResponseBody) SetCode(v string) *GetRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepositoryResponseBody) SetCreateTime(v int64) *GetRepositoryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRepositoryResponseBody) SetDetail(v string) *GetRepositoryResponseBody {
	s.Detail = &v
	return s
}

func (s *GetRepositoryResponseBody) SetInstanceId(v string) *GetRepositoryResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetIsSuccess(v bool) *GetRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepositoryResponseBody) SetModifiedTime(v int64) *GetRepositoryResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoBuildType(v string) *GetRepositoryResponseBody {
	s.RepoBuildType = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoId(v string) *GetRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoName(v string) *GetRepositoryResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoNamespaceName(v string) *GetRepositoryResponseBody {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoStatus(v string) *GetRepositoryResponseBody {
	s.RepoStatus = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoType(v string) *GetRepositoryResponseBody {
	s.RepoType = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRequestId(v string) *GetRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetResourceGroupId(v string) *GetRepositoryResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetSummary(v string) *GetRepositoryResponseBody {
	s.Summary = &v
	return s
}

func (s *GetRepositoryResponseBody) SetTagImmutability(v bool) *GetRepositoryResponseBody {
	s.TagImmutability = &v
	return s
}

func (s *GetRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
