// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTenantDirectoryResponseBody
	GetCode() *string
	SetDescription(v string) *CreateTenantDirectoryResponseBody
	GetDescription() *string
	SetDirectoryId(v string) *CreateTenantDirectoryResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v int64) *CreateTenantDirectoryResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *CreateTenantDirectoryResponseBody
	GetGmtModified() *int64
	SetId(v int64) *CreateTenantDirectoryResponseBody
	GetId() *int64
	SetMessage(v string) *CreateTenantDirectoryResponseBody
	GetMessage() *string
	SetName(v string) *CreateTenantDirectoryResponseBody
	GetName() *string
	SetOperatingObjectName(v string) *CreateTenantDirectoryResponseBody
	GetOperatingObjectName() *string
	SetParentId(v int64) *CreateTenantDirectoryResponseBody
	GetParentId() *int64
	SetPath(v string) *CreateTenantDirectoryResponseBody
	GetPath() *string
	SetRequestId(v string) *CreateTenantDirectoryResponseBody
	GetRequestId() *string
	SetTenantId(v int64) *CreateTenantDirectoryResponseBody
	GetTenantId() *int64
	SetUserId(v int64) *CreateTenantDirectoryResponseBody
	GetUserId() *int64
}

type CreateTenantDirectoryResponseBody struct {
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
	// 目录唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 创建时间戳
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间戳
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 目录内部主键
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
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
	// 历史运营对象名称
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 父目录内部主键
	//
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件 OSS URL
	//
	// This parameter is required.
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
	// 租户 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 创建人用户 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateTenantDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTenantDirectoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateTenantDirectoryResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateTenantDirectoryResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreateTenantDirectoryResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateTenantDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTenantDirectoryResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateTenantDirectoryResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateTenantDirectoryResponseBody) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateTenantDirectoryResponseBody) GetPath() *string {
	return s.Path
}

func (s *CreateTenantDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTenantDirectoryResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *CreateTenantDirectoryResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateTenantDirectoryResponseBody) SetCode(v string) *CreateTenantDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetDescription(v string) *CreateTenantDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetDirectoryId(v string) *CreateTenantDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetGmtCreate(v int64) *CreateTenantDirectoryResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetGmtModified(v int64) *CreateTenantDirectoryResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetId(v int64) *CreateTenantDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetMessage(v string) *CreateTenantDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetName(v string) *CreateTenantDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetOperatingObjectName(v string) *CreateTenantDirectoryResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetParentId(v int64) *CreateTenantDirectoryResponseBody {
	s.ParentId = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetPath(v string) *CreateTenantDirectoryResponseBody {
	s.Path = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetRequestId(v string) *CreateTenantDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetTenantId(v int64) *CreateTenantDirectoryResponseBody {
	s.TenantId = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) SetUserId(v int64) *CreateTenantDirectoryResponseBody {
	s.UserId = &v
	return s
}

func (s *CreateTenantDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
