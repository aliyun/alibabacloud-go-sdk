// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSettled interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *Settled
	GetAccountId() *string
	SetApplication(v string) *Settled
	GetApplication() *string
	SetDeveloperName(v string) *Settled
	GetDeveloperName() *string
	SetEmail(v string) *Settled
	GetEmail() *string
	SetEnterpriseName(v string) *Settled
	GetEnterpriseName() *string
	SetGmtCreate(v string) *Settled
	GetGmtCreate() *string
	SetId(v int64) *Settled
	GetId() *int64
	SetPhone(v string) *Settled
	GetPhone() *string
	SetRequestId(v string) *Settled
	GetRequestId() *string
	SetStatus(v string) *Settled
	GetStatus() *string
	SetType(v string) *Settled
	GetType() *string
	SetUsage(v string) *Settled
	GetUsage() *string
}

type Settled struct {
	AccountId      *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Application    *string `json:"application,omitempty" xml:"application,omitempty"`
	DeveloperName  *string `json:"developerName,omitempty" xml:"developerName,omitempty"`
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	EnterpriseName *string `json:"enterpriseName,omitempty" xml:"enterpriseName,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
	Usage          *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s Settled) String() string {
	return dara.Prettify(s)
}

func (s Settled) GoString() string {
	return s.String()
}

func (s *Settled) GetAccountId() *string {
	return s.AccountId
}

func (s *Settled) GetApplication() *string {
	return s.Application
}

func (s *Settled) GetDeveloperName() *string {
	return s.DeveloperName
}

func (s *Settled) GetEmail() *string {
	return s.Email
}

func (s *Settled) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *Settled) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Settled) GetId() *int64 {
	return s.Id
}

func (s *Settled) GetPhone() *string {
	return s.Phone
}

func (s *Settled) GetRequestId() *string {
	return s.RequestId
}

func (s *Settled) GetStatus() *string {
	return s.Status
}

func (s *Settled) GetType() *string {
	return s.Type
}

func (s *Settled) GetUsage() *string {
	return s.Usage
}

func (s *Settled) SetAccountId(v string) *Settled {
	s.AccountId = &v
	return s
}

func (s *Settled) SetApplication(v string) *Settled {
	s.Application = &v
	return s
}

func (s *Settled) SetDeveloperName(v string) *Settled {
	s.DeveloperName = &v
	return s
}

func (s *Settled) SetEmail(v string) *Settled {
	s.Email = &v
	return s
}

func (s *Settled) SetEnterpriseName(v string) *Settled {
	s.EnterpriseName = &v
	return s
}

func (s *Settled) SetGmtCreate(v string) *Settled {
	s.GmtCreate = &v
	return s
}

func (s *Settled) SetId(v int64) *Settled {
	s.Id = &v
	return s
}

func (s *Settled) SetPhone(v string) *Settled {
	s.Phone = &v
	return s
}

func (s *Settled) SetRequestId(v string) *Settled {
	s.RequestId = &v
	return s
}

func (s *Settled) SetStatus(v string) *Settled {
	s.Status = &v
	return s
}

func (s *Settled) SetType(v string) *Settled {
	s.Type = &v
	return s
}

func (s *Settled) SetUsage(v string) *Settled {
	s.Usage = &v
	return s
}

func (s *Settled) Validate() error {
	return dara.Validate(s)
}
