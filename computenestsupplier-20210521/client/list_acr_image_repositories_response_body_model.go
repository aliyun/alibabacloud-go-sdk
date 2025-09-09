// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcrImageRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAcrImageRepositoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAcrImageRepositoriesResponseBody
	GetNextToken() *string
	SetRepositories(v []*ListAcrImageRepositoriesResponseBodyRepositories) *ListAcrImageRepositoriesResponseBody
	GetRepositories() []*ListAcrImageRepositoriesResponseBodyRepositories
	SetRequestId(v string) *ListAcrImageRepositoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAcrImageRepositoriesResponseBody
	GetTotalCount() *int32
}

type ListAcrImageRepositoriesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The images.
	Repositories []*ListAcrImageRepositoriesResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C4A145D8-6F6C-532A-9001-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcrImageRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAcrImageRepositoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAcrImageRepositoriesResponseBody) GetRepositories() []*ListAcrImageRepositoriesResponseBodyRepositories {
	return s.Repositories
}

func (s *ListAcrImageRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAcrImageRepositoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAcrImageRepositoriesResponseBody) SetMaxResults(v int32) *ListAcrImageRepositoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetNextToken(v string) *ListAcrImageRepositoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetRepositories(v []*ListAcrImageRepositoriesResponseBodyRepositories) *ListAcrImageRepositoriesResponseBody {
	s.Repositories = v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetRequestId(v string) *ListAcrImageRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetTotalCount(v int32) *ListAcrImageRepositoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAcrImageRepositoriesResponseBodyRepositories struct {
	// The time when the image was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the image was modified.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The namespace of the repository
	//
	// example:
	//
	// computenest
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The image repo ID.
	//
	// example:
	//
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The image repo name.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `Private`: a private repository
	//
	// 	- `Public`: a public repository
	//
	// example:
	//
	// Private
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
}

func (s ListAcrImageRepositoriesResponseBodyRepositories) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageRepositoriesResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) GetRepoId() *string {
	return s.RepoId
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) GetRepoName() *string {
	return s.RepoName
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) GetRepoType() *string {
	return s.RepoType
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetCreateTime(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetModifiedTime(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetNamespace(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.Namespace = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoId(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoName(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoName = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoType(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoType = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) Validate() error {
	return dara.Validate(s)
}
