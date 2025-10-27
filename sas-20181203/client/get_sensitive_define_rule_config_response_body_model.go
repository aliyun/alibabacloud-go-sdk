// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveDefineRuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSensitiveDefineRuleConfigResponseBody
	GetCode() *string
	SetData(v *GetSensitiveDefineRuleConfigResponseBodyData) *GetSensitiveDefineRuleConfigResponseBody
	GetData() *GetSensitiveDefineRuleConfigResponseBodyData
	SetMessage(v string) *GetSensitiveDefineRuleConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSensitiveDefineRuleConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSensitiveDefineRuleConfigResponseBody
	GetSuccess() *bool
}

type GetSensitiveDefineRuleConfigResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetSensitiveDefineRuleConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSensitiveDefineRuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDefineRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSensitiveDefineRuleConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSensitiveDefineRuleConfigResponseBody) GetData() *GetSensitiveDefineRuleConfigResponseBodyData {
	return s.Data
}

func (s *GetSensitiveDefineRuleConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSensitiveDefineRuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSensitiveDefineRuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSensitiveDefineRuleConfigResponseBody) SetCode(v string) *GetSensitiveDefineRuleConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBody) SetData(v *GetSensitiveDefineRuleConfigResponseBodyData) *GetSensitiveDefineRuleConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBody) SetMessage(v string) *GetSensitiveDefineRuleConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBody) SetRequestId(v string) *GetSensitiveDefineRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBody) SetSuccess(v bool) *GetSensitiveDefineRuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSensitiveDefineRuleConfigResponseBodyData struct {
	// Indicates whether the new rule is enabled for automatic check only on agentless detection. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 0
	EnableNewRule *int32 `json:"EnableNewRule,omitempty" xml:"EnableNewRule,omitempty"`
	// The custom configuration ID.
	//
	// example:
	//
	// 44616
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The total number of check rules.
	//
	// example:
	//
	// 100
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The tree of the check rules.
	RuleTree []*GetSensitiveDefineRuleConfigResponseBodyDataRuleTree `json:"RuleTree,omitempty" xml:"RuleTree,omitempty" type:"Repeated"`
	// The number of selected check rules.
	//
	// example:
	//
	// 99
	SelectedCount *int32 `json:"SelectedCount,omitempty" xml:"SelectedCount,omitempty"`
}

func (s GetSensitiveDefineRuleConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDefineRuleConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) GetEnableNewRule() *int32 {
	return s.EnableNewRule
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) GetRuleTree() []*GetSensitiveDefineRuleConfigResponseBodyDataRuleTree {
	return s.RuleTree
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) GetSelectedCount() *int32 {
	return s.SelectedCount
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) SetEnableNewRule(v int32) *GetSensitiveDefineRuleConfigResponseBodyData {
	s.EnableNewRule = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) SetId(v int64) *GetSensitiveDefineRuleConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) SetRuleCount(v int32) *GetSensitiveDefineRuleConfigResponseBodyData {
	s.RuleCount = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) SetRuleTree(v []*GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) *GetSensitiveDefineRuleConfigResponseBodyData {
	s.RuleTree = v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) SetSelectedCount(v int32) *GetSensitiveDefineRuleConfigResponseBodyData {
	s.SelectedCount = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyData) Validate() error {
	if s.RuleTree != nil {
		for _, item := range s.RuleTree {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSensitiveDefineRuleConfigResponseBodyDataRuleTree struct {
	// The category keyword of the check rule.
	//
	// example:
	//
	// password
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The category name of the check rule.
	//
	// example:
	//
	// password
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The list of check rules.
	RuleList []*GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) GoString() string {
	return s.String()
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) GetClassKey() *string {
	return s.ClassKey
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) GetClassName() *string {
	return s.ClassName
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) GetRuleList() []*GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	return s.RuleList
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) SetClassKey(v string) *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree {
	s.ClassKey = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) SetClassName(v string) *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree {
	s.ClassName = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) SetRuleList(v []*GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree {
	s.RuleList = v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTree) Validate() error {
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

type GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList struct {
	// The keyword of the check rule.
	//
	// example:
	//
	// huaweicloud_ak
	RuleKey *string `json:"RuleKey,omitempty" xml:"RuleKey,omitempty"`
	// The name of the check rule.
	//
	// example:
	//
	// huaweicloud_ak
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the check rule is selected. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) GoString() string {
	return s.String()
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) GetRuleKey() *string {
	return s.RuleKey
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) GetSelected() *bool {
	return s.Selected
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) SetRuleKey(v string) *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	s.RuleKey = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) SetRuleName(v string) *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	s.RuleName = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) SetSelected(v bool) *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	s.Selected = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponseBodyDataRuleTreeRuleList) Validate() error {
	return dara.Validate(s)
}
