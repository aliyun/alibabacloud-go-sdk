// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalAlidingDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAlidingDocResponseBody
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreatePersonalAlidingDocResponseBody
	GetFilePublicUrl() *string
	SetGmtCreate(v string) *CreatePersonalAlidingDocResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalAlidingDocResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAlidingDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalAlidingDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalAlidingDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalAlidingDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalAlidingDocResponseBody
	GetStatus() *string
}

type CreatePersonalAlidingDocResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 绑定的目录 ID（请求体传入时 echo 回；缺省走默认根目录时为 null）
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 文档公开 URL（echo 回入参，便于调用方对齐）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 创建时间 ISO8601
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 资源 scope，固定为 PERSONAL
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 新建资源 ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 资源状态（创建链路初始多为 PENDING；on_create 失败则为 FAILED）
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalAlidingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAlidingDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingDocResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreatePersonalAlidingDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAlidingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAlidingDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAlidingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAlidingDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalAlidingDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalAlidingDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAlidingDocResponseBody) SetCode(v string) *CreatePersonalAlidingDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetDirectoryId(v string) *CreatePersonalAlidingDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetFilePublicUrl(v string) *CreatePersonalAlidingDocResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetGmtCreate(v string) *CreatePersonalAlidingDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetMessage(v string) *CreatePersonalAlidingDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetName(v string) *CreatePersonalAlidingDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetRequestId(v string) *CreatePersonalAlidingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetScope(v string) *CreatePersonalAlidingDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetSourceId(v string) *CreatePersonalAlidingDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetStatus(v string) *CreatePersonalAlidingDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) Validate() error {
	return dara.Validate(s)
}
