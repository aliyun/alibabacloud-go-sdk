// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalVoiceMeetingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalVoiceMeetingResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalVoiceMeetingResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalVoiceMeetingResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalVoiceMeetingResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalVoiceMeetingResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalVoiceMeetingResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalVoiceMeetingResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalVoiceMeetingResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalVoiceMeetingResponseBody
	GetStatus() *string
}

type CreatePersonalVoiceMeetingResponseBody struct {
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

func (s CreatePersonalVoiceMeetingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalVoiceMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetCode(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetDirectoryId(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetGmtCreate(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetMessage(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetName(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetRequestId(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetScope(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetSourceId(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetStatus(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) Validate() error {
	return dara.Validate(s)
}
