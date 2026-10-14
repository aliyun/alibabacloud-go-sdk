// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebCCRuleV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigWebCCRuleV2Request
	GetDomain() *string
	SetExpires(v int64) *ConfigWebCCRuleV2Request
	GetExpires() *int64
	SetRuleList(v string) *ConfigWebCCRuleV2Request
	GetRuleList() *string
}

type ConfigWebCCRuleV2Request struct {
	// The domain name of the website service.
	//
	// > The domain name must have a website service forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The validity period of the rule. Unit: seconds. Default value: 0, which indicates that the rule is permanently effective.
	//
	// example:
	//
	// 0
	Expires *int64 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The configuration of the frequency control rule, expressed as a JSON-formatted string with the following structure:
	//
	// - **action**: String type. Required. The matching action. Valid values:
	//
	//     - **block**: Block.
	//
	//     - **challenge**: JavaScript Challenge.
	//
	//     - **watch**: Monitor.
	//
	// - **name**: String type. Required. The rule name.
	//
	// - **condition**: Map type. Required. The matching condition. The structure is as follows:
	//
	//     > If multiple matching conditions are specified, the conditions are evaluated using the AND logic.
	//
	//     - **field**: String type. Required. The matching field.
	//
	//     - **match_method**: String type. Required. The matching method.
	//
	//
	//
	//         > For the valid values of **field*	- and **match_method**, see the supplementary description below the request parameters table.
	//
	//     - **header_name**: String type. Optional. The name of the custom HTTP header field.
	//
	//         > This parameter takes effect only when **field*	- is set to **header**.
	//
	//     - **content**: String type. Required. The matching content.
	//
	// - **ratelimit**: String type. Required. The frequency control field. You can collect frequency statistics by IP address or custom header. The structure is as follows:
	//
	//   - **interval**: Integer type. Required. The statistical period (seconds).
	//
	//   - **ttl**: Integer type. Required. The action duration (seconds).
	//
	//   - **threshold**: Integer type. Required. The threshold (count).
	//
	//   - **subkey**: String type. Optional. The field name. Set this parameter only when the statistical source is header.
	//
	//   - **target**: String type. Required. The statistical source. Valid values: ip and header.
	//
	// - **status_code**: String type. Optional. The frequency control field. You can collect statistics based on the count or ratio of status codes. The structure is as follows:
	//
	//   - **enabled**: Boolean type. Required. Specifies whether to enable status code statistics.
	//
	//   - **code**: Integer type. Required. Enabling status code. Valid values: **100*	- to **599**.
	//
	//   - **use_ratio**: Boolean type. Required. Specifies whether to use the ratio. If set to true, the ratio is used.
	//
	//   - **ratio_threshold**: Integer type. Optional. The ratio at which enabling status code appears. When the ratio is used, the action is triggered only when the corresponding status code reaches **ratio_threshold**. Valid values: **1*	- to **100**.
	//
	//   - **count_threshold**: Integer type. Optional. The count at which enabling status code appears. When the ratio is not used, the action is triggered only when the corresponding status code reaches **count_threshold**. Valid values: **2*	- to **50000**.
	//
	// - **statistics**: String type. Optional. The deduplication statistics. This parameter can be left empty, which indicates that deduplication is not performed. The structure is as follows:
	//
	//   - **mode**: String type. Required. Specifies whether to enable deduplication statistics. Valid values:
	//
	//     - **count**: deduplication disabled
	//
	//     - **distinct**: deduplication enabled
	//
	//   - **field**: String type. Required. The statistical source for deduplication. Valid values: ip, header, and uri.
	//
	//   - **header_name**: String type. Optional. The field name. Set this parameter only when the statistical source is header.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"action":"block","name":"trdsss","statistics":{"mode":"count","field":"header","header_name":"hahah"},"ratelimit":{"interval":60,"ttl":300,"threshold":70,"target":"ip"},"status_code":{"enabled":true,"code":404,"use_ratio":true,"ratio_threshold":80,"count_threshold":2},"condition":[{"field":"ip","match_method":"belong","content":"1.1.1.1"}]}]
	RuleList *string `json:"RuleList,omitempty" xml:"RuleList,omitempty"`
}

func (s ConfigWebCCRuleV2Request) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebCCRuleV2Request) GoString() string {
	return s.String()
}

func (s *ConfigWebCCRuleV2Request) GetDomain() *string {
	return s.Domain
}

func (s *ConfigWebCCRuleV2Request) GetExpires() *int64 {
	return s.Expires
}

func (s *ConfigWebCCRuleV2Request) GetRuleList() *string {
	return s.RuleList
}

func (s *ConfigWebCCRuleV2Request) SetDomain(v string) *ConfigWebCCRuleV2Request {
	s.Domain = &v
	return s
}

func (s *ConfigWebCCRuleV2Request) SetExpires(v int64) *ConfigWebCCRuleV2Request {
	s.Expires = &v
	return s
}

func (s *ConfigWebCCRuleV2Request) SetRuleList(v string) *ConfigWebCCRuleV2Request {
	s.RuleList = &v
	return s
}

func (s *ConfigWebCCRuleV2Request) Validate() error {
	return dara.Validate(s)
}
