// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFeishuDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetDocUrl() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseFeishuDocResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseFeishuDocResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_tenant_kb_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-08-26T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The mirror name.
	//
	// example:
	//
	// Enterprise Policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// TENANT
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src_feishu_doc_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source status after re-parsing.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBaseFeishuDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFeishuDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetCode(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetDocUrl(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.DocUrl = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetMessage(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetName(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetRequestId(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetScope(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetSourceId(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) SetStatus(v string) *CreateKnowledgeBaseFeishuDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponseBody) Validate() error {
	return dara.Validate(s)
}
