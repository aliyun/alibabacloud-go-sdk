// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCommentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCommentResponse
	GetStatusCode() *int32
	SetBody(v *CreateCommentResponseBody) *CreateCommentResponse
	GetBody() *CreateCommentResponseBody
}

type CreateCommentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCommentResponse) GoString() string {
	return s.String()
}

func (s *CreateCommentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCommentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCommentResponse) GetBody() *CreateCommentResponseBody {
	return s.Body
}

func (s *CreateCommentResponse) SetHeaders(v map[string]*string) *CreateCommentResponse {
	s.Headers = v
	return s
}

func (s *CreateCommentResponse) SetStatusCode(v int32) *CreateCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommentResponse) SetBody(v *CreateCommentResponseBody) *CreateCommentResponse {
	s.Body = v
	return s
}

func (s *CreateCommentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
