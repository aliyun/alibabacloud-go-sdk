// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupTextResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupTextResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupTextResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupTextResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupTextResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupTextResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupTextResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupTextResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupTextResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupTextResponseBody
	GetStatus() *string
}

type CreateGroupTextResponseBody struct {
	// 业务状态码，成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 解析并绑定的真实目录ID
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 创建时间，ISO8601格式
	//
	// example:
	//
	// example
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 协作空间ID
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 错误描述
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Provider处理后的实际资料名称
	//
	// example:
	//
	// 项目资料
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 资料范围，固定GROUP
	//
	// example:
	//
	// example
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 新建资料ID
	//
	// example:
	//
	// example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 实际资料状态；RUNNING表示处理中，FAILED表示创建处理失败
	//
	// example:
	//
	// example
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupTextResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupTextResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupTextResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupTextResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupTextResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupTextResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupTextResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupTextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupTextResponseBody) SetCode(v string) *CreateGroupTextResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetDirectoryId(v string) *CreateGroupTextResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetGmtCreate(v string) *CreateGroupTextResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetGroupId(v string) *CreateGroupTextResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetMessage(v string) *CreateGroupTextResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetName(v string) *CreateGroupTextResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetRequestId(v string) *CreateGroupTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetScope(v string) *CreateGroupTextResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetSourceId(v string) *CreateGroupTextResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupTextResponseBody) SetStatus(v string) *CreateGroupTextResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupTextResponseBody) Validate() error {
	return dara.Validate(s)
}
