// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseSourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetCode() *string
	SetFilePath(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetFilePath() *string
	SetFilePublicUrl(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetFileRecordId() *string
	SetMessage(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetMessage() *string
	SetName(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetName() *string
	SetRequestId(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetSourceId() *string
	SetSourceType(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetSourceType() *string
	SetStatus(v string) *ReplaceKnowledgeBaseSourceFileResponseBody
	GetStatus() *string
}

type ReplaceKnowledgeBaseSourceFileResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 替换后的文件 OSS 持久化地址
	//
	// example:
	//
	// string_value
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 替换后的文件公开访问 URL
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 替换后的文件记录 ID
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
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
	// 数据源类型，固定为 FILE
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 重新解析后的数据源状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetFilePath() *string {
	return s.FilePath
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetName() *string {
	return s.Name
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetCode(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetFilePath(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.FilePath = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetFilePublicUrl(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetFileRecordId(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.FileRecordId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetMessage(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetName(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.Name = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetRequestId(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetSourceId(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetSourceType(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.SourceType = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) SetStatus(v string) *ReplaceKnowledgeBaseSourceFileResponseBody {
	s.Status = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
