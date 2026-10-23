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
	// The short-term preview URL of the attachment. You can send a GET request to this URL before it expires.
	//
	// example:
	//
	// example
	AccessUrl *string `json:"accessUrl,omitempty" xml:"accessUrl,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The expiration time of the URL in ISO 8601 format.
	//
	// example:
	//
	// example
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
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
