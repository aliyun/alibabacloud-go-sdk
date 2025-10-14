// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMetricRuleTemplateListResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMetricRuleTemplateListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricRuleTemplateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricRuleTemplateListResponseBody
	GetSuccess() *bool
	SetTemplates(v *DescribeMetricRuleTemplateListResponseBodyTemplates) *DescribeMetricRuleTemplateListResponseBody
	GetTemplates() *DescribeMetricRuleTemplateListResponseBodyTemplates
	SetTotal(v int64) *DescribeMetricRuleTemplateListResponseBody
	GetTotal() *int64
}

type DescribeMetricRuleTemplateListResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 659401C0-6214-5C02-972A-CFA929D717B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The queried alert templates.
	Templates *DescribeMetricRuleTemplateListResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMetricRuleTemplateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleTemplateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleTemplateListResponseBody) GetTemplates() *DescribeMetricRuleTemplateListResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeMetricRuleTemplateListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetCode(v int32) *DescribeMetricRuleTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetMessage(v string) *DescribeMetricRuleTemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetRequestId(v string) *DescribeMetricRuleTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetSuccess(v bool) *DescribeMetricRuleTemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetTemplates(v *DescribeMetricRuleTemplateListResponseBodyTemplates) *DescribeMetricRuleTemplateListResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetTotal(v int64) *DescribeMetricRuleTemplateListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) Validate() error {
	if s.Templates != nil {
		if err := s.Templates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleTemplateListResponseBodyTemplates struct {
	Template []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplates) GetTemplate() []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	return s.Template
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplates) SetTemplate(v []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) *DescribeMetricRuleTemplateListResponseBodyTemplates {
	s.Template = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplates) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate struct {
	// The history of applying the alert templates to application groups.
	ApplyHistories *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories `json:"ApplyHistories,omitempty" xml:"ApplyHistories,omitempty" type:"Struct"`
	// The description of the alert template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the alert template was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1646018798000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the alert template was modified.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1646054798000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the alert template.
	//
	// example:
	//
	// ECS_Template1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the alert template.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	RestVersion *int64 `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	// The ID of the alert template.
	//
	// example:
	//
	// 70****
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetApplyHistories() *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories {
	return s.ApplyHistories
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetDescription() *string {
	return s.Description
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetName() *string {
	return s.Name
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetRestVersion() *int64 {
	return s.RestVersion
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetApplyHistories(v *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.ApplyHistories = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetDescription(v string) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.Description = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetGmtCreate(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetGmtModified(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.GmtModified = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetName(v string) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetRestVersion(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.RestVersion = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetTemplateId(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.TemplateId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) Validate() error {
	if s.ApplyHistories != nil {
		if err := s.ApplyHistories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories struct {
	ApplyHistory []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory `json:"ApplyHistory,omitempty" xml:"ApplyHistory,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) GetApplyHistory() []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	return s.ApplyHistory
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) SetApplyHistory(v []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories {
	s.ApplyHistory = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) Validate() error {
	if s.ApplyHistory != nil {
		for _, item := range s.ApplyHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory struct {
	// The timestamp when the alert template was applied to the application group.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1646055898000
	ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 3607****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// ECS_Group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) SetApplyTime(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	s.ApplyTime = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) SetGroupId(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) SetGroupName(v string) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	s.GroupName = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) Validate() error {
	return dara.Validate(s)
}
