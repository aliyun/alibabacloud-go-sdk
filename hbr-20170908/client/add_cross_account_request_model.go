// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCrossAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AddCrossAccountRequest
	GetAlias() *string
	SetCrossAccountRoleName(v string) *AddCrossAccountRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *AddCrossAccountRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *AddCrossAccountRequest
	GetCrossAccountUserId() *int64
}

type AddCrossAccountRequest struct {
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// example:
	//
	// 1589753xxxxxx625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s AddCrossAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCrossAccountRequest) GoString() string {
	return s.String()
}

func (s *AddCrossAccountRequest) GetAlias() *string {
	return s.Alias
}

func (s *AddCrossAccountRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *AddCrossAccountRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *AddCrossAccountRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *AddCrossAccountRequest) SetAlias(v string) *AddCrossAccountRequest {
	s.Alias = &v
	return s
}

func (s *AddCrossAccountRequest) SetCrossAccountRoleName(v string) *AddCrossAccountRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *AddCrossAccountRequest) SetCrossAccountType(v string) *AddCrossAccountRequest {
	s.CrossAccountType = &v
	return s
}

func (s *AddCrossAccountRequest) SetCrossAccountUserId(v int64) *AddCrossAccountRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *AddCrossAccountRequest) Validate() error {
	return dara.Validate(s)
}
