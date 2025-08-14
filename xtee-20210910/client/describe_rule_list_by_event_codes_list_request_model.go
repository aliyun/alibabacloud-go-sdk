// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleListByEventCodesListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleListByEventCodesListRequest
	GetLang() *string
	SetEventCodes(v string) *DescribeRuleListByEventCodesListRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeRuleListByEventCodesListRequest
	GetRegId() *string
}

type DescribeRuleListByEventCodesListRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event codes, JSON array string
	//
	// This parameter is required.
	//
	// example:
	//
	// ["de_asssce8122"]
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeRuleListByEventCodesListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleListByEventCodesListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleListByEventCodesListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleListByEventCodesListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeRuleListByEventCodesListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleListByEventCodesListRequest) SetLang(v string) *DescribeRuleListByEventCodesListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleListByEventCodesListRequest) SetEventCodes(v string) *DescribeRuleListByEventCodesListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeRuleListByEventCodesListRequest) SetRegId(v string) *DescribeRuleListByEventCodesListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleListByEventCodesListRequest) Validate() error {
	return dara.Validate(s)
}
