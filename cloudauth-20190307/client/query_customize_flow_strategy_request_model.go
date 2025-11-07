// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomizeFlowStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *QueryCustomizeFlowStrategyRequest
	GetProductType() *string
	SetRegionId(v string) *QueryCustomizeFlowStrategyRequest
	GetRegionId() *string
	SetUserId(v int64) *QueryCustomizeFlowStrategyRequest
	GetUserId() *int64
}

type QueryCustomizeFlowStrategyRequest struct {
	// Product type, currently only supports **ANT_CLOUD_AUTH*	- (Financial-grade Real Person), all others have been phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// regionId
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// User ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 126005125163xxxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryCustomizeFlowStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomizeFlowStrategyRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomizeFlowStrategyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryCustomizeFlowStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCustomizeFlowStrategyRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryCustomizeFlowStrategyRequest) SetProductType(v string) *QueryCustomizeFlowStrategyRequest {
	s.ProductType = &v
	return s
}

func (s *QueryCustomizeFlowStrategyRequest) SetRegionId(v string) *QueryCustomizeFlowStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCustomizeFlowStrategyRequest) SetUserId(v int64) *QueryCustomizeFlowStrategyRequest {
	s.UserId = &v
	return s
}

func (s *QueryCustomizeFlowStrategyRequest) Validate() error {
	return dara.Validate(s)
}
