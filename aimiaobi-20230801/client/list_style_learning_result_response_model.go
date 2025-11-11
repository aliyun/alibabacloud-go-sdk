// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStyleLearningResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStyleLearningResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStyleLearningResultResponse
	GetStatusCode() *int32
	SetBody(v *ListStyleLearningResultResponseBody) *ListStyleLearningResultResponse
	GetBody() *ListStyleLearningResultResponseBody
}

type ListStyleLearningResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStyleLearningResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStyleLearningResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStyleLearningResultResponse) GoString() string {
	return s.String()
}

func (s *ListStyleLearningResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStyleLearningResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStyleLearningResultResponse) GetBody() *ListStyleLearningResultResponseBody {
	return s.Body
}

func (s *ListStyleLearningResultResponse) SetHeaders(v map[string]*string) *ListStyleLearningResultResponse {
	s.Headers = v
	return s
}

func (s *ListStyleLearningResultResponse) SetStatusCode(v int32) *ListStyleLearningResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStyleLearningResultResponse) SetBody(v *ListStyleLearningResultResponseBody) *ListStyleLearningResultResponse {
	s.Body = v
	return s
}

func (s *ListStyleLearningResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
