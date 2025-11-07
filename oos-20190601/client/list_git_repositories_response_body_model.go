// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListGitRepositoriesResponseBody
	GetCount() *int32
	SetGitRepos(v []*ListGitRepositoriesResponseBodyGitRepos) *ListGitRepositoriesResponseBody
	GetGitRepos() []*ListGitRepositoriesResponseBodyGitRepos
	SetRequestId(v string) *ListGitRepositoriesResponseBody
	GetRequestId() *string
}

type ListGitRepositoriesResponseBody struct {
	// example:
	//
	// 1
	Count    *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	GitRepos []*ListGitRepositoriesResponseBodyGitRepos `json:"GitRepos,omitempty" xml:"GitRepos,omitempty" type:"Repeated"`
	// example:
	//
	// DBA6E6C8-F75D-41DE-AFF5-1FA03F551CA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGitRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListGitRepositoriesResponseBody) GetGitRepos() []*ListGitRepositoriesResponseBodyGitRepos {
	return s.GitRepos
}

func (s *ListGitRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGitRepositoriesResponseBody) SetCount(v int32) *ListGitRepositoriesResponseBody {
	s.Count = &v
	return s
}

func (s *ListGitRepositoriesResponseBody) SetGitRepos(v []*ListGitRepositoriesResponseBodyGitRepos) *ListGitRepositoriesResponseBody {
	s.GitRepos = v
	return s
}

func (s *ListGitRepositoriesResponseBody) SetRequestId(v string) *ListGitRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGitRepositoriesResponseBody) Validate() error {
	if s.GitRepos != nil {
		for _, item := range s.GitRepos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGitRepositoriesResponseBodyGitRepos struct {
	// example:
	//
	// Test secret parameter for vulnerability testing - 1757298077.453695
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// aliyun/repo
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// example:
	//
	// https://github.com/alibaba/fastjson
	HtmlUrl *string `json:"HtmlUrl,omitempty" xml:"HtmlUrl,omitempty"`
	// example:
	//
	// False
	IsPrivate *bool `json:"IsPrivate,omitempty" xml:"IsPrivate,omitempty"`
	// example:
	//
	// 22351123
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListGitRepositoriesResponseBodyGitRepos) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoriesResponseBodyGitRepos) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesResponseBodyGitRepos) GetDescription() *string {
	return s.Description
}

func (s *ListGitRepositoriesResponseBodyGitRepos) GetFullName() *string {
	return s.FullName
}

func (s *ListGitRepositoriesResponseBodyGitRepos) GetHtmlUrl() *string {
	return s.HtmlUrl
}

func (s *ListGitRepositoriesResponseBodyGitRepos) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *ListGitRepositoriesResponseBodyGitRepos) GetRepoId() *int64 {
	return s.RepoId
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetDescription(v string) *ListGitRepositoriesResponseBodyGitRepos {
	s.Description = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetFullName(v string) *ListGitRepositoriesResponseBodyGitRepos {
	s.FullName = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetHtmlUrl(v string) *ListGitRepositoriesResponseBodyGitRepos {
	s.HtmlUrl = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetIsPrivate(v bool) *ListGitRepositoriesResponseBodyGitRepos {
	s.IsPrivate = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetRepoId(v int64) *ListGitRepositoriesResponseBodyGitRepos {
	s.RepoId = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) Validate() error {
	return dara.Validate(s)
}
