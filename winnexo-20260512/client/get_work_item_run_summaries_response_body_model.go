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
	// 公共业务状态码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前计划、当前轮次已经生成的运行总结；重跑受影响节点的新总结生成前不会返回旧总结
	//
	// example:
	//
	// [{"attemptId":"attempt_example","planId":"plan_example","nodeKey":"query.amount","nodeStatus":"COMPLETED"}]
	Items []*GetWorkItemRunSummariesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 公共错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
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
	// 节点尝试 ID
	//
	// example:
	//
	// attempt_1
	AttemptId *string `json:"attemptId,omitempty" xml:"attemptId,omitempty"`
	// 节点尝试序号，从 1 开始
	//
	// example:
	//
	// 1
	AttemptSeq *int64 `json:"attemptSeq,omitempty" xml:"attemptSeq,omitempty"`
	// 执行器 ID
	//
	// example:
	//
	// agent_1001
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 执行器名称
	//
	// example:
	//
	// 销售分析助手
	ExecutorName *string `json:"executorName,omitempty" xml:"executorName,omitempty"`
	// 执行器类型，如 AGENT、HUMAN
	//
	// example:
	//
	// AGENT
	ExecutorType *string `json:"executorType,omitempty" xml:"executorType,omitempty"`
	// 节点结束时间，ISO8601 格式
	//
	// example:
	//
	// 2026-11-05T08:05:00Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// 节点 key
	//
	// example:
	//
	// query.amount
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
	// 节点名称
	//
	// example:
	//
	// 查询销售金额
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 节点状态，如 COMPLETED、FAILED
	//
	// example:
	//
	// COMPLETED
	NodeStatus *string `json:"nodeStatus,omitempty" xml:"nodeStatus,omitempty"`
	// 所属物理计划 ID
	//
	// example:
	//
	// plan_v3
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// 节点执行结果，JSON 对象
	//
	// example:
	//
	// {"rows":3}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// 运行总结正文；未生成时为空
	//
	// example:
	//
	// 已完成销售金额查询，共返回 3 行数据
	RunSummary interface{} `json:"runSummary,omitempty" xml:"runSummary,omitempty"`
	// 运行总结失败原因；成功时为空
	//
	// example:
	//
	// 模型调用超时
	RunSummaryError *string `json:"runSummaryError,omitempty" xml:"runSummaryError,omitempty"`
	// 运行总结生成时间，ISO8601 格式
	//
	// example:
	//
	// 2026-11-05T08:06:00Z
	RunSummaryGeneratedAt *string `json:"runSummaryGeneratedAt,omitempty" xml:"runSummaryGeneratedAt,omitempty"`
	// 运行总结引用的数据来源列表
	//
	// example:
	//
	// [{"name":"report.md"}]
	RunSummarySources []interface{} `json:"runSummarySources,omitempty" xml:"runSummarySources,omitempty" type:"Repeated"`
	// 运行总结生成状态，如 COMPLETED、FAILED、PROCESSING
	//
	// example:
	//
	// COMPLETED
	RunSummaryStatus *string `json:"runSummaryStatus,omitempty" xml:"runSummaryStatus,omitempty"`
	// 运行总结使用的工具列表
	//
	// example:
	//
	// [{"type":"SKILL","name":"sql_query"}]
	RunSummaryTools []*GetWorkItemRunSummariesResponseBodyItemsRunSummaryTools `json:"runSummaryTools,omitempty" xml:"runSummaryTools,omitempty" type:"Repeated"`
	// 节点开始时间，ISO8601 格式
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// 节点轮次序号，从 1 开始
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
	// 工具名称
	//
	// example:
	//
	// 销售分析助手
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工具类型，如 SKILL、MCP
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
