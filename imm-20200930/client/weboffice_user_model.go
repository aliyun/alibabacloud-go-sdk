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
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
