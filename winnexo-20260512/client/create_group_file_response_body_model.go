// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupFileResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupFileResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupFileResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupFileResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupFileResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupFileResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupFileResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupFileResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupFileResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupFileResponseBody
	GetStatus() *string
}

type CreateGroupFileResponseBody struct {
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
	// ok
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
	// E68654BD-F7BA-5837-8686-5645D739A47C
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

func (s CreateGroupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupFileResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFileResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupFileResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupFileResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupFileResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupFileResponseBody) SetCode(v string) *CreateGroupFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetDirectoryId(v string) *CreateGroupFileResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetGmtCreate(v string) *CreateGroupFileResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetGroupId(v string) *CreateGroupFileResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetMessage(v string) *CreateGroupFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetName(v string) *CreateGroupFileResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetRequestId(v string) *CreateGroupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetScope(v string) *CreateGroupFileResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetSourceId(v string) *CreateGroupFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupFileResponseBody) SetStatus(v string) *CreateGroupFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupFileResponseBody) Validate() error {
	return dara.Validate(s)
}
