// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedNodeKeys(v []*string) *RerunWorkflowNodeResponseBody
	GetAffectedNodeKeys() []*string
	SetCode(v string) *RerunWorkflowNodeResponseBody
	GetCode() *string
	SetMessage(v string) *RerunWorkflowNodeResponseBody
	GetMessage() *string
	SetPlanId(v string) *RerunWorkflowNodeResponseBody
	GetPlanId() *string
	SetRequestId(v string) *RerunWorkflowNodeResponseBody
	GetRequestId() *string
	SetRevision(v int64) *RerunWorkflowNodeResponseBody
	GetRevision() *int64
	SetRunId(v string) *RerunWorkflowNodeResponseBody
	GetRunId() *string
	SetStateVersion(v int64) *RerunWorkflowNodeResponseBody
	GetStateVersion() *int64
}

type RerunWorkflowNodeResponseBody struct {
	// 本次重跑影响的节点 key 列表，供调用方刷新；成功响应恒包含，错误响应可缺省
	//
	// example:
	//
	// ["query.amount"]
	AffectedNodeKeys []*string `json:"affectedNodeKeys,omitempty" xml:"affectedNodeKeys,omitempty" type:"Repeated"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 重跑后的当前物理计划ID；固定规划节点重跑可能生成新计划，后续查询应使用此值
	//
	// example:
	//
	// plan_new
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 重跑后的当前计划修订号
	//
	// example:
	//
	// 2
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// 当前工作流运行ID，固定节点生成新 revision 时保持不变
	//
	// example:
	//
	// run_20261105_01
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// 重跑被接受后的工作流状态版本
	//
	// example:
	//
	// 8
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
}

func (s RerunWorkflowNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RerunWorkflowNodeResponseBody) GetAffectedNodeKeys() []*string {
	return s.AffectedNodeKeys
}

func (s *RerunWorkflowNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *RerunWorkflowNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RerunWorkflowNodeResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *RerunWorkflowNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RerunWorkflowNodeResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *RerunWorkflowNodeResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *RerunWorkflowNodeResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *RerunWorkflowNodeResponseBody) SetAffectedNodeKeys(v []*string) *RerunWorkflowNodeResponseBody {
	s.AffectedNodeKeys = v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetCode(v string) *RerunWorkflowNodeResponseBody {
	s.Code = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetMessage(v string) *RerunWorkflowNodeResponseBody {
	s.Message = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetPlanId(v string) *RerunWorkflowNodeResponseBody {
	s.PlanId = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetRequestId(v string) *RerunWorkflowNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetRevision(v int64) *RerunWorkflowNodeResponseBody {
	s.Revision = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetRunId(v string) *RerunWorkflowNodeResponseBody {
	s.RunId = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) SetStateVersion(v int64) *RerunWorkflowNodeResponseBody {
	s.StateVersion = &v
	return s
}

func (s *RerunWorkflowNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
