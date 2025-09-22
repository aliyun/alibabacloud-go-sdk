// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventoryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInventoryListRequest
	GetCode() *string
	SetEmissionType(v string) *GetInventoryListRequest
	GetEmissionType() *string
	SetGroup(v string) *GetInventoryListRequest
	GetGroup() *string
	SetMethodType(v string) *GetInventoryListRequest
	GetMethodType() *string
	SetProductId(v int64) *GetInventoryListRequest
	GetProductId() *int64
	SetProductType(v int64) *GetInventoryListRequest
	GetProductType() *int64
}

type GetInventoryListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Type of emission
	//
	// >  Valid values: footprint | emission. Meaning: footprint: all inventories are involved in the calculation; emission: only inventories with positive and zero emissions are involved in the calculation, and negative numbers are not involved in the calculation.
	//
	// This parameter is required.
	//
	// example:
	//
	// footprint
	EmissionType *string `json:"emissionType,omitempty" xml:"emissionType,omitempty"`
	// Group by
	//
	// >  Valid values: resource | process | resourceType | processType. Meaning: resource: aggregation by inventory group, process: aggregation by operation group, resourceType: aggregation by inventory type, processType: aggregation by phase group
	//
	// This parameter is required.
	//
	// example:
	//
	// resource
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The type of the obtained environmental impact: gwp indicates the carbon footprint of climate change.
	//
	// <props="intl">[For more information, see the environment impact category enumeration.](https://www.alibabacloud.com/help/en/energy-expert/developer-reference/enumerated-values-of-energy-expert#RhGn7)
	//
	// This parameter is required.
	//
	// example:
	//
	// gwp
	MethodType *string `json:"methodType,omitempty" xml:"methodType,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetInventoryListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInventoryListRequest) GoString() string {
	return s.String()
}

func (s *GetInventoryListRequest) GetCode() *string {
	return s.Code
}

func (s *GetInventoryListRequest) GetEmissionType() *string {
	return s.EmissionType
}

func (s *GetInventoryListRequest) GetGroup() *string {
	return s.Group
}

func (s *GetInventoryListRequest) GetMethodType() *string {
	return s.MethodType
}

func (s *GetInventoryListRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetInventoryListRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetInventoryListRequest) SetCode(v string) *GetInventoryListRequest {
	s.Code = &v
	return s
}

func (s *GetInventoryListRequest) SetEmissionType(v string) *GetInventoryListRequest {
	s.EmissionType = &v
	return s
}

func (s *GetInventoryListRequest) SetGroup(v string) *GetInventoryListRequest {
	s.Group = &v
	return s
}

func (s *GetInventoryListRequest) SetMethodType(v string) *GetInventoryListRequest {
	s.MethodType = &v
	return s
}

func (s *GetInventoryListRequest) SetProductId(v int64) *GetInventoryListRequest {
	s.ProductId = &v
	return s
}

func (s *GetInventoryListRequest) SetProductType(v int64) *GetInventoryListRequest {
	s.ProductType = &v
	return s
}

func (s *GetInventoryListRequest) Validate() error {
	return dara.Validate(s)
}
