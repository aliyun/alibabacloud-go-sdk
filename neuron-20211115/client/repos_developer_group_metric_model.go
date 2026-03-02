// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReposDeveloperGroupMetric interface {
	dara.Model
	String() string
	GoString() string
	SetAddLines(v int64) *ReposDeveloperGroupMetric
	GetAddLines() *int64
	SetCommentLines(v int64) *ReposDeveloperGroupMetric
	GetCommentLines() *int64
	SetCommitCnt(v int64) *ReposDeveloperGroupMetric
	GetCommitCnt() *int64
	SetDelLines(v int64) *ReposDeveloperGroupMetric
	GetDelLines() *int64
	SetModLines(v int64) *ReposDeveloperGroupMetric
	GetModLines() *int64
	SetShowName(v string) *ReposDeveloperGroupMetric
	GetShowName() *string
	SetWorkNo(v string) *ReposDeveloperGroupMetric
	GetWorkNo() *string
}

type ReposDeveloperGroupMetric struct {
	// example:
	//
	// 342
	AddLines *int64 `json:"addLines,omitempty" xml:"addLines,omitempty"`
	// example:
	//
	// 242
	CommentLines *int64 `json:"commentLines,omitempty" xml:"commentLines,omitempty"`
	// example:
	//
	// 14
	CommitCnt *int64 `json:"commitCnt,omitempty" xml:"commitCnt,omitempty"`
	// example:
	//
	// 23
	DelLines *int64 `json:"delLines,omitempty" xml:"delLines,omitempty"`
	// example:
	//
	// 129
	ModLines *int64 `json:"modLines,omitempty" xml:"modLines,omitempty"`
	// example:
	//
	// 旭坤
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// example:
	//
	// 64634
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s ReposDeveloperGroupMetric) String() string {
	return dara.Prettify(s)
}

func (s ReposDeveloperGroupMetric) GoString() string {
	return s.String()
}

func (s *ReposDeveloperGroupMetric) GetAddLines() *int64 {
	return s.AddLines
}

func (s *ReposDeveloperGroupMetric) GetCommentLines() *int64 {
	return s.CommentLines
}

func (s *ReposDeveloperGroupMetric) GetCommitCnt() *int64 {
	return s.CommitCnt
}

func (s *ReposDeveloperGroupMetric) GetDelLines() *int64 {
	return s.DelLines
}

func (s *ReposDeveloperGroupMetric) GetModLines() *int64 {
	return s.ModLines
}

func (s *ReposDeveloperGroupMetric) GetShowName() *string {
	return s.ShowName
}

func (s *ReposDeveloperGroupMetric) GetWorkNo() *string {
	return s.WorkNo
}

func (s *ReposDeveloperGroupMetric) SetAddLines(v int64) *ReposDeveloperGroupMetric {
	s.AddLines = &v
	return s
}

func (s *ReposDeveloperGroupMetric) SetCommentLines(v int64) *ReposDeveloperGroupMetric {
	s.CommentLines = &v
	return s
}

func (s *ReposDeveloperGroupMetric) SetCommitCnt(v int64) *ReposDeveloperGroupMetric {
	s.CommitCnt = &v
	return s
}

func (s *ReposDeveloperGroupMetric) SetDelLines(v int64) *ReposDeveloperGroupMetric {
	s.DelLines = &v
	return s
}

func (s *ReposDeveloperGroupMetric) SetModLines(v int64) *ReposDeveloperGroupMetric {
	s.ModLines = &v
	return s
}

func (s *ReposDeveloperGroupMetric) SetShowName(v string) *ReposDeveloperGroupMetric {
	s.ShowName = &v
	return s
}

func (s *ReposDeveloperGroupMetric) SetWorkNo(v string) *ReposDeveloperGroupMetric {
	s.WorkNo = &v
	return s
}

func (s *ReposDeveloperGroupMetric) Validate() error {
	return dara.Validate(s)
}
