// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StopCdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StopCdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StopCdnDomainRequest
	GetSecurityToken() *string
}

type StopCdnDomainRequest struct {
	// The accelerated domain name that you want to disable. You can specify only one domain name in each request.
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

func (s StopCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StopCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StopCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StopCdnDomainRequest) SetDomainName(v string) *StopCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopCdnDomainRequest) SetOwnerId(v int64) *StopCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopCdnDomainRequest) SetSecurityToken(v string) *StopCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
