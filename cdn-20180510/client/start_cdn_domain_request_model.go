// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StartCdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StartCdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StartCdnDomainRequest
	GetSecurityToken() *string
}

type StartCdnDomainRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StartCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StartCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StartCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartCdnDomainRequest) SetDomainName(v string) *StartCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartCdnDomainRequest) SetOwnerId(v int64) *StartCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartCdnDomainRequest) SetSecurityToken(v string) *StartCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
