// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateErAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateErAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateErAttachmentResponseBody) *UpdateErAttachmentResponse
	GetBody() *UpdateErAttachmentResponseBody
}

type UpdateErAttachmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateErAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateErAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateErAttachmentResponse) GetBody() *UpdateErAttachmentResponseBody {
	return s.Body
}

func (s *UpdateErAttachmentResponse) SetHeaders(v map[string]*string) *UpdateErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateErAttachmentResponse) SetStatusCode(v int32) *UpdateErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErAttachmentResponse) SetBody(v *UpdateErAttachmentResponseBody) *UpdateErAttachmentResponse {
	s.Body = v
	return s
}

func (s *UpdateErAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
