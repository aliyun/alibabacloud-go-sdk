// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardLogStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *GetGuardLogStatsRequest
	GetCommodityCode() *string
}

type GetGuardLogStatsRequest struct {
	// The commodity code.
	//
	// example:
	//
	// lvwang_guardrail_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s GetGuardLogStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGuardLogStatsRequest) GoString() string {
	return s.String()
}

func (s *GetGuardLogStatsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetGuardLogStatsRequest) SetCommodityCode(v string) *GetGuardLogStatsRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetGuardLogStatsRequest) Validate() error {
	return dara.Validate(s)
}
