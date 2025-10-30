// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v *DescribeInstanceResponseBodyDetails) *DescribeInstanceResponseBody
	GetDetails() *DescribeInstanceResponseBodyDetails
	SetEdition(v string) *DescribeInstanceResponseBody
	GetEdition() *string
	SetEndTime(v int64) *DescribeInstanceResponseBody
	GetEndTime() *int64
	SetInDebt(v string) *DescribeInstanceResponseBody
	GetInDebt() *string
	SetInstanceId(v string) *DescribeInstanceResponseBody
	GetInstanceId() *string
	SetPayType(v string) *DescribeInstanceResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *DescribeInstanceResponseBody
	GetStartTime() *int64
	SetStatus(v int32) *DescribeInstanceResponseBody
	GetStatus() *int32
}

type DescribeInstanceResponseBody struct {
	// The details of the WAF instance.
	Details *DescribeInstanceResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	// The edition of the WAF instance.
	//
	// example:
	//
	// default_version
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The expiration time of the WAF instance.
	//
	// example:
	//
	// 4809859200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the WAF instance has overdue payments. Valid values:
	//
	// 	- **0**: The WAF instance does not have overdue payments.
	//
	// 	- **1**: The WAF instance has overdue payments.
	//
	// example:
	//
	// 1
	InDebt *string `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	// The ID of the WAF instance.
	//
	// example:
	//
	// waf-cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The billing method of the WAF instance. Valid values:
	//
	// 	- **POSTPAY:*	- The WAF instance uses the pay-as-you-go billing method.
	//
	// 	- **PREPAY:*	- The WAF instance uses the subscription billing method.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The purchase time of the WAF instance. The time is in the UNIX timestamp format. The time is displayed in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1668496310000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the WAF instance. Valid values:
	//
	// 	- **1:*	- The WAF instance is in a normal state.
	//
	// 	- **2:*	- The WAF instance has expired.
	//
	// 	- **3:*	- The WAF instance has been released.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetDetails() *DescribeInstanceResponseBodyDetails {
	return s.Details
}

func (s *DescribeInstanceResponseBody) GetEdition() *string {
	return s.Edition
}

func (s *DescribeInstanceResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeInstanceResponseBody) GetInDebt() *string {
	return s.InDebt
}

func (s *DescribeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeInstanceResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeInstanceResponseBody) SetDetails(v *DescribeInstanceResponseBodyDetails) *DescribeInstanceResponseBody {
	s.Details = v
	return s
}

func (s *DescribeInstanceResponseBody) SetEdition(v string) *DescribeInstanceResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEndTime(v int64) *DescribeInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInDebt(v string) *DescribeInstanceResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPayType(v string) *DescribeInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStartTime(v int64) *DescribeInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v int32) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	if s.Details != nil {
		if err := s.Details.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceResponseBodyDetails struct {
	// The maximum number of IP addresses that can be added to the match content of a match condition. For more information, see [Match conditions](https://help.aliyun.com/document_detail/374354.html).
	//
	// example:
	//
	// 100
	AclRuleMaxIpCount *int64 `json:"AclRuleMaxIpCount,omitempty" xml:"AclRuleMaxIpCount,omitempty"`
	// Indicates whether the scan protection module is supported. Valid values:
	//
	// 	- **true:*	- The scan protection module is supported.
	//
	// 	- **false:*	- The scan protection module is not supported.
	//
	// example:
	//
	// true
	AntiScan *bool `json:"AntiScan,omitempty" xml:"AntiScan,omitempty"`
	// The maximum number of scan protection rule templates that can be configured.
	//
	// example:
	//
	// 20
	AntiScanTemplateMaxCount *int64 `json:"AntiScanTemplateMaxCount,omitempty" xml:"AntiScanTemplateMaxCount,omitempty"`
	// The maximum number of back-to-origin IP addresses that can be configured.
	//
	// example:
	//
	// 20
	BackendMaxCount *int64 `json:"BackendMaxCount,omitempty" xml:"BackendMaxCount,omitempty"`
	// Indicates whether the basic protection rule module is supported. Valid values:
	//
	// 	- **true:*	- The basic protection rule module is supported.
	//
	// 	- **false:*	- The basic protection rule module is not supported.
	//
	// example:
	//
	// true
	BaseWafGroup *bool `json:"BaseWafGroup,omitempty" xml:"BaseWafGroup,omitempty"`
	// The maximum number of protection rules that can be included in a basic protection rule template.
	//
	// example:
	//
	// 100
	BaseWafGroupRuleInTemplateMaxCount *int64 `json:"BaseWafGroupRuleInTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleInTemplateMaxCount,omitempty"`
	// The maximum number of basic protection rule templates that can be configured.
	//
	// example:
	//
	// 20
	BaseWafGroupRuleTemplateMaxCount *int64 `json:"BaseWafGroupRuleTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleTemplateMaxCount,omitempty"`
	// Indicates whether the bot management module is supported. Valid values:
	//
	// 	- **true:*	- The bot management module is supported.
	//
	// 	- **false:*	- The bot management module is not supported.
	//
	// example:
	//
	// true
	Bot *bool `json:"Bot,omitempty" xml:"Bot,omitempty"`
	// Indicates whether bot management for app protection is supported. Valid values:
	//
	// 	- **true:*	- Bot management for app protection is supported.
	//
	// 	- **false:*	- Bot management for app protection is not supported.
	//
	// example:
	//
	// true
	BotApp *string `json:"BotApp,omitempty" xml:"BotApp,omitempty"`
	// The maximum number of bot management rule templates that can be configured.
	//
	// example:
	//
	// 50
	BotTemplateMaxCount *int64 `json:"BotTemplateMaxCount,omitempty" xml:"BotTemplateMaxCount,omitempty"`
	// Indicates whether bot management for website protection is supported. Valid values:
	//
	// 	- **true:*	- Bot management for website protection is supported.
	//
	// 	- **false:*	- Bot management for website protection is not supported.
	//
	// example:
	//
	// true
	BotWeb *string `json:"BotWeb,omitempty" xml:"BotWeb,omitempty"`
	// The maximum number of CNAMEs that can be added.
	//
	// example:
	//
	// 1000
	CnameResourceMaxCount *int64 `json:"CnameResourceMaxCount,omitempty" xml:"CnameResourceMaxCount,omitempty"`
	// Indicates whether the custom response module is supported. Valid values:
	//
	// 	- **true:*	- The custom response module is supported.
	//
	// 	- **false:*	- The custom response module is not supported.
	//
	// example:
	//
	// true
	CustomResponse *bool `json:"CustomResponse,omitempty" xml:"CustomResponse,omitempty"`
	// The maximum number of rules that can be included in a custom response rule template.
	//
	// example:
	//
	// 100
	CustomResponseRuleInTemplateMaxCount *int64 `json:"CustomResponseRuleInTemplateMaxCount,omitempty" xml:"CustomResponseRuleInTemplateMaxCount,omitempty"`
	// The maximum number of custom response rule templates that can be configured.
	//
	// example:
	//
	// 20
	CustomResponseTemplateMaxCount *int64 `json:"CustomResponseTemplateMaxCount,omitempty" xml:"CustomResponseTemplateMaxCount,omitempty"`
	// Indicates whether the custom rule module is supported. Valid values:
	//
	// 	- **true:*	- The custom rule module is supported.
	//
	// 	- **false:*	- The custom rule module is not supported.
	//
	// example:
	//
	// true
	CustomRule *bool `json:"CustomRule,omitempty" xml:"CustomRule,omitempty"`
	// The action that can be included in a custom rule.
	//
	// example:
	//
	// block
	CustomRuleAction *string `json:"CustomRuleAction,omitempty" xml:"CustomRuleAction,omitempty"`
	// The match conditions that can be used in a custom rule. For more information, see **Match condition parameters*	- in the "**Parameters of custom rules (custom_acl)**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	//
	// example:
	//
	// URL
	CustomRuleCondition *string `json:"CustomRuleCondition,omitempty" xml:"CustomRuleCondition,omitempty"`
	// The maximum number of rules that can be included in a custom rule template.
	//
	// example:
	//
	// 100
	CustomRuleInTemplateMaxCount *int64 `json:"CustomRuleInTemplateMaxCount,omitempty" xml:"CustomRuleInTemplateMaxCount,omitempty"`
	// The statistical object for rate limiting in a custom rule.
	//
	// example:
	//
	// header
	CustomRuleRatelimitor *string `json:"CustomRuleRatelimitor,omitempty" xml:"CustomRuleRatelimitor,omitempty"`
	// The maximum number of custom rule templates that can be configured.
	//
	// example:
	//
	// 20
	CustomRuleTemplateMaxCount *int64 `json:"CustomRuleTemplateMaxCount,omitempty" xml:"CustomRuleTemplateMaxCount,omitempty"`
	// The maximum number of protected object groups that can be configured.
	//
	// example:
	//
	// 100
	DefenseGroupMaxCount *int64 `json:"DefenseGroupMaxCount,omitempty" xml:"DefenseGroupMaxCount,omitempty"`
	// The maximum number of protected objects that can be included in a protected object group.
	//
	// example:
	//
	// 100
	DefenseObjectInGroupMaxCount *int64 `json:"DefenseObjectInGroupMaxCount,omitempty" xml:"DefenseObjectInGroupMaxCount,omitempty"`
	// The maximum number of protected objects to which a protection rule template can be applied.
	//
	// example:
	//
	// 100
	DefenseObjectInTemplateMaxCount *int64 `json:"DefenseObjectInTemplateMaxCount,omitempty" xml:"DefenseObjectInTemplateMaxCount,omitempty"`
	// The maximum number of protected objects that can be configured.
	//
	// example:
	//
	// 20,000
	DefenseObjectMaxCount *int64 `json:"DefenseObjectMaxCount,omitempty" xml:"DefenseObjectMaxCount,omitempty"`
	// Indicates whether the data leakage prevention module is supported. Valid values:
	//
	// 	- **true:*	- The data leakage prevention module is supported.
	//
	// 	- **false:*	- The data leakage prevention module is not supported.
	//
	// example:
	//
	// true
	Dlp *bool `json:"Dlp,omitempty" xml:"Dlp,omitempty"`
	// The maximum number of rules that can be included in a data leakage prevention rule template.
	//
	// example:
	//
	// 50
	DlpRuleInTemplateMaxCount *int64 `json:"DlpRuleInTemplateMaxCount,omitempty" xml:"DlpRuleInTemplateMaxCount,omitempty"`
	// The maximum number of data leakage prevention rule templates that can be configured.
	//
	// example:
	//
	// 50
	DlpTemplateMaxCount *int64 `json:"DlpTemplateMaxCount,omitempty" xml:"DlpTemplateMaxCount,omitempty"`
	// example:
	//
	// 2000
	ElasticQps *int32 `json:"ElasticQps,omitempty" xml:"ElasticQps,omitempty"`
	// Indicates whether exclusive IP addresses are supported. Valid values:
	//
	// 	- **true:*	- Exclusive IP addresses are supported.
	//
	// 	- **false:*	- Exclusive IP addresses are not supported.
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// example:
	//
	// 10000
	ExtendQps *int32 `json:"ExtendQps,omitempty" xml:"ExtendQps,omitempty"`
	// example:
	//
	// 1000
	FreeQps *int32 `json:"FreeQps,omitempty" xml:"FreeQps,omitempty"`
	// Indicates whether global server load balancing (GSLB) is supported. Valid values:
	//
	// 	- **true:*	- GSLB is supported.
	//
	// 	- **false:*	- GSLB is not supported.
	//
	// example:
	//
	// true
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// The HTTP port range that is supported. For more information, see [View supported ports](https://help.aliyun.com/document_detail/385578.html).
	//
	// example:
	//
	// 80
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS port range that is supported. For more information, see [View supported ports](https://help.aliyun.com/document_detail/385578.html).
	//
	// example:
	//
	// 443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Indicates whether the IP address blacklist module is supported. Valid values:
	//
	// 	- **true:*	- The IP address blacklist module is supported.
	//
	// 	- **false:*	- The IP address blacklist module is not supported.
	//
	// example:
	//
	// true
	IpBlacklist *bool `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The maximum number of IP addresses that can be added to an IP address blacklist rule.
	//
	// example:
	//
	// 200
	IpBlacklistIpInRuleMaxCount *int64 `json:"IpBlacklistIpInRuleMaxCount,omitempty" xml:"IpBlacklistIpInRuleMaxCount,omitempty"`
	// The maximum number of rules that can be included in an IP address blacklist rule template.
	//
	// example:
	//
	// 100
	IpBlacklistRuleInTemplateMaxCount *int64 `json:"IpBlacklistRuleInTemplateMaxCount,omitempty" xml:"IpBlacklistRuleInTemplateMaxCount,omitempty"`
	// The maximum number of IP address blacklist rule templates that can be configured.
	//
	// example:
	//
	// 20
	IpBlacklistTemplateMaxCount *int64 `json:"IpBlacklistTemplateMaxCount,omitempty" xml:"IpBlacklistTemplateMaxCount,omitempty"`
	// Indicates whether IPv6 is supported. Valid values:
	//
	// 	- **true:*	- IPv6 is supported.
	//
	// 	- **false:*	- IPv6 is not supported.
	//
	// example:
	//
	// true
	Ipv6 *bool `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Indicates whether the log collection feature is supported. Valid values:
	//
	// 	- **true:*	- The log collection feature is supported.
	//
	// 	- **false:*	- The log collection feature is not supported.
	//
	// example:
	//
	// true
	LogService *bool `json:"LogService,omitempty" xml:"LogService,omitempty"`
	// Indicates whether major event protection is supported. Valid values:
	//
	// 	- **true:*	- Major event protection is supported.
	//
	// 	- **false:*	- Major event protection is not supported.
	//
	// example:
	//
	// true
	MajorProtection *bool `json:"MajorProtection,omitempty" xml:"MajorProtection,omitempty"`
	// The maximum number of major event protection rule templates that can be configured.
	//
	// example:
	//
	// 20
	MajorProtectionTemplateMaxCount *int64 `json:"MajorProtectionTemplateMaxCount,omitempty" xml:"MajorProtectionTemplateMaxCount,omitempty"`
	// example:
	//
	// 2000
	QpsBillingCap *int32 `json:"QpsBillingCap,omitempty" xml:"QpsBillingCap,omitempty"`
	// Indicates whether the website tamper-proofing module is supported. Valid values:
	//
	// 	- **true:*	- The website tamper-proofing module is supported.
	//
	// 	- **false:*	- The website tamper-proofing module is not supported.
	//
	// example:
	//
	// true
	Tamperproof *bool `json:"Tamperproof,omitempty" xml:"Tamperproof,omitempty"`
	// The maximum number of rules that can be included in a website tamper-proofing rule template.
	//
	// example:
	//
	// 50
	TamperproofRuleInTemplateMaxCount *int64 `json:"TamperproofRuleInTemplateMaxCount,omitempty" xml:"TamperproofRuleInTemplateMaxCount,omitempty"`
	// The maximum number of website tamper-proofing rule templates that can be configured.
	//
	// example:
	//
	// 50
	TamperproofTemplateMaxCount *int64 `json:"TamperproofTemplateMaxCount,omitempty" xml:"TamperproofTemplateMaxCount,omitempty"`
	// The maximum number of IP addresses or CIDR blocks that can be added to an IP address blacklist in a batch.
	//
	// example:
	//
	// 2,000
	VastIpBlacklistInFileMaxCount *int64 `json:"VastIpBlacklistInFileMaxCount,omitempty" xml:"VastIpBlacklistInFileMaxCount,omitempty"`
	// The maximum number of IP addresses or CIDR blocks that can be added to an IP address blacklist on a page.
	//
	// example:
	//
	// 500
	VastIpBlacklistInOperationMaxCount *int64 `json:"VastIpBlacklistInOperationMaxCount,omitempty" xml:"VastIpBlacklistInOperationMaxCount,omitempty"`
	// The maximum number of IP addresses or CIDR blocks that can be added to an IP address blacklist per Alibaba Cloud account.
	//
	// example:
	//
	// 50,000
	VastIpBlacklistMaxCount *int64 `json:"VastIpBlacklistMaxCount,omitempty" xml:"VastIpBlacklistMaxCount,omitempty"`
	// Indicates whether the whitelist module is supported. Valid values:
	//
	// 	- **true:*	- The whitelist module is supported.
	//
	// 	- **false:*	- The whitelist module is not supported.
	//
	// example:
	//
	// true
	Whitelist *bool `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// The logical operators that can be used in a whitelist rule. For more information, see **Match condition parameters*	- in the "**Parameters of whitelist rules (whitelist)**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	//
	// example:
	//
	// contain
	WhitelistLogical *string `json:"WhitelistLogical,omitempty" xml:"WhitelistLogical,omitempty"`
	// The match fields that can be used in a whitelist rule. For more information, see **Match condition parameters*	- in the "**Parameters of whitelist rules (whitelist)**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	//
	// example:
	//
	// URL
	WhitelistRuleCondition *string `json:"WhitelistRuleCondition,omitempty" xml:"WhitelistRuleCondition,omitempty"`
	// The maximum number of rules that can be included in a whitelist rule template.
	//
	// example:
	//
	// 100
	WhitelistRuleInTemplateMaxCount *int64 `json:"WhitelistRuleInTemplateMaxCount,omitempty" xml:"WhitelistRuleInTemplateMaxCount,omitempty"`
	// The maximum number of whitelist rule templates that can be configured.
	//
	// example:
	//
	// 20
	WhitelistTemplateMaxCount *int64 `json:"WhitelistTemplateMaxCount,omitempty" xml:"WhitelistTemplateMaxCount,omitempty"`
}

func (s DescribeInstanceResponseBodyDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyDetails) GetAclRuleMaxIpCount() *int64 {
	return s.AclRuleMaxIpCount
}

func (s *DescribeInstanceResponseBodyDetails) GetAntiScan() *bool {
	return s.AntiScan
}

func (s *DescribeInstanceResponseBodyDetails) GetAntiScanTemplateMaxCount() *int64 {
	return s.AntiScanTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetBackendMaxCount() *int64 {
	return s.BackendMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetBaseWafGroup() *bool {
	return s.BaseWafGroup
}

func (s *DescribeInstanceResponseBodyDetails) GetBaseWafGroupRuleInTemplateMaxCount() *int64 {
	return s.BaseWafGroupRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetBaseWafGroupRuleTemplateMaxCount() *int64 {
	return s.BaseWafGroupRuleTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetBot() *bool {
	return s.Bot
}

func (s *DescribeInstanceResponseBodyDetails) GetBotApp() *string {
	return s.BotApp
}

func (s *DescribeInstanceResponseBodyDetails) GetBotTemplateMaxCount() *int64 {
	return s.BotTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetBotWeb() *string {
	return s.BotWeb
}

func (s *DescribeInstanceResponseBodyDetails) GetCnameResourceMaxCount() *int64 {
	return s.CnameResourceMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomResponse() *bool {
	return s.CustomResponse
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomResponseRuleInTemplateMaxCount() *int64 {
	return s.CustomResponseRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomResponseTemplateMaxCount() *int64 {
	return s.CustomResponseTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomRule() *bool {
	return s.CustomRule
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomRuleAction() *string {
	return s.CustomRuleAction
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomRuleCondition() *string {
	return s.CustomRuleCondition
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomRuleInTemplateMaxCount() *int64 {
	return s.CustomRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomRuleRatelimitor() *string {
	return s.CustomRuleRatelimitor
}

func (s *DescribeInstanceResponseBodyDetails) GetCustomRuleTemplateMaxCount() *int64 {
	return s.CustomRuleTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetDefenseGroupMaxCount() *int64 {
	return s.DefenseGroupMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetDefenseObjectInGroupMaxCount() *int64 {
	return s.DefenseObjectInGroupMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetDefenseObjectInTemplateMaxCount() *int64 {
	return s.DefenseObjectInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetDefenseObjectMaxCount() *int64 {
	return s.DefenseObjectMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetDlp() *bool {
	return s.Dlp
}

func (s *DescribeInstanceResponseBodyDetails) GetDlpRuleInTemplateMaxCount() *int64 {
	return s.DlpRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetDlpTemplateMaxCount() *int64 {
	return s.DlpTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetElasticQps() *int32 {
	return s.ElasticQps
}

func (s *DescribeInstanceResponseBodyDetails) GetExclusiveIp() *bool {
	return s.ExclusiveIp
}

func (s *DescribeInstanceResponseBodyDetails) GetExtendQps() *int32 {
	return s.ExtendQps
}

func (s *DescribeInstanceResponseBodyDetails) GetFreeQps() *int32 {
	return s.FreeQps
}

func (s *DescribeInstanceResponseBodyDetails) GetGslb() *bool {
	return s.Gslb
}

func (s *DescribeInstanceResponseBodyDetails) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *DescribeInstanceResponseBodyDetails) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *DescribeInstanceResponseBodyDetails) GetIpBlacklist() *bool {
	return s.IpBlacklist
}

func (s *DescribeInstanceResponseBodyDetails) GetIpBlacklistIpInRuleMaxCount() *int64 {
	return s.IpBlacklistIpInRuleMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetIpBlacklistRuleInTemplateMaxCount() *int64 {
	return s.IpBlacklistRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetIpBlacklistTemplateMaxCount() *int64 {
	return s.IpBlacklistTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetIpv6() *bool {
	return s.Ipv6
}

func (s *DescribeInstanceResponseBodyDetails) GetLogService() *bool {
	return s.LogService
}

func (s *DescribeInstanceResponseBodyDetails) GetMajorProtection() *bool {
	return s.MajorProtection
}

func (s *DescribeInstanceResponseBodyDetails) GetMajorProtectionTemplateMaxCount() *int64 {
	return s.MajorProtectionTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetQpsBillingCap() *int32 {
	return s.QpsBillingCap
}

func (s *DescribeInstanceResponseBodyDetails) GetTamperproof() *bool {
	return s.Tamperproof
}

func (s *DescribeInstanceResponseBodyDetails) GetTamperproofRuleInTemplateMaxCount() *int64 {
	return s.TamperproofRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetTamperproofTemplateMaxCount() *int64 {
	return s.TamperproofTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetVastIpBlacklistInFileMaxCount() *int64 {
	return s.VastIpBlacklistInFileMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetVastIpBlacklistInOperationMaxCount() *int64 {
	return s.VastIpBlacklistInOperationMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetVastIpBlacklistMaxCount() *int64 {
	return s.VastIpBlacklistMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetWhitelist() *bool {
	return s.Whitelist
}

func (s *DescribeInstanceResponseBodyDetails) GetWhitelistLogical() *string {
	return s.WhitelistLogical
}

func (s *DescribeInstanceResponseBodyDetails) GetWhitelistRuleCondition() *string {
	return s.WhitelistRuleCondition
}

func (s *DescribeInstanceResponseBodyDetails) GetWhitelistRuleInTemplateMaxCount() *int64 {
	return s.WhitelistRuleInTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) GetWhitelistTemplateMaxCount() *int64 {
	return s.WhitelistTemplateMaxCount
}

func (s *DescribeInstanceResponseBodyDetails) SetAclRuleMaxIpCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AclRuleMaxIpCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScan(v bool) *DescribeInstanceResponseBodyDetails {
	s.AntiScan = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScanTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AntiScanTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBackendMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BackendMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroup(v bool) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroup = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBot(v bool) *DescribeInstanceResponseBodyDetails {
	s.Bot = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotApp(v string) *DescribeInstanceResponseBodyDetails {
	s.BotApp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BotTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotWeb(v string) *DescribeInstanceResponseBodyDetails {
	s.BotWeb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCnameResourceMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CnameResourceMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponse(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomResponse = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRule(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomRule = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleAction(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleAction = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleRatelimitor(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleRatelimitor = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlp(v bool) *DescribeInstanceResponseBodyDetails {
	s.Dlp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlpRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DlpRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlpTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DlpTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetElasticQps(v int32) *DescribeInstanceResponseBodyDetails {
	s.ElasticQps = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetExclusiveIp(v bool) *DescribeInstanceResponseBodyDetails {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetExtendQps(v int32) *DescribeInstanceResponseBodyDetails {
	s.ExtendQps = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetFreeQps(v int32) *DescribeInstanceResponseBodyDetails {
	s.FreeQps = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetGslb(v bool) *DescribeInstanceResponseBodyDetails {
	s.Gslb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpsPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklist(v bool) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistIpInRuleMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistIpInRuleMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpv6(v bool) *DescribeInstanceResponseBodyDetails {
	s.Ipv6 = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetLogService(v bool) *DescribeInstanceResponseBodyDetails {
	s.LogService = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtection(v bool) *DescribeInstanceResponseBodyDetails {
	s.MajorProtection = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtectionTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.MajorProtectionTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetQpsBillingCap(v int32) *DescribeInstanceResponseBodyDetails {
	s.QpsBillingCap = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproof(v bool) *DescribeInstanceResponseBodyDetails {
	s.Tamperproof = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproofRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.TamperproofRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproofTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.TamperproofTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInFileMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInFileMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInOperationMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInOperationMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelist(v bool) *DescribeInstanceResponseBodyDetails {
	s.Whitelist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistLogical(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistLogical = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) Validate() error {
	return dara.Validate(s)
}
