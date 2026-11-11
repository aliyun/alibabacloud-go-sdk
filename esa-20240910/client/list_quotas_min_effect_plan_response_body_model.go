// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasMinEffectPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotas(v []*ListQuotasMinEffectPlanResponseBodyQuotas) *ListQuotasMinEffectPlanResponseBody
	GetQuotas() []*ListQuotasMinEffectPlanResponseBodyQuotas
	SetRequestId(v string) *ListQuotasMinEffectPlanResponseBody
	GetRequestId() *string
}

type ListQuotasMinEffectPlanResponseBody struct {
	// The details of the minimum required plan versions for the quotas.
	Quotas []*ListQuotasMinEffectPlanResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQuotasMinEffectPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasMinEffectPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasMinEffectPlanResponseBody) GetQuotas() []*ListQuotasMinEffectPlanResponseBodyQuotas {
	return s.Quotas
}

func (s *ListQuotasMinEffectPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotasMinEffectPlanResponseBody) SetQuotas(v []*ListQuotasMinEffectPlanResponseBodyQuotas) *ListQuotasMinEffectPlanResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBody) SetRequestId(v string) *ListQuotasMinEffectPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBody) Validate() error {
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQuotasMinEffectPlanResponseBodyQuotas struct {
	// The name of the quota. Valid values:
	//
	// - **waf:phase:http_anti_scan:actions**: WAF scan protection - action enumeration
	//
	// - **waf:phase:http_bot:actions**: WAF bot management - all action enumerations
	//
	// - **waf:phase:http_bot:http_custom_cc_dev:characteristic:fields**: WAF bot management - statistical object enumeration for custom device rate limiting
	//
	// - **waf:phase:http_bot:http_custom_cc_ip:characteristic:fields**: WAF bot management - statistical object enumeration for custom IP rate limiting
	//
	// - ****waf:phase:http_bot:match:symbols****: WAF bot management - match symbol enumeration
	//
	// - **waf:phase:http_bot:http_custom_cc:characteristic:fields**: WAF bot management - statistical object enumeration for custom session rate limiting
	//
	// - **waf:phase:http_bot:match:fields**: WAF bot management phase - match field enumeration
	//
	// - **waf:phase:http_whitelist:match:symbols**: WAF whitelist - match symbol enumeration
	//
	// - **waf:phase:http_whitelist:match:fields**: WAF whitelist - match field enumeration
	//
	// - **waf:phase:http_anti_scan:http_directory_traversal:characteristic:fields**: WAF scan protection - statistical object enumeration for folder traverse blocking
	//
	// - **waf:phase:http_anti_scan:http_high_frequency:characteristic:fields**: WAF scan protection - statistical object enumeration for high-frequency scan blocking
	//
	// - **waf:phase:http_anti_scan:match:symbols**: WAF scan protection - match symbol enumeration
	//
	// - **waf:phase:http_anti_scan:match:fields**: WAF scan protection - match field enumeration
	//
	// - **waf:phase:http_managed:actions**: WAF managed rules - action enumeration
	//
	// - **waf:phase:http_managed:group:reference:ids**: WAF managed rules - referenced rule group enumeration
	//
	// - **waf:phase:http_ratelimit:actions**: WAF rate limiting - action enumeration
	//
	// - **waf:phase:http_ratelimit:ttls**: WAF rate limiting - duration enumeration
	//
	// - **waf:phase:http_ratelimit:intervals**: WAF rate limiting - statistical duration
	//
	// - **waf:phase:http_ratelimit:http_ratelimit:characteristic:fields**: WAF rate limiting - match characteristic enumeration for control type rules
	//
	// - **waf:phase:http_ratelimit:match:symbols**: WAF rate limiting rule phase - match symbol enumeration
	//
	// - **waf:phase:http_ratelimit:match:fields**: WAF rate limiting rule phase - match field enumeration
	//
	// - **waf:phase:http_custom:actions**: WAF custom rule phase - action enumeration
	//
	// - **waf:phase:http_custom:match:symbols**: WAF custom rule phase - match symbol enumeration
	//
	// - **waf:phase:http_custom:match:fields**: WAF custom rule phase - match fields
	//
	// - **waitingroom|wr_queueing_method**: Waiting room - queuing method
	//
	// - **origin_rules|origin_scheme**: Back-to-origin rules - back-to-origin protocol
	//
	// - **origin_rules|origin_sni**: Back-to-origin rules - back-to-origin SNI
	//
	// - **origin_rules|origin_host**: Back-to-origin rules - back-to-origin host
	//
	// - **fourlayeracceleration**: Layer 4 acceleration
	//
	// - **rtlog_service**: Real-time log feature switch
	//
	// - **dashboard_traffic**: Value-added capability for network traffic analysis
	//
	// - **custom_name_server**: Custom NS name
	//
	// - **waf:phase:http_bot:enable**: WAF bot management switch
	//
	// - **waf:phase:http_whitelist:enable**: WAF whitelist switch
	//
	// - **instantlog**: Whether instant logs are active
	//
	// - **waf:phase:http_anti_scan:enable**: WAF scan protection switch
	//
	// - **waf:phase:http_managed:group:reference:enable**: WAF managed rules - configure referenced rule group switch
	//
	// - **waf:phase:http_managed:enable**: WAF managed rules switch
	//
	// - **waf:phase:http_ratelimit:on_hit:enable**: WAF rate limiting - apply to cached requests switch
	//
	// - **ddos**: DDoS instance
	//
	// - **waf:phase:http_ratelimit:enable**: WAF rate limiting rule phase switch
	//
	// - **waf:phase:http_custom:enable**: WAF custom rule phase switch
	//
	// - **waf:phase:all:page:reference:enable**: WAF custom response page switch
	//
	// - **rules_support_regex**: Whether the rule DPI engine supports regular expression operations
	//
	// - **waitingroom_event**: Waiting room - scheduled events
	//
	// - **waitingroom_rule**: Waiting room - bypass waiting room
	//
	// - **waitingroom|wr_json_response**: Waiting room - enable JSON response
	//
	// - **waitingroom|wr_disable_session_renewal**: Waiting room - disable session renewal
	//
	// - **origin_rules|dns_record**: Back-to-origin rules - DNS record
	//
	// - **managed_transforms|add_client_geolocation_headers**: Whether the real client IP header is active in transformation rules
	//
	// - **tiered_cache|regional_enable**: Area cache
	//
	// - **real_client_ip_header**: Client IP header
	//
	// example:
	//
	// waf:phase:http_anti_scan:actions
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The threshold details of the quota.
	QuotaValueDetail []*ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail `json:"QuotaValueDetail,omitempty" xml:"QuotaValueDetail,omitempty" type:"Repeated"`
	// The threshold type of the quota. Valid values:
	//
	// - value: enumeration type. A quota of the enumeration type has multiple enumeration thresholds. For each enumeration threshold, the response indicates whether it is available in the current plan version. If it is unavailable, the response indicates the minimum required plan version.
	//
	// - bool: Boolean type. A quota of the Boolean type is abstracted into two enumeration thresholds: true and false. To check whether the quota is available in the current plan version and identify the minimum required plan version, you only need to focus on the quota details where the enumeration value is true.
	//
	// example:
	//
	// value
	QuotaValueType *string `json:"QuotaValueType,omitempty" xml:"QuotaValueType,omitempty"`
}

func (s ListQuotasMinEffectPlanResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasMinEffectPlanResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) GetQuotaValueDetail() []*ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail {
	return s.QuotaValueDetail
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) GetQuotaValueType() *string {
	return s.QuotaValueType
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) SetQuotaName(v string) *ListQuotasMinEffectPlanResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) SetQuotaValueDetail(v []*ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) *ListQuotasMinEffectPlanResponseBodyQuotas {
	s.QuotaValueDetail = v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) SetQuotaValueType(v string) *ListQuotasMinEffectPlanResponseBodyQuotas {
	s.QuotaValueType = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotas) Validate() error {
	if s.QuotaValueDetail != nil {
		for _, item := range s.QuotaValueDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail struct {
	// Specifies whether the quota value is available in the plan version corresponding to the current site. Valid values:
	//
	// - ture: Available.
	//
	// - false: Unavailable.
	//
	// example:
	//
	// true
	IsEffect *string `json:"IsEffect,omitempty" xml:"IsEffect,omitempty"`
	// The minimum required plan version for the quota value. Valid values:
	//
	// - basic: Basic Edition.
	//
	// - medium: Standard Edition.
	//
	// - high: Advanced Edition.
	//
	// - enterprise_standard_cn: Enterprise Edition.
	//
	// The return value of this parameter follows these rules:
	//
	// - If the quota value is already available in the current plan version, this parameter is left empty, meaning the minimum required plan version is not displayed.
	//
	// - If the quota value is unavailable in the current plan version, the minimum required plan version is displayed.
	//
	// - If the current plan is already the Enterprise Edition, this parameter is always empty.
	//
	// example:
	//
	// basic
	MinEffectPlan *string `json:"MinEffectPlan,omitempty" xml:"MinEffectPlan,omitempty"`
	// The quota value.
	//
	// example:
	//
	// quotaValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) GoString() string {
	return s.String()
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) GetIsEffect() *string {
	return s.IsEffect
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) GetMinEffectPlan() *string {
	return s.MinEffectPlan
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) GetValue() *string {
	return s.Value
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) SetIsEffect(v string) *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail {
	s.IsEffect = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) SetMinEffectPlan(v string) *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail {
	s.MinEffectPlan = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) SetValue(v string) *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail {
	s.Value = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail) Validate() error {
	return dara.Validate(s)
}
