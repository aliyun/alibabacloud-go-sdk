// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseGroupSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReparseGroupSourceResponseBody
	GetCode() *string
	SetMessage(v string) *ReparseGroupSourceResponseBody
	GetMessage() *string
	SetName(v string) *ReparseGroupSourceResponseBody
	GetName() *string
	SetRequestId(v string) *ReparseGroupSourceResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReparseGroupSourceResponseBody
	GetSourceId() *string
	SetSourceType(v string) *ReparseGroupSourceResponseBody
	GetSourceType() *string
	SetStatus(v string) *ReparseGroupSourceResponseBody
	GetStatus() *string
}

type ReparseGroupSourceResponseBody struct {
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
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作后的资料名称，沿用已有名称维护规则
	//
	// example:
	//
	// 项目资料
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 资料 ID；替换、编辑、重新解析均保持该 ID
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 资料类型
	//
	// example:
	//
	// example
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 当前资料状态；RUNNING 表示处理中，异步受理不代表解析完成
	//
	// example:
	//
	// example
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReparseGroupSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReparseGroupSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ReparseGroupSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReparseGroupSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReparseGroupSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *ReparseGroupSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReparseGroupSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseGroupSourceResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *ReparseGroupSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReparseGroupSourceResponseBody) SetCode(v string) *ReparseGroupSourceResponseBody {
	s.Code = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) SetMessage(v string) *ReparseGroupSourceResponseBody {
	s.Message = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) SetName(v string) *ReparseGroupSourceResponseBody {
	s.Name = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) SetRequestId(v string) *ReparseGroupSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) SetSourceId(v string) *ReparseGroupSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) SetSourceType(v string) *ReparseGroupSourceResponseBody {
	s.SourceType = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) SetStatus(v string) *ReparseGroupSourceResponseBody {
	s.Status = &v
	return s
}

func (s *ReparseGroupSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
