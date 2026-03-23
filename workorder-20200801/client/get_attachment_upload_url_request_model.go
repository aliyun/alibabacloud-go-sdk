// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetAttachmentUploadUrlRequest
	GetFileName() *string
}

type GetAttachmentUploadUrlRequest struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetAttachmentUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetAttachmentUploadUrlRequest) SetFileName(v string) *GetAttachmentUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetAttachmentUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
