// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotRuleLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeBotRuleLabelsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeBotRuleLabelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeBotRuleLabelsResponseBody
	GetRequestId() *string
	SetRuleLabels(v []*DescribeBotRuleLabelsResponseBodyRuleLabels) *DescribeBotRuleLabelsResponseBody
	GetRuleLabels() []*DescribeBotRuleLabelsResponseBodyRuleLabels
	SetTotalCount(v int32) *DescribeBotRuleLabelsResponseBody
	GetTotalCount() *int32
}

type DescribeBotRuleLabelsResponseBody struct {
	// The number of entries per page for paging. Valid values: 1 to 200. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. If a next page exists, this field has a return value.
	//
	// > If this parameter has a return value, a next page exists. You can use the returned **NextToken*	- as a request parameter to obtain the data on the next page. Repeat this process until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of bot management rule tags.
	RuleLabels []*DescribeBotRuleLabelsResponseBodyRuleLabels `json:"RuleLabels,omitempty" xml:"RuleLabels,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBotRuleLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotRuleLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBotRuleLabelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeBotRuleLabelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeBotRuleLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBotRuleLabelsResponseBody) GetRuleLabels() []*DescribeBotRuleLabelsResponseBodyRuleLabels {
	return s.RuleLabels
}

func (s *DescribeBotRuleLabelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBotRuleLabelsResponseBody) SetMaxResults(v int32) *DescribeBotRuleLabelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBody) SetNextToken(v string) *DescribeBotRuleLabelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBody) SetRequestId(v string) *DescribeBotRuleLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBody) SetRuleLabels(v []*DescribeBotRuleLabelsResponseBodyRuleLabels) *DescribeBotRuleLabelsResponseBody {
	s.RuleLabels = v
	return s
}

func (s *DescribeBotRuleLabelsResponseBody) SetTotalCount(v int32) *DescribeBotRuleLabelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBody) Validate() error {
	if s.RuleLabels != nil {
		for _, item := range s.RuleLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBotRuleLabelsResponseBodyRuleLabels struct {
	// The crawler behavior corresponding to the rule tag.
	//
	// - **malicious**: malicious crawler.
	//
	// - **suspicious**: suspected crawler.
	//
	// - **normal**: normal crawler.
	//
	// example:
	//
	// malicious
	BotBehavior *string `json:"BotBehavior,omitempty" xml:"BotBehavior,omitempty"`
	// The default action. Valid values:
	//
	// - **block**: Block.
	//
	// - **monitor**: Monitor.
	//
	// - **js**: JavaScript verification.
	//
	// - **captcha**: Slider CAPTCHA.
	//
	// - **captcha_strict**: Strict slider CAPTCHA.
	//
	// - **bypass**: Allow.
	//
	// example:
	//
	// block
	DefaultAction *string `json:"DefaultAction,omitempty" xml:"DefaultAction,omitempty"`
	// The default configurations corresponding to the label.
	//
	// example:
	//
	// {"crawlerStatusMap":{"360":1,"bytedance":1}}
	DefaultConfig *string `json:"DefaultConfig,omitempty" xml:"DefaultConfig,omitempty"`
	// The default status of the tag rule.
	//
	// - **1**: The rule is enabled.
	//
	// - **0**: The rule is disabled.
	//
	// example:
	//
	// 1
	DefaultStatus *int32 `json:"DefaultStatus,omitempty" xml:"DefaultStatus,omitempty"`
	// The bot management rule tag.
	//
	// example:
	//
	// malicious_crawler_python
	LabelKey *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	// The tag status.
	//
	// - **online**: Online.
	//
	// - **wait_offline**: Pending offline.
	//
	// example:
	//
	// online
	LabelStatus *string `json:"LabelStatus,omitempty" xml:"LabelStatus,omitempty"`
	// The type of the bot rule tag.
	//
	// example:
	//
	// human_machine_challenge
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The set of bot management protection scenarios to which the rule belongs. Multiple scenarios are separated by commas (,). Valid values:
	//
	// - **web**: Web protection scenario.
	//
	// - **app**: App protection scenario.
	//
	// example:
	//
	// Web,app
	SubScene *string `json:"SubScene,omitempty" xml:"SubScene,omitempty"`
}

func (s DescribeBotRuleLabelsResponseBodyRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotRuleLabelsResponseBodyRuleLabels) GoString() string {
	return s.String()
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetBotBehavior() *string {
	return s.BotBehavior
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetDefaultAction() *string {
	return s.DefaultAction
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetDefaultConfig() *string {
	return s.DefaultConfig
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetDefaultStatus() *int32 {
	return s.DefaultStatus
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetLabelKey() *string {
	return s.LabelKey
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetLabelStatus() *string {
	return s.LabelStatus
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetLabelType() *string {
	return s.LabelType
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetSubScene() *string {
	return s.SubScene
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetBotBehavior(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.BotBehavior = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetDefaultAction(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.DefaultAction = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetDefaultConfig(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.DefaultConfig = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetDefaultStatus(v int32) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.DefaultStatus = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetLabelKey(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.LabelKey = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetLabelStatus(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.LabelStatus = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetLabelType(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.LabelType = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetSubScene(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.SubScene = &v
	return s
}

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) Validate() error {
	return dara.Validate(s)
}
