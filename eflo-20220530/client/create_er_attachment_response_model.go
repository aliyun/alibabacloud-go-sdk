// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateErAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateErAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateErAttachmentResponseBody) *CreateErAttachmentResponse
	GetBody() *CreateErAttachmentResponseBody
}

type CreateErAttachmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateErAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateErAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateErAttachmentResponse) GetBody() *CreateErAttachmentResponseBody {
	return s.Body
}

func (s *CreateErAttachmentResponse) SetHeaders(v map[string]*string) *CreateErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateErAttachmentResponse) SetStatusCode(v int32) *CreateErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErAttachmentResponse) SetBody(v *CreateErAttachmentResponseBody) *CreateErAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateErAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
