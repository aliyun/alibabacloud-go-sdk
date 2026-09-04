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
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The tenant folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (the operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The ID of the parent node.
	//
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The path of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the tenant for which the operation takes effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The user ID of the creator.
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
