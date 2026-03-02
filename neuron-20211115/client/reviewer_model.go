// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewer interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *Reviewer
	GetAccountId() *string
	SetEnterpriseId(v int64) *Reviewer
	GetEnterpriseId() *int64
	SetId(v int64) *Reviewer
	GetId() *int64
	SetName(v string) *Reviewer
	GetName() *string
}

type Reviewer struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	EnterpriseId *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Reviewer) String() string {
	return dara.Prettify(s)
}

func (s Reviewer) GoString() string {
	return s.String()
}

func (s *Reviewer) GetAccountId() *string {
	return s.AccountId
}

func (s *Reviewer) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *Reviewer) GetId() *int64 {
	return s.Id
}

func (s *Reviewer) GetName() *string {
	return s.Name
}

func (s *Reviewer) SetAccountId(v string) *Reviewer {
	s.AccountId = &v
	return s
}

func (s *Reviewer) SetEnterpriseId(v int64) *Reviewer {
	s.EnterpriseId = &v
	return s
}

func (s *Reviewer) SetId(v int64) *Reviewer {
	s.Id = &v
	return s
}

func (s *Reviewer) SetName(v string) *Reviewer {
	s.Name = &v
	return s
}

func (s *Reviewer) Validate() error {
	return dara.Validate(s)
}
