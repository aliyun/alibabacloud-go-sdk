// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTs(v int64) *DescribeMeterImsSummaryRequest
	GetEndTs() *int64
	SetRegion(v string) *DescribeMeterImsSummaryRequest
	GetRegion() *string
	SetStartTs(v int64) *DescribeMeterImsSummaryRequest
	GetStartTs() *int64
}

type DescribeMeterImsSummaryRequest struct {
	// The end of the time range to query. The value is a 10-digit timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656995036
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// This parameter does not take effect. By default, the usage data of all regions is returned.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. The value is a 10-digit timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1654403036
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterImsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsSummaryRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeMeterImsSummaryRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeMeterImsSummaryRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeMeterImsSummaryRequest) SetEndTs(v int64) *DescribeMeterImsSummaryRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterImsSummaryRequest) SetRegion(v string) *DescribeMeterImsSummaryRequest {
	s.Region = &v
	return s
}

func (s *DescribeMeterImsSummaryRequest) SetStartTs(v int64) *DescribeMeterImsSummaryRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeMeterImsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
