// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGitRepositoryExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRepoExists(v bool) *CheckGitRepositoryExistsResponseBody
	GetRepoExists() *bool
	SetRequestId(v string) *CheckGitRepositoryExistsResponseBody
	GetRequestId() *string
}

type CheckGitRepositoryExistsResponseBody struct {
	// example:
	//
	// true
	RepoExists *bool `json:"RepoExists,omitempty" xml:"RepoExists,omitempty"`
	// example:
	//
	// AA9FA778-AE4B-55EC-81CC-C46BAF08A166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckGitRepositoryExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckGitRepositoryExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGitRepositoryExistsResponseBody) GetRepoExists() *bool {
	return s.RepoExists
}

func (s *CheckGitRepositoryExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckGitRepositoryExistsResponseBody) SetRepoExists(v bool) *CheckGitRepositoryExistsResponseBody {
	s.RepoExists = &v
	return s
}

func (s *CheckGitRepositoryExistsResponseBody) SetRequestId(v string) *CheckGitRepositoryExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGitRepositoryExistsResponseBody) Validate() error {
	return dara.Validate(s)
}
