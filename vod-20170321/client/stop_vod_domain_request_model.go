// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StopVodDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StopVodDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StopVodDomainRequest
	GetSecurityToken() *string
}

type StopVodDomainRequest struct {
	// This parameter is required.
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StopVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StopVodDomainRequest) GoString() string {
	return s.String()
}

func (s *StopVodDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StopVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StopVodDomainRequest) SetDomainName(v string) *StopVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopVodDomainRequest) SetOwnerId(v int64) *StopVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopVodDomainRequest) SetSecurityToken(v string) *StopVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
