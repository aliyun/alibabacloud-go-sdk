// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommercialStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *GetCommercialStatusRequest
	GetCommodityCode() *string
	SetRegionId(v string) *GetCommercialStatusRequest
	GetRegionId() *string
}

type GetCommercialStatusRequest struct {
	// The product code.
	//
	// 	- arms_app_post
	//
	// 	- arms_web_post
	//
	// 	- arms_promethues_public_cn
	//
	// 	- prometheus_pay_public_cn
	//
	// 	- xtrace
	//
	// 	- arms_serverless_public_cn
	//
	// 	- arms_rumserverless_public_cn
	//
	// 	- prometheus_serverless_public_cn
	//
	// 	- xtrace_serverless_public_cn
	//
	// This parameter is required.
	//
	// example:
	//
	// arms_app_post
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCommercialStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommercialStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCommercialStatusRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetCommercialStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCommercialStatusRequest) SetCommodityCode(v string) *GetCommercialStatusRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetCommercialStatusRequest) SetRegionId(v string) *GetCommercialStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetCommercialStatusRequest) Validate() error {
	return dara.Validate(s)
}
