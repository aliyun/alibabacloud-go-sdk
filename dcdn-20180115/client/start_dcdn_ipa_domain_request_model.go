// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDcdnIpaDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StartDcdnIpaDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StartDcdnIpaDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StartDcdnIpaDomainRequest
	GetSecurityToken() *string
}

type StartDcdnIpaDomainRequest struct {
	// The name of the accelerated domain to be enabled. You can specify only one accelerated domain name at a time.
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

func (s StartDcdnIpaDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *StartDcdnIpaDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StartDcdnIpaDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartDcdnIpaDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartDcdnIpaDomainRequest) SetDomainName(v string) *StartDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartDcdnIpaDomainRequest) SetOwnerId(v int64) *StartDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartDcdnIpaDomainRequest) SetSecurityToken(v string) *StartDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartDcdnIpaDomainRequest) Validate() error {
	return dara.Validate(s)
}
