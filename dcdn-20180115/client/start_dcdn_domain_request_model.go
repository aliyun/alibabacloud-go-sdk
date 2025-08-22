// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StartDcdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StartDcdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StartDcdnDomainRequest
	GetSecurityToken() *string
}

type StartDcdnDomainRequest struct {
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

func (s StartDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StartDcdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StartDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartDcdnDomainRequest) SetDomainName(v string) *StartDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartDcdnDomainRequest) SetOwnerId(v int64) *StartDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartDcdnDomainRequest) SetSecurityToken(v string) *StartDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
