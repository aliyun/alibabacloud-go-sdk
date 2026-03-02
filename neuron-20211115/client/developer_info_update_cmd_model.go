// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeveloperInfoUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeveloperInfoUpdateCmd
	GetAccountId() *string
	SetEmail(v string) *DeveloperInfoUpdateCmd
	GetEmail() *string
	SetEnterpriseId(v int64) *DeveloperInfoUpdateCmd
	GetEnterpriseId() *int64
	SetName(v string) *DeveloperInfoUpdateCmd
	GetName() *string
	SetPhone(v string) *DeveloperInfoUpdateCmd
	GetPhone() *string
}

type DeveloperInfoUpdateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// xxx@alibaba.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 尚仁
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 15113456789
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s DeveloperInfoUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s DeveloperInfoUpdateCmd) GoString() string {
	return s.String()
}

func (s *DeveloperInfoUpdateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *DeveloperInfoUpdateCmd) GetEmail() *string {
	return s.Email
}

func (s *DeveloperInfoUpdateCmd) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *DeveloperInfoUpdateCmd) GetName() *string {
	return s.Name
}

func (s *DeveloperInfoUpdateCmd) GetPhone() *string {
	return s.Phone
}

func (s *DeveloperInfoUpdateCmd) SetAccountId(v string) *DeveloperInfoUpdateCmd {
	s.AccountId = &v
	return s
}

func (s *DeveloperInfoUpdateCmd) SetEmail(v string) *DeveloperInfoUpdateCmd {
	s.Email = &v
	return s
}

func (s *DeveloperInfoUpdateCmd) SetEnterpriseId(v int64) *DeveloperInfoUpdateCmd {
	s.EnterpriseId = &v
	return s
}

func (s *DeveloperInfoUpdateCmd) SetName(v string) *DeveloperInfoUpdateCmd {
	s.Name = &v
	return s
}

func (s *DeveloperInfoUpdateCmd) SetPhone(v string) *DeveloperInfoUpdateCmd {
	s.Phone = &v
	return s
}

func (s *DeveloperInfoUpdateCmd) Validate() error {
	return dara.Validate(s)
}
