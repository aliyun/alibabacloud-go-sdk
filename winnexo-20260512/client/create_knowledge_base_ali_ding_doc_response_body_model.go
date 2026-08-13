// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseAliDingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetFilePublicUrl() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseAliDingDocResponseBody struct {
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
	// 文档公开 URL（echo 回入参）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 创建时间 ISO8601
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 资源显示名称
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

func (s CreateKnowledgeBaseAliDingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseAliDingDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetCode(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetFilePublicUrl(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetMessage(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetName(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetRequestId(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetScope(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetSourceId(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetStatus(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) Validate() error {
	return dara.Validate(s)
}
