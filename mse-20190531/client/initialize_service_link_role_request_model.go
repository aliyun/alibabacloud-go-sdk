// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeServiceLinkRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *InitializeServiceLinkRoleRequest
	GetAcceptLanguage() *string
	SetRoleName(v string) *InitializeServiceLinkRoleRequest
	GetRoleName() *string
	SetToken(v string) *InitializeServiceLinkRoleRequest
	GetToken() *string
}

type InitializeServiceLinkRoleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// AliyunServiceRoleForMSE
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// ""
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s InitializeServiceLinkRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeServiceLinkRoleRequest) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkRoleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *InitializeServiceLinkRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *InitializeServiceLinkRoleRequest) GetToken() *string {
	return s.Token
}

func (s *InitializeServiceLinkRoleRequest) SetAcceptLanguage(v string) *InitializeServiceLinkRoleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *InitializeServiceLinkRoleRequest) SetRoleName(v string) *InitializeServiceLinkRoleRequest {
	s.RoleName = &v
	return s
}

func (s *InitializeServiceLinkRoleRequest) SetToken(v string) *InitializeServiceLinkRoleRequest {
	s.Token = &v
	return s
}

func (s *InitializeServiceLinkRoleRequest) Validate() error {
	return dara.Validate(s)
}
