// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessFileUrl(v string) *GetAttachmentUploadUrlResponseBody
	GetAccessFileUrl() *string
	SetObjectKey(v string) *GetAttachmentUploadUrlResponseBody
	GetObjectKey() *string
	SetPutObjectUrl(v string) *GetAttachmentUploadUrlResponseBody
	GetPutObjectUrl() *string
	SetRequestId(v string) *GetAttachmentUploadUrlResponseBody
	GetRequestId() *string
}

type GetAttachmentUploadUrlResponseBody struct {
	AccessFileUrl *string `json:"AccessFileUrl,omitempty" xml:"AccessFileUrl,omitempty"`
	ObjectKey     *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	PutObjectUrl  *string `json:"PutObjectUrl,omitempty" xml:"PutObjectUrl,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttachmentUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponseBody) GetAccessFileUrl() *string {
	return s.AccessFileUrl
}

func (s *GetAttachmentUploadUrlResponseBody) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *GetAttachmentUploadUrlResponseBody) GetPutObjectUrl() *string {
	return s.PutObjectUrl
}

func (s *GetAttachmentUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttachmentUploadUrlResponseBody) SetAccessFileUrl(v string) *GetAttachmentUploadUrlResponseBody {
	s.AccessFileUrl = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetObjectKey(v string) *GetAttachmentUploadUrlResponseBody {
	s.ObjectKey = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetPutObjectUrl(v string) *GetAttachmentUploadUrlResponseBody {
	s.PutObjectUrl = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetRequestId(v string) *GetAttachmentUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
