// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStyleLearningResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStyleLearningResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStyleLearningResultResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStyleLearningResultResponseBody) *DeleteStyleLearningResultResponse
	GetBody() *DeleteStyleLearningResultResponseBody
}

type DeleteStyleLearningResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStyleLearningResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStyleLearningResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStyleLearningResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteStyleLearningResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStyleLearningResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStyleLearningResultResponse) GetBody() *DeleteStyleLearningResultResponseBody {
	return s.Body
}

func (s *DeleteStyleLearningResultResponse) SetHeaders(v map[string]*string) *DeleteStyleLearningResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteStyleLearningResultResponse) SetStatusCode(v int32) *DeleteStyleLearningResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStyleLearningResultResponse) SetBody(v *DeleteStyleLearningResultResponseBody) *DeleteStyleLearningResultResponse {
	s.Body = v
	return s
}

func (s *DeleteStyleLearningResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
