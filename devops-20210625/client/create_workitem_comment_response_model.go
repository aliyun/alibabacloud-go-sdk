// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemCommentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkitemCommentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkitemCommentResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkitemCommentResponseBody) *CreateWorkitemCommentResponse
	GetBody() *CreateWorkitemCommentResponseBody
}

type CreateWorkitemCommentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkitemCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkitemCommentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemCommentResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkitemCommentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkitemCommentResponse) GetBody() *CreateWorkitemCommentResponseBody {
	return s.Body
}

func (s *CreateWorkitemCommentResponse) SetHeaders(v map[string]*string) *CreateWorkitemCommentResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemCommentResponse) SetStatusCode(v int32) *CreateWorkitemCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemCommentResponse) SetBody(v *CreateWorkitemCommentResponseBody) *CreateWorkitemCommentResponse {
	s.Body = v
	return s
}

func (s *CreateWorkitemCommentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
