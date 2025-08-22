// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnRefreshQuotaRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnRefreshQuotaRequest
	GetSecurityToken() *string
}

type DescribeDcdnRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnRefreshQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnRefreshQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnRefreshQuotaRequest) SetOwnerId(v int64) *DescribeDcdnRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaRequest) SetSecurityToken(v string) *DescribeDcdnRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaRequest) Validate() error {
	return dara.Validate(s)
}
