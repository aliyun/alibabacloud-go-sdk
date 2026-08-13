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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 分类描述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 新建分类 ID
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 目录 KB 归属类型：normal / aliding_kb_root / aliding_kb_internal
	//
	// example:
	//
	// string_value
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// 创建时间戳（毫秒）
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间戳（毫秒）
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父分类 ID；新分类挂在租户根目录下时返回 null
	//
	// example:
	//
	// exampleParentDirectoryId
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// 文件 OSS URL
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
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
