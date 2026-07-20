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
	// The list of minimum effective plan editions for quotas.
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
	// The quota name. Valid values:
	//
	// - **waf:phase:http_anti_scan:actions**: WAF scan protection - action enumeration.
	//
	// - **waf:phase:http_bot:actions**: WAF bot management - all action enumeration.
	//
	// - **waf:phase:http_bot:http_custom_cc_dev:characteristic:fields**: WAF bot management - custom device rate limiting statistical object enumeration.
	//
	// - **waf:phase:http_bot:http_custom_cc_ip:characteristic:fields**: WAF bot management - custom IP rate limiting statistical object enumeration.
	//
	// - ****waf:phase:http_bot:match:symbols****: WAF bot management - match operator enumeration.
	//
	// - **waf:phase:http_bot:http_custom_cc:characteristic:fields**: WAF bot management - custom session rate limiting statistical object enumeration.
	//
	// - **waf:phase:http_bot:match:fields**: WAF bot management - match field enumeration.
	//
	// - **waf:phase:http_whitelist:match:symbols**: WAF whitelist - match operator enumeration.
	//
	// - **waf:phase:http_whitelist:match:fields**: WAF whitelist - match field enumeration.
	//
	// - **waf:phase:http_anti_scan:http_directory_traversal:characteristic:fields**: WAF scan protection - folder traverse blocking statistical object enumeration.
	//
	// - **waf:phase:http_anti_scan:http_high_frequency:characteristic:fields**: WAF scan protection - high-frequency scan blocking statistical object enumeration.
	//
	// - **waf:phase:http_anti_scan:match:symbols**: WAF scan protection - match operator enumeration.
	//
	// - **waf:phase:http_anti_scan:match:fields**: WAF scan protection - match field enumeration.
	//
	// - **waf:phase:http_managed:actions**: WAF managed rules - action enumeration.
	//
	// - **waf:phase:http_managed:group:reference:ids**: WAF managed rules - referenced rule group enumeration.
	//
	// - **waf:phase:http_ratelimit:actions**: WAF rate limiting - action enumeration.
	//
	// - **waf:phase:http_ratelimit:ttls**: WAF rate limiting - duration enumeration.
	//
	// - **waf:phase:http_ratelimit:intervals**: WAF rate limiting - statistical period.
	//
	// - **waf:phase:http_ratelimit:http_ratelimit:characteristic:fields**: WAF rate limiting - control type rule match characteristic enumeration.
	//
	// - **waf:phase:http_ratelimit:match:symbols**: WAF rate limiting rule phase match operator enumeration.
	//
	// - **waf:phase:http_ratelimit:match:fields**: WAF rate limiting rule phase match field enumeration.
	//
	// - **waf:phase:http_custom:actions**: WAF custom rule phase action enumeration.
	//
	// - **waf:phase:http_custom:match:symbols**: WAF custom rule phase match operator enumeration.
	//
	// - **waf:phase:http_custom:match:fields**: WAF custom rule phase match field.
	//
	// - **waitingroom|wr_queueing_method**: waiting room - queuing method.
	//
	// - **origin_rules|origin_scheme**: back-to-origin rule - back-to-origin protocol.
	//
	// - **origin_rules|origin_sni**: back-to-origin rule - back-to-origin SNI.
	//
	// - **origin_rules|origin_host**: back-to-origin rule - back-to-origin host.
	//
	// - **fourlayeracceleration**: Layer 4 acceleration.
	//
	// - **rtlog_service**: real-time log feature switch.
	//
	// - **dashboard_traffic**: value-added capability for network traffic analysis.
	//
	// - **custom_name_server**: custom NS name.
	//
	// - **waf:phase:http_bot:enable**: WAF bot management switch.
	//
	// - **waf:phase:http_whitelist:enable**: WAF whitelist switch.
	//
	// - **instantlog**: instant log active or not.
	//
	// - **waf:phase:http_anti_scan:enable**: WAF scan protection switch.
	//
	// - **waf:phase:http_managed:group:reference:enable**: WAF managed rules - referenced rule group configuration switch.
	//
	// - **waf:phase:http_managed:enable**: WAF managed rules switch.
	//
	// - **waf:phase:http_ratelimit:on_hit:enable**: WAF rate limiting - apply to cache-hit requests switch.
	//
	// - **ddos**: DDoS instance.
	//
	// - **waf:phase:http_ratelimit:enable**: WAF rate limiting rule phase switch.
	//
	// - **waf:phase:http_custom:enable**: WAF custom rule phase switch.
	//
	// - **waf:phase:all:page:reference:enable**: WAF custom response page switch.
	//
	// - **rules_support_regex**: whether the DPI engine supports regular expressions.
	//
	// - **waitingroom_event**: waiting room - scheduled event.
	//
	// - **waitingroom_rule**: waiting room - bypass waiting room.
	//
	// - **waitingroom|wr_json_response**: waiting room - enable JSON response.
	//
	// - **waitingroom|wr_disable_session_renewal**: waiting room - disable session renewal.
	//
	// - **origin_rules|dns_record**: back-to-origin rule - DNS record.
	//
	// - **managed_transforms|add_client_geolocation_headers**: whether the real client IP header is active in transform rules.
	//
	// - **tiered_cache|regional_enable**: area cache.
	//
	// - **real_client_ip_header**: client IP header.
	//
	// example:
	//
	// waf:phase:http_anti_scan:actions
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The list of quota threshold details.
	QuotaValueDetail []*ListQuotasMinEffectPlanResponseBodyQuotasQuotaValueDetail `json:"QuotaValueDetail,omitempty" xml:"QuotaValueDetail,omitempty" type:"Repeated"`
	// The threshold type of the quota. Valid values:
	//
	// - value: enumeration type. An enumeration-type quota has multiple enumeration thresholds. For each enumeration threshold, the system indicates whether it is available in the current edition and, if not, the minimum plan edition in which it becomes available.
	//
	// - bool: Boolean type. A Boolean-type quota is abstracted into two enumeration thresholds: true and false. To check whether the quota is available in the current plan edition and the minimum available plan edition, you only need to check the quota details for the true enumeration in the quota details list.
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
	// Indicates whether the quota value is available in the plan edition associated with the current site. Valid values:
	//
	// - true: Available.
	//
	// - false: Not available.
	//
	// example:
	//
	// true
	IsEffect *string `json:"IsEffect,omitempty" xml:"IsEffect,omitempty"`
	// The minimum plan edition in which the quota value is available. Valid values:
	//
	// - basic: Basic Edition.
	//
	// - medium: Standard Edition.
	//
	// - high: Pro Edition.
	//
	// - enterprise_standard_cn: Enterprise Edition.
	//
	// This parameter follows these rules:
	//
	// - If the quota value is already available in the current edition, this parameter is empty, meaning the minimum available plan edition is not displayed.
	//
	// - If the quota value is not available in the current edition, the minimum available edition is displayed.
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
