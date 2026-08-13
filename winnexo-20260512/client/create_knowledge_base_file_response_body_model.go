// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseFileResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseFileResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseFileResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseFileResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseFileResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseFileResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseFileResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseFileResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseFileResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseFileResponseBody struct {
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

func (s CreateKnowledgeBaseFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseFileResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseFileResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseFileResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseFileResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseFileResponseBody) SetCode(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseFileResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseFileResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetMessage(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetName(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetRequestId(v string) *CreateKnowledgeBaseFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetScope(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetSourceId(v string) *CreateKnowledgeBaseFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetStatus(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) Validate() error {
	return dara.Validate(s)
}
