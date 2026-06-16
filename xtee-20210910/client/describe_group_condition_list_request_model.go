// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupConditionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupConditionListRequest
	GetLang() *string
	SetRegId(v string) *DescribeGroupConditionListRequest
	GetRegId() *string
}

type DescribeGroupConditionListRequest struct {
	// The language type for the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeGroupConditionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupConditionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupConditionListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupConditionListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeGroupConditionListRequest) SetLang(v string) *DescribeGroupConditionListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupConditionListRequest) SetRegId(v string) *DescribeGroupConditionListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeGroupConditionListRequest) Validate() error {
	return dara.Validate(s)
}
