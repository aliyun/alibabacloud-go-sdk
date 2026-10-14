// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCCRulesV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeWebCCRulesV2ResponseBody
	GetDomain() *string
	SetRequestId(v string) *DescribeWebCCRulesV2ResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeWebCCRulesV2ResponseBody
	GetTotalCount() *string
	SetWebCCRules(v []*DescribeWebCCRulesV2ResponseBodyWebCCRules) *DescribeWebCCRulesV2ResponseBody
	GetWebCCRules() []*DescribeWebCCRulesV2ResponseBodyWebCCRules
}

type DescribeWebCCRulesV2ResponseBody struct {
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of frequency control rules.
	//
	// example:
	//
	// 12
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The frequency control (HTTP flood mitigation) custom rules.
	WebCCRules []*DescribeWebCCRulesV2ResponseBodyWebCCRules `json:"WebCCRules,omitempty" xml:"WebCCRules,omitempty" type:"Repeated"`
}

func (s DescribeWebCCRulesV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebCCRulesV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebCCRulesV2ResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeWebCCRulesV2ResponseBody) GetWebCCRules() []*DescribeWebCCRulesV2ResponseBodyWebCCRules {
	return s.WebCCRules
}

func (s *DescribeWebCCRulesV2ResponseBody) SetDomain(v string) *DescribeWebCCRulesV2ResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBody) SetRequestId(v string) *DescribeWebCCRulesV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBody) SetTotalCount(v string) *DescribeWebCCRulesV2ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBody) SetWebCCRules(v []*DescribeWebCCRulesV2ResponseBodyWebCCRules) *DescribeWebCCRulesV2ResponseBody {
	s.WebCCRules = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBody) Validate() error {
	if s.WebCCRules != nil {
		for _, item := range s.WebCCRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebCCRulesV2ResponseBodyWebCCRules struct {
	// The validity period of the rule. Unit: seconds. This parameter takes effect when the matching action is set to block (action is block). Access requests are blocked within the validity period. A value of 0 indicates that the rule is permanently effective.
	//
	// example:
	//
	// 0
	Expires *int64 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// wq
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the rule. Valid values:
	//
	// example:
	//
	// manual
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The rule details.
	RuleDetail *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail `json:"RuleDetail,omitempty" xml:"RuleDetail,omitempty" type:"Struct"`
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRules) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) GetExpires() *int64 {
	return s.Expires
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) GetName() *string {
	return s.Name
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) GetOwner() *string {
	return s.Owner
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) GetRuleDetail() *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	return s.RuleDetail
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) SetExpires(v int64) *DescribeWebCCRulesV2ResponseBodyWebCCRules {
	s.Expires = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) SetName(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRules {
	s.Name = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) SetOwner(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRules {
	s.Owner = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) SetRuleDetail(v *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) *DescribeWebCCRulesV2ResponseBodyWebCCRules {
	s.RuleDetail = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRules) Validate() error {
	if s.RuleDetail != nil {
		if err := s.RuleDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail struct {
	// The matching action. Valid values:
	//
	// - **accept**: Allow.
	//
	// - **block**: Block.
	//
	// - **challenge**: JavaScript Challenge.
	//
	// - **watch**: Monitor.
	//
	// example:
	//
	// block
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The list of matching conditions.
	Condition []*DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Repeated"`
	// The legacy rule format. This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The legacy rule format. This parameter is deprecated.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The legacy rule format. This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// ccauto14
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The frequency statistics.
	RateLimit *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit `json:"RateLimit,omitempty" xml:"RateLimit,omitempty" type:"Struct"`
	// The deduplication statistics. This parameter is optional. If this parameter is not specified, deduplication is not performed.
	Statistics *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The status code.
	StatusCode *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode `json:"StatusCode,omitempty" xml:"StatusCode,omitempty" type:"Struct"`
	// The legacy rule format. This parameter is deprecated.
	//
	// example:
	//
	// 300
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The legacy rule format. This parameter is deprecated.
	//
	// example:
	//
	// /p3shijihao
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetAction() *string {
	return s.Action
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetCondition() []*DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition {
	return s.Condition
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetMode() *string {
	return s.Mode
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetName() *string {
	return s.Name
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetRateLimit() *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit {
	return s.RateLimit
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetStatistics() *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics {
	return s.Statistics
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetStatusCode() *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode {
	return s.StatusCode
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) GetUri() *string {
	return s.Uri
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetAction(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Action = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetCondition(v []*DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Condition = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetCount(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Count = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetInterval(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Interval = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetMode(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Mode = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetName(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Name = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetRateLimit(v *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.RateLimit = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetStatistics(v *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Statistics = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetStatusCode(v *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.StatusCode = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetTtl(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Ttl = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) SetUri(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail {
	s.Uri = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetail) Validate() error {
	if s.Condition != nil {
		for _, item := range s.Condition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RateLimit != nil {
		if err := s.RateLimit.Validate(); err != nil {
			return err
		}
	}
	if s.Statistics != nil {
		if err := s.Statistics.Validate(); err != nil {
			return err
		}
	}
	if s.StatusCode != nil {
		if err := s.StatusCode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition struct {
	// The matching content.
	//
	// example:
	//
	// 192.0.XX.XX
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The matching content when the matching method is set to equal to one of multiple values.
	//
	// example:
	//
	// ["2","3","ad"]
	ContentList []*string `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// The matching field.
	//
	// example:
	//
	// ip
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the custom HTTP header field.
	//
	// example:
	//
	// null
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The matching method.
	//
	// example:
	//
	// belong
	MatchMethod *string `json:"MatchMethod,omitempty" xml:"MatchMethod,omitempty"`
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) GetContent() *string {
	return s.Content
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) GetContentList() []*string {
	return s.ContentList
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) GetField() *string {
	return s.Field
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) GetMatchMethod() *string {
	return s.MatchMethod
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) SetContent(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition {
	s.Content = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) SetContentList(v []*string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition {
	s.ContentList = v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) SetField(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition {
	s.Field = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) SetHeaderName(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition {
	s.HeaderName = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) SetMatchMethod(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition {
	s.MatchMethod = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition) Validate() error {
	return dara.Validate(s)
}

type DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit struct {
	// The statistical period. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The field name. This parameter is required only when the statistics source is set to header.
	//
	// example:
	//
	// action
	SubKey *string `json:"SubKey,omitempty" xml:"SubKey,omitempty"`
	// The statistics source. Valid values:
	//
	// example:
	//
	// ip
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The trigger threshold.
	//
	// example:
	//
	// 20
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The blocking duration. Unit: seconds.
	//
	// example:
	//
	// 15
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) GetSubKey() *string {
	return s.SubKey
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) GetTarget() *string {
	return s.Target
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) GetThreshold() *int32 {
	return s.Threshold
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) SetInterval(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit {
	s.Interval = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) SetSubKey(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit {
	s.SubKey = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) SetTarget(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit {
	s.Target = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) SetThreshold(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit {
	s.Threshold = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) SetTtl(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit {
	s.Ttl = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics struct {
	// The statistics source. Valid values:
	//
	// - **ip**: by IP address
	//
	// - **header**: by header
	//
	// - **uri**: by URI
	//
	// example:
	//
	// uri
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// This parameter is required only when the statistics source is set to header.
	//
	// example:
	//
	// hello
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The deduplication mode. Valid values:
	//
	// example:
	//
	// count
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) GetField() *string {
	return s.Field
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) GetMode() *string {
	return s.Mode
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) SetField(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics {
	s.Field = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) SetHeaderName(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics {
	s.HeaderName = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) SetMode(v string) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics {
	s.Mode = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode struct {
	// The status code. Valid values: **100*	- to **599**.
	//
	// - **200**: Successful.
	//
	// - Other values: An error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// When the ratio is not used, the action is triggered only when the corresponding status code reaches **CountThreshold**. Valid range: **2*	- to **50000**.
	//
	// example:
	//
	// 10
	CountThreshold *int32 `json:"CountThreshold,omitempty" xml:"CountThreshold,omitempty"`
	// Indicates whether the feature is enabled. Valid values:
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// When the ratio is used, the action is triggered only when the corresponding status code reaches **RatioThreshold**. Valid range: **1*	- to **100**.
	//
	// example:
	//
	// 10
	RatioThreshold *int32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// Indicates whether to use the ratio:
	//
	// example:
	//
	// true
	UseRatio *bool `json:"UseRatio,omitempty" xml:"UseRatio,omitempty"`
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) GetCode() *int32 {
	return s.Code
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) GetCountThreshold() *int32 {
	return s.CountThreshold
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) GetRatioThreshold() *int32 {
	return s.RatioThreshold
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) GetUseRatio() *bool {
	return s.UseRatio
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) SetCode(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode {
	s.Code = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) SetCountThreshold(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode {
	s.CountThreshold = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) SetEnabled(v bool) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode {
	s.Enabled = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) SetRatioThreshold(v int32) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode {
	s.RatioThreshold = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) SetUseRatio(v bool) *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode {
	s.UseRatio = &v
	return s
}

func (s *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode) Validate() error {
	return dara.Validate(s)
}
