// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnswerSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAnswerSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAnswerSampleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAnswerSampleResponseBody) *DeleteAnswerSampleResponse
	GetBody() *DeleteAnswerSampleResponseBody
}

type DeleteAnswerSampleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAnswerSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAnswerSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnswerSampleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnswerSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAnswerSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAnswerSampleResponse) GetBody() *DeleteAnswerSampleResponseBody {
	return s.Body
}

func (s *DeleteAnswerSampleResponse) SetHeaders(v map[string]*string) *DeleteAnswerSampleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnswerSampleResponse) SetStatusCode(v int32) *DeleteAnswerSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAnswerSampleResponse) SetBody(v *DeleteAnswerSampleResponseBody) *DeleteAnswerSampleResponse {
	s.Body = v
	return s
}

func (s *DeleteAnswerSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
