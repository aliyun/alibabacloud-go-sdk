// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatDetail interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ChatDetail
	GetContent() *string
	SetGmtCreateTime(v string) *ChatDetail
	GetGmtCreateTime() *string
	SetUserInfo(v *ChatDetailUserInfo) *ChatDetail
	GetUserInfo() *ChatDetailUserInfo
}

type ChatDetail struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreateTime *string             `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	UserInfo      *ChatDetailUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ChatDetail) String() string {
	return dara.Prettify(s)
}

func (s ChatDetail) GoString() string {
	return s.String()
}

func (s *ChatDetail) GetContent() *string {
	return s.Content
}

func (s *ChatDetail) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ChatDetail) GetUserInfo() *ChatDetailUserInfo {
	return s.UserInfo
}

func (s *ChatDetail) SetContent(v string) *ChatDetail {
	s.Content = &v
	return s
}

func (s *ChatDetail) SetGmtCreateTime(v string) *ChatDetail {
	s.GmtCreateTime = &v
	return s
}

func (s *ChatDetail) SetUserInfo(v *ChatDetailUserInfo) *ChatDetail {
	s.UserInfo = v
	return s
}

func (s *ChatDetail) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatDetailUserInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ChatDetailUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatDetailUserInfo) GoString() string {
	return s.String()
}

func (s *ChatDetailUserInfo) GetName() *string {
	return s.Name
}

func (s *ChatDetailUserInfo) GetRole() *string {
	return s.Role
}

func (s *ChatDetailUserInfo) SetName(v string) *ChatDetailUserInfo {
	s.Name = &v
	return s
}

func (s *ChatDetailUserInfo) SetRole(v string) *ChatDetailUserInfo {
	s.Role = &v
	return s
}

func (s *ChatDetailUserInfo) Validate() error {
	return dara.Validate(s)
}
