// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StartVodDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StartVodDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StartVodDomainRequest
	GetSecurityToken() *string
}

type StartVodDomainRequest struct {
	// This parameter is required.
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StartVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StartVodDomainRequest) GoString() string {
	return s.String()
}

func (s *StartVodDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StartVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartVodDomainRequest) SetDomainName(v string) *StartVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartVodDomainRequest) SetOwnerId(v int64) *StartVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartVodDomainRequest) SetSecurityToken(v string) *StartVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
