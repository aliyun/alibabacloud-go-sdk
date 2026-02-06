// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGlobalQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGlobalQuestionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGlobalQuestionResponseBody) *DeleteGlobalQuestionResponse
	GetBody() *DeleteGlobalQuestionResponseBody
}

type DeleteGlobalQuestionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalQuestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGlobalQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGlobalQuestionResponse) GetBody() *DeleteGlobalQuestionResponseBody {
	return s.Body
}

func (s *DeleteGlobalQuestionResponse) SetHeaders(v map[string]*string) *DeleteGlobalQuestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalQuestionResponse) SetStatusCode(v int32) *DeleteGlobalQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalQuestionResponse) SetBody(v *DeleteGlobalQuestionResponseBody) *DeleteGlobalQuestionResponse {
	s.Body = v
	return s
}

func (s *DeleteGlobalQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
