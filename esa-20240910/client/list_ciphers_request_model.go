// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCiphersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphersGroup(v string) *ListCiphersRequest
	GetCiphersGroup() *string
}

type ListCiphersRequest struct {
	// The name of the cipher suite group, which can be: all, strict, custom.
	//
	// This parameter is required.
	//
	// example:
	//
	// strict
	CiphersGroup *string `json:"CiphersGroup,omitempty" xml:"CiphersGroup,omitempty"`
}

func (s ListCiphersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCiphersRequest) GoString() string {
	return s.String()
}

func (s *ListCiphersRequest) GetCiphersGroup() *string {
	return s.CiphersGroup
}

func (s *ListCiphersRequest) SetCiphersGroup(v string) *ListCiphersRequest {
	s.CiphersGroup = &v
	return s
}

func (s *ListCiphersRequest) Validate() error {
	return dara.Validate(s)
}
