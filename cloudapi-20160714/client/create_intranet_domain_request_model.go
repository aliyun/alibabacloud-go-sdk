// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntranetDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateIntranetDomainRequest
	GetGroupId() *string
	SetSecurityToken(v string) *CreateIntranetDomainRequest
	GetSecurityToken() *string
}

type CreateIntranetDomainRequest struct {
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 927d50c0f2e54b359919923d908bb015
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateIntranetDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntranetDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateIntranetDomainRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateIntranetDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateIntranetDomainRequest) SetGroupId(v string) *CreateIntranetDomainRequest {
	s.GroupId = &v
	return s
}

func (s *CreateIntranetDomainRequest) SetSecurityToken(v string) *CreateIntranetDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateIntranetDomainRequest) Validate() error {
	return dara.Validate(s)
}
