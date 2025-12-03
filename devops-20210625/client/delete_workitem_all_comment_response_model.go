// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemAllCommentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkitemAllCommentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkitemAllCommentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkitemAllCommentResponseBody) *DeleteWorkitemAllCommentResponse
	GetBody() *DeleteWorkitemAllCommentResponseBody
}

type DeleteWorkitemAllCommentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkitemAllCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkitemAllCommentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemAllCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemAllCommentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkitemAllCommentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkitemAllCommentResponse) GetBody() *DeleteWorkitemAllCommentResponseBody {
	return s.Body
}

func (s *DeleteWorkitemAllCommentResponse) SetHeaders(v map[string]*string) *DeleteWorkitemAllCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkitemAllCommentResponse) SetStatusCode(v int32) *DeleteWorkitemAllCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponse) SetBody(v *DeleteWorkitemAllCommentResponseBody) *DeleteWorkitemAllCommentResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkitemAllCommentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
