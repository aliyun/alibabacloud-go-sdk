// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameKnowledgeBaseSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenameKnowledgeBaseSourceResponseBody
	GetCode() *string
	SetGmtModified(v string) *RenameKnowledgeBaseSourceResponseBody
	GetGmtModified() *string
	SetMessage(v string) *RenameKnowledgeBaseSourceResponseBody
	GetMessage() *string
	SetName(v string) *RenameKnowledgeBaseSourceResponseBody
	GetName() *string
	SetRequestId(v string) *RenameKnowledgeBaseSourceResponseBody
	GetRequestId() *string
	SetSourceId(v string) *RenameKnowledgeBaseSourceResponseBody
	GetSourceId() *string
	SetStatus(v string) *RenameKnowledgeBaseSourceResponseBody
	GetStatus() *string
}

type RenameKnowledgeBaseSourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 最近修改时间，ISO8601 格式
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
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
	// 数据源 ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 数据源状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RenameKnowledgeBaseSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameKnowledgeBaseSourceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *RenameKnowledgeBaseSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetCode(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.Code = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetGmtModified(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetMessage(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.Message = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetName(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.Name = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetRequestId(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetSourceId(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) SetStatus(v string) *RenameKnowledgeBaseSourceResponseBody {
	s.Status = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
