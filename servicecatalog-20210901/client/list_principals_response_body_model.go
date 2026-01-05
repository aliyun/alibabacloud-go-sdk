// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrincipalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrincipals(v []*ListPrincipalsResponseBodyPrincipals) *ListPrincipalsResponseBody
	GetPrincipals() []*ListPrincipalsResponseBodyPrincipals
	SetRequestId(v string) *ListPrincipalsResponseBody
	GetRequestId() *string
}

type ListPrincipalsResponseBody struct {
	// The RAM entities.
	Principals []*ListPrincipalsResponseBodyPrincipals `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrincipalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrincipalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrincipalsResponseBody) GetPrincipals() []*ListPrincipalsResponseBodyPrincipals {
	return s.Principals
}

func (s *ListPrincipalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrincipalsResponseBody) SetPrincipals(v []*ListPrincipalsResponseBodyPrincipals) *ListPrincipalsResponseBody {
	s.Principals = v
	return s
}

func (s *ListPrincipalsResponseBody) SetRequestId(v string) *ListPrincipalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrincipalsResponseBody) Validate() error {
	if s.Principals != nil {
		for _, item := range s.Principals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrincipalsResponseBodyPrincipals struct {
	// The ID of the RAM entity.
	//
	// example:
	//
	// 24477111603637****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the RAM entity. Valid values:
	//
	// 	- RamUser: a RAM user
	//
	// 	- RamRole: a RAM role
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListPrincipalsResponseBodyPrincipals) String() string {
	return dara.Prettify(s)
}

func (s ListPrincipalsResponseBodyPrincipals) GoString() string {
	return s.String()
}

func (s *ListPrincipalsResponseBodyPrincipals) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListPrincipalsResponseBodyPrincipals) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPrincipalsResponseBodyPrincipals) SetPrincipalId(v string) *ListPrincipalsResponseBodyPrincipals {
	s.PrincipalId = &v
	return s
}

func (s *ListPrincipalsResponseBodyPrincipals) SetPrincipalType(v string) *ListPrincipalsResponseBodyPrincipals {
	s.PrincipalType = &v
	return s
}

func (s *ListPrincipalsResponseBodyPrincipals) Validate() error {
	return dara.Validate(s)
}
