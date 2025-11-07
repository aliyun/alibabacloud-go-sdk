// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeFlowStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DeleteCustomizeFlowStrategyRequest
	GetApiName() *string
	SetId(v int64) *DeleteCustomizeFlowStrategyRequest
	GetId() *int64
	SetProductType(v string) *DeleteCustomizeFlowStrategyRequest
	GetProductType() *string
	SetRegionId(v string) *DeleteCustomizeFlowStrategyRequest
	GetRegionId() *string
	SetUserId(v int64) *DeleteCustomizeFlowStrategyRequest
	GetUserId() *int64
}

type DeleteCustomizeFlowStrategyRequest struct {
	// API name, same as **ProductCode**.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Policy ID.
	//
	// This parameter is required.
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
	// regionId
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// User ID
	//
	// example:
	//
	// 126005125163xxxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteCustomizeFlowStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeFlowStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeFlowStrategyRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DeleteCustomizeFlowStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCustomizeFlowStrategyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DeleteCustomizeFlowStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomizeFlowStrategyRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DeleteCustomizeFlowStrategyRequest) SetApiName(v string) *DeleteCustomizeFlowStrategyRequest {
	s.ApiName = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyRequest) SetId(v int64) *DeleteCustomizeFlowStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyRequest) SetProductType(v string) *DeleteCustomizeFlowStrategyRequest {
	s.ProductType = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyRequest) SetRegionId(v string) *DeleteCustomizeFlowStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyRequest) SetUserId(v int64) *DeleteCustomizeFlowStrategyRequest {
	s.UserId = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyRequest) Validate() error {
	return dara.Validate(s)
}
