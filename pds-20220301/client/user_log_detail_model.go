// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserLogDetail interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *UserLogDetail
	GetEmail() *string
	SetExpiredAt(v int64) *UserLogDetail
	GetExpiredAt() *int64
	SetName(v string) *UserLogDetail
	GetName() *string
	SetPhone(v string) *UserLogDetail
	GetPhone() *string
	SetRoleId(v string) *UserLogDetail
	GetRoleId() *string
	SetUpdateTo(v *UserLogDetailUpdateTo) *UserLogDetail
	GetUpdateTo() *UserLogDetailUpdateTo
}

type UserLogDetail struct {
	Email     *string                `json:"email,omitempty" xml:"email,omitempty"`
	ExpiredAt *int64                 `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	Name      *string                `json:"name,omitempty" xml:"name,omitempty"`
	Phone     *string                `json:"phone,omitempty" xml:"phone,omitempty"`
	RoleId    *string                `json:"role_id,omitempty" xml:"role_id,omitempty"`
	UpdateTo  *UserLogDetailUpdateTo `json:"update_to,omitempty" xml:"update_to,omitempty" type:"Struct"`
}

func (s UserLogDetail) String() string {
	return dara.Prettify(s)
}

func (s UserLogDetail) GoString() string {
	return s.String()
}

func (s *UserLogDetail) GetEmail() *string {
	return s.Email
}

func (s *UserLogDetail) GetExpiredAt() *int64 {
	return s.ExpiredAt
}

func (s *UserLogDetail) GetName() *string {
	return s.Name
}

func (s *UserLogDetail) GetPhone() *string {
	return s.Phone
}

func (s *UserLogDetail) GetRoleId() *string {
	return s.RoleId
}

func (s *UserLogDetail) GetUpdateTo() *UserLogDetailUpdateTo {
	return s.UpdateTo
}

func (s *UserLogDetail) SetEmail(v string) *UserLogDetail {
	s.Email = &v
	return s
}

func (s *UserLogDetail) SetExpiredAt(v int64) *UserLogDetail {
	s.ExpiredAt = &v
	return s
}

func (s *UserLogDetail) SetName(v string) *UserLogDetail {
	s.Name = &v
	return s
}

func (s *UserLogDetail) SetPhone(v string) *UserLogDetail {
	s.Phone = &v
	return s
}

func (s *UserLogDetail) SetRoleId(v string) *UserLogDetail {
	s.RoleId = &v
	return s
}

func (s *UserLogDetail) SetUpdateTo(v *UserLogDetailUpdateTo) *UserLogDetail {
	s.UpdateTo = v
	return s
}

func (s *UserLogDetail) Validate() error {
	if s.UpdateTo != nil {
		if err := s.UpdateTo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UserLogDetailUpdateTo struct {
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	ExpiredAt *int64  `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone     *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RoleId    *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s UserLogDetailUpdateTo) String() string {
	return dara.Prettify(s)
}

func (s UserLogDetailUpdateTo) GoString() string {
	return s.String()
}

func (s *UserLogDetailUpdateTo) GetEmail() *string {
	return s.Email
}

func (s *UserLogDetailUpdateTo) GetExpiredAt() *int64 {
	return s.ExpiredAt
}

func (s *UserLogDetailUpdateTo) GetName() *string {
	return s.Name
}

func (s *UserLogDetailUpdateTo) GetPhone() *string {
	return s.Phone
}

func (s *UserLogDetailUpdateTo) GetRoleId() *string {
	return s.RoleId
}

func (s *UserLogDetailUpdateTo) SetEmail(v string) *UserLogDetailUpdateTo {
	s.Email = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetExpiredAt(v int64) *UserLogDetailUpdateTo {
	s.ExpiredAt = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetName(v string) *UserLogDetailUpdateTo {
	s.Name = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetPhone(v string) *UserLogDetailUpdateTo {
	s.Phone = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetRoleId(v string) *UserLogDetailUpdateTo {
	s.RoleId = &v
	return s
}

func (s *UserLogDetailUpdateTo) Validate() error {
	return dara.Validate(s)
}
