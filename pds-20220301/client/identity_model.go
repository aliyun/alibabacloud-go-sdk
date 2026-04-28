// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentity interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityId(v string) *Identity
	GetIdentityId() *string
	SetIdentityType(v string) *Identity
	GetIdentityType() *string
}

type Identity struct {
	// The ID of the user or the group.
	//
	// example:
	//
	// 16435bdf934248b788b7b3771ee9a3dw
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the identity. Valid values:
	//
	// 	- IT_User
	//
	// 	- IT_Group
	//
	// example:
	//
	// IT_User
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s Identity) String() string {
	return dara.Prettify(s)
}

func (s Identity) GoString() string {
	return s.String()
}

func (s *Identity) GetIdentityId() *string {
	return s.IdentityId
}

func (s *Identity) GetIdentityType() *string {
	return s.IdentityType
}

func (s *Identity) SetIdentityId(v string) *Identity {
	s.IdentityId = &v
	return s
}

func (s *Identity) SetIdentityType(v string) *Identity {
	s.IdentityType = &v
	return s
}

func (s *Identity) Validate() error {
	return dara.Validate(s)
}
