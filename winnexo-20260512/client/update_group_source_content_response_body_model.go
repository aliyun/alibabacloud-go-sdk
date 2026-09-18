// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupSourceContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGroupSourceContentResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGroupSourceContentResponseBody
	GetMessage() *string
	SetName(v string) *UpdateGroupSourceContentResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateGroupSourceContentResponseBody
	GetRequestId() *string
	SetSourceId(v string) *UpdateGroupSourceContentResponseBody
	GetSourceId() *string
	SetSourceType(v string) *UpdateGroupSourceContentResponseBody
	GetSourceType() *string
	SetStatus(v string) *UpdateGroupSourceContentResponseBody
	GetStatus() *string
}

type UpdateGroupSourceContentResponseBody struct {
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
	// C474BFC7-7B11-5D92-971E-74AA82EC495B
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

func (s UpdateGroupSourceContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupSourceContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupSourceContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGroupSourceContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGroupSourceContentResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateGroupSourceContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupSourceContentResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateGroupSourceContentResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateGroupSourceContentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateGroupSourceContentResponseBody) SetCode(v string) *UpdateGroupSourceContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetMessage(v string) *UpdateGroupSourceContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetName(v string) *UpdateGroupSourceContentResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetRequestId(v string) *UpdateGroupSourceContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetSourceId(v string) *UpdateGroupSourceContentResponseBody {
	s.SourceId = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetSourceType(v string) *UpdateGroupSourceContentResponseBody {
	s.SourceType = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetStatus(v string) *UpdateGroupSourceContentResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) Validate() error {
	return dara.Validate(s)
}
