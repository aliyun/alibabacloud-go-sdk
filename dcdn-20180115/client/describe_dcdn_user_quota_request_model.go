// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnUserQuotaRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnUserQuotaRequest
	GetSecurityToken() *string
}

type DescribeDcdnUserQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnUserQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnUserQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnUserQuotaRequest) SetOwnerId(v int64) *DescribeDcdnUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserQuotaRequest) SetSecurityToken(v string) *DescribeDcdnUserQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnUserQuotaRequest) Validate() error {
	return dara.Validate(s)
}
