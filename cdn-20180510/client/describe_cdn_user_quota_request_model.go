// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeCdnUserQuotaRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnUserQuotaRequest
	GetSecurityToken() *string
}

type DescribeCdnUserQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnUserQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnUserQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnUserQuotaRequest) SetOwnerId(v int64) *DescribeCdnUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserQuotaRequest) SetSecurityToken(v string) *DescribeCdnUserQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnUserQuotaRequest) Validate() error {
	return dara.Validate(s)
}
