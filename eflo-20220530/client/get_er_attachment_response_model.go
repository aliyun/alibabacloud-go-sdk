// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *GetErAttachmentResponseBody) *GetErAttachmentResponse
	GetBody() *GetErAttachmentResponseBody
}

type GetErAttachmentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErAttachmentResponse) GetBody() *GetErAttachmentResponseBody {
	return s.Body
}

func (s *GetErAttachmentResponse) SetHeaders(v map[string]*string) *GetErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetErAttachmentResponse) SetStatusCode(v int32) *GetErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErAttachmentResponse) SetBody(v *GetErAttachmentResponseBody) *GetErAttachmentResponse {
	s.Body = v
	return s
}

func (s *GetErAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
