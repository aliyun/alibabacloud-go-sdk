// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalTextResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalTextResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalTextResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalTextResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalTextResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalTextResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalTextResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalTextResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalTextResponseBody
	GetStatus() *string
}

type CreatePersonalTextResponseBody struct {
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

func (s CreatePersonalTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTextResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalTextResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalTextResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalTextResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalTextResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalTextResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalTextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalTextResponseBody) SetCode(v string) *CreatePersonalTextResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetDirectoryId(v string) *CreatePersonalTextResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetGmtCreate(v string) *CreatePersonalTextResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetMessage(v string) *CreatePersonalTextResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetName(v string) *CreatePersonalTextResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetRequestId(v string) *CreatePersonalTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetScope(v string) *CreatePersonalTextResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetSourceId(v string) *CreatePersonalTextResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetStatus(v string) *CreatePersonalTextResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalTextResponseBody) Validate() error {
	return dara.Validate(s)
}
