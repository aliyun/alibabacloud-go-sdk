// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomType(v string) *DescribeStrategyRequest
	GetCustomType() *string
	SetLang(v string) *DescribeStrategyRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeStrategyRequest
	GetSourceIp() *string
	SetStrategyIds(v string) *DescribeStrategyRequest
	GetStrategyIds() *string
}

type DescribeStrategyRequest struct {
	// The type of the baseline check policy that you want to query. Valid values:
	//
	// 	- **common**: standard baseline check policy
	//
	// 	- **custom**: custom baseline check policy
	//
	// example:
	//
	// custom
	CustomType *string `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
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
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the baseline check policy that you want to query. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 8164248
	StrategyIds *string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty"`
}

func (s DescribeStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeStrategyRequest) GetCustomType() *string {
	return s.CustomType
}

func (s *DescribeStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeStrategyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeStrategyRequest) GetStrategyIds() *string {
	return s.StrategyIds
}

func (s *DescribeStrategyRequest) SetCustomType(v string) *DescribeStrategyRequest {
	s.CustomType = &v
	return s
}

func (s *DescribeStrategyRequest) SetLang(v string) *DescribeStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStrategyRequest) SetSourceIp(v string) *DescribeStrategyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStrategyRequest) SetStrategyIds(v string) *DescribeStrategyRequest {
	s.StrategyIds = &v
	return s
}

func (s *DescribeStrategyRequest) Validate() error {
	return dara.Validate(s)
}
