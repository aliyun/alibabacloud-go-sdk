// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyExecDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeStrategyExecDetailRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeStrategyExecDetailRequest
	GetPageSize() *int32
	SetSourceIp(v string) *DescribeStrategyExecDetailRequest
	GetSourceIp() *string
	SetStrategyId(v int32) *DescribeStrategyExecDetailRequest
	GetStrategyId() *int32
}

type DescribeStrategyExecDetailRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 10. If you leave this parameter empty, 10 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the baseline check policy.
	//
	// >  You can call the [DescribeStrategy](~~DescribeStrategy~~) operation to query the IDs of baseline check policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8437592
	StrategyId *int32 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeStrategyExecDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyExecDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeStrategyExecDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStrategyExecDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeStrategyExecDetailRequest) GetStrategyId() *int32 {
	return s.StrategyId
}

func (s *DescribeStrategyExecDetailRequest) SetCurrentPage(v int32) *DescribeStrategyExecDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeStrategyExecDetailRequest) SetPageSize(v int32) *DescribeStrategyExecDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStrategyExecDetailRequest) SetSourceIp(v string) *DescribeStrategyExecDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStrategyExecDetailRequest) SetStrategyId(v int32) *DescribeStrategyExecDetailRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeStrategyExecDetailRequest) Validate() error {
	return dara.Validate(s)
}
