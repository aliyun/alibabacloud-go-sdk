// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalFileResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalFileResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalFileResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalFileResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalFileResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalFileResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalFileResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalFileResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalFileResponseBody
	GetStatus() *string
}

type CreatePersonalFileResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 绑定的目录 ID
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
	// 资源状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalFileResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFileResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalFileResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalFileResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalFileResponseBody) SetCode(v string) *CreatePersonalFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetDirectoryId(v string) *CreatePersonalFileResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetGmtCreate(v string) *CreatePersonalFileResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetMessage(v string) *CreatePersonalFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetName(v string) *CreatePersonalFileResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetRequestId(v string) *CreatePersonalFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetScope(v string) *CreatePersonalFileResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetSourceId(v string) *CreatePersonalFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetStatus(v string) *CreatePersonalFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalFileResponseBody) Validate() error {
	return dara.Validate(s)
}
