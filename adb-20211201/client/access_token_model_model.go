// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessTokenModel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *AccessTokenModel
	GetAccessToken() *string
	SetComment(v string) *AccessTokenModel
	GetComment() *string
	SetCreatedAt(v string) *AccessTokenModel
	GetCreatedAt() *string
	SetExpiredAt(v string) *AccessTokenModel
	GetExpiredAt() *string
	SetStatus(v string) *AccessTokenModel
	GetStatus() *string
}

type AccessTokenModel struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Comment     *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ExpiredAt   *string `json:"ExpiredAt,omitempty" xml:"ExpiredAt,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AccessTokenModel) String() string {
	return dara.Prettify(s)
}

func (s AccessTokenModel) GoString() string {
	return s.String()
}

func (s *AccessTokenModel) GetAccessToken() *string {
	return s.AccessToken
}

func (s *AccessTokenModel) GetComment() *string {
	return s.Comment
}

func (s *AccessTokenModel) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AccessTokenModel) GetExpiredAt() *string {
	return s.ExpiredAt
}

func (s *AccessTokenModel) GetStatus() *string {
	return s.Status
}

func (s *AccessTokenModel) SetAccessToken(v string) *AccessTokenModel {
	s.AccessToken = &v
	return s
}

func (s *AccessTokenModel) SetComment(v string) *AccessTokenModel {
	s.Comment = &v
	return s
}

func (s *AccessTokenModel) SetCreatedAt(v string) *AccessTokenModel {
	s.CreatedAt = &v
	return s
}

func (s *AccessTokenModel) SetExpiredAt(v string) *AccessTokenModel {
	s.ExpiredAt = &v
	return s
}

func (s *AccessTokenModel) SetStatus(v string) *AccessTokenModel {
	s.Status = &v
	return s
}

func (s *AccessTokenModel) Validate() error {
	return dara.Validate(s)
}
