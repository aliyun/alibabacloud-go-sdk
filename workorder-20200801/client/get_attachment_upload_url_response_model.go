// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttachmentUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttachmentUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAttachmentUploadUrlResponseBody) *GetAttachmentUploadUrlResponse
	GetBody() *GetAttachmentUploadUrlResponseBody
}

type GetAttachmentUploadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttachmentUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttachmentUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttachmentUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttachmentUploadUrlResponse) GetBody() *GetAttachmentUploadUrlResponseBody {
	return s.Body
}

func (s *GetAttachmentUploadUrlResponse) SetHeaders(v map[string]*string) *GetAttachmentUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAttachmentUploadUrlResponse) SetStatusCode(v int32) *GetAttachmentUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachmentUploadUrlResponse) SetBody(v *GetAttachmentUploadUrlResponseBody) *GetAttachmentUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GetAttachmentUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
