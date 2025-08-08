// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectionSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *GitAccount) *ConnectionSpec
	GetAccount() *GitAccount
	SetGitlabConfig(v *GitLabConfig) *ConnectionSpec
	GetGitlabConfig() *GitLabConfig
	SetPlatform(v string) *ConnectionSpec
	GetPlatform() *string
}

type ConnectionSpec struct {
	Account      *GitAccount   `json:"account,omitempty" xml:"account,omitempty"`
	GitlabConfig *GitLabConfig `json:"gitlabConfig,omitempty" xml:"gitlabConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// github
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s ConnectionSpec) String() string {
	return dara.Prettify(s)
}

func (s ConnectionSpec) GoString() string {
	return s.String()
}

func (s *ConnectionSpec) GetAccount() *GitAccount {
	return s.Account
}

func (s *ConnectionSpec) GetGitlabConfig() *GitLabConfig {
	return s.GitlabConfig
}

func (s *ConnectionSpec) GetPlatform() *string {
	return s.Platform
}

func (s *ConnectionSpec) SetAccount(v *GitAccount) *ConnectionSpec {
	s.Account = v
	return s
}

func (s *ConnectionSpec) SetGitlabConfig(v *GitLabConfig) *ConnectionSpec {
	s.GitlabConfig = v
	return s
}

func (s *ConnectionSpec) SetPlatform(v string) *ConnectionSpec {
	s.Platform = &v
	return s
}

func (s *ConnectionSpec) Validate() error {
	return dara.Validate(s)
}
