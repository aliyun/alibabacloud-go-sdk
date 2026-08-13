// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *UploadChatFileRequest
	GetContentType() *string
	SetFileName(v string) *UploadChatFileRequest
	GetFileName() *string
	SetFileUrl(v string) *UploadChatFileRequest
	GetFileUrl() *string
	SetOperatingObjectName(v string) *UploadChatFileRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *UploadChatFileRequest
	GetTenantId() *string
}

type UploadChatFileRequest struct {
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
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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

func (s UploadChatFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileRequest) GoString() string {
	return s.String()
}

func (s *UploadChatFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadChatFileRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UploadChatFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadChatFileRequest) SetContentType(v string) *UploadChatFileRequest {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileRequest) SetFileName(v string) *UploadChatFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadChatFileRequest) SetFileUrl(v string) *UploadChatFileRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadChatFileRequest) SetOperatingObjectName(v string) *UploadChatFileRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UploadChatFileRequest) SetTenantId(v string) *UploadChatFileRequest {
	s.TenantId = &v
	return s
}

func (s *UploadChatFileRequest) Validate() error {
	return dara.Validate(s)
}
