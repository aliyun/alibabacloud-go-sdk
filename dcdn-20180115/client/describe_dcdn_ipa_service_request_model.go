// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnIpaServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnIpaServiceRequest
	GetSecurityToken() *string
}

type DescribeDcdnIpaServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnIpaServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnIpaServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnIpaServiceRequest) SetOwnerId(v int64) *DescribeDcdnIpaServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaServiceRequest) SetSecurityToken(v string) *DescribeDcdnIpaServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnIpaServiceRequest) Validate() error {
	return dara.Validate(s)
}
