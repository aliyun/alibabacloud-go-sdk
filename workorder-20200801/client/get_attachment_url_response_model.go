// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttachmentUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttachmentUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAttachmentUrlResponseBody) *GetAttachmentUrlResponse
	GetBody() *GetAttachmentUrlResponseBody
}

type GetAttachmentUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttachmentUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttachmentUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAttachmentUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttachmentUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttachmentUrlResponse) GetBody() *GetAttachmentUrlResponseBody {
	return s.Body
}

func (s *GetAttachmentUrlResponse) SetHeaders(v map[string]*string) *GetAttachmentUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAttachmentUrlResponse) SetStatusCode(v int32) *GetAttachmentUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachmentUrlResponse) SetBody(v *GetAttachmentUrlResponseBody) *GetAttachmentUrlResponse {
	s.Body = v
	return s
}

func (s *GetAttachmentUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
