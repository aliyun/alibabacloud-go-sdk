// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGitEventSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *GitEventSnapshot
	GetBranch() *string
	SetCommitID(v string) *GitEventSnapshot
	GetCommitID() *string
	SetTag(v string) *GitEventSnapshot
	GetTag() *string
}

type GitEventSnapshot struct {
	// example:
	//
	// main
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 12721ec262d03a93809ba2bbc717963cb298ceca
	CommitID *string `json:"commitID,omitempty" xml:"commitID,omitempty"`
	// example:
	//
	// 1.0
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s GitEventSnapshot) String() string {
	return dara.Prettify(s)
}

func (s GitEventSnapshot) GoString() string {
	return s.String()
}

func (s *GitEventSnapshot) GetBranch() *string {
	return s.Branch
}

func (s *GitEventSnapshot) GetCommitID() *string {
	return s.CommitID
}

func (s *GitEventSnapshot) GetTag() *string {
	return s.Tag
}

func (s *GitEventSnapshot) SetBranch(v string) *GitEventSnapshot {
	s.Branch = &v
	return s
}

func (s *GitEventSnapshot) SetCommitID(v string) *GitEventSnapshot {
	s.CommitID = &v
	return s
}

func (s *GitEventSnapshot) SetTag(v string) *GitEventSnapshot {
	s.Tag = &v
	return s
}

func (s *GitEventSnapshot) Validate() error {
	return dara.Validate(s)
}
