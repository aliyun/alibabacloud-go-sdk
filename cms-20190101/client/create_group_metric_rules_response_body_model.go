// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMetricRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateGroupMetricRulesResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateGroupMetricRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGroupMetricRulesResponseBody
	GetRequestId() *string
	SetResources(v *CreateGroupMetricRulesResponseBodyResources) *CreateGroupMetricRulesResponseBody
	GetResources() *CreateGroupMetricRulesResponseBodyResources
	SetSuccess(v bool) *CreateGroupMetricRulesResponseBody
	GetSuccess() *bool
}

type CreateGroupMetricRulesResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call is successful.
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
	// The ID of the request.
	//
	// example:
	//
	// 461CF2CD-2FC3-4B26-8645-7BD27E7D0F1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the alert rules.
	Resources *CreateGroupMetricRulesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the call is successful. Valid value:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGroupMetricRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateGroupMetricRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupMetricRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupMetricRulesResponseBody) GetResources() *CreateGroupMetricRulesResponseBodyResources {
	return s.Resources
}

func (s *CreateGroupMetricRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGroupMetricRulesResponseBody) SetCode(v int32) *CreateGroupMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetMessage(v string) *CreateGroupMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetRequestId(v string) *CreateGroupMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetResources(v *CreateGroupMetricRulesResponseBodyResources) *CreateGroupMetricRulesResponseBody {
	s.Resources = v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetSuccess(v bool) *CreateGroupMetricRulesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesResponseBodyResources struct {
	AlertResult []*CreateGroupMetricRulesResponseBodyResourcesAlertResult `json:"AlertResult,omitempty" xml:"AlertResult,omitempty" type:"Repeated"`
}

func (s CreateGroupMetricRulesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponseBodyResources) GetAlertResult() []*CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	return s.AlertResult
}

func (s *CreateGroupMetricRulesResponseBodyResources) SetAlertResult(v []*CreateGroupMetricRulesResponseBodyResourcesAlertResult) *CreateGroupMetricRulesResponseBodyResources {
	s.AlertResult = v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesResponseBodyResourcesAlertResult struct {
	// The status code that is returned for the alert rule.
	//
	// >  The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned for the alert rule.
	//
	// example:
	//
	// Metric not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// 456789
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// ECS_Rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the alert rule was created. Valid value:
	//
	// - true: The alert rule was created.
	//
	// - false: The alert rule failed to be created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGroupMetricRulesResponseBodyResourcesAlertResult) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesResponseBodyResourcesAlertResult) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) GetCode() *int32 {
	return s.Code
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetCode(v int32) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.Code = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetMessage(v string) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.Message = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetRuleId(v string) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.RuleId = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetRuleName(v string) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.RuleName = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetSuccess(v bool) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.Success = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) Validate() error {
	return dara.Validate(s)
}
