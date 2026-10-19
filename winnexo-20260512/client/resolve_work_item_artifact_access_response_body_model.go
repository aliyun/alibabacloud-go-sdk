// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemArtifactAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResolveWorkItemArtifactAccessResponseBody
	GetCode() *string
	SetDownloadUrl(v string) *ResolveWorkItemArtifactAccessResponseBody
	GetDownloadUrl() *string
	SetExpiresAt(v string) *ResolveWorkItemArtifactAccessResponseBody
	GetExpiresAt() *string
	SetMessage(v string) *ResolveWorkItemArtifactAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResolveWorkItemArtifactAccessResponseBody
	GetRequestId() *string
}

type ResolveWorkItemArtifactAccessResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 产物短期签名下载地址，过期前可直接 GET
	//
	// example:
	//
	// example
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// 地址过期时间(ISO8601)
	//
	// example:
	//
	// example
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ResolveWorkItemArtifactAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemArtifactAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemArtifactAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResolveWorkItemArtifactAccessResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ResolveWorkItemArtifactAccessResponseBody) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *ResolveWorkItemArtifactAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResolveWorkItemArtifactAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResolveWorkItemArtifactAccessResponseBody) SetCode(v string) *ResolveWorkItemArtifactAccessResponseBody {
	s.Code = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponseBody) SetDownloadUrl(v string) *ResolveWorkItemArtifactAccessResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponseBody) SetExpiresAt(v string) *ResolveWorkItemArtifactAccessResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponseBody) SetMessage(v string) *ResolveWorkItemArtifactAccessResponseBody {
	s.Message = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponseBody) SetRequestId(v string) *ResolveWorkItemArtifactAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
