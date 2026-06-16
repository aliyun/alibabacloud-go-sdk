// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariableConfigListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustVariableConfigListRequest
	GetLang() *string
	SetBizType(v string) *DescribeCustVariableConfigListRequest
	GetBizType() *string
	SetRegId(v string) *DescribeCustVariableConfigListRequest
	GetRegId() *string
	SetTimeType(v string) *DescribeCustVariableConfigListRequest
	GetTimeType() *string
}

type DescribeCustVariableConfigListRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The configuration type. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// TIME_TYPE
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The time type. Valid values:
	//
	// example:
	//
	// CURRENT
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
}

func (s DescribeCustVariableConfigListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableConfigListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableConfigListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustVariableConfigListRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeCustVariableConfigListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeCustVariableConfigListRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribeCustVariableConfigListRequest) SetLang(v string) *DescribeCustVariableConfigListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustVariableConfigListRequest) SetBizType(v string) *DescribeCustVariableConfigListRequest {
	s.BizType = &v
	return s
}

func (s *DescribeCustVariableConfigListRequest) SetRegId(v string) *DescribeCustVariableConfigListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeCustVariableConfigListRequest) SetTimeType(v string) *DescribeCustVariableConfigListRequest {
	s.TimeType = &v
	return s
}

func (s *DescribeCustVariableConfigListRequest) Validate() error {
	return dara.Validate(s)
}
