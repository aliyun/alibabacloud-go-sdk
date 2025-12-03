// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemCommentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkitemCommentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkitemCommentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkitemCommentResponseBody) *DeleteWorkitemCommentResponse
	GetBody() *DeleteWorkitemCommentResponseBody
}

type DeleteWorkitemCommentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkitemCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkitemCommentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemCommentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkitemCommentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkitemCommentResponse) GetBody() *DeleteWorkitemCommentResponseBody {
	return s.Body
}

func (s *DeleteWorkitemCommentResponse) SetHeaders(v map[string]*string) *DeleteWorkitemCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkitemCommentResponse) SetStatusCode(v int32) *DeleteWorkitemCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkitemCommentResponse) SetBody(v *DeleteWorkitemCommentResponseBody) *DeleteWorkitemCommentResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkitemCommentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
