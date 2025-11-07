// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackListStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteBlackListStrategyRequest
	GetId() *int64
	SetProductName(v string) *DeleteBlackListStrategyRequest
	GetProductName() *string
	SetRegionId(v string) *DeleteBlackListStrategyRequest
	GetRegionId() *string
}

type DeleteBlackListStrategyRequest struct {
	// Rule ID.
	//
	// example:
	//
	// 38
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Product Name:
	//
	// - **id2meta**: ID card two-factor verification
	//
	// - **mobile3Meta**: Mobile phone number factor verification
	//
	// - **bankcardMeta**: Bank card factor verification
	//
	// example:
	//
	// id2meta
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBlackListStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackListStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackListStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteBlackListStrategyRequest) GetProductName() *string {
	return s.ProductName
}

func (s *DeleteBlackListStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBlackListStrategyRequest) SetId(v int64) *DeleteBlackListStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteBlackListStrategyRequest) SetProductName(v string) *DeleteBlackListStrategyRequest {
	s.ProductName = &v
	return s
}

func (s *DeleteBlackListStrategyRequest) SetRegionId(v string) *DeleteBlackListStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBlackListStrategyRequest) Validate() error {
	return dara.Validate(s)
}
