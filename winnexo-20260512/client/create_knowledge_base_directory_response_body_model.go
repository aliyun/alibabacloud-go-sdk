// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetCode() *string
	SetDescription(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetDirectoryId() *string
	SetDirectoryKind(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetDirectoryKind() *string
	SetGmtCreate(v int64) *CreateKnowledgeBaseDirectoryResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *CreateKnowledgeBaseDirectoryResponseBody
	GetGmtModified() *int64
	SetMessage(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetName() *string
	SetParentDirectoryId(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetParentDirectoryId() *string
	SetPath(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetPath() *string
	SetRequestId(v string) *CreateKnowledgeBaseDirectoryResponseBody
	GetRequestId() *string
}

type CreateKnowledgeBaseDirectoryResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// This is default function description by fc-deploy component
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The directory type.
	//
	// example:
	//
	// string_value
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-11-14T02:18:27Z
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2026-01-19T01:48:56Z
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name.
	//
	// example:
	//
	// p-toolset-89550434-4e20-4e4e-bcac-9ab81b82c5b3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// wd-lxykjnnw4lyl9eq
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// UVCIpI0siUski9iw
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C474BFC7-7B11-5D92-971E-74AA82EC495B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateKnowledgeBaseDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetPath() *string {
	return s.Path
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetCode(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetDescription(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetDirectoryKind(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.DirectoryKind = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetGmtCreate(v int64) *CreateKnowledgeBaseDirectoryResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetGmtModified(v int64) *CreateKnowledgeBaseDirectoryResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetMessage(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetName(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetParentDirectoryId(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetPath(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.Path = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) SetRequestId(v string) *CreateKnowledgeBaseDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
