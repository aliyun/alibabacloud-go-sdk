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
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleLabels []*DescribeBotRuleLabelsResponseBodyRuleLabels `json:"RuleLabels,omitempty" xml:"RuleLabels,omitempty" type:"Repeated"`
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
	// example:
	//
	// malicious
	BotBehavior *string `json:"BotBehavior,omitempty" xml:"BotBehavior,omitempty"`
	// example:
	//
	// malicious_crawler_python
	LabelKey *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	// example:
	//
	// human_machine_challenge
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
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

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) GetLabelKey() *string {
	return s.LabelKey
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

func (s *DescribeBotRuleLabelsResponseBodyRuleLabels) SetLabelKey(v string) *DescribeBotRuleLabelsResponseBodyRuleLabels {
	s.LabelKey = &v
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
