// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeWhiteListStrategyListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeWhiteListStrategyListRequest
	GetSourceIp() *string
	SetStrategyIds(v string) *DescribeWhiteListStrategyListRequest
	GetStrategyIds() *string
}

type DescribeWhiteListStrategyListRequest struct {
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
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 116.88.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 1,2
	StrategyIds *string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty"`
}

func (s DescribeWhiteListStrategyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListStrategyListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListStrategyListRequest) GetStrategyIds() *string {
	return s.StrategyIds
}

func (s *DescribeWhiteListStrategyListRequest) SetLang(v string) *DescribeWhiteListStrategyListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListStrategyListRequest) SetSourceIp(v string) *DescribeWhiteListStrategyListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListStrategyListRequest) SetStrategyIds(v string) *DescribeWhiteListStrategyListRequest {
	s.StrategyIds = &v
	return s
}

func (s *DescribeWhiteListStrategyListRequest) Validate() error {
	return dara.Validate(s)
}
