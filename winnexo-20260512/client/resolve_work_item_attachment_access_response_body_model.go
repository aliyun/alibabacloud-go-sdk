// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemAttachmentAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessUrl(v string) *ResolveWorkItemAttachmentAccessResponseBody
	GetAccessUrl() *string
	SetCode(v string) *ResolveWorkItemAttachmentAccessResponseBody
	GetCode() *string
	SetExpiresAt(v string) *ResolveWorkItemAttachmentAccessResponseBody
	GetExpiresAt() *string
	SetMessage(v string) *ResolveWorkItemAttachmentAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResolveWorkItemAttachmentAccessResponseBody
	GetRequestId() *string
}

type ResolveWorkItemAttachmentAccessResponseBody struct {
	// 附件短期预览地址，过期前可直接 GET
	//
	// example:
	//
	// example
	AccessUrl *string `json:"accessUrl,omitempty" xml:"accessUrl,omitempty"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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

func (s ResolveWorkItemAttachmentAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemAttachmentAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) GetAccessUrl() *string {
	return s.AccessUrl
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) SetAccessUrl(v string) *ResolveWorkItemAttachmentAccessResponseBody {
	s.AccessUrl = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) SetCode(v string) *ResolveWorkItemAttachmentAccessResponseBody {
	s.Code = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) SetExpiresAt(v string) *ResolveWorkItemAttachmentAccessResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) SetMessage(v string) *ResolveWorkItemAttachmentAccessResponseBody {
	s.Message = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) SetRequestId(v string) *ResolveWorkItemAttachmentAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
