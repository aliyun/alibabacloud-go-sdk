// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSettledCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *SettledCreateCmd
	GetAccountId() *string
	SetApplication(v string) *SettledCreateCmd
	GetApplication() *string
	SetDeveloperName(v string) *SettledCreateCmd
	GetDeveloperName() *string
	SetEmail(v string) *SettledCreateCmd
	GetEmail() *string
	SetEnterpriseName(v string) *SettledCreateCmd
	GetEnterpriseName() *string
	SetPhone(v string) *SettledCreateCmd
	GetPhone() *string
	SetUsage(v string) *SettledCreateCmd
	GetUsage() *string
}

type SettledCreateCmd struct {
	AccountId      *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Application    *string `json:"application,omitempty" xml:"application,omitempty"`
	DeveloperName  *string `json:"developerName,omitempty" xml:"developerName,omitempty"`
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	EnterpriseName *string `json:"enterpriseName,omitempty" xml:"enterpriseName,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Usage          *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s SettledCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s SettledCreateCmd) GoString() string {
	return s.String()
}

func (s *SettledCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *SettledCreateCmd) GetApplication() *string {
	return s.Application
}

func (s *SettledCreateCmd) GetDeveloperName() *string {
	return s.DeveloperName
}

func (s *SettledCreateCmd) GetEmail() *string {
	return s.Email
}

func (s *SettledCreateCmd) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *SettledCreateCmd) GetPhone() *string {
	return s.Phone
}

func (s *SettledCreateCmd) GetUsage() *string {
	return s.Usage
}

func (s *SettledCreateCmd) SetAccountId(v string) *SettledCreateCmd {
	s.AccountId = &v
	return s
}

func (s *SettledCreateCmd) SetApplication(v string) *SettledCreateCmd {
	s.Application = &v
	return s
}

func (s *SettledCreateCmd) SetDeveloperName(v string) *SettledCreateCmd {
	s.DeveloperName = &v
	return s
}

func (s *SettledCreateCmd) SetEmail(v string) *SettledCreateCmd {
	s.Email = &v
	return s
}

func (s *SettledCreateCmd) SetEnterpriseName(v string) *SettledCreateCmd {
	s.EnterpriseName = &v
	return s
}

func (s *SettledCreateCmd) SetPhone(v string) *SettledCreateCmd {
	s.Phone = &v
	return s
}

func (s *SettledCreateCmd) SetUsage(v string) *SettledCreateCmd {
	s.Usage = &v
	return s
}

func (s *SettledCreateCmd) Validate() error {
	return dara.Validate(s)
}
