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
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The short-term signed download URL of the artifact. You can send a GET request directly to this URL before it expires.
	//
	// example:
	//
	// example
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
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
	// The request trace ID.
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
