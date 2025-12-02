// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBuyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *GetUserBuyStatusRequest
	GetCommodityCode() *string
	SetRegionId(v string) *GetUserBuyStatusRequest
	GetRegionId() *string
}

type GetUserBuyStatusRequest struct {
	// Commodity code.
	//
	// example:
	//
	// lvwang_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUserBuyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserBuyStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetUserBuyStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserBuyStatusRequest) SetCommodityCode(v string) *GetUserBuyStatusRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUserBuyStatusRequest) SetRegionId(v string) *GetUserBuyStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserBuyStatusRequest) Validate() error {
	return dara.Validate(s)
}
