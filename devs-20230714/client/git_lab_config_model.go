// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGitLabConfig interface {
	dara.Model
	String() string
	GoString() string
	SetIsFixedIP(v bool) *GitLabConfig
	GetIsFixedIP() *bool
	SetToken(v string) *GitLabConfig
	GetToken() *string
	SetUri(v string) *GitLabConfig
	GetUri() *string
}

type GitLabConfig struct {
	IsFixedIP *bool `json:"isFixedIP,omitempty" xml:"isFixedIP,omitempty"`
	// example:
	//
	// your-token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// http://gitlab.c16194660f14898a0810408171302ac.cn-shanghai.alicontainer.com/
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s GitLabConfig) String() string {
	return dara.Prettify(s)
}

func (s GitLabConfig) GoString() string {
	return s.String()
}

func (s *GitLabConfig) GetIsFixedIP() *bool {
	return s.IsFixedIP
}

func (s *GitLabConfig) GetToken() *string {
	return s.Token
}

func (s *GitLabConfig) GetUri() *string {
	return s.Uri
}

func (s *GitLabConfig) SetIsFixedIP(v bool) *GitLabConfig {
	s.IsFixedIP = &v
	return s
}

func (s *GitLabConfig) SetToken(v string) *GitLabConfig {
	s.Token = &v
	return s
}

func (s *GitLabConfig) SetUri(v string) *GitLabConfig {
	s.Uri = &v
	return s
}

func (s *GitLabConfig) Validate() error {
	return dara.Validate(s)
}
