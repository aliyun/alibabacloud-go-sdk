// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertingMetricRuleResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertingMetricRuleResourcesResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeAlertingMetricRuleResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertingMetricRuleResourcesResponseBody
	GetRequestId() *string
	SetResources(v *DescribeAlertingMetricRuleResourcesResponseBodyResources) *DescribeAlertingMetricRuleResourcesResponseBody
	GetResources() *DescribeAlertingMetricRuleResourcesResponseBodyResources
	SetSuccess(v bool) *DescribeAlertingMetricRuleResourcesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeAlertingMetricRuleResourcesResponseBody
	GetTotal() *int32
}

type DescribeAlertingMetricRuleResourcesResponseBody struct {
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
	// The request processing has failed due to some unknown error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0724011B-D9E0-4B2F-8C51-F17A894CC42C
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources *DescribeAlertingMetricRuleResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) GetResources() *DescribeAlertingMetricRuleResourcesResponseBodyResources {
	return s.Resources
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) SetCode(v int32) *DescribeAlertingMetricRuleResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) SetMessage(v string) *DescribeAlertingMetricRuleResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) SetRequestId(v string) *DescribeAlertingMetricRuleResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) SetResources(v *DescribeAlertingMetricRuleResourcesResponseBodyResources) *DescribeAlertingMetricRuleResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) SetSuccess(v bool) *DescribeAlertingMetricRuleResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) SetTotal(v int32) *DescribeAlertingMetricRuleResourcesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlertingMetricRuleResourcesResponseBodyResources struct {
	Resource []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResources) GetResource() []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResources) SetResource(v []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) *DescribeAlertingMetricRuleResourcesResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource struct {
	Dimensions      *string                                                                     `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Enable          *string                                                                     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Escalation      *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation `json:"Escalation,omitempty" xml:"Escalation,omitempty" type:"Struct"`
	GroupId         *string                                                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	LastAlertTime   *string                                                                     `json:"LastAlertTime,omitempty" xml:"LastAlertTime,omitempty"`
	LastModifyTime  *string                                                                     `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	Level           *int32                                                                      `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName      *string                                                                     `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricValues    *string                                                                     `json:"MetricValues,omitempty" xml:"MetricValues,omitempty"`
	Namespace       *string                                                                     `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProductCategory *string                                                                     `json:"ProductCategory,omitempty" xml:"ProductCategory,omitempty"`
	Resource        *string                                                                     `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RetryTimes      *string                                                                     `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	RuleId          *string                                                                     `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string                                                                     `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	StartTime       *string                                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Statistics      *string                                                                     `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold       *string                                                                     `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetEnable() *string {
	return s.Enable
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetEscalation() *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation {
	return s.Escalation
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetLastAlertTime() *string {
	return s.LastAlertTime
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetLevel() *int32 {
	return s.Level
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetMetricValues() *string {
	return s.MetricValues
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetResource() *string {
	return s.Resource
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetRetryTimes() *string {
	return s.RetryTimes
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetDimensions(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetEnable(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Enable = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetEscalation(v *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Escalation = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetGroupId(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetLastAlertTime(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.LastAlertTime = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetLastModifyTime(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.LastModifyTime = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetLevel(v int32) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Level = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetMetricName(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetMetricValues(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.MetricValues = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetNamespace(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetProductCategory(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.ProductCategory = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetResource(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Resource = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetRetryTimes(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.RetryTimes = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetRuleId(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetRuleName(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetStartTime(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetStatistics(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Statistics = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) SetThreshold(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource {
	s.Threshold = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResource) Validate() error {
	if s.Escalation != nil {
		if err := s.Escalation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation struct {
	Resource []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation) GetResource() []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	return s.Resource
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation) SetResource(v []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation {
	s.Resource = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalation) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource struct {
	ComparisonOperator *string                                                                                           `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Expression         *string                                                                                           `json:"Expression,omitempty" xml:"Expression,omitempty"`
	ExpressionList     *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Struct"`
	ExpressionListJoin *string                                                                                           `json:"ExpressionListJoin,omitempty" xml:"ExpressionListJoin,omitempty"`
	ExpressionRaw      *string                                                                                           `json:"ExpressionRaw,omitempty" xml:"ExpressionRaw,omitempty"`
	Level              *int32                                                                                            `json:"Level,omitempty" xml:"Level,omitempty"`
	PreCondition       *string                                                                                           `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Tag                *string                                                                                           `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Threshold          *string                                                                                           `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32                                                                                            `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetExpression() *string {
	return s.Expression
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetExpressionList() *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList {
	return s.ExpressionList
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetExpressionListJoin() *string {
	return s.ExpressionListJoin
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetExpressionRaw() *string {
	return s.ExpressionRaw
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetLevel() *int32 {
	return s.Level
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetTag() *string {
	return s.Tag
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetComparisonOperator(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetExpression(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.Expression = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetExpressionList(v *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.ExpressionList = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetExpressionListJoin(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.ExpressionListJoin = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetExpressionRaw(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.ExpressionRaw = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetLevel(v int32) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.Level = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetPreCondition(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.PreCondition = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetTag(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.Tag = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetThreshold(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.Threshold = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) SetTimes(v int32) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource {
	s.Times = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResource) Validate() error {
	if s.ExpressionList != nil {
		if err := s.ExpressionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList struct {
	ExpressionList []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Repeated"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList) GetExpressionList() []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList {
	return s.ExpressionList
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList) SetExpressionList(v []*DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList {
	s.ExpressionList = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionList) Validate() error {
	if s.ExpressionList != nil {
		for _, item := range s.ExpressionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) GetPeriod() *string {
	return s.Period
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) SetComparisonOperator(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) SetMetricName(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) SetPeriod(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList {
	s.Period = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) SetStatistics(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList {
	s.Statistics = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) SetThreshold(v string) *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList {
	s.Threshold = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponseBodyResourcesResourceEscalationResourceExpressionListExpressionList) Validate() error {
	return dara.Validate(s)
}
