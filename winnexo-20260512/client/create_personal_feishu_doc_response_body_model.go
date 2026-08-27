// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalFeishuDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalFeishuDocResponseBody
	GetDirectoryId() *string
	SetDocUrl(v string) *CreatePersonalFeishuDocResponseBody
	GetDocUrl() *string
	SetGmtCreate(v string) *CreatePersonalFeishuDocResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalFeishuDocResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalFeishuDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalFeishuDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalFeishuDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalFeishuDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalFeishuDocResponseBody
	GetStatus() *string
}

type CreatePersonalFeishuDocResponseBody struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_personal_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-26T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the AI assistant.
	//
	// example:
	//
	// ProjectPlan
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
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src_feishu_doc_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalFeishuDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalFeishuDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuDocResponseBody) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreatePersonalFeishuDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalFeishuDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalFeishuDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalFeishuDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalFeishuDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalFeishuDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalFeishuDocResponseBody) SetCode(v string) *CreatePersonalFeishuDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetDirectoryId(v string) *CreatePersonalFeishuDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetDocUrl(v string) *CreatePersonalFeishuDocResponseBody {
	s.DocUrl = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetGmtCreate(v string) *CreatePersonalFeishuDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetMessage(v string) *CreatePersonalFeishuDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetName(v string) *CreatePersonalFeishuDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetRequestId(v string) *CreatePersonalFeishuDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetScope(v string) *CreatePersonalFeishuDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetSourceId(v string) *CreatePersonalFeishuDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) SetStatus(v string) *CreatePersonalFeishuDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalFeishuDocResponseBody) Validate() error {
	return dara.Validate(s)
}
