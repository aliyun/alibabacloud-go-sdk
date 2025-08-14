// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSupportRuleListRequest
	GetLang() *string
	SetEventCode(v string) *DescribeSupportRuleListRequest
	GetEventCode() *string
	SetRegId(v string) *DescribeSupportRuleListRequest
	GetRegId() *string
}

type DescribeSupportRuleListRequest struct {
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
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ahpayh4121
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSupportRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportRuleListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSupportRuleListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSupportRuleListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSupportRuleListRequest) SetLang(v string) *DescribeSupportRuleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSupportRuleListRequest) SetEventCode(v string) *DescribeSupportRuleListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeSupportRuleListRequest) SetRegId(v string) *DescribeSupportRuleListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSupportRuleListRequest) Validate() error {
	return dara.Validate(s)
}
