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
	// The business status code. A value of 200 indicates success. If the request fails, a specific business error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the created directory. If no description is set, the value is empty.
	//
	// example:
	//
	// Team product materials and operation guides
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The business ID of the newly created directory. You can use this ID to read or update the knowledge base configuration for standard enterprise root directories.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The time when the directory was created. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1699000000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the directory was last modified. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1699000010
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The internal primary key of the directory. You can use this key as the parentId when you create a subdirectory. This key is different from the business ID directoryId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The response message. The value is success if the request is successful. If the request fails, this parameter describes the cause of the error.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the created directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Product knowledge base
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The legacy single digital employee association field. You cannot use this field to replace the mode and association list returned by GetKnowledgeBaseConfiguration.
	//
	// example:
	//
	// opt_1
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The internal primary key of the parent directory. This value is empty for enterprise root directories.
	//
	// example:
	//
	// 1001
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The path of the created directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /Product knowledge base
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The request tracking ID. You can use this ID to troubleshoot issues related to this call.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the tenant to which the directory belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the platform user who created the directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
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
