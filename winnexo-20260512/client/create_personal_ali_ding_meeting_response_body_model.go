// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingMeetingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalAliDingMeetingResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAliDingMeetingResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalAliDingMeetingResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalAliDingMeetingResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAliDingMeetingResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalAliDingMeetingResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalAliDingMeetingResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalAliDingMeetingResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalAliDingMeetingResponseBody
	GetStatus() *string
}

type CreatePersonalAliDingMeetingResponseBody struct {
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

func (s CreatePersonalAliDingMeetingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetCode(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetDirectoryId(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetGmtCreate(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetMessage(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetName(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetRequestId(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetScope(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetSourceId(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetStatus(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) Validate() error {
	return dara.Validate(s)
}
