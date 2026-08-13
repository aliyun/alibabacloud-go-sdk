// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetMessage() *string
	SetName(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetRequestId() *string
	SetSourceId(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetSourceId() *string
	SetSourceType(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetSourceType() *string
	SetStatus(v string) *UpdateKnowledgeBaseSourceContentResponseBody
	GetStatus() *string
}

type UpdateKnowledgeBaseSourceContentResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 数据源 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 数据源类型
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 重新解析后的数据源状态
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateKnowledgeBaseSourceContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetCode(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetMessage(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetName(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetSourceId(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetSourceType(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.SourceType = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) SetStatus(v string) *UpdateKnowledgeBaseSourceContentResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponseBody) Validate() error {
	return dara.Validate(s)
}
