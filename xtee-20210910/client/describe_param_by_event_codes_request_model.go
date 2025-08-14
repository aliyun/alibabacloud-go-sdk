// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParamByEventCodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeParamByEventCodesRequest
	GetLang() *string
	SetEventCodes(v string) *DescribeParamByEventCodesRequest
	GetEventCodes() *string
	SetParma(v string) *DescribeParamByEventCodesRequest
	GetParma() *string
	SetRegId(v string) *DescribeParamByEventCodesRequest
	GetRegId() *string
}

type DescribeParamByEventCodesRequest struct {
	// Set the language type for request and response, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// account_abuse_pro,account_abuse
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Query condition
	//
	// example:
	//
	// 标题/描述
	Parma *string `json:"parma,omitempty" xml:"parma,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeParamByEventCodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParamByEventCodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParamByEventCodesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeParamByEventCodesRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeParamByEventCodesRequest) GetParma() *string {
	return s.Parma
}

func (s *DescribeParamByEventCodesRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeParamByEventCodesRequest) SetLang(v string) *DescribeParamByEventCodesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeParamByEventCodesRequest) SetEventCodes(v string) *DescribeParamByEventCodesRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeParamByEventCodesRequest) SetParma(v string) *DescribeParamByEventCodesRequest {
	s.Parma = &v
	return s
}

func (s *DescribeParamByEventCodesRequest) SetRegId(v string) *DescribeParamByEventCodesRequest {
	s.RegId = &v
	return s
}

func (s *DescribeParamByEventCodesRequest) Validate() error {
	return dara.Validate(s)
}
