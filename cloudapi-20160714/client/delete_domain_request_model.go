// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteDomainRequest
	GetDomainName() *string
	SetGroupId(v string) *DeleteDomainRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DeleteDomainRequest
	GetSecurityToken() *string
}

type DeleteDomainRequest struct {
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.demo.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the API group to which the domain name to be unbound is bound. This ID is generated by the system and globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 927d50c0f2e54b359919923d908bb015
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDomainRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainRequest) SetGroupId(v string) *DeleteDomainRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDomainRequest) SetSecurityToken(v string) *DeleteDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
