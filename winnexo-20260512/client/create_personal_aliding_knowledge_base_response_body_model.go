// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetGmtCreate() *string
	SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetKbUrl() *string
	SetMessage(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetOperatingObjectName() *string
	SetRequestId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetStatus() *string
}

type CreatePersonalAlidingKnowledgeBaseResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 新建知识库根目录 ID
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
	// 知识库 URL（echo 回入参，便于调用方对齐）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	KbUrl *string `json:"kbUrl,omitempty" xml:"kbUrl,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所属数字员工名称（echo 回入参，可为 null）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 知识库根目录状态（创建后为 RUNNING；后台同步完成后转 READY 或 FAILED）
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetKbUrl() *string {
	return s.KbUrl
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetCode(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetGmtCreate(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.KbUrl = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetMessage(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetRequestId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetStatus(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
