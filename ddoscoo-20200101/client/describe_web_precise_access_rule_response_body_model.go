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
	// The precise access control rules for the website.
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
	if s.PreciseAccessConfigList != nil {
		for _, item := range s.PreciseAccessConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList struct {
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The list of rules.
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
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList struct {
	// The action that is performed on requests that match the rule. Valid values:
	//
	// - **accept**: allows the requests.
	//
	// - **block**: blocks the requests.
	//
	// - **challenge**: triggers a JavaScript Challenge for the requests.
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The list of match conditions.
	ConditionList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The validity period of the rule. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1598889600
	ExpirePeriod *int64 `json:"ExpirePeriod,omitempty" xml:"ExpirePeriod,omitempty"`
	// The validity period of the rule. Unit: seconds. This parameter is required only when the **action*	- of the rule is **block**. Requests that match the rule are blocked within the validity period. A value of **0*	- indicates that the rule is permanently valid.
	//
	// example:
	//
	// 0
	Expires *int64 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// testrule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the rule. Valid values:
	//
	// - **manual*	- (default): The rule is manually created.
	//
	// - **clover**: The rule is automatically generated.
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

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GetExpirePeriod() *int64 {
	return s.ExpirePeriod
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

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetExpirePeriod(v int64) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.ExpirePeriod = &v
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
	if s.ConditionList != nil {
		for _, item := range s.ConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList struct {
	// The match content.
	//
	// example:
	//
	// 1.1.1.1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The match conditions of the rule. The value is a JSON string. The value contains the following fields:
	//
	// > If you specify multiple match conditions, the logical relation between these conditions is AND.
	//
	// - **field**: The match field. This parameter is of the string type.
	//
	// - **match_method**: The match method. This parameter is of the string type.
	//
	//   > For more information about the valid values of **field*	- and **match_method**, see the "Additional information about field and match_method" section in this topic.
	//
	// - **content**: The match content. This parameter is of the string type.
	//
	// - **header_name**: The name of the header field. This parameter is of the string type. This parameter is required only when **field*	- is set to **header**.
	ContentList []*string `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// The match field.
	//
	// example:
	//
	// ip
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the custom HTTP header field.
	//
	// > This parameter is required only when **Field*	- is set to **header**.
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
