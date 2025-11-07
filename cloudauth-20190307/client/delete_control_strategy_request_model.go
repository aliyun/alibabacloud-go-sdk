// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DeleteControlStrategyRequest
	GetApiName() *string
	SetId(v int64) *DeleteControlStrategyRequest
	GetId() *int64
	SetProductType(v string) *DeleteControlStrategyRequest
	GetProductType() *string
	SetRegionId(v string) *DeleteControlStrategyRequest
	GetRegionId() *string
}

type DeleteControlStrategyRequest struct {
	// API name, same as **ProductCode**.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Security rule ID.
	//
	// example:
	//
	// 38
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Product type, currently only supports **ANT_CLOUD_AUTH*	- (Financial-grade Real Person), all others have been phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteControlStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlStrategyRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DeleteControlStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteControlStrategyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DeleteControlStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteControlStrategyRequest) SetApiName(v string) *DeleteControlStrategyRequest {
	s.ApiName = &v
	return s
}

func (s *DeleteControlStrategyRequest) SetId(v int64) *DeleteControlStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteControlStrategyRequest) SetProductType(v string) *DeleteControlStrategyRequest {
	s.ProductType = &v
	return s
}

func (s *DeleteControlStrategyRequest) SetRegionId(v string) *DeleteControlStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteControlStrategyRequest) Validate() error {
	return dara.Validate(s)
}
