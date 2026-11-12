// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplateRulesApplyResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFirewallTemplateRulesApplyResultResponseBody
	GetRequestId() *string
	SetData(v []*DescribeFirewallTemplateRulesApplyResultResponseBodyData) *DescribeFirewallTemplateRulesApplyResultResponseBody
	GetData() []*DescribeFirewallTemplateRulesApplyResultResponseBodyData
}

type DescribeFirewallTemplateRulesApplyResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Data []*DescribeFirewallTemplateRulesApplyResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) GetData() []*DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	return s.Data
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) SetRequestId(v string) *DescribeFirewallTemplateRulesApplyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) SetData(v []*DescribeFirewallTemplateRulesApplyResultResponseBodyData) *DescribeFirewallTemplateRulesApplyResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallTemplateRulesApplyResultResponseBodyData struct {
	// The error codes when the firewall template rule fails to be applied.
	//
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is displayed when the firewall template rule fails to be applied.
	//
	// example:
	//
	// An error occurred while processing your request.
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// > If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// - TCP
	//
	// - UDP
	//
	// - TCP+UDP
	//
	// - ICMP
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 119.145.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The status of applying the firewall template rule to the simple application servers.
	//
	// - Pending: The template rule does not start to be applied to the simple application servers.
	//
	// - Applying: The template rule is being applied to the simple application servers.
	//
	// - Success: The template rule is successfully applied to the simple application servers.
	//
	// - PartFailed: The template rule fails to be applied to some simple application servers.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetPort() *string {
	return s.Port
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetErrorCode(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetErrorInfo(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetPort(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.Port = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetRemark(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetRuleProtocol(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.RuleProtocol = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetSourceCidrIp(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetSuccess(v bool) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.Success = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
