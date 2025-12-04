// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteErAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteErAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteErAttachmentResponseBody) *DeleteErAttachmentResponse
	GetBody() *DeleteErAttachmentResponseBody
}

type DeleteErAttachmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteErAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteErAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteErAttachmentResponse) GetBody() *DeleteErAttachmentResponseBody {
	return s.Body
}

func (s *DeleteErAttachmentResponse) SetHeaders(v map[string]*string) *DeleteErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteErAttachmentResponse) SetStatusCode(v int32) *DeleteErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErAttachmentResponse) SetBody(v *DeleteErAttachmentResponseBody) *DeleteErAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteErAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
