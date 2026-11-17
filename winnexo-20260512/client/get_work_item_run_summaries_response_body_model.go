// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemRunSummariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkItemRunSummariesResponseBody
	GetCode() *string
	SetItems(v []*GetWorkItemRunSummariesResponseBodyItems) *GetWorkItemRunSummariesResponseBody
	GetItems() []*GetWorkItemRunSummariesResponseBodyItems
	SetMessage(v string) *GetWorkItemRunSummariesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkItemRunSummariesResponseBody
	GetRequestId() *string
}

type GetWorkItemRunSummariesResponseBody struct {
	// The common business status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The execution summaries generated for the current plan and current turn. Previous summaries are not returned until new summaries are generated for affected nodes that are re-run.
	//
	// example:
	//
	// [{"attemptId":"attempt_1","planId":"plan_v3","nodeKey":"query.amount","nodeStatus":"COMPLETED"}]
	Items []*GetWorkItemRunSummariesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The common error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetWorkItemRunSummariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRunSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemRunSummariesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkItemRunSummariesResponseBody) GetItems() []*GetWorkItemRunSummariesResponseBodyItems {
	return s.Items
}

func (s *GetWorkItemRunSummariesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkItemRunSummariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemRunSummariesResponseBody) SetCode(v string) *GetWorkItemRunSummariesResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBody) SetItems(v []*GetWorkItemRunSummariesResponseBodyItems) *GetWorkItemRunSummariesResponseBody {
	s.Items = v
	return s
}

func (s *GetWorkItemRunSummariesResponseBody) SetMessage(v string) *GetWorkItemRunSummariesResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBody) SetRequestId(v string) *GetWorkItemRunSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemRunSummariesResponseBodyItems struct {
	// The ID of the node attempt.
	//
	// example:
	//
	// attempt_1
	AttemptId *string `json:"attemptId,omitempty" xml:"attemptId,omitempty"`
	// The sequence number of the node attempt, starting from 1.
	//
	// example:
	//
	// 1
	AttemptSeq *int64 `json:"attemptSeq,omitempty" xml:"attemptSeq,omitempty"`
	// The ID of the executor.
	//
	// example:
	//
	// agent_1001
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// The name of the executor.
	//
	// example:
	//
	// Sales Analysis Assistant
	ExecutorName *string `json:"executorName,omitempty" xml:"executorName,omitempty"`
	// The executor type, such as AGENT or HUMAN.
	//
	// example:
	//
	// AGENT
	ExecutorType *string `json:"executorType,omitempty" xml:"executorType,omitempty"`
	// The end time of the node, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-05T08:05:00Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The node key.
	//
	// example:
	//
	// query.amount
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Query sales amount
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// The node status, such as COMPLETED or FAILED.
	//
	// example:
	//
	// COMPLETED
	NodeStatus *string `json:"nodeStatus,omitempty" xml:"nodeStatus,omitempty"`
	// The ID of the associated physical plan.
	//
	// example:
	//
	// plan_v3
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The execution result of the node, as a JSON object.
	//
	// example:
	//
	// {"rows":3}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// The content of the execution summary. This field is empty if the summary is not generated.
	//
	// example:
	//
	// Completed the sales amount query, returning a total of 3 rows of data
	RunSummary interface{} `json:"runSummary,omitempty" xml:"runSummary,omitempty"`
	// The reason for the execution summary failure. This field is empty if the generation is successful.
	//
	// example:
	//
	// Model invocation timed out
	RunSummaryError *string `json:"runSummaryError,omitempty" xml:"runSummaryError,omitempty"`
	// The generation time of the execution summary, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-05T08:06:00Z
	RunSummaryGeneratedAt *string `json:"runSummaryGeneratedAt,omitempty" xml:"runSummaryGeneratedAt,omitempty"`
	// The list of data sources referenced in the execution summary.
	//
	// example:
	//
	// [{"name":"report.md"}]
	RunSummarySources []interface{} `json:"runSummarySources,omitempty" xml:"runSummarySources,omitempty" type:"Repeated"`
	// The generation status of the execution summary, such as COMPLETED, FAILED, or PROCESSING.
	//
	// example:
	//
	// COMPLETED
	RunSummaryStatus *string `json:"runSummaryStatus,omitempty" xml:"runSummaryStatus,omitempty"`
	// The list of tools used in the execution summary.
	//
	// example:
	//
	// [{"type":"SKILL","name":"sql_query"}]
	RunSummaryTools []*GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools `json:"runSummaryTools,omitempty" xml:"runSummaryTools,omitempty" type:"Repeated"`
	// The start time of the node, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The sequence number of the node turn, starting from 1.
	//
	// example:
	//
	// 1
	TurnSeq *int64 `json:"turnSeq,omitempty" xml:"turnSeq,omitempty"`
}

func (s GetWorkItemRunSummariesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRunSummariesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetAttemptId() *string {
	return s.AttemptId
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetAttemptSeq() *int64 {
	return s.AttemptSeq
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetExecutorId() *string {
	return s.ExecutorId
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetExecutorName() *string {
	return s.ExecutorName
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetExecutorType() *string {
	return s.ExecutorType
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetNodeKey() *string {
	return s.NodeKey
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetNodeName() *string {
	return s.NodeName
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetPlanId() *string {
	return s.PlanId
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetResult() interface{} {
	return s.Result
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetRunSummary() interface{} {
	return s.RunSummary
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetRunSummaryError() *string {
	return s.RunSummaryError
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetRunSummaryGeneratedAt() *string {
	return s.RunSummaryGeneratedAt
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetRunSummarySources() []interface{} {
	return s.RunSummarySources
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetRunSummaryStatus() *string {
	return s.RunSummaryStatus
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetRunSummaryTools() []*GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools {
	return s.RunSummaryTools
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetWorkItemRunSummariesResponseBodyItems) GetTurnSeq() *int64 {
	return s.TurnSeq
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetAttemptId(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.AttemptId = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetAttemptSeq(v int64) *GetWorkItemRunSummariesResponseBodyItems {
	s.AttemptSeq = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetExecutorId(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.ExecutorId = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetExecutorName(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.ExecutorName = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetExecutorType(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.ExecutorType = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetFinishedAt(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.FinishedAt = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetNodeKey(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.NodeKey = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetNodeName(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.NodeName = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetNodeStatus(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.NodeStatus = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetPlanId(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.PlanId = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetResult(v interface{}) *GetWorkItemRunSummariesResponseBodyItems {
	s.Result = v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetRunSummary(v interface{}) *GetWorkItemRunSummariesResponseBodyItems {
	s.RunSummary = v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetRunSummaryError(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.RunSummaryError = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetRunSummaryGeneratedAt(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.RunSummaryGeneratedAt = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetRunSummarySources(v []interface{}) *GetWorkItemRunSummariesResponseBodyItems {
	s.RunSummarySources = v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetRunSummaryStatus(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.RunSummaryStatus = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetRunSummaryTools(v []*GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) *GetWorkItemRunSummariesResponseBodyItems {
	s.RunSummaryTools = v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetStartedAt(v string) *GetWorkItemRunSummariesResponseBodyItems {
	s.StartedAt = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) SetTurnSeq(v int64) *GetWorkItemRunSummariesResponseBodyItems {
	s.TurnSeq = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItems) Validate() error {
	if s.RunSummaryTools != nil {
		for _, item := range s.RunSummaryTools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools struct {
	// The name of the tool.
	//
	// example:
	//
	// Sales Analysis Assistant
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tool type, such as SKILL or MCP.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) GoString() string {
	return s.String()
}

func (s *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) GetName() *string {
	return s.Name
}

func (s *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) GetType() *string {
	return s.Type
}

func (s *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) SetName(v string) *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools {
	s.Name = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) SetType(v string) *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools {
	s.Type = &v
	return s
}

func (s *GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools) Validate() error {
	return dara.Validate(s)
}
