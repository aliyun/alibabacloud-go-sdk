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
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description.
	//
	// example:
	//
	// PublicApplication
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The folder type.
	//
	// example:
	//
	// string_value
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-03-04 13:54:52
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2025-11-14T02:18:27Z
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the worksheet.
	//
	// example:
	//
	// conn_ip_101
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital human (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// wd-lxykjnnw4lyl9eq
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The path.
	//
	// example:
	//
	// oss://clg-paimon-4a00f1ac43464714b86fb54ca41a84c9/db-abc73646-6a08-4b96-820f-3d1d547a1e3b.db/tbl-c8a33522-5398-4f8e-9a2a-fba1efad94d1
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 375701FC-2FB9-5782-BE8F-A3F5E2F2158C
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
