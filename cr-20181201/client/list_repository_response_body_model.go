// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepositoryResponseBody
	GetIsSuccess() *bool
	SetMaxResults(v int32) *ListRepositoryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRepositoryResponseBody
	GetNextToken() *string
	SetPageNo(v int32) *ListRepositoryResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepositoryResponseBody
	GetPageSize() *int32
	SetRepositories(v []*ListRepositoryResponseBodyRepositories) *ListRepositoryResponseBody
	GetRepositories() []*ListRepositoryResponseBodyRepositories
	SetRequestId(v string) *ListRepositoryResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRepositoryResponseBody
	GetTotalCount() *string
}

type ListRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess  *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of repositories.
	Repositories []*ListRepositoryResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5241C090-DA69-4B0F-8E3F-2F24FDE1110E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepositoryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRepositoryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRepositoryResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepositoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepositoryResponseBody) GetRepositories() []*ListRepositoryResponseBodyRepositories {
	return s.Repositories
}

func (s *ListRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRepositoryResponseBody) SetCode(v string) *ListRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepositoryResponseBody) SetIsSuccess(v bool) *ListRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepositoryResponseBody) SetMaxResults(v int32) *ListRepositoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRepositoryResponseBody) SetNextToken(v string) *ListRepositoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRepositoryResponseBody) SetPageNo(v int32) *ListRepositoryResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepositoryResponseBody) SetPageSize(v int32) *ListRepositoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryResponseBody) SetRepositories(v []*ListRepositoryResponseBodyRepositories) *ListRepositoryResponseBody {
	s.Repositories = v
	return s
}

func (s *ListRepositoryResponseBody) SetRequestId(v string) *ListRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryResponseBody) SetTotalCount(v string) *ListRepositoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepositoryResponseBody) Validate() error {
	if s.Repositories != nil {
		for _, item := range s.Repositories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepositoryResponseBodyRepositories struct {
	// The creation time.
	//
	// example:
	//
	// 1564153576000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-kmsiwlxxdcv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1564153576000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The repository build type. Valid values:
	//
	// - `AUTO`: Automatically triggered build.
	//
	// - `MANUAL`: Manually triggered build.
	//
	// example:
	//
	// MANUAL
	RepoBuildType *string `json:"RepoBuildType,omitempty" xml:"RepoBuildType,omitempty"`
	// The repository ID.
	//
	// example:
	//
	// crr-03cuozrsqhkw****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The repository namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The repository status.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The repository type. Valid values:
	//
	// - `PUBLIC`: Public.
	//
	// - `PRIVATE`: Private.
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm4n5kzyfxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary information.
	//
	// example:
	//
	// test OK
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The image tag immutability.
	//
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s ListRepositoryResponseBodyRepositories) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListRepositoryResponseBodyRepositories) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListRepositoryResponseBodyRepositories) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepositoryResponseBodyRepositories) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListRepositoryResponseBodyRepositories) GetRepoBuildType() *string {
	return s.RepoBuildType
}

func (s *ListRepositoryResponseBodyRepositories) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepositoryResponseBodyRepositories) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepositoryResponseBodyRepositories) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListRepositoryResponseBodyRepositories) GetRepoStatus() *string {
	return s.RepoStatus
}

func (s *ListRepositoryResponseBodyRepositories) GetRepoType() *string {
	return s.RepoType
}

func (s *ListRepositoryResponseBodyRepositories) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRepositoryResponseBodyRepositories) GetSummary() *string {
	return s.Summary
}

func (s *ListRepositoryResponseBodyRepositories) GetTagImmutability() *bool {
	return s.TagImmutability
}

func (s *ListRepositoryResponseBodyRepositories) SetCreateTime(v int64) *ListRepositoryResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetInstanceId(v string) *ListRepositoryResponseBodyRepositories {
	s.InstanceId = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetModifiedTime(v int64) *ListRepositoryResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoBuildType(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoBuildType = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoId(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoName(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoName = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoNamespaceName(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoStatus(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoStatus = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoType(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoType = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetResourceGroupId(v string) *ListRepositoryResponseBodyRepositories {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetSummary(v string) *ListRepositoryResponseBodyRepositories {
	s.Summary = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetTagImmutability(v bool) *ListRepositoryResponseBodyRepositories {
	s.TagImmutability = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) Validate() error {
	return dara.Validate(s)
}
