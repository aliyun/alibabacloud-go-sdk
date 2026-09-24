// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatedDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationFactor(v string) *GetAggregatedDesktopsRequest
	GetAggregationFactor() *string
	SetRegionId(v string) *GetAggregatedDesktopsRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *GetAggregatedDesktopsRequest
	GetSearchRegionId() *string
}

type GetAggregatedDesktopsRequest struct {
	// The aggregation factor.
	//
	// example:
	//
	// STATUS
	AggregationFactor *string `json:"AggregationFactor,omitempty" xml:"AggregationFactor,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the list of regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The search region ID. Used to filter desktop information for a specified region.
	//
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
}

func (s GetAggregatedDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatedDesktopsRequest) GoString() string {
	return s.String()
}

func (s *GetAggregatedDesktopsRequest) GetAggregationFactor() *string {
	return s.AggregationFactor
}

func (s *GetAggregatedDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAggregatedDesktopsRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *GetAggregatedDesktopsRequest) SetAggregationFactor(v string) *GetAggregatedDesktopsRequest {
	s.AggregationFactor = &v
	return s
}

func (s *GetAggregatedDesktopsRequest) SetRegionId(v string) *GetAggregatedDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAggregatedDesktopsRequest) SetSearchRegionId(v string) *GetAggregatedDesktopsRequest {
	s.SearchRegionId = &v
	return s
}

func (s *GetAggregatedDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
