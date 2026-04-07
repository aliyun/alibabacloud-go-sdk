// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemOnly(v bool) *ResumeSessionRequest
	GetFileSystemOnly() *bool
	SetQualifier(v string) *ResumeSessionRequest
	GetQualifier() *string
}

type ResumeSessionRequest struct {
	FileSystemOnly *bool `json:"fileSystemOnly,omitempty" xml:"fileSystemOnly,omitempty"`
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s ResumeSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeSessionRequest) GoString() string {
	return s.String()
}

func (s *ResumeSessionRequest) GetFileSystemOnly() *bool {
	return s.FileSystemOnly
}

func (s *ResumeSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ResumeSessionRequest) SetFileSystemOnly(v bool) *ResumeSessionRequest {
	s.FileSystemOnly = &v
	return s
}

func (s *ResumeSessionRequest) SetQualifier(v string) *ResumeSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *ResumeSessionRequest) Validate() error {
	return dara.Validate(s)
}
