// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyUuidCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeWhiteListStrategyUuidCountRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeWhiteListStrategyUuidCountRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *DescribeWhiteListStrategyUuidCountRequest
	GetStrategyId() *int64
	SetType(v int32) *DescribeWhiteListStrategyUuidCountRequest
	GetType() *int32
}

type DescribeWhiteListStrategyUuidCountRequest struct {
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
	// 42.120.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyStatistics](~~DescribeWhiteListStrategyStatistics~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8516
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **1**: learning policy
	//
	// 	- **2**: application policy
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeWhiteListStrategyUuidCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyUuidCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyUuidCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListStrategyUuidCountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListStrategyUuidCountRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListStrategyUuidCountRequest) GetType() *int32 {
	return s.Type
}

func (s *DescribeWhiteListStrategyUuidCountRequest) SetLang(v string) *DescribeWhiteListStrategyUuidCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountRequest) SetSourceIp(v string) *DescribeWhiteListStrategyUuidCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountRequest) SetStrategyId(v int64) *DescribeWhiteListStrategyUuidCountRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountRequest) SetType(v int32) *DescribeWhiteListStrategyUuidCountRequest {
	s.Type = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountRequest) Validate() error {
	return dara.Validate(s)
}
