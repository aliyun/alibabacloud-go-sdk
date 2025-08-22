// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnWafServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnWafServiceRequest
	GetSecurityToken() *string
}

type DescribeDcdnWafServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnWafServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnWafServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnWafServiceRequest) SetOwnerId(v int64) *DescribeDcdnWafServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnWafServiceRequest) SetSecurityToken(v string) *DescribeDcdnWafServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnWafServiceRequest) Validate() error {
	return dara.Validate(s)
}
