// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUser interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *User
	GetCreatedAt() *int64
	SetCreatedBy(v string) *User
	GetCreatedBy() *string
	SetDisplayName(v string) *User
	GetDisplayName() *string
	SetType(v string) *User
	GetType() *string
	SetUpdatedAt(v int64) *User
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *User
	GetUpdatedBy() *string
	SetUserId(v string) *User
	GetUserId() *string
	SetUserName(v string) *User
	GetUserName() *string
	SetUserPrincipal(v string) *User
	GetUserPrincipal() *string
}

type User struct {
	CreatedAt     *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy     *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	DisplayName   *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt     *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy     *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName      *string `json:"userName,omitempty" xml:"userName,omitempty"`
	UserPrincipal *string `json:"userPrincipal,omitempty" xml:"userPrincipal,omitempty"`
}

func (s User) String() string {
	return dara.Prettify(s)
}

func (s User) GoString() string {
	return s.String()
}

func (s *User) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *User) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *User) GetDisplayName() *string {
	return s.DisplayName
}

func (s *User) GetType() *string {
	return s.Type
}

func (s *User) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *User) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *User) GetUserId() *string {
	return s.UserId
}

func (s *User) GetUserName() *string {
	return s.UserName
}

func (s *User) GetUserPrincipal() *string {
	return s.UserPrincipal
}

func (s *User) SetCreatedAt(v int64) *User {
	s.CreatedAt = &v
	return s
}

func (s *User) SetCreatedBy(v string) *User {
	s.CreatedBy = &v
	return s
}

func (s *User) SetDisplayName(v string) *User {
	s.DisplayName = &v
	return s
}

func (s *User) SetType(v string) *User {
	s.Type = &v
	return s
}

func (s *User) SetUpdatedAt(v int64) *User {
	s.UpdatedAt = &v
	return s
}

func (s *User) SetUpdatedBy(v string) *User {
	s.UpdatedBy = &v
	return s
}

func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}

func (s *User) SetUserName(v string) *User {
	s.UserName = &v
	return s
}

func (s *User) SetUserPrincipal(v string) *User {
	s.UserPrincipal = &v
	return s
}

func (s *User) Validate() error {
	return dara.Validate(s)
}
