// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepoItem interface {
	dara.Model
	String() string
	GoString() string
	SetCodeLines(v int64) *RepoItem
	GetCodeLines() *int64
	SetGitProjectUrl(v string) *RepoItem
	GetGitProjectUrl() *string
	SetOwner(v string) *RepoItem
	GetOwner() *string
	SetRepoShortUrl(v string) *RepoItem
	GetRepoShortUrl() *string
	SetRepoUrl(v string) *RepoItem
	GetRepoUrl() *string
	SetSummary(v string) *RepoItem
	GetSummary() *string
}

type RepoItem struct {
	// example:
	//
	// 2034
	CodeLines *int64 `json:"codeLines,omitempty" xml:"codeLines,omitempty"`
	// example:
	//
	// https://code.aone.alibaba-inc.com/yunmall/yunmall-custome
	GitProjectUrl *string `json:"gitProjectUrl,omitempty" xml:"gitProjectUrl,omitempty"`
	// example:
	//
	// 中驿
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// yunmall/product.git
	RepoShortUrl *string `json:"repoShortUrl,omitempty" xml:"repoShortUrl,omitempty"`
	// example:
	//
	// git@gitlab.alibaba-inc.com:yunmall/product.git
	RepoUrl *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
	// example:
	//
	// 类目
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s RepoItem) String() string {
	return dara.Prettify(s)
}

func (s RepoItem) GoString() string {
	return s.String()
}

func (s *RepoItem) GetCodeLines() *int64 {
	return s.CodeLines
}

func (s *RepoItem) GetGitProjectUrl() *string {
	return s.GitProjectUrl
}

func (s *RepoItem) GetOwner() *string {
	return s.Owner
}

func (s *RepoItem) GetRepoShortUrl() *string {
	return s.RepoShortUrl
}

func (s *RepoItem) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *RepoItem) GetSummary() *string {
	return s.Summary
}

func (s *RepoItem) SetCodeLines(v int64) *RepoItem {
	s.CodeLines = &v
	return s
}

func (s *RepoItem) SetGitProjectUrl(v string) *RepoItem {
	s.GitProjectUrl = &v
	return s
}

func (s *RepoItem) SetOwner(v string) *RepoItem {
	s.Owner = &v
	return s
}

func (s *RepoItem) SetRepoShortUrl(v string) *RepoItem {
	s.RepoShortUrl = &v
	return s
}

func (s *RepoItem) SetRepoUrl(v string) *RepoItem {
	s.RepoUrl = &v
	return s
}

func (s *RepoItem) SetSummary(v string) *RepoItem {
	s.Summary = &v
	return s
}

func (s *RepoItem) Validate() error {
	return dara.Validate(s)
}
