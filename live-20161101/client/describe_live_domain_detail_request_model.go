// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainDetailRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveDomainDetailRequest
	GetSecurityToken() *string
}

type DescribeLiveDomainDetailRequest struct {
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

func (s DescribeLiveDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveDomainDetailRequest) SetDomainName(v string) *DescribeLiveDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainDetailRequest) SetOwnerId(v int64) *DescribeLiveDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainDetailRequest) SetSecurityToken(v string) *DescribeLiveDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
