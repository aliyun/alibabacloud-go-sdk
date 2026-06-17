// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDefaultIPSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasicRules(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody
	GetBasicRules() *int32
	SetEnableAllPatch(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody
	GetEnableAllPatch() *int32
	SetRequestId(v string) *DescribeVpcFirewallDefaultIPSConfigResponseBody
	GetRequestId() *string
	SetRuleClass(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody
	GetRuleClass() *int32
	SetRunMode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody
	GetRunMode() *int32
}

type DescribeVpcFirewallDefaultIPSConfigResponseBody struct {
	// Indicates whether basic policies are enabled. Valid values:
	//
	// - **1**: On.
	//
	// - **0**: Off.
	//
	// example:
	//
	// 1
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// - **1**: On.
	//
	// - **0**: Off.
	//
	// example:
	//
	// 1
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-****-00090125adf1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IPS rule group. Valid values:
	//
	// - **1**: Loose rule group.
	//
	// - **2**: Medium rule group.
	//
	// - **3**: Strict rule group.
	//
	// example:
	//
	// 1
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The mode of the intrusion prevention system (IPS). Valid values:
	//
	// - **1**: Block Mode.
	//
	// - **0**: Monitor Mode.
	//
	// example:
	//
	// 0
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) GetEnableAllPatch() *int32 {
	return s.EnableAllPatch
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) GetRunMode() *int32 {
	return s.RunMode
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetBasicRules(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetEnableAllPatch(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRequestId(v string) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRuleClass(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RuleClass = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRunMode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RunMode = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
