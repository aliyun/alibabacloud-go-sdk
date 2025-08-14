// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleCountByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleCountByUserIdRequest
	GetLang() *string
	SetCreateType(v string) *DescribeRuleCountByUserIdRequest
	GetCreateType() *string
	SetRegId(v string) *DescribeRuleCountByUserIdRequest
	GetRegId() *string
}

type DescribeRuleCountByUserIdRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeRuleCountByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleCountByUserIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleCountByUserIdRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleCountByUserIdRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeRuleCountByUserIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleCountByUserIdRequest) SetLang(v string) *DescribeRuleCountByUserIdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleCountByUserIdRequest) SetCreateType(v string) *DescribeRuleCountByUserIdRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeRuleCountByUserIdRequest) SetRegId(v string) *DescribeRuleCountByUserIdRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleCountByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
