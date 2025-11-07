// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryControlStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *QueryControlStrategyRequest
	GetProductType() *string
	SetRegionId(v string) *QueryControlStrategyRequest
	GetRegionId() *string
}

type QueryControlStrategyRequest struct {
	// Product type, currently only supports ANT_CLOUD_AUTH (financial-grade real person), all others have been phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryControlStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryControlStrategyRequest) GoString() string {
	return s.String()
}

func (s *QueryControlStrategyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryControlStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryControlStrategyRequest) SetProductType(v string) *QueryControlStrategyRequest {
	s.ProductType = &v
	return s
}

func (s *QueryControlStrategyRequest) SetRegionId(v string) *QueryControlStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *QueryControlStrategyRequest) Validate() error {
	return dara.Validate(s)
}
