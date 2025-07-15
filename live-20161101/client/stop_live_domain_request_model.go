// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StopLiveDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StopLiveDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StopLiveDomainRequest
	GetSecurityToken() *string
}

type StopLiveDomainRequest struct {
	// The streaming domain or ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StopLiveDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *StopLiveDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StopLiveDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopLiveDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StopLiveDomainRequest) SetDomainName(v string) *StopLiveDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopLiveDomainRequest) SetOwnerId(v int64) *StopLiveDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopLiveDomainRequest) SetSecurityToken(v string) *StopLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopLiveDomainRequest) Validate() error {
	return dara.Validate(s)
}
