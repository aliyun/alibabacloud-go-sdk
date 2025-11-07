// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGitBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommit(v *GetGitBranchResponseBodyCommit) *GetGitBranchResponseBody
	GetCommit() *GetGitBranchResponseBodyCommit
	SetRequestId(v string) *GetGitBranchResponseBody
	GetRequestId() *string
}

type GetGitBranchResponseBody struct {
	Commit *GetGitBranchResponseBodyCommit `json:"Commit,omitempty" xml:"Commit,omitempty" type:"Struct"`
	// example:
	//
	// AA9FA778-AE4B-55EC-81CC-C46BAF08A166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGitBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGitBranchResponseBody) GoString() string {
	return s.String()
}

func (s *GetGitBranchResponseBody) GetCommit() *GetGitBranchResponseBodyCommit {
	return s.Commit
}

func (s *GetGitBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGitBranchResponseBody) SetCommit(v *GetGitBranchResponseBodyCommit) *GetGitBranchResponseBody {
	s.Commit = v
	return s
}

func (s *GetGitBranchResponseBody) SetRequestId(v string) *GetGitBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGitBranchResponseBody) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGitBranchResponseBodyCommit struct {
	// git commit message
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// git commit hash
	//
	// example:
	//
	// 51417755
	Sha *string `json:"Sha,omitempty" xml:"Sha,omitempty"`
}

func (s GetGitBranchResponseBodyCommit) String() string {
	return dara.Prettify(s)
}

func (s GetGitBranchResponseBodyCommit) GoString() string {
	return s.String()
}

func (s *GetGitBranchResponseBodyCommit) GetMessage() *string {
	return s.Message
}

func (s *GetGitBranchResponseBodyCommit) GetSha() *string {
	return s.Sha
}

func (s *GetGitBranchResponseBodyCommit) SetMessage(v string) *GetGitBranchResponseBodyCommit {
	s.Message = &v
	return s
}

func (s *GetGitBranchResponseBodyCommit) SetSha(v string) *GetGitBranchResponseBodyCommit {
	s.Sha = &v
	return s
}

func (s *GetGitBranchResponseBodyCommit) Validate() error {
	return dara.Validate(s)
}
