// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeModules(v []*DescribeChargeModuleResponseBodyChargeModules) *DescribeChargeModuleResponseBody
	GetChargeModules() []*DescribeChargeModuleResponseBodyChargeModules
	SetRequestId(v string) *DescribeChargeModuleResponseBody
	GetRequestId() *string
}

type DescribeChargeModuleResponseBody struct {
	// The list of WAF pricing module information.
	ChargeModules []*DescribeChargeModuleResponseBodyChargeModules `json:"ChargeModules,omitempty" xml:"ChargeModules,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChargeModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChargeModuleResponseBody) GetChargeModules() []*DescribeChargeModuleResponseBodyChargeModules {
	return s.ChargeModules
}

func (s *DescribeChargeModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChargeModuleResponseBody) SetChargeModules(v []*DescribeChargeModuleResponseBodyChargeModules) *DescribeChargeModuleResponseBody {
	s.ChargeModules = v
	return s
}

func (s *DescribeChargeModuleResponseBody) SetRequestId(v string) *DescribeChargeModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChargeModuleResponseBody) Validate() error {
	if s.ChargeModules != nil {
		for _, item := range s.ChargeModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChargeModuleResponseBodyChargeModules struct {
	// The pricing mode of the pricing module. Valid values:
	//
	// - **NORMAL_PRICE**: standard pricing.
	//
	// - **STEP_ACCUMULATION**: tiered pricing.
	//
	// example:
	//
	// NORMAL_PRICE
	ChargeMode *string `json:"ChargeMode,omitempty" xml:"ChargeMode,omitempty"`
	// The pricing details of the pricing module.
	ChargeModeDetails []*string `json:"ChargeModeDetails,omitempty" xml:"ChargeModeDetails,omitempty" type:"Repeated"`
	// The pricing unit.
	//
	// example:
	//
	// SeCU
	ChargeUnit *string `json:"ChargeUnit,omitempty" xml:"ChargeUnit,omitempty"`
	// The pricing module identity. Valid values:
	//
	// - **domainCount**: the number of CNAME-connected domain names.
	//
	// - **qps**: the peak QPS.
	//
	// - **request**: the basic traffic fee.
	//
	// - **ipBlacklistRuleCount**: the number of IP blacklist rules.
	//
	// - **customAclBaseRuleCount**: the number of Basic Policies in custom rules.
	//
	// - **customAclAdvanceRuleCount**: the number of advanced rules in custom rules.
	//
	// - **antiScanRuleCount**: the number of scan protection rules.
	//
	// - **customResponseRuleCount**: the number of custom response rules.
	//
	// - **ipv6**: IPv6.
	//
	// - **gslb**: intelligent load balancing.
	//
	// - **exclusiveIpCount**: the number of exclusive IP addresses.
	//
	// - **ccRuleCount**: the number of HTTP flood mitigation rules.
	//
	// - **regionBlockRuleCount**: the number of Location Blacklist rules.
	//
	// - **tamperproofRuleCount**: the number of web tamper proofing rules.
	//
	// - **dlpRuleCount**: the number of information leak prevention rules.
	//
	// - **botTraffic**: the Bot management traffic fee.
	//
	// - **aiWhiteListTemplateCount**: the number of intelligent whitelist templates.
	//
	// - **apisecResourceCount**: the number of protected objects with API security enabled.
	//
	// - **apisecTraffic**: the API security traffic fee.
	//
	// - **compliance**: the number of protocol compliance templates.
	//
	// - **riskTraffic**: the number of risk identification hits in Bot management.
	//
	// - **assetStatus**: the asset center.
	//
	// - **nonPort**: non-standard ports.
	//
	// - **customAclCaptcha**: the number of custom rule slider verification attempts.
	//
	// - **wafBaseTemplateCount**: the number of web core protection rules.
	//
	// - **instanceFee**: the WAF instance fee.
	//
	// - **spikeThrottleRuleCount**: the number of peak traffic throttling rules.
	//
	// - **botWebTemplateCount**: the number of web protection templates in Bot management.
	//
	// - **botAppTemplateCount**: the number of app protection templates in Bot management.
	//
	// - **customAclBotRuleCount**: the number of advanced custom rules in Bot management.
	//
	// example:
	//
	// domainCount
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The billing period type of the pricing module. Valid values:
	//
	// - **Hour**: hourly billing.
	//
	// example:
	//
	// Hour
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The usage type of the pricing module. Valid values:
	//
	// - **template**: template.
	//
	// - **qps**: QPS.
	//
	// - **domain**: domain name.
	//
	// - **rule**: rule.
	//
	// - **ip**: IP address.
	//
	// - **resource**: protected object.
	//
	// - **reqest**: request.
	//
	// - **function**: feature enablement.
	//
	// - **time**: number of times.
	//
	// example:
	//
	// domain
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	// The billing unit factor of the pricing module.
	//
	// > The billing unit factor **UsageUnitFactor*	- multiplied by the usage type **UsageType*	- forms the billing unit of the module.
	//
	// example:
	//
	// 1
	UsageUnitFactor *int32 `json:"UsageUnitFactor,omitempty" xml:"UsageUnitFactor,omitempty"`
}

func (s DescribeChargeModuleResponseBodyChargeModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeModuleResponseBodyChargeModules) GoString() string {
	return s.String()
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetChargeMode() *string {
	return s.ChargeMode
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetChargeModeDetails() []*string {
	return s.ChargeModeDetails
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetChargeUnit() *string {
	return s.ChargeUnit
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetUsageUnitFactor() *int32 {
	return s.UsageUnitFactor
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetChargeMode(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ChargeMode = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetChargeModeDetails(v []*string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ChargeModeDetails = v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetChargeUnit(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ChargeUnit = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetModuleCode(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ModuleCode = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetPeriodType(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.PeriodType = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetUsageType(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.UsageType = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetUsageUnitFactor(v int32) *DescribeChargeModuleResponseBodyChargeModules {
	s.UsageUnitFactor = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) Validate() error {
	return dara.Validate(s)
}
