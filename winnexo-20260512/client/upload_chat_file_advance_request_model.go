// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadChatFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *UploadChatFileAdvanceRequest
	GetContentType() *string
	SetFileName(v string) *UploadChatFileAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *UploadChatFileAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOperatingObjectName(v string) *UploadChatFileAdvanceRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *UploadChatFileAdvanceRequest
	GetTenantId() *string
}

type UploadChatFileAdvanceRequest struct {
	// 文件 MIME 类型（可选，不传时按 application/octet-stream 处理）
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 原始文件名（含后缀，如 report.pdf）。中转生成的 OSS 地址不携带原始文件名，后端据此确定文件后缀与展示名
	//
	// This parameter is required.
	//
	// example:
	//
	// report.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件的 OSS 地址。使用 SDK 的 UploadChatFileAdvance 方法时由 SDK 中转上传后自动回填；直接调用本 API 时需自行传入可被服务端访问的 OSS 地址
	//
	// This parameter is required.
	//
	// example:
	//
	// http://winnexo-file-transfer.oss-cn-hangzhou.aliyuncs.com/openapi/2026-08-06/9f8c2a1b
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Agent 命名空间标识
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UploadChatFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadChatFileAdvanceRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *UploadChatFileAdvanceRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UploadChatFileAdvanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadChatFileAdvanceRequest) SetContentType(v string) *UploadChatFileAdvanceRequest {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetFileName(v string) *UploadChatFileAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetFileUrlObject(v io.Reader) *UploadChatFileAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetOperatingObjectName(v string) *UploadChatFileAdvanceRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetTenantId(v string) *UploadChatFileAdvanceRequest {
	s.TenantId = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
