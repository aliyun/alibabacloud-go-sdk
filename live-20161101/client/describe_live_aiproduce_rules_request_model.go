// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAIProduceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeLiveAIProduceRulesRequest
	GetApp() *string
	SetDomain(v string) *DescribeLiveAIProduceRulesRequest
	GetDomain() *string
	SetOwnerId(v int64) *DescribeLiveAIProduceRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeLiveAIProduceRulesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeLiveAIProduceRulesRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeLiveAIProduceRulesRequest
	GetRegionId() *string
	SetRulesId(v string) *DescribeLiveAIProduceRulesRequest
	GetRulesId() *string
	SetSuffixName(v string) *DescribeLiveAIProduceRulesRequest
	GetSuffixName() *string
}

type DescribeLiveAIProduceRulesRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// AppName
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: [1,100].
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: [1,100].
	//
	// example:
	//
	// 100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the subtitle rule.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec****
	RulesId *string `json:"RulesId,omitempty" xml:"RulesId,omitempty"`
	// The suffix of the subtitle rule.
	//
	// > Set the value to the name of the subtitle template.
	//
	// example:
	//
	// sub01
	SuffixName *string `json:"SuffixName,omitempty" xml:"SuffixName,omitempty"`
}

func (s DescribeLiveAIProduceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIProduceRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIProduceRulesRequest) GetApp() *string {
	return s.App
}

func (s *DescribeLiveAIProduceRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveAIProduceRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveAIProduceRulesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeLiveAIProduceRulesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLiveAIProduceRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveAIProduceRulesRequest) GetRulesId() *string {
	return s.RulesId
}

func (s *DescribeLiveAIProduceRulesRequest) GetSuffixName() *string {
	return s.SuffixName
}

func (s *DescribeLiveAIProduceRulesRequest) SetApp(v string) *DescribeLiveAIProduceRulesRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetDomain(v string) *DescribeLiveAIProduceRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetOwnerId(v int64) *DescribeLiveAIProduceRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetPageNumber(v string) *DescribeLiveAIProduceRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetPageSize(v string) *DescribeLiveAIProduceRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetRegionId(v string) *DescribeLiveAIProduceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetRulesId(v string) *DescribeLiveAIProduceRulesRequest {
	s.RulesId = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) SetSuffixName(v string) *DescribeLiveAIProduceRulesRequest {
	s.SuffixName = &v
	return s
}

func (s *DescribeLiveAIProduceRulesRequest) Validate() error {
	return dara.Validate(s)
}
