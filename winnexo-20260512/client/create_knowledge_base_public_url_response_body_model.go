// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBasePublicUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetName() *string
	SetOriginalUrl(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetOriginalUrl() *string
	SetRequestId(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBasePublicUrlResponseBody
	GetStatus() *string
}

type CreateKnowledgeBasePublicUrlResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the requested target directory in the enterprise knowledge base.
	//
	// example:
	//
	// dir_tenant_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-14T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The error message. This parameter is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// Project plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The URL of the submitted public web page.
	//
	// example:
	//
	// https://example.com
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scope of the resource. The value is fixed to TENANT.
	//
	// example:
	//
	// TENANT
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The ID of the newly created source.
	//
	// example:
	//
	// src_public_url_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The status of the resource. A value of RUNNING indicates that the request is accepted but the crawling and parsing are not completed.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBasePublicUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBasePublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetCode(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetDirectoryId(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetGmtCreate(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetMessage(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetName(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetOriginalUrl(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.OriginalUrl = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetRequestId(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetScope(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetSourceId(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) SetStatus(v string) *CreateKnowledgeBasePublicUrlResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
