// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalDirectoryResponseBody
	GetCode() *string
	SetDescription(v string) *CreatePersonalDirectoryResponseBody
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDirectoryResponseBody
	GetDirectoryId() *string
	SetDirectoryKind(v string) *CreatePersonalDirectoryResponseBody
	GetDirectoryKind() *string
	SetGmtCreate(v int64) *CreatePersonalDirectoryResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *CreatePersonalDirectoryResponseBody
	GetGmtModified() *int64
	SetMessage(v string) *CreatePersonalDirectoryResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalDirectoryResponseBody
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalDirectoryResponseBody
	GetOperatingObjectName() *string
	SetParentDirectoryId(v string) *CreatePersonalDirectoryResponseBody
	GetParentDirectoryId() *string
	SetPath(v string) *CreatePersonalDirectoryResponseBody
	GetPath() *string
	SetRequestId(v string) *CreatePersonalDirectoryResponseBody
	GetRequestId() *string
}

type CreatePersonalDirectoryResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 目录描述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 新建目录 ID
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
	// 所属数字员工名称
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 父目录 ID（service 若回填默认根目录，这里返回回填后的父目录 ID）
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

func (s CreatePersonalDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalDirectoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDirectoryResponseBody) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *CreatePersonalDirectoryResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreatePersonalDirectoryResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreatePersonalDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalDirectoryResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDirectoryResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDirectoryResponseBody) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreatePersonalDirectoryResponseBody) GetPath() *string {
	return s.Path
}

func (s *CreatePersonalDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalDirectoryResponseBody) SetCode(v string) *CreatePersonalDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetDescription(v string) *CreatePersonalDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetDirectoryId(v string) *CreatePersonalDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetDirectoryKind(v string) *CreatePersonalDirectoryResponseBody {
	s.DirectoryKind = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetGmtCreate(v int64) *CreatePersonalDirectoryResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetGmtModified(v int64) *CreatePersonalDirectoryResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetMessage(v string) *CreatePersonalDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetName(v string) *CreatePersonalDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetOperatingObjectName(v string) *CreatePersonalDirectoryResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetParentDirectoryId(v string) *CreatePersonalDirectoryResponseBody {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetPath(v string) *CreatePersonalDirectoryResponseBody {
	s.Path = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) SetRequestId(v string) *CreatePersonalDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
