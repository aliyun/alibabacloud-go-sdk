// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBuildRiskDefineRuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBuildRiskDefineRuleConfigResponseBody
	GetCode() *string
	SetData(v *GetBuildRiskDefineRuleConfigResponseBodyData) *GetBuildRiskDefineRuleConfigResponseBody
	GetData() *GetBuildRiskDefineRuleConfigResponseBodyData
	SetMessage(v string) *GetBuildRiskDefineRuleConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBuildRiskDefineRuleConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBuildRiskDefineRuleConfigResponseBody
	GetSuccess() *bool
}

type GetBuildRiskDefineRuleConfigResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetBuildRiskDefineRuleConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 69BFFCDE-37D6-5A49-A8BC-BB03AC83****
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

func (s GetBuildRiskDefineRuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBuildRiskDefineRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) GetData() *GetBuildRiskDefineRuleConfigResponseBodyData {
	return s.Data
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) SetCode(v string) *GetBuildRiskDefineRuleConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) SetData(v *GetBuildRiskDefineRuleConfigResponseBodyData) *GetBuildRiskDefineRuleConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) SetMessage(v string) *GetBuildRiskDefineRuleConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) SetRequestId(v string) *GetBuildRiskDefineRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) SetSuccess(v bool) *GetBuildRiskDefineRuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBuildRiskDefineRuleConfigResponseBodyData struct {
	// The configuration ID for scanning image build command risks.
	//
	// example:
	//
	// 273698***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The total number of check items.
	//
	// example:
	//
	// 100
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The details of all check items.
	RuleTree []*GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree `json:"RuleTree,omitempty" xml:"RuleTree,omitempty" type:"Repeated"`
	// The number of selected check items.
	//
	// example:
	//
	// 99
	SelectedCount *int32 `json:"SelectedCount,omitempty" xml:"SelectedCount,omitempty"`
}

func (s GetBuildRiskDefineRuleConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBuildRiskDefineRuleConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) GetRuleTree() []*GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree {
	return s.RuleTree
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) GetSelectedCount() *int32 {
	return s.SelectedCount
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) SetId(v int64) *GetBuildRiskDefineRuleConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) SetRuleCount(v int32) *GetBuildRiskDefineRuleConfigResponseBodyData {
	s.RuleCount = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) SetRuleTree(v []*GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) *GetBuildRiskDefineRuleConfigResponseBodyData {
	s.RuleTree = v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) SetSelectedCount(v int32) *GetBuildRiskDefineRuleConfigResponseBodyData {
	s.SelectedCount = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyData) Validate() error {
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

type GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree struct {
	// The check item type.
	//
	// example:
	//
	// other
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The name of the check item type.
	//
	// example:
	//
	// other
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The check items of the type.
	RuleList []*GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) String() string {
	return dara.Prettify(s)
}

func (s GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) GoString() string {
	return s.String()
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) GetClassKey() *string {
	return s.ClassKey
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) GetClassName() *string {
	return s.ClassName
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) GetRuleList() []*GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	return s.RuleList
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) SetClassKey(v string) *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree {
	s.ClassKey = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) SetClassName(v string) *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree {
	s.ClassName = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) SetRuleList(v []*GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree {
	s.RuleList = v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTree) Validate() error {
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

type GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList struct {
	// The check item.
	//
	// example:
	//
	// add
	RuleKey *string `json:"RuleKey,omitempty" xml:"RuleKey,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// used ADD
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the check item is selected. Valid values:
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

func (s GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) GoString() string {
	return s.String()
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) GetRuleKey() *string {
	return s.RuleKey
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) GetSelected() *bool {
	return s.Selected
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) SetRuleKey(v string) *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	s.RuleKey = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) SetRuleName(v string) *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	s.RuleName = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) SetSelected(v bool) *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList {
	s.Selected = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponseBodyDataRuleTreeRuleList) Validate() error {
	return dara.Validate(s)
}
