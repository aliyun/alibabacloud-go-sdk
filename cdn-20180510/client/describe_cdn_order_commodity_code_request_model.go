// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnOrderCommodityCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeCdnOrderCommodityCodeRequest
	GetCommodityCode() *string
	SetOwnerId(v int64) *DescribeCdnOrderCommodityCodeRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnOrderCommodityCodeRequest
	GetSecurityToken() *string
}

type DescribeCdnOrderCommodityCodeRequest struct {
	// The original commodity code.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnOrderCommodityCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnOrderCommodityCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnOrderCommodityCodeRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCdnOrderCommodityCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnOrderCommodityCodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnOrderCommodityCodeRequest) SetCommodityCode(v string) *DescribeCdnOrderCommodityCodeRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCdnOrderCommodityCodeRequest) SetOwnerId(v int64) *DescribeCdnOrderCommodityCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnOrderCommodityCodeRequest) SetSecurityToken(v string) *DescribeCdnOrderCommodityCodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnOrderCommodityCodeRequest) Validate() error {
	return dara.Validate(s)
}
