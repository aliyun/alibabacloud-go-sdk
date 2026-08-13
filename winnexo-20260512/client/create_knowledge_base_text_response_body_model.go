// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseTextResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseTextResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseTextResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseTextResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseTextResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseTextResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseTextResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseTextResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseTextResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseTextResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 绑定的目录 ID
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 创建时间 ISO8601
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
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
	// 资源 scope，固定为 TENANT
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 新建资源 ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 资源状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBaseTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseTextResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseTextResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseTextResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseTextResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseTextResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseTextResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseTextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseTextResponseBody) SetCode(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseTextResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseTextResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetMessage(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetName(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetRequestId(v string) *CreateKnowledgeBaseTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetScope(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetSourceId(v string) *CreateKnowledgeBaseTextResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetStatus(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) Validate() error {
	return dara.Validate(s)
}
