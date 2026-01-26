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
	// The request ID.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned custom frequency control rules.
	//
	// example:
	//
	// 12
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The custom frequency control rules.
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
	// The validity period of the rule. Unit: seconds. If the Action parameter is set to block, the system blocks the requests that match the rule within the validity period of the rule. The value 0 indicates that the rule is permanently valid.
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
	// The method used to create the rule. Valid values:
	//
	// 	- **manual*	- (default): manually created.
	//
	// 	- **clover**: automatically created.
	//
	// example:
	//
	// manual
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The details of the rule.
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
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **accept**: The requests that match the rule are allowed.
	//
	// 	- **block**: The requests that match the rule are blocked.
	//
	// 	- **challenge**: Completely Automated Public Turing test to tell Computers and Humans Apart (CAPTCHA) verification for the requests that match the rule is implemented.
	//
	// 	- **watch**: The requests that match the rule are recorded in logs and allowed.
	//
	// example:
	//
	// block
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	Condition []*DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Repeated"`
	// The parameter is deprecated.
	//
	// example:
	//
	// 废弃
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The parameter is deprecated.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The parameter is deprecated.
	//
	// example:
	//
	// 废弃
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// ccauto14
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The frequency statistics.
	RateLimit *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailRateLimit `json:"RateLimit,omitempty" xml:"RateLimit,omitempty" type:"Struct"`
	// The statistics after deduplication. By default, the system collects statistics before deduplication.
	Statistics *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The status codes.
	StatusCode *DescribeWebCCRulesV2ResponseBodyWebCCRulesRuleDetailStatusCode `json:"StatusCode,omitempty" xml:"StatusCode,omitempty" type:"Struct"`
	// The parameter is deprecated.
	//
	// example:
	//
	// 300
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The parameter is deprecated.
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
	// The match content.
	//
	// example:
	//
	// 192.0.XX.XX
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The match content when the match method is Equals to One of Multiple Values.
	//
	// example:
	//
	// ["2","3","ad"]
	ContentList []*string `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// The match field.
	//
	// example:
	//
	// ip
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The custom HTTP request header.
	//
	// >  This parameter takes effect only when **Field*	- is set to **header**.
	//
	// example:
	//
	// null
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The match method.
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
	// The name of the field. This parameter is required only when the Target parameter is set to header.
	//
	// example:
	//
	// action
	SubKey *string `json:"SubKey,omitempty" xml:"SubKey,omitempty"`
	// The statistical method. Valid values:
	//
	// 	- **ip**
	//
	// 	- **header**
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
	// The statistical method. Valid values:
	//
	// 	- **ip**
	//
	// 	- **header**
	//
	// 	- **uri**
	//
	// example:
	//
	// uri
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the header. This parameter is required only when the Field parameter is set to header.
	//
	// example:
	//
	// hello
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// Indicates whether the system collects statistics after deduplication. Valid values:
	//
	// 	- **count**: The system collects statistics before deduplication.
	//
	// 	- **distinct**: The system collects statistics after deduplication.
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
	// 	- **200**: The request was successful.
	//
	// 	- Other codes: The request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// If a ratio is not used, the handling action is triggered only when the number of requests of the corresponding status code reaches the value of **CountThreshold**. Valid values: **2*	- to **50000**.
	//
	// example:
	//
	// 10
	CountThreshold *int32 `json:"CountThreshold,omitempty" xml:"CountThreshold,omitempty"`
	// Indicates whether the status code is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// If a ratio is used, the handling action is triggered only when the number of requests of the corresponding status code reaches the value of **RatioThreshold**. Valid values: **1*	- to **100**.
	//
	// example:
	//
	// 10
	RatioThreshold *int32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// Indicates whether to use a ratio.
	//
	// 	- **true**
	//
	// 	- **false**
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
