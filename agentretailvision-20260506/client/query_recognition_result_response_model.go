// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognitionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecognitionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecognitionResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecognitionResultResponseBody) *QueryRecognitionResultResponse
	GetBody() *QueryRecognitionResultResponseBody
}

type QueryRecognitionResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecognitionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecognitionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultResponse) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecognitionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecognitionResultResponse) GetBody() *QueryRecognitionResultResponseBody {
	return s.Body
}

func (s *QueryRecognitionResultResponse) SetHeaders(v map[string]*string) *QueryRecognitionResultResponse {
	s.Headers = v
	return s
}

func (s *QueryRecognitionResultResponse) SetStatusCode(v int32) *QueryRecognitionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecognitionResultResponse) SetBody(v *QueryRecognitionResultResponseBody) *QueryRecognitionResultResponse {
	s.Body = v
	return s
}

func (s *QueryRecognitionResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
