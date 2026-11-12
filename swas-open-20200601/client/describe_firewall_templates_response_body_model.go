// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallTemplates(v []*DescribeFirewallTemplatesResponseBodyFirewallTemplates) *DescribeFirewallTemplatesResponseBody
	GetFirewallTemplates() []*DescribeFirewallTemplatesResponseBodyFirewallTemplates
	SetPageNumber(v int32) *DescribeFirewallTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFirewallTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFirewallTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFirewallTemplatesResponseBody
	GetTotalCount() *int32
}

type DescribeFirewallTemplatesResponseBody struct {
	// The information about the queried firewall templates.
	FirewallTemplates []*DescribeFirewallTemplatesResponseBodyFirewallTemplates `json:"FirewallTemplates,omitempty" xml:"FirewallTemplates,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFirewallTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponseBody) GetFirewallTemplates() []*DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	return s.FirewallTemplates
}

func (s *DescribeFirewallTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFirewallTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFirewallTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFirewallTemplatesResponseBody) SetFirewallTemplates(v []*DescribeFirewallTemplatesResponseBodyFirewallTemplates) *DescribeFirewallTemplatesResponseBody {
	s.FirewallTemplates = v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetPageNumber(v int32) *DescribeFirewallTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetPageSize(v int32) *DescribeFirewallTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetRequestId(v string) *DescribeFirewallTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetTotalCount(v int32) *DescribeFirewallTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) Validate() error {
	if s.FirewallTemplates != nil {
		for _, item := range s.FirewallTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallTemplatesResponseBodyFirewallTemplates struct {
	// The time when the firewall template was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// example:
	//
	// 2024-04-29T02:01:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the firewall template was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-10T02:10:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the firewall template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the firewall template.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The details of the firewall template rules.
	FirewallTemplateRules []*DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules `json:"FirewallTemplateRules,omitempty" xml:"FirewallTemplateRules,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplates) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) GetFirewallTemplateRules() []*DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	return s.FirewallTemplateRules
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetCreateTime(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetCreationTime(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.CreationTime = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetDescription(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.Description = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetFirewallTemplateId(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetFirewallTemplateRules(v []*DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.FirewallTemplateRules = v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetName(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.Name = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) Validate() error {
	if s.FirewallTemplateRules != nil {
		for _, item := range s.FirewallTemplateRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules struct {
	// The ID of the firewall template rule.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	FirewallTemplateRuleId *string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty"`
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall template rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
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
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GetFirewallTemplateRuleId() *string {
	return s.FirewallTemplateRuleId
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GetPort() *string {
	return s.Port
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GetRemark() *string {
	return s.Remark
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetFirewallTemplateRuleId(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.FirewallTemplateRuleId = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetPort(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.Port = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetRemark(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.Remark = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetRuleProtocol(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.RuleProtocol = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetSourceCidrIp(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) Validate() error {
	return dara.Validate(s)
}
