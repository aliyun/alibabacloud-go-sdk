// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChartRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListChartRepositoryResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListChartRepositoryResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChartRepositoryResponseBody
	GetPageSize() *int32
	SetRepositories(v []*ListChartRepositoryResponseBodyRepositories) *ListChartRepositoryResponseBody
	GetRepositories() []*ListChartRepositoryResponseBodyRepositories
	SetRequestId(v string) *ListChartRepositoryResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListChartRepositoryResponseBody
	GetTotalCount() *string
}

type ListChartRepositoryResponseBody struct {
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
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried repositories.
	Repositories []*ListChartRepositoryResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0AB62FB8-6873-4032-8515-4578D27523B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChartRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChartRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListChartRepositoryResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChartRepositoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChartRepositoryResponseBody) GetRepositories() []*ListChartRepositoryResponseBodyRepositories {
	return s.Repositories
}

func (s *ListChartRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChartRepositoryResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListChartRepositoryResponseBody) SetCode(v string) *ListChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetIsSuccess(v bool) *ListChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetPageNo(v int32) *ListChartRepositoryResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetPageSize(v int32) *ListChartRepositoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetRepositories(v []*ListChartRepositoryResponseBodyRepositories) *ListChartRepositoryResponseBody {
	s.Repositories = v
	return s
}

func (s *ListChartRepositoryResponseBody) SetRequestId(v string) *ListChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetTotalCount(v string) *ListChartRepositoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChartRepositoryResponseBody) Validate() error {
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

type ListChartRepositoryResponseBodyRepositories struct {
	// The time when the repository was created.
	//
	// example:
	//
	// 1571926644000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the repository was last modified.
	//
	// example:
	//
	// 1571930329000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crcr-gpsu7b8chmxk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository. Valid values:
	//
	// 	- `NORMAL`: The repository is normal.
	//
	// 	- `DELETING`: The repository is being deleted.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PRIVATE`: a private repository
	//
	// 	- `PUBLIC`: a public repository
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the resource group to which the repository belongs.
	//
	// example:
	//
	// rg-aek2ikd5rxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary about the repository.
	//
	// example:
	//
	// test
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListChartRepositoryResponseBodyRepositories) String() string {
	return dara.Prettify(s)
}

func (s ListChartRepositoryResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryResponseBodyRepositories) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListChartRepositoryResponseBodyRepositories) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChartRepositoryResponseBodyRepositories) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListChartRepositoryResponseBodyRepositories) GetRepoId() *string {
	return s.RepoId
}

func (s *ListChartRepositoryResponseBodyRepositories) GetRepoName() *string {
	return s.RepoName
}

func (s *ListChartRepositoryResponseBodyRepositories) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListChartRepositoryResponseBodyRepositories) GetRepoStatus() *string {
	return s.RepoStatus
}

func (s *ListChartRepositoryResponseBodyRepositories) GetRepoType() *string {
	return s.RepoType
}

func (s *ListChartRepositoryResponseBodyRepositories) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListChartRepositoryResponseBodyRepositories) GetSummary() *string {
	return s.Summary
}

func (s *ListChartRepositoryResponseBodyRepositories) SetCreateTime(v int64) *ListChartRepositoryResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetInstanceId(v string) *ListChartRepositoryResponseBodyRepositories {
	s.InstanceId = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetModifiedTime(v int64) *ListChartRepositoryResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoId(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoName(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoName = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoNamespaceName(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoStatus(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoStatus = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoType(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoType = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetResourceGroupId(v string) *ListChartRepositoryResponseBodyRepositories {
	s.ResourceGroupId = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetSummary(v string) *ListChartRepositoryResponseBodyRepositories {
	s.Summary = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) Validate() error {
	return dara.Validate(s)
}
