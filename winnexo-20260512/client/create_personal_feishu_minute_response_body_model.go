// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuMinuteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalFeishuMinuteResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalFeishuMinuteResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalFeishuMinuteResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalFeishuMinuteResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalFeishuMinuteResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalFeishuMinuteResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalFeishuMinuteResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalFeishuMinuteResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalFeishuMinuteResponseBody
	GetStatus() *string
}

type CreatePersonalFeishuMinuteResponseBody struct {
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

func (s CreatePersonalFeishuMinuteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuMinuteResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetCode(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetDirectoryId(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetGmtCreate(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetMessage(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetName(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetRequestId(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetScope(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetSourceId(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetStatus(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) Validate() error {
	return dara.Validate(s)
}
