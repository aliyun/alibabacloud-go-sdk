// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillsMarketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *DescribePolarClawSkillsMarketRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribePolarClawSkillsMarketRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarClawSkillsMarketRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePolarClawSkillsMarketRequest
	GetRegionId() *string
	SetSearchBySkillName(v string) *DescribePolarClawSkillsMarketRequest
	GetSearchBySkillName() *string
	SetSearchByTagName(v string) *DescribePolarClawSkillsMarketRequest
	GetSearchByTagName() *string
	SetSource(v string) *DescribePolarClawSkillsMarketRequest
	GetSource() *string
}

type DescribePolarClawSkillsMarketRequest struct {
	// The sorting condition. DOWNLOAD_COUNT indicates sorting by download count.
	//
	// example:
	//
	// DOWNLOAD_COUNT
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all available regions for the current account, including region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The skill name to search by.
	//
	// example:
	//
	// polarclawhub
	SearchBySkillName *string `json:"SearchBySkillName,omitempty" xml:"SearchBySkillName,omitempty"`
	// The tag name to search by.
	//
	// example:
	//
	// PolarDB
	SearchByTagName *string `json:"SearchByTagName,omitempty" xml:"SearchByTagName,omitempty"`
	// The source of the skill. SELF indicates private skills of the current user, and PUBLIC indicates public skills.
	//
	// example:
	//
	// PUBLIC
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarClawSkillsMarketRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsMarketRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsMarketRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribePolarClawSkillsMarketRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarClawSkillsMarketRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarClawSkillsMarketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarClawSkillsMarketRequest) GetSearchBySkillName() *string {
	return s.SearchBySkillName
}

func (s *DescribePolarClawSkillsMarketRequest) GetSearchByTagName() *string {
	return s.SearchByTagName
}

func (s *DescribePolarClawSkillsMarketRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawSkillsMarketRequest) SetOrderBy(v string) *DescribePolarClawSkillsMarketRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) SetPageNumber(v int32) *DescribePolarClawSkillsMarketRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) SetPageSize(v int32) *DescribePolarClawSkillsMarketRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) SetRegionId(v string) *DescribePolarClawSkillsMarketRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) SetSearchBySkillName(v string) *DescribePolarClawSkillsMarketRequest {
	s.SearchBySkillName = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) SetSearchByTagName(v string) *DescribePolarClawSkillsMarketRequest {
	s.SearchByTagName = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) SetSource(v string) *DescribePolarClawSkillsMarketRequest {
	s.Source = &v
	return s
}

func (s *DescribePolarClawSkillsMarketRequest) Validate() error {
	return dara.Validate(s)
}
