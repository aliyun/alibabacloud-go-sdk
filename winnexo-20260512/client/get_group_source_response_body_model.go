// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGroupSourceResponseBody
	GetCode() *string
	SetDescription(v string) *GetGroupSourceResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *GetGroupSourceResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetGroupSourceResponseBody
	GetGmtModified() *string
	SetGroupId(v string) *GetGroupSourceResponseBody
	GetGroupId() *string
	SetMessage(v string) *GetGroupSourceResponseBody
	GetMessage() *string
	SetName(v string) *GetGroupSourceResponseBody
	GetName() *string
	SetRequestId(v string) *GetGroupSourceResponseBody
	GetRequestId() *string
	SetScope(v string) *GetGroupSourceResponseBody
	GetScope() *string
	SetSourceId(v string) *GetGroupSourceResponseBody
	GetSourceId() *string
	SetSourceKind(v string) *GetGroupSourceResponseBody
	GetSourceKind() *string
	SetSourceTags(v string) *GetGroupSourceResponseBody
	GetSourceTags() *string
	SetSourceType(v string) *GetGroupSourceResponseBody
	GetSourceType() *string
	SetStatus(v string) *GetGroupSourceResponseBody
	GetStatus() *string
}

type GetGroupSourceResponseBody struct {
	// 业务状态码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 资料描述
	//
	// example:
	//
	// recorder function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间，ISO8601格式
	//
	// example:
	//
	// 2026-08-26T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间，ISO8601格式
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 本次授权读取的协作空间ID
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 错误描述
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 资料名称
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 资料实际范围；引用资料保留 PERSONAL 或 TENANT
	//
	// example:
	//
	// GROUP
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 资料ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 知识归属类型，沿用 Source 分类
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// 资料标签JSON字符串列表
	//
	// example:
	//
	// ["重点","文档"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// 资料类型，例如 TEXT、FILE、ONLINE_DOC、FEISHU
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 当前资料状态，例如 READY、RUNNING、FAILED
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetGroupSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGroupSourceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetGroupSourceResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGroupSourceResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGroupSourceResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGroupSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetGroupSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupSourceResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetGroupSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetGroupSourceResponseBody) GetSourceKind() *string {
	return s.SourceKind
}

func (s *GetGroupSourceResponseBody) GetSourceTags() *string {
	return s.SourceTags
}

func (s *GetGroupSourceResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetGroupSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetGroupSourceResponseBody) SetCode(v string) *GetGroupSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetDescription(v string) *GetGroupSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetGmtCreate(v string) *GetGroupSourceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetGmtModified(v string) *GetGroupSourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetGroupId(v string) *GetGroupSourceResponseBody {
	s.GroupId = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetMessage(v string) *GetGroupSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetName(v string) *GetGroupSourceResponseBody {
	s.Name = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetRequestId(v string) *GetGroupSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetScope(v string) *GetGroupSourceResponseBody {
	s.Scope = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceId(v string) *GetGroupSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceKind(v string) *GetGroupSourceResponseBody {
	s.SourceKind = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceTags(v string) *GetGroupSourceResponseBody {
	s.SourceTags = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceType(v string) *GetGroupSourceResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetStatus(v string) *GetGroupSourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetGroupSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
