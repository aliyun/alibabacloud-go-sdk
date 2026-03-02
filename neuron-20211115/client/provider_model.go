// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvider interface {
	dara.Model
	String() string
	GoString() string
	SetCompany(v string) *Provider
	GetCompany() *string
	SetCompanyId(v int64) *Provider
	GetCompanyId() *int64
	SetContact(v string) *Provider
	GetContact() *string
	SetEmail(v string) *Provider
	GetEmail() *string
	SetTeam(v string) *Provider
	GetTeam() *string
	SetTeamId(v int64) *Provider
	GetTeamId() *int64
}

type Provider struct {
	Company   *string `json:"company,omitempty" xml:"company,omitempty"`
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Contact   *string `json:"contact,omitempty" xml:"contact,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	Team      *string `json:"team,omitempty" xml:"team,omitempty"`
	TeamId    *int64  `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s Provider) String() string {
	return dara.Prettify(s)
}

func (s Provider) GoString() string {
	return s.String()
}

func (s *Provider) GetCompany() *string {
	return s.Company
}

func (s *Provider) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *Provider) GetContact() *string {
	return s.Contact
}

func (s *Provider) GetEmail() *string {
	return s.Email
}

func (s *Provider) GetTeam() *string {
	return s.Team
}

func (s *Provider) GetTeamId() *int64 {
	return s.TeamId
}

func (s *Provider) SetCompany(v string) *Provider {
	s.Company = &v
	return s
}

func (s *Provider) SetCompanyId(v int64) *Provider {
	s.CompanyId = &v
	return s
}

func (s *Provider) SetContact(v string) *Provider {
	s.Contact = &v
	return s
}

func (s *Provider) SetEmail(v string) *Provider {
	s.Email = &v
	return s
}

func (s *Provider) SetTeam(v string) *Provider {
	s.Team = &v
	return s
}

func (s *Provider) SetTeamId(v int64) *Provider {
	s.TeamId = &v
	return s
}

func (s *Provider) Validate() error {
	return dara.Validate(s)
}
