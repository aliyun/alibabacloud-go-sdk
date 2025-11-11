// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveStyleLearningResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveStyleLearningResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveStyleLearningResultResponse
	GetStatusCode() *int32
	SetBody(v *SaveStyleLearningResultResponseBody) *SaveStyleLearningResultResponse
	GetBody() *SaveStyleLearningResultResponseBody
}

type SaveStyleLearningResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveStyleLearningResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveStyleLearningResultResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveStyleLearningResultResponse) GoString() string {
	return s.String()
}

func (s *SaveStyleLearningResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveStyleLearningResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveStyleLearningResultResponse) GetBody() *SaveStyleLearningResultResponseBody {
	return s.Body
}

func (s *SaveStyleLearningResultResponse) SetHeaders(v map[string]*string) *SaveStyleLearningResultResponse {
	s.Headers = v
	return s
}

func (s *SaveStyleLearningResultResponse) SetStatusCode(v int32) *SaveStyleLearningResultResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveStyleLearningResultResponse) SetBody(v *SaveStyleLearningResultResponseBody) *SaveStyleLearningResultResponse {
	s.Body = v
	return s
}

func (s *SaveStyleLearningResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
