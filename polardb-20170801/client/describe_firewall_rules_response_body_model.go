// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeFirewallRulesResponseBody
	GetDBClusterId() *string
	SetData(v *DescribeFirewallRulesResponseBodyData) *DescribeFirewallRulesResponseBody
	GetData() *DescribeFirewallRulesResponseBodyData
	SetMessage(v string) *DescribeFirewallRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeFirewallRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeFirewallRulesResponseBody
	GetSuccess() *bool
}

type DescribeFirewallRulesResponseBody struct {
	// example:
	//
	// pc-*****************
	DBClusterId *string                                `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Data        *DescribeFirewallRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 84D7DCD2-54F7-5BD2-B055-F5DE9D2B5264
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallRulesResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeFirewallRulesResponseBody) GetData() *DescribeFirewallRulesResponseBodyData {
	return s.Data
}

func (s *DescribeFirewallRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFirewallRulesResponseBody) SetDBClusterId(v string) *DescribeFirewallRulesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeFirewallRulesResponseBody) SetData(v *DescribeFirewallRulesResponseBodyData) *DescribeFirewallRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFirewallRulesResponseBody) SetMessage(v string) *DescribeFirewallRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFirewallRulesResponseBody) SetRequestId(v string) *DescribeFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallRulesResponseBody) SetSuccess(v bool) *DescribeFirewallRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFirewallRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFirewallRulesResponseBodyData struct {
	RuleList []*string `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeFirewallRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFirewallRulesResponseBodyData) GetRuleList() []*string {
	return s.RuleList
}

func (s *DescribeFirewallRulesResponseBodyData) SetRuleList(v []*string) *DescribeFirewallRulesResponseBodyData {
	s.RuleList = v
	return s
}

func (s *DescribeFirewallRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
