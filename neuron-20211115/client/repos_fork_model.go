// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReposFork interface {
	dara.Model
	String() string
	GoString() string
	SetGitGroup(v string) *ReposFork
	GetGitGroup() *string
	SetGmtCreate(v string) *ReposFork
	GetGmtCreate() *string
	SetId(v int64) *ReposFork
	GetId() *int64
	SetIsOrigin(v bool) *ReposFork
	GetIsOrigin() *bool
	SetPbcRepoItems(v []*RepoItem) *ReposFork
	GetPbcRepoItems() []*RepoItem
	SetRequestId(v string) *ReposFork
	GetRequestId() *string
	SetUsage(v string) *ReposFork
	GetUsage() *string
}

type ReposFork struct {
	// example:
	//
	// global-mall
	GitGroup *string `json:"gitGroup,omitempty" xml:"gitGroup,omitempty"`
	// example:
	//
	// 2022-04-03T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsOrigin     *bool       `json:"isOrigin,omitempty" xml:"isOrigin,omitempty"`
	PbcRepoItems []*RepoItem `json:"pbcRepoItems,omitempty" xml:"pbcRepoItems,omitempty" type:"Repeated"`
	RequestId    *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 商城国际版本
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ReposFork) String() string {
	return dara.Prettify(s)
}

func (s ReposFork) GoString() string {
	return s.String()
}

func (s *ReposFork) GetGitGroup() *string {
	return s.GitGroup
}

func (s *ReposFork) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ReposFork) GetId() *int64 {
	return s.Id
}

func (s *ReposFork) GetIsOrigin() *bool {
	return s.IsOrigin
}

func (s *ReposFork) GetPbcRepoItems() []*RepoItem {
	return s.PbcRepoItems
}

func (s *ReposFork) GetRequestId() *string {
	return s.RequestId
}

func (s *ReposFork) GetUsage() *string {
	return s.Usage
}

func (s *ReposFork) SetGitGroup(v string) *ReposFork {
	s.GitGroup = &v
	return s
}

func (s *ReposFork) SetGmtCreate(v string) *ReposFork {
	s.GmtCreate = &v
	return s
}

func (s *ReposFork) SetId(v int64) *ReposFork {
	s.Id = &v
	return s
}

func (s *ReposFork) SetIsOrigin(v bool) *ReposFork {
	s.IsOrigin = &v
	return s
}

func (s *ReposFork) SetPbcRepoItems(v []*RepoItem) *ReposFork {
	s.PbcRepoItems = v
	return s
}

func (s *ReposFork) SetRequestId(v string) *ReposFork {
	s.RequestId = &v
	return s
}

func (s *ReposFork) SetUsage(v string) *ReposFork {
	s.Usage = &v
	return s
}

func (s *ReposFork) Validate() error {
	if s.PbcRepoItems != nil {
		for _, item := range s.PbcRepoItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
