// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeVersionReference interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *CodeVersionReference
	GetBranch() *string
	SetCommitID(v string) *CodeVersionReference
	GetCommitID() *string
}

type CodeVersionReference struct {
	// example:
	//
	// main
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 12721ec262d03a93809ba2bbc717963cb298ceca
	CommitID *string `json:"commitID,omitempty" xml:"commitID,omitempty"`
}

func (s CodeVersionReference) String() string {
	return dara.Prettify(s)
}

func (s CodeVersionReference) GoString() string {
	return s.String()
}

func (s *CodeVersionReference) GetBranch() *string {
	return s.Branch
}

func (s *CodeVersionReference) GetCommitID() *string {
	return s.CommitID
}

func (s *CodeVersionReference) SetBranch(v string) *CodeVersionReference {
	s.Branch = &v
	return s
}

func (s *CodeVersionReference) SetCommitID(v string) *CodeVersionReference {
	s.CommitID = &v
	return s
}

func (s *CodeVersionReference) Validate() error {
	return dara.Validate(s)
}
