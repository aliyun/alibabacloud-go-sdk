// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeHttpDDoSAttackRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHttpDDoSAttackRulesRequest
	GetPageSize() *int32
	SetSiteId(v int64) *DescribeHttpDDoSAttackRulesRequest
	GetSiteId() *int64
}

type DescribeHttpDDoSAttackRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHttpDDoSAttackRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHttpDDoSAttackRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSAttackRulesRequest) SetPageNumber(v int32) *DescribeHttpDDoSAttackRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesRequest) SetPageSize(v int32) *DescribeHttpDDoSAttackRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesRequest) SetSiteId(v int64) *DescribeHttpDDoSAttackRulesRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesRequest) Validate() error {
	return dara.Validate(s)
}
