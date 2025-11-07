// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGitRepoFileExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileExists(v bool) *CheckGitRepoFileExistsResponseBody
	GetFileExists() *bool
	SetRequestId(v string) *CheckGitRepoFileExistsResponseBody
	GetRequestId() *string
}

type CheckGitRepoFileExistsResponseBody struct {
	// example:
	//
	// true
	FileExists *bool `json:"FileExists,omitempty" xml:"FileExists,omitempty"`
	// example:
	//
	// A5320F1D-92D9-44BB-A416-5FC525ED6D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckGitRepoFileExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckGitRepoFileExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGitRepoFileExistsResponseBody) GetFileExists() *bool {
	return s.FileExists
}

func (s *CheckGitRepoFileExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckGitRepoFileExistsResponseBody) SetFileExists(v bool) *CheckGitRepoFileExistsResponseBody {
	s.FileExists = &v
	return s
}

func (s *CheckGitRepoFileExistsResponseBody) SetRequestId(v string) *CheckGitRepoFileExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGitRepoFileExistsResponseBody) Validate() error {
	return dara.Validate(s)
}
