// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebofficeUser interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *WebofficeUser
	GetAvatar() *string
	SetId(v string) *WebofficeUser
	GetId() *string
	SetName(v string) *WebofficeUser
	GetName() *string
}

type WebofficeUser struct {
	// The custom URL of the avatar picture. The avatar picture is displayed on the WebOffice page.
	//
	// example:
	//
	// http://example.com/?id=user1
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The custom user ID. The user ID is displayed on the WebOffice page. A user ID can contain letters and digits and cannot exceed 15 characters in length.
	//
	// example:
	//
	// user1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The custom username. The username is displayed on the WebOffice page. The username must meet the following requirements:
	//
	// 	- A username can contain digits, letters, hyphens (-), underscores (_), plus signs (+), forward slashes (/), equal signs (=), and at signs (@).
	//
	// 	- A username can contain up to 32 characters.
	//
	// example:
	//
	// test-user1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s WebofficeUser) String() string {
	return dara.Prettify(s)
}

func (s WebofficeUser) GoString() string {
	return s.String()
}

func (s *WebofficeUser) GetAvatar() *string {
	return s.Avatar
}

func (s *WebofficeUser) GetId() *string {
	return s.Id
}

func (s *WebofficeUser) GetName() *string {
	return s.Name
}

func (s *WebofficeUser) SetAvatar(v string) *WebofficeUser {
	s.Avatar = &v
	return s
}

func (s *WebofficeUser) SetId(v string) *WebofficeUser {
	s.Id = &v
	return s
}

func (s *WebofficeUser) SetName(v string) *WebofficeUser {
	s.Name = &v
	return s
}

func (s *WebofficeUser) Validate() error {
	return dara.Validate(s)
}
