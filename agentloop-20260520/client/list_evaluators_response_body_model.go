// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluators(v []*ListEvaluatorsResponseBodyEvaluators) *ListEvaluatorsResponseBody
	GetEvaluators() []*ListEvaluatorsResponseBodyEvaluators
	SetMaxResults(v int32) *ListEvaluatorsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluatorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEvaluatorsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListEvaluatorsResponseBody
	GetTotal() *int32
}

type ListEvaluatorsResponseBody struct {
	// The list of evaluator summaries.
	//
	// example:
	//
	// [{"name":"trace_task_completion","type":"AGENT","latestVersion":"1.0.0"}]
	Evaluators []*ListEvaluatorsResponseBodyEvaluators `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The number of entries per page used in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEzM30=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of evaluators that match the filter conditions.
	//
	// example:
	//
	// 12
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEvaluatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluatorsResponseBody) GetEvaluators() []*ListEvaluatorsResponseBodyEvaluators {
	return s.Evaluators
}

func (s *ListEvaluatorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluatorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluatorsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListEvaluatorsResponseBody) SetEvaluators(v []*ListEvaluatorsResponseBodyEvaluators) *ListEvaluatorsResponseBody {
	s.Evaluators = v
	return s
}

func (s *ListEvaluatorsResponseBody) SetMaxResults(v int32) *ListEvaluatorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluatorsResponseBody) SetNextToken(v string) *ListEvaluatorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEvaluatorsResponseBody) SetRequestId(v string) *ListEvaluatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluatorsResponseBody) SetTotal(v int32) *ListEvaluatorsResponseBody {
	s.Total = &v
	return s
}

func (s *ListEvaluatorsResponseBody) Validate() error {
	if s.Evaluators != nil {
		for _, item := range s.Evaluators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluatorsResponseBodyEvaluators struct {
	// The list of annotations.
	//
	// example:
	//
	// ["__en"]
	Annotations []*string `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	// The creation time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The evaluator description.
	//
	// example:
	//
	// 判断 Agent 是否完成用户任务
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name.
	//
	// example:
	//
	// 链路任务完成度
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// 1.0.0
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// The metric name.
	//
	// example:
	//
	// agent_task_completion
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The evaluator name.
	//
	// example:
	//
	// trace_task_completion
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The evaluator properties.
	//
	// example:
	//
	// {"agentEvaluatorMode":"raw_prompt"}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// The evaluator type.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The update time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816600
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListEvaluatorsResponseBodyEvaluators) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorsResponseBodyEvaluators) GoString() string {
	return s.String()
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetAnnotations() []*string {
	return s.Annotations
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetDescription() *string {
	return s.Description
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetMetricName() *string {
	return s.MetricName
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetName() *string {
	return s.Name
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetType() *string {
	return s.Type
}

func (s *ListEvaluatorsResponseBodyEvaluators) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetAnnotations(v []*string) *ListEvaluatorsResponseBodyEvaluators {
	s.Annotations = v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetCreatedAt(v int64) *ListEvaluatorsResponseBodyEvaluators {
	s.CreatedAt = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetDescription(v string) *ListEvaluatorsResponseBodyEvaluators {
	s.Description = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetDisplayName(v string) *ListEvaluatorsResponseBodyEvaluators {
	s.DisplayName = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetLatestVersion(v string) *ListEvaluatorsResponseBodyEvaluators {
	s.LatestVersion = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetMetricName(v string) *ListEvaluatorsResponseBodyEvaluators {
	s.MetricName = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetName(v string) *ListEvaluatorsResponseBodyEvaluators {
	s.Name = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetProperties(v map[string]interface{}) *ListEvaluatorsResponseBodyEvaluators {
	s.Properties = v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetType(v string) *ListEvaluatorsResponseBodyEvaluators {
	s.Type = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) SetUpdatedAt(v int64) *ListEvaluatorsResponseBodyEvaluators {
	s.UpdatedAt = &v
	return s
}

func (s *ListEvaluatorsResponseBodyEvaluators) Validate() error {
	return dara.Validate(s)
}
