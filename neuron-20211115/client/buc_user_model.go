// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucUser interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *BucUser
	GetAccount() *string
	SetEmpId(v string) *BucUser
	GetEmpId() *string
	SetEmpType(v string) *BucUser
	GetEmpType() *string
	SetEnterpriseId(v int64) *BucUser
	GetEnterpriseId() *int64
	SetName(v string) *BucUser
	GetName() *string
	SetNickName(v string) *BucUser
	GetNickName() *string
	SetPersonalPhotoUrl(v string) *BucUser
	GetPersonalPhotoUrl() *string
	SetRequestId(v string) *BucUser
	GetRequestId() *string
}

type BucUser struct {
	Account          *string `json:"account,omitempty" xml:"account,omitempty"`
	EmpId            *string `json:"empId,omitempty" xml:"empId,omitempty"`
	EmpType          *string `json:"empType,omitempty" xml:"empType,omitempty"`
	EnterpriseId     *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NickName         *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	PersonalPhotoUrl *string `json:"personalPhotoUrl,omitempty" xml:"personalPhotoUrl,omitempty"`
	RequestId        *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BucUser) String() string {
	return dara.Prettify(s)
}

func (s BucUser) GoString() string {
	return s.String()
}

func (s *BucUser) GetAccount() *string {
	return s.Account
}

func (s *BucUser) GetEmpId() *string {
	return s.EmpId
}

func (s *BucUser) GetEmpType() *string {
	return s.EmpType
}

func (s *BucUser) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *BucUser) GetName() *string {
	return s.Name
}

func (s *BucUser) GetNickName() *string {
	return s.NickName
}

func (s *BucUser) GetPersonalPhotoUrl() *string {
	return s.PersonalPhotoUrl
}

func (s *BucUser) GetRequestId() *string {
	return s.RequestId
}

func (s *BucUser) SetAccount(v string) *BucUser {
	s.Account = &v
	return s
}

func (s *BucUser) SetEmpId(v string) *BucUser {
	s.EmpId = &v
	return s
}

func (s *BucUser) SetEmpType(v string) *BucUser {
	s.EmpType = &v
	return s
}

func (s *BucUser) SetEnterpriseId(v int64) *BucUser {
	s.EnterpriseId = &v
	return s
}

func (s *BucUser) SetName(v string) *BucUser {
	s.Name = &v
	return s
}

func (s *BucUser) SetNickName(v string) *BucUser {
	s.NickName = &v
	return s
}

func (s *BucUser) SetPersonalPhotoUrl(v string) *BucUser {
	s.PersonalPhotoUrl = &v
	return s
}

func (s *BucUser) SetRequestId(v string) *BucUser {
	s.RequestId = &v
	return s
}

func (s *BucUser) Validate() error {
	return dara.Validate(s)
}
