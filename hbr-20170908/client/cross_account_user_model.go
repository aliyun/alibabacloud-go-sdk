// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCrossAccountUser interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *CrossAccountUser
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CrossAccountUser
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CrossAccountUser
	GetCrossAccountUserId() *int64
}

type CrossAccountUser struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s CrossAccountUser) String() string {
	return dara.Prettify(s)
}

func (s CrossAccountUser) GoString() string {
	return s.String()
}

func (s *CrossAccountUser) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CrossAccountUser) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CrossAccountUser) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CrossAccountUser) SetCrossAccountRoleName(v string) *CrossAccountUser {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CrossAccountUser) SetCrossAccountType(v string) *CrossAccountUser {
	s.CrossAccountType = &v
	return s
}

func (s *CrossAccountUser) SetCrossAccountUserId(v int64) *CrossAccountUser {
	s.CrossAccountUserId = &v
	return s
}

func (s *CrossAccountUser) Validate() error {
	return dara.Validate(s)
}
