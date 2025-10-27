// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAegisContainerPluginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListAegisContainerPluginRuleResponseBodyPageInfo) *ListAegisContainerPluginRuleResponseBody
	GetPageInfo() *ListAegisContainerPluginRuleResponseBodyPageInfo
	SetRequestId(v string) *ListAegisContainerPluginRuleResponseBody
	GetRequestId() *string
	SetRuleList(v []*ListAegisContainerPluginRuleResponseBodyRuleList) *ListAegisContainerPluginRuleResponseBody
	GetRuleList() []*ListAegisContainerPluginRuleResponseBodyRuleList
}

type ListAegisContainerPluginRuleResponseBody struct {
	// The pagination information.
	PageInfo *ListAegisContainerPluginRuleResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rules.
	RuleList []*ListAegisContainerPluginRuleResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s ListAegisContainerPluginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAegisContainerPluginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListAegisContainerPluginRuleResponseBody) GetPageInfo() *ListAegisContainerPluginRuleResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAegisContainerPluginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAegisContainerPluginRuleResponseBody) GetRuleList() []*ListAegisContainerPluginRuleResponseBodyRuleList {
	return s.RuleList
}

func (s *ListAegisContainerPluginRuleResponseBody) SetPageInfo(v *ListAegisContainerPluginRuleResponseBodyPageInfo) *ListAegisContainerPluginRuleResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBody) SetRequestId(v string) *ListAegisContainerPluginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBody) SetRuleList(v []*ListAegisContainerPluginRuleResponseBodyRuleList) *ListAegisContainerPluginRuleResponseBody {
	s.RuleList = v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
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

type ListAegisContainerPluginRuleResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAegisContainerPluginRuleResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAegisContainerPluginRuleResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) SetCurrentPage(v int32) *ListAegisContainerPluginRuleResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) SetPageSize(v int32) *ListAegisContainerPluginRuleResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) SetTotalCount(v int32) *ListAegisContainerPluginRuleResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListAegisContainerPluginRuleResponseBodyRuleList struct {
	// The time when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1676355025000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the rule was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1681985833000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The action of the rule. Valid values:
	//
	// 	- **1**: Alert
	//
	// 	- **2**: Block
	//
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// An array that consists of policies.
	Policies []*ListAegisContainerPluginRuleResponseBodyRuleListPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// system_call
	RuleDescription *string `json:"RuleDescription,omitempty" xml:"RuleDescription,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 30****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-18****
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the rule template. The ListSystemClientRules operation returns the ID of a rule template.
	//
	// example:
	//
	// 868**
	RuleTemplateId *string `json:"RuleTemplateId,omitempty" xml:"RuleTemplateId,omitempty"`
	// The name of the rule template.
	//
	// example:
	//
	// system_call
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" xml:"RuleTemplateName,omitempty"`
	// The fields in the value of the rule subtype.
	SelectedPolicy []*string `json:"SelectedPolicy,omitempty" xml:"SelectedPolicy,omitempty" type:"Repeated"`
	// The switch ID of the rule.
	//
	// example:
	//
	// USER-ENABLE-SWITCH-TYPE_****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The images that are added to the whitelist of the rule.
	WhiteImages []*string `json:"WhiteImages,omitempty" xml:"WhiteImages,omitempty" type:"Repeated"`
}

func (s ListAegisContainerPluginRuleResponseBodyRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListAegisContainerPluginRuleResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetMode() *int32 {
	return s.Mode
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetPolicies() []*ListAegisContainerPluginRuleResponseBodyRuleListPolicies {
	return s.Policies
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetRuleTemplateId() *string {
	return s.RuleTemplateId
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetRuleTemplateName() *string {
	return s.RuleTemplateName
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetSelectedPolicy() []*string {
	return s.SelectedPolicy
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) GetWhiteImages() []*string {
	return s.WhiteImages
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetGmtCreate(v int64) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.GmtCreate = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetGmtModified(v int64) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.GmtModified = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetMode(v int32) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.Mode = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetPolicies(v []*ListAegisContainerPluginRuleResponseBodyRuleListPolicies) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.Policies = v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetRuleDescription(v string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.RuleDescription = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetRuleId(v int64) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.RuleId = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetRuleName(v string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.RuleName = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetRuleTemplateId(v string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.RuleTemplateId = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetRuleTemplateName(v string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.RuleTemplateName = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetSelectedPolicy(v []*string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.SelectedPolicy = v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetSwitchId(v string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.SwitchId = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) SetWhiteImages(v []*string) *ListAegisContainerPluginRuleResponseBodyRuleList {
	s.WhiteImages = v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleList) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAegisContainerPluginRuleResponseBodyRuleListPolicies struct {
	// The policy key.
	//
	// example:
	//
	// system_auto_****
	PolicyKey *string `json:"PolicyKey,omitempty" xml:"PolicyKey,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// System self-starting task\\*\\*\\*\\*
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s ListAegisContainerPluginRuleResponseBodyRuleListPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListAegisContainerPluginRuleResponseBodyRuleListPolicies) GoString() string {
	return s.String()
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleListPolicies) GetPolicyKey() *string {
	return s.PolicyKey
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleListPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleListPolicies) SetPolicyKey(v string) *ListAegisContainerPluginRuleResponseBodyRuleListPolicies {
	s.PolicyKey = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleListPolicies) SetPolicyName(v string) *ListAegisContainerPluginRuleResponseBodyRuleListPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponseBodyRuleListPolicies) Validate() error {
	return dara.Validate(s)
}
