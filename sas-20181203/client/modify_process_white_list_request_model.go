// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProcessWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyProcessWhiteListRequest
	GetLang() *string
	SetMd5s(v string) *ModifyProcessWhiteListRequest
	GetMd5s() *string
	SetSourceIp(v string) *ModifyProcessWhiteListRequest
	GetSourceIp() *string
	SetStatus(v int32) *ModifyProcessWhiteListRequest
	GetStatus() *int32
	SetStrategyId(v int64) *ModifyProcessWhiteListRequest
	GetStrategyId() *int64
}

type ModifyProcessWhiteListRequest struct {
	// Specifies the language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The MD5 hash value of the process startup file.
	//
	// > You can call the [DescribeWhiteListProcess](~~DescribeWhiteListProcess~~) operation to query the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 001d7f68c3b44147988f0dc81192****
	Md5s *string `json:"Md5s,omitempty" xml:"Md5s,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. The system automatically retrieves the value.
	//
	// example:
	//
	// 173.128.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Specifies the whitelist status of the process. Valid values:
	//
	// - **1**: remove from the whitelist
	//
	// - **2**: add to the whitelist
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the policy.
	//
	// > You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ModifyProcessWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProcessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyProcessWhiteListRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyProcessWhiteListRequest) GetMd5s() *string {
	return s.Md5s
}

func (s *ModifyProcessWhiteListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyProcessWhiteListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyProcessWhiteListRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *ModifyProcessWhiteListRequest) SetLang(v string) *ModifyProcessWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *ModifyProcessWhiteListRequest) SetMd5s(v string) *ModifyProcessWhiteListRequest {
	s.Md5s = &v
	return s
}

func (s *ModifyProcessWhiteListRequest) SetSourceIp(v string) *ModifyProcessWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyProcessWhiteListRequest) SetStatus(v int32) *ModifyProcessWhiteListRequest {
	s.Status = &v
	return s
}

func (s *ModifyProcessWhiteListRequest) SetStrategyId(v int64) *ModifyProcessWhiteListRequest {
	s.StrategyId = &v
	return s
}

func (s *ModifyProcessWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
