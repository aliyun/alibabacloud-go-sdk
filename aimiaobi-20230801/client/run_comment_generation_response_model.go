// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommentGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCommentGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCommentGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunCommentGenerationResponseBody) *RunCommentGenerationResponse
	GetBody() *RunCommentGenerationResponseBody
}

type RunCommentGenerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommentGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommentGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCommentGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCommentGenerationResponse) GetBody() *RunCommentGenerationResponseBody {
	return s.Body
}

func (s *RunCommentGenerationResponse) SetHeaders(v map[string]*string) *RunCommentGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunCommentGenerationResponse) SetStatusCode(v int32) *RunCommentGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommentGenerationResponse) SetBody(v *RunCommentGenerationResponseBody) *RunCommentGenerationResponse {
	s.Body = v
	return s
}

func (s *RunCommentGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
