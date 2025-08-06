// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateVodDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *UpdateVodDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *UpdateVodDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *UpdateVodDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *UpdateVodDomainRequest
	GetTopLevelDomain() *string
}

type UpdateVodDomainRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of origin servers.
	//
	// example:
	//
	// [{"content":"1.1.1.1","type":"ipaddr","priority":"20","port":80}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The top-level domain.
	//
	// example:
	//
	// example.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s UpdateVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVodDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateVodDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateVodDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *UpdateVodDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *UpdateVodDomainRequest) SetDomainName(v string) *UpdateVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateVodDomainRequest) SetOwnerId(v int64) *UpdateVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVodDomainRequest) SetSecurityToken(v string) *UpdateVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateVodDomainRequest) SetSources(v string) *UpdateVodDomainRequest {
	s.Sources = &v
	return s
}

func (s *UpdateVodDomainRequest) SetTopLevelDomain(v string) *UpdateVodDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *UpdateVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
