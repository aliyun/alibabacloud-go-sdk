// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyMetricRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ApplyMetricRuleTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *ApplyMetricRuleTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyMetricRuleTemplateResponseBody
	GetRequestId() *string
	SetResource(v *ApplyMetricRuleTemplateResponseBodyResource) *ApplyMetricRuleTemplateResponseBody
	GetResource() *ApplyMetricRuleTemplateResponseBodyResource
	SetSuccess(v bool) *ApplyMetricRuleTemplateResponseBody
	GetSuccess() *bool
}

type ApplyMetricRuleTemplateResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F897F3C-020A-4993-95B4-63ABB84F83E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that are affected by the alert rule.
	Resource *ApplyMetricRuleTemplateResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
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
}

func (s ApplyMetricRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ApplyMetricRuleTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyMetricRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyMetricRuleTemplateResponseBody) GetResource() *ApplyMetricRuleTemplateResponseBodyResource {
	return s.Resource
}

func (s *ApplyMetricRuleTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyMetricRuleTemplateResponseBody) SetCode(v int32) *ApplyMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetMessage(v string) *ApplyMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetRequestId(v string) *ApplyMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetResource(v *ApplyMetricRuleTemplateResponseBodyResource) *ApplyMetricRuleTemplateResponseBody {
	s.Resource = v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetSuccess(v bool) *ApplyMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyMetricRuleTemplateResponseBodyResource struct {
	// The details of the generated alert rule.
	AlertResults []*ApplyMetricRuleTemplateResponseBodyResourceAlertResults `json:"AlertResults,omitempty" xml:"AlertResults,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ApplyMetricRuleTemplateResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponseBodyResource) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) GetAlertResults() []*ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	return s.AlertResults
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) SetAlertResults(v []*ApplyMetricRuleTemplateResponseBodyResourceAlertResults) *ApplyMetricRuleTemplateResponseBodyResource {
	s.AlertResults = v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) SetGroupId(v int64) *ApplyMetricRuleTemplateResponseBodyResource {
	s.GroupId = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) Validate() error {
	return dara.Validate(s)
}

type ApplyMetricRuleTemplateResponseBodyResourceAlertResults struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// alert rule is creating, please wait a few minutes.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// applyTemplate8ab74c6b-9f27-47ab-8841-de01dc08****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
}

func (s ApplyMetricRuleTemplateResponseBodyResourceAlertResults) String() string {
	return dara.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GetCode() *string {
	return s.Code
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GetMessage() *string {
	return s.Message
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GetRuleId() *string {
	return s.RuleId
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GetRuleName() *string {
	return s.RuleName
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetCode(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.Code = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetMessage(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.Message = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetRuleId(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.RuleId = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetRuleName(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.RuleName = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetSuccess(v bool) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.Success = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) Validate() error {
	return dara.Validate(s)
}
