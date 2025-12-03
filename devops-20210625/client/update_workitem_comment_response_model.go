// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkitemCommentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkitemCommentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkitemCommentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkitemCommentResponseBody) *UpdateWorkitemCommentResponse
	GetBody() *UpdateWorkitemCommentResponseBody
}

type UpdateWorkitemCommentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkitemCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkitemCommentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemCommentResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkitemCommentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkitemCommentResponse) GetBody() *UpdateWorkitemCommentResponseBody {
	return s.Body
}

func (s *UpdateWorkitemCommentResponse) SetHeaders(v map[string]*string) *UpdateWorkitemCommentResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkitemCommentResponse) SetStatusCode(v int32) *UpdateWorkitemCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkitemCommentResponse) SetBody(v *UpdateWorkitemCommentResponseBody) *UpdateWorkitemCommentResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkitemCommentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
