// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserSummaryModel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokens(v []*AccessTokenModel) *UserSummaryModel
	GetAccessTokens() []*AccessTokenModel
	SetHost(v string) *UserSummaryModel
	GetHost() *string
	SetRamUser(v string) *UserSummaryModel
	GetRamUser() *string
	SetStatus(v string) *UserSummaryModel
	GetStatus() *string
	SetUserName(v string) *UserSummaryModel
	GetUserName() *string
}

type UserSummaryModel struct {
	AccessTokens []*AccessTokenModel `json:"AccessTokens,omitempty" xml:"AccessTokens,omitempty" type:"Repeated"`
	Host         *string             `json:"Host,omitempty" xml:"Host,omitempty"`
	RamUser      *string             `json:"RamUser,omitempty" xml:"RamUser,omitempty"`
	Status       *string             `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName     *string             `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UserSummaryModel) String() string {
	return dara.Prettify(s)
}

func (s UserSummaryModel) GoString() string {
	return s.String()
}

func (s *UserSummaryModel) GetAccessTokens() []*AccessTokenModel {
	return s.AccessTokens
}

func (s *UserSummaryModel) GetHost() *string {
	return s.Host
}

func (s *UserSummaryModel) GetRamUser() *string {
	return s.RamUser
}

func (s *UserSummaryModel) GetStatus() *string {
	return s.Status
}

func (s *UserSummaryModel) GetUserName() *string {
	return s.UserName
}

func (s *UserSummaryModel) SetAccessTokens(v []*AccessTokenModel) *UserSummaryModel {
	s.AccessTokens = v
	return s
}

func (s *UserSummaryModel) SetHost(v string) *UserSummaryModel {
	s.Host = &v
	return s
}

func (s *UserSummaryModel) SetRamUser(v string) *UserSummaryModel {
	s.RamUser = &v
	return s
}

func (s *UserSummaryModel) SetStatus(v string) *UserSummaryModel {
	s.Status = &v
	return s
}

func (s *UserSummaryModel) SetUserName(v string) *UserSummaryModel {
	s.UserName = &v
	return s
}

func (s *UserSummaryModel) Validate() error {
	if s.AccessTokens != nil {
		for _, item := range s.AccessTokens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
