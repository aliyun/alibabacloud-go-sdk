// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListEffectiveAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWhiteListEffectiveAssetsRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeWhiteListEffectiveAssetsRequest
	GetLang() *string
	SetNeedStatistics(v int32) *DescribeWhiteListEffectiveAssetsRequest
	GetNeedStatistics() *int32
	SetPageSize(v int32) *DescribeWhiteListEffectiveAssetsRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeWhiteListEffectiveAssetsRequest
	GetRemark() *string
	SetSourceIp(v string) *DescribeWhiteListEffectiveAssetsRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *DescribeWhiteListEffectiveAssetsRequest
	GetStrategyId() *int64
}

type DescribeWhiteListEffectiveAssetsRequest struct {
	// The page number. Default value: **1**. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to return the number of **untrusted programs**. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	NeedStatistics *int32 `json:"NeedStatistics,omitempty" xml:"NeedStatistics,omitempty"`
	// The number of entries per page. Maximum value: **1000**. Default value: 20. If you leave this parameter empty, 20 data entries are returned per page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The condition that is used to query assets. You can enter an IP address, a public IP address, an private IP address, or an asset name for fuzzy match.
	//
	// example:
	//
	// 222.185.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 27.212.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain the ID.
	//
	// example:
	//
	// 8437592
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeWhiteListEffectiveAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListEffectiveAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetNeedStatistics() *int32 {
	return s.NeedStatistics
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListEffectiveAssetsRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetCurrentPage(v int32) *DescribeWhiteListEffectiveAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetLang(v string) *DescribeWhiteListEffectiveAssetsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetNeedStatistics(v int32) *DescribeWhiteListEffectiveAssetsRequest {
	s.NeedStatistics = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetPageSize(v int32) *DescribeWhiteListEffectiveAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetRemark(v string) *DescribeWhiteListEffectiveAssetsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetSourceIp(v string) *DescribeWhiteListEffectiveAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) SetStrategyId(v int64) *DescribeWhiteListEffectiveAssetsRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsRequest) Validate() error {
	return dara.Validate(s)
}
