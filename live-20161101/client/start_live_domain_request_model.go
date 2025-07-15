// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *StartLiveDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *StartLiveDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *StartLiveDomainRequest
	GetSecurityToken() *string
}

type StartLiveDomainRequest struct {
	// The ingest domain or streaming domain.
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

func (s StartLiveDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *StartLiveDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *StartLiveDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartLiveDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartLiveDomainRequest) SetDomainName(v string) *StartLiveDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartLiveDomainRequest) SetOwnerId(v int64) *StartLiveDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartLiveDomainRequest) SetSecurityToken(v string) *StartLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartLiveDomainRequest) Validate() error {
	return dara.Validate(s)
}
