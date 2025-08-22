// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDcdnIpaDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StopDcdnIpaDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StopDcdnIpaDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StopDcdnIpaDomainRequest
	GetSecurityToken() *string
}

type StopDcdnIpaDomainRequest struct {
	// The name of the accelerated domain that you want to disable. You can specify only one domain name at a time.
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

func (s StopDcdnIpaDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *StopDcdnIpaDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StopDcdnIpaDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopDcdnIpaDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StopDcdnIpaDomainRequest) SetDomainName(v string) *StopDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopDcdnIpaDomainRequest) SetOwnerId(v int64) *StopDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDcdnIpaDomainRequest) SetSecurityToken(v string) *StopDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopDcdnIpaDomainRequest) Validate() error {
	return dara.Validate(s)
}
