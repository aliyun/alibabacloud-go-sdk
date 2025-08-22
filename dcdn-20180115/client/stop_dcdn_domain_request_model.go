// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StopDcdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StopDcdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StopDcdnDomainRequest
	GetSecurityToken() *string
}

type StopDcdnDomainRequest struct {
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

func (s StopDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StopDcdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StopDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StopDcdnDomainRequest) SetDomainName(v string) *StopDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopDcdnDomainRequest) SetOwnerId(v int64) *StopDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDcdnDomainRequest) SetSecurityToken(v string) *StopDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
