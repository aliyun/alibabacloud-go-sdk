// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkitemAttachmentCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WorkitemAttachmentCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WorkitemAttachmentCreateResponse
	GetStatusCode() *int32
	SetBody(v *WorkitemAttachmentCreateResponseBody) *WorkitemAttachmentCreateResponse
	GetBody() *WorkitemAttachmentCreateResponseBody
}

type WorkitemAttachmentCreateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WorkitemAttachmentCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WorkitemAttachmentCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s WorkitemAttachmentCreateResponse) GoString() string {
	return s.String()
}

func (s *WorkitemAttachmentCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WorkitemAttachmentCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WorkitemAttachmentCreateResponse) GetBody() *WorkitemAttachmentCreateResponseBody {
	return s.Body
}

func (s *WorkitemAttachmentCreateResponse) SetHeaders(v map[string]*string) *WorkitemAttachmentCreateResponse {
	s.Headers = v
	return s
}

func (s *WorkitemAttachmentCreateResponse) SetStatusCode(v int32) *WorkitemAttachmentCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *WorkitemAttachmentCreateResponse) SetBody(v *WorkitemAttachmentCreateResponseBody) *WorkitemAttachmentCreateResponse {
	s.Body = v
	return s
}

func (s *WorkitemAttachmentCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
