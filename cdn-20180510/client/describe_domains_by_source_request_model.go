// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsBySourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDomainsBySourceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDomainsBySourceRequest
	GetSecurityToken() *string
	SetSources(v string) *DescribeDomainsBySourceRequest
	GetSources() *string
}

type DescribeDomainsBySourceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The origin servers. Separate multiple origin servers with commas (,). Fuzzy match is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeDomainsBySourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDomainsBySourceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDomainsBySourceRequest) GetSources() *string {
	return s.Sources
}

func (s *DescribeDomainsBySourceRequest) SetOwnerId(v int64) *DescribeDomainsBySourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) SetSecurityToken(v string) *DescribeDomainsBySourceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) SetSources(v string) *DescribeDomainsBySourceRequest {
	s.Sources = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) Validate() error {
	return dara.Validate(s)
}
