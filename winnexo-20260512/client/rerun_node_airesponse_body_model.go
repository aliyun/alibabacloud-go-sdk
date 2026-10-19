// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunNodeAIResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RerunNodeAIResponseBody
	GetCode() *string
	SetInstanceId(v string) *RerunNodeAIResponseBody
	GetInstanceId() *string
	SetMessage(v string) *RerunNodeAIResponseBody
	GetMessage() *string
	SetNodeId(v string) *RerunNodeAIResponseBody
	GetNodeId() *string
	SetRequestId(v string) *RerunNodeAIResponseBody
	GetRequestId() *string
	SetRerunAccepted(v bool) *RerunNodeAIResponseBody
	GetRerunAccepted() *bool
	SetStateVersion(v int64) *RerunNodeAIResponseBody
	GetStateVersion() *int64
}

type RerunNodeAIResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 物理实例ID
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 已安排重跑的节点 key
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 重跑事实接受为 true，不表示 AI 已完成；失败通过 Code 和 Message 返回
	//
	// example:
	//
	// true
	RerunAccepted *bool `json:"rerunAccepted,omitempty" xml:"rerunAccepted,omitempty"`
	// 状态版本
	//
	// example:
	//
	// 2
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
}

func (s RerunNodeAIResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RerunNodeAIResponseBody) GoString() string {
	return s.String()
}

func (s *RerunNodeAIResponseBody) GetCode() *string {
	return s.Code
}

func (s *RerunNodeAIResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RerunNodeAIResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RerunNodeAIResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *RerunNodeAIResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RerunNodeAIResponseBody) GetRerunAccepted() *bool {
	return s.RerunAccepted
}

func (s *RerunNodeAIResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *RerunNodeAIResponseBody) SetCode(v string) *RerunNodeAIResponseBody {
	s.Code = &v
	return s
}

func (s *RerunNodeAIResponseBody) SetInstanceId(v string) *RerunNodeAIResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RerunNodeAIResponseBody) SetMessage(v string) *RerunNodeAIResponseBody {
	s.Message = &v
	return s
}

func (s *RerunNodeAIResponseBody) SetNodeId(v string) *RerunNodeAIResponseBody {
	s.NodeId = &v
	return s
}

func (s *RerunNodeAIResponseBody) SetRequestId(v string) *RerunNodeAIResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerunNodeAIResponseBody) SetRerunAccepted(v bool) *RerunNodeAIResponseBody {
	s.RerunAccepted = &v
	return s
}

func (s *RerunNodeAIResponseBody) SetStateVersion(v int64) *RerunNodeAIResponseBody {
	s.StateVersion = &v
	return s
}

func (s *RerunNodeAIResponseBody) Validate() error {
	return dara.Validate(s)
}
