// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarkRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeLiveStreamWatermarkRulesRequest
	GetDomain() *string
	SetOwnerId(v int64) *DescribeLiveStreamWatermarkRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveStreamWatermarkRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveStreamWatermarkRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLiveStreamWatermarkRulesRequest
	GetRegionId() *string
}

type DescribeLiveStreamWatermarkRulesRequest struct {
	// The main streaming domain.
	//
	// example:
	//
	// live.yourdomain.com
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. If you do not specify this parameter, the default value 1 is used.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. If you do not specify this parameter, the default value 100 is used.
	//
	// example:
	//
	// 100
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveStreamWatermarkRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveStreamWatermarkRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamWatermarkRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveStreamWatermarkRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamWatermarkRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamWatermarkRulesRequest) SetDomain(v string) *DescribeLiveStreamWatermarkRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesRequest) SetOwnerId(v int64) *DescribeLiveStreamWatermarkRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesRequest) SetPageNumber(v int32) *DescribeLiveStreamWatermarkRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesRequest) SetPageSize(v int32) *DescribeLiveStreamWatermarkRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesRequest) SetRegionId(v string) *DescribeLiveStreamWatermarkRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesRequest) Validate() error {
	return dara.Validate(s)
}
