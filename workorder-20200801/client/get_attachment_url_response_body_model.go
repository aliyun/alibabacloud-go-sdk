// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessFileUrl(v string) *GetAttachmentUrlResponseBody
	GetAccessFileUrl() *string
	SetRequestId(v string) *GetAttachmentUrlResponseBody
	GetRequestId() *string
}

type GetAttachmentUrlResponseBody struct {
	AccessFileUrl *string `json:"AccessFileUrl,omitempty" xml:"AccessFileUrl,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttachmentUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachmentUrlResponseBody) GetAccessFileUrl() *string {
	return s.AccessFileUrl
}

func (s *GetAttachmentUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttachmentUrlResponseBody) SetAccessFileUrl(v string) *GetAttachmentUrlResponseBody {
	s.AccessFileUrl = &v
	return s
}

func (s *GetAttachmentUrlResponseBody) SetRequestId(v string) *GetAttachmentUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttachmentUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
