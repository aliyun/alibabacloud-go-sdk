// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGitAccount interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *GitAccount
	GetAvatar() *string
	SetDisplayName(v string) *GitAccount
	GetDisplayName() *string
	SetId(v string) *GitAccount
	GetId() *string
	SetName(v string) *GitAccount
	GetName() *string
	SetUri(v string) *GitAccount
	GetUri() *string
}

type GitAccount struct {
	// example:
	//
	// https://gitee.com/assets/no_portrait.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// your_displayname
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// your_username
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://gitlab.com
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s GitAccount) String() string {
	return dara.Prettify(s)
}

func (s GitAccount) GoString() string {
	return s.String()
}

func (s *GitAccount) GetAvatar() *string {
	return s.Avatar
}

func (s *GitAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GitAccount) GetId() *string {
	return s.Id
}

func (s *GitAccount) GetName() *string {
	return s.Name
}

func (s *GitAccount) GetUri() *string {
	return s.Uri
}

func (s *GitAccount) SetAvatar(v string) *GitAccount {
	s.Avatar = &v
	return s
}

func (s *GitAccount) SetDisplayName(v string) *GitAccount {
	s.DisplayName = &v
	return s
}

func (s *GitAccount) SetId(v string) *GitAccount {
	s.Id = &v
	return s
}

func (s *GitAccount) SetName(v string) *GitAccount {
	s.Name = &v
	return s
}

func (s *GitAccount) SetUri(v string) *GitAccount {
	s.Uri = &v
	return s
}

func (s *GitAccount) Validate() error {
	return dara.Validate(s)
}
