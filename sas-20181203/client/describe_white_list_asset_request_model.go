// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeWhiteListAssetRequest
	GetLang() *string
	SetLastMaxId(v int64) *DescribeWhiteListAssetRequest
	GetLastMaxId() *int64
	SetPageSize(v int32) *DescribeWhiteListAssetRequest
	GetPageSize() *int32
	SetSourceIp(v string) *DescribeWhiteListAssetRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *DescribeWhiteListAssetRequest
	GetStrategyId() *int64
	SetType(v int32) *DescribeWhiteListAssetRequest
	GetType() *int32
}

type DescribeWhiteListAssetRequest struct {
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
	// The maximum asset ID of the most recent query.
	//
	// example:
	//
	// 1001
	LastMaxId *int64 `json:"LastMaxId,omitempty" xml:"LastMaxId,omitempty"`
	// The number of entries per page. Maximum value: **500**. Default value: **500**. This value indicates that 500 entries are displayed on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 180.119.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain the ID.
	//
	// example:
	//
	// 2730
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy type of the asset that you want to query. Valid values:
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

func (s DescribeWhiteListAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAssetRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListAssetRequest) GetLastMaxId() *int64 {
	return s.LastMaxId
}

func (s *DescribeWhiteListAssetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListAssetRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListAssetRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListAssetRequest) GetType() *int32 {
	return s.Type
}

func (s *DescribeWhiteListAssetRequest) SetLang(v string) *DescribeWhiteListAssetRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListAssetRequest) SetLastMaxId(v int64) *DescribeWhiteListAssetRequest {
	s.LastMaxId = &v
	return s
}

func (s *DescribeWhiteListAssetRequest) SetPageSize(v int32) *DescribeWhiteListAssetRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListAssetRequest) SetSourceIp(v string) *DescribeWhiteListAssetRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListAssetRequest) SetStrategyId(v int64) *DescribeWhiteListAssetRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListAssetRequest) SetType(v int32) *DescribeWhiteListAssetRequest {
	s.Type = &v
	return s
}

func (s *DescribeWhiteListAssetRequest) Validate() error {
	return dara.Validate(s)
}
