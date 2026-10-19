// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseWorkflowTodosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCancelledCount(v int64) *CloseWorkflowTodosResponseBody
	GetCancelledCount() *int64
	SetCode(v string) *CloseWorkflowTodosResponseBody
	GetCode() *string
	SetMessage(v string) *CloseWorkflowTodosResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloseWorkflowTodosResponseBody
	GetRequestId() *string
	SetSourceWorkflowId(v string) *CloseWorkflowTodosResponseBody
	GetSourceWorkflowId() *string
}

type CloseWorkflowTodosResponseBody struct {
	// 关闭条数（幂等重复调用为 0）
	//
	// example:
	//
	// 3
	CancelledCount *int64 `json:"cancelledCount,omitempty" xml:"cancelledCount,omitempty"`
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
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 工作流 ID（原样回显）
	//
	// example:
	//
	// wf-1
	SourceWorkflowId *string `json:"sourceWorkflowId,omitempty" xml:"sourceWorkflowId,omitempty"`
}

func (s CloseWorkflowTodosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseWorkflowTodosResponseBody) GoString() string {
	return s.String()
}

func (s *CloseWorkflowTodosResponseBody) GetCancelledCount() *int64 {
	return s.CancelledCount
}

func (s *CloseWorkflowTodosResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloseWorkflowTodosResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloseWorkflowTodosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseWorkflowTodosResponseBody) GetSourceWorkflowId() *string {
	return s.SourceWorkflowId
}

func (s *CloseWorkflowTodosResponseBody) SetCancelledCount(v int64) *CloseWorkflowTodosResponseBody {
	s.CancelledCount = &v
	return s
}

func (s *CloseWorkflowTodosResponseBody) SetCode(v string) *CloseWorkflowTodosResponseBody {
	s.Code = &v
	return s
}

func (s *CloseWorkflowTodosResponseBody) SetMessage(v string) *CloseWorkflowTodosResponseBody {
	s.Message = &v
	return s
}

func (s *CloseWorkflowTodosResponseBody) SetRequestId(v string) *CloseWorkflowTodosResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseWorkflowTodosResponseBody) SetSourceWorkflowId(v string) *CloseWorkflowTodosResponseBody {
	s.SourceWorkflowId = &v
	return s
}

func (s *CloseWorkflowTodosResponseBody) Validate() error {
	return dara.Validate(s)
}
