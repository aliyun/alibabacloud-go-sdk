// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebPreciseAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreciseAccessConfigList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) *DescribeWebPreciseAccessRuleResponseBody
	GetPreciseAccessConfigList() []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList
	SetRequestId(v string) *DescribeWebPreciseAccessRuleResponseBody
	GetRequestId() *string
}

type DescribeWebPreciseAccessRuleResponseBody struct {
	// The configuration of the accurate access control rule that is created for the website.
	PreciseAccessConfigList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList `json:"PreciseAccessConfigList,omitempty" xml:"PreciseAccessConfigList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 209EEFBF-B0C7-441E-8C28-D0945A57A638
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebPreciseAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBody) GetPreciseAccessConfigList() []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList {
	return s.PreciseAccessConfigList
}

func (s *DescribeWebPreciseAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebPreciseAccessRuleResponseBody) SetPreciseAccessConfigList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) *DescribeWebPreciseAccessRuleResponseBody {
	s.PreciseAccessConfigList = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBody) SetRequestId(v string) *DescribeWebPreciseAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList struct {
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The scheduling rules.
	RuleList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) GetRuleList() []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	return s.RuleList
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) SetDomain(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList {
	s.Domain = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) SetRuleList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList {
	s.RuleList = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList struct {
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **accept**: The requests that match the rule are allowed.
	//
	// 	- **block**: The requests that match the rule are blocked.
	//
	// 	- **challenge**: Completely Automated Public Turing test to tell Computers and Humans Apart (CAPTCHA) verification for the requests that match the rule is implemented.
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	ConditionList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The validity period of the rule. Unit: seconds. This parameter takes effect only when **action*	- of a rule is **block**. Access requests that match the rule are blocked within the specified validity period of the rule. The value **0*	- indicates that the whitelist takes effect all the time.
	//
	// example:
	//
	// 0
	Expires *int64 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The name of the scheduling rule.
	//
	// example:
	//
	// testrule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the rule. Valid values:
	//
	// 	- **manual*	- (default): manually created.
	//
	// 	- **auto**: automatically generated.
	//
	// example:
	//
	// manual
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GetAction() *string {
	return s.Action
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GetConditionList() []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	return s.ConditionList
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GetExpires() *int64 {
	return s.Expires
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GetOwner() *string {
	return s.Owner
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetAction(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Action = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetConditionList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.ConditionList = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetExpires(v int64) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Expires = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetName(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Name = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetOwner(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Owner = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) Validate() error {
	return dara.Validate(s)
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList struct {
	// The match content.
	//
	// example:
	//
	// 1.1.1.1
	Content     *string   `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// The logical operator.
	//
	// example:
	//
	// belong
	MatchMethod *string `json:"MatchMethod,omitempty" xml:"MatchMethod,omitempty"`
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GetContent() *string {
	return s.Content
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GetContentList() []*string {
	return s.ContentList
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GetField() *string {
	return s.Field
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GetMatchMethod() *string {
	return s.MatchMethod
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetContent(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.Content = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetContentList(v []*string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.ContentList = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetField(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.Field = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetHeaderName(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.HeaderName = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetMatchMethod(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.MatchMethod = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) Validate() error {
	return dara.Validate(s)
}
