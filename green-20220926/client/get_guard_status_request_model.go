// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *GetGuardStatusRequest
	GetCommodityCode() *string
	SetRegionId(v string) *GetGuardStatusRequest
	GetRegionId() *string
}

type GetGuardStatusRequest struct {
	// The commodity code.
	//
	// example:
	//
	// xxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetGuardStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGuardStatusRequest) GoString() string {
	return s.String()
}

func (s *GetGuardStatusRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetGuardStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGuardStatusRequest) SetCommodityCode(v string) *GetGuardStatusRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetGuardStatusRequest) SetRegionId(v string) *GetGuardStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetGuardStatusRequest) Validate() error {
	return dara.Validate(s)
}
