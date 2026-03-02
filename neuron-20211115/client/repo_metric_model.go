// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepoMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCodeLines(v int64) *RepoMetric
	GetCodeLines() *int64
	SetCommitCnt(v int64) *RepoMetric
	GetCommitCnt() *int64
	SetDeveloperCnt(v int64) *RepoMetric
	GetDeveloperCnt() *int64
	SetRefreshDate(v string) *RepoMetric
	GetRefreshDate() *string
	SetRepoShortUrl(v string) *RepoMetric
	GetRepoShortUrl() *string
	SetRepoUrl(v string) *RepoMetric
	GetRepoUrl() *string
}

type RepoMetric struct {
	// example:
	//
	// 294
	CodeLines *int64 `json:"codeLines,omitempty" xml:"codeLines,omitempty"`
	// example:
	//
	// 23
	CommitCnt *int64 `json:"commitCnt,omitempty" xml:"commitCnt,omitempty"`
	// example:
	//
	// 3
	DeveloperCnt *int64 `json:"developerCnt,omitempty" xml:"developerCnt,omitempty"`
	// example:
	//
	// 2022-05-03T00:00:00.000Z
	RefreshDate *string `json:"refreshDate,omitempty" xml:"refreshDate,omitempty"`
	// example:
	//
	// neuron/catalog.git
	RepoShortUrl *string `json:"repoShortUrl,omitempty" xml:"repoShortUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// git@gitlab.alibaba-inc.com:neuron/catalog.git
	RepoUrl *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
}

func (s RepoMetric) String() string {
	return dara.Prettify(s)
}

func (s RepoMetric) GoString() string {
	return s.String()
}

func (s *RepoMetric) GetCodeLines() *int64 {
	return s.CodeLines
}

func (s *RepoMetric) GetCommitCnt() *int64 {
	return s.CommitCnt
}

func (s *RepoMetric) GetDeveloperCnt() *int64 {
	return s.DeveloperCnt
}

func (s *RepoMetric) GetRefreshDate() *string {
	return s.RefreshDate
}

func (s *RepoMetric) GetRepoShortUrl() *string {
	return s.RepoShortUrl
}

func (s *RepoMetric) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *RepoMetric) SetCodeLines(v int64) *RepoMetric {
	s.CodeLines = &v
	return s
}

func (s *RepoMetric) SetCommitCnt(v int64) *RepoMetric {
	s.CommitCnt = &v
	return s
}

func (s *RepoMetric) SetDeveloperCnt(v int64) *RepoMetric {
	s.DeveloperCnt = &v
	return s
}

func (s *RepoMetric) SetRefreshDate(v string) *RepoMetric {
	s.RefreshDate = &v
	return s
}

func (s *RepoMetric) SetRepoShortUrl(v string) *RepoMetric {
	s.RepoShortUrl = &v
	return s
}

func (s *RepoMetric) SetRepoUrl(v string) *RepoMetric {
	s.RepoUrl = &v
	return s
}

func (s *RepoMetric) Validate() error {
	return dara.Validate(s)
}
