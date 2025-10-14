// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarxDataNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeType(v string) *DescribePolarxDataNodesRequest
	GetNodeType() *string
	SetPageNumber(v int32) *DescribePolarxDataNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarxDataNodesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePolarxDataNodesRequest
	GetRegionId() *string
	SetSearchKey(v string) *DescribePolarxDataNodesRequest
	GetSearchKey() *string
}

type DescribePolarxDataNodesRequest struct {
	// example:
	//
	// dn
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// pc-bp1c5w4fx****2274
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s DescribePolarxDataNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarxDataNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribePolarxDataNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarxDataNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarxDataNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarxDataNodesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribePolarxDataNodesRequest) SetNodeType(v string) *DescribePolarxDataNodesRequest {
	s.NodeType = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) SetPageNumber(v int32) *DescribePolarxDataNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) SetPageSize(v int32) *DescribePolarxDataNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) SetRegionId(v string) *DescribePolarxDataNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) SetSearchKey(v string) *DescribePolarxDataNodesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) Validate() error {
	return dara.Validate(s)
}
