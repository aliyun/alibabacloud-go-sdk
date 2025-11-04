// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionSamplesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecognitionSamplesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecognitionSamplesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecognitionSamplesResponseBody) *ListRecognitionSamplesResponse
	GetBody() *ListRecognitionSamplesResponseBody
}

type ListRecognitionSamplesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecognitionSamplesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecognitionSamplesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionSamplesResponse) GoString() string {
	return s.String()
}

func (s *ListRecognitionSamplesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecognitionSamplesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecognitionSamplesResponse) GetBody() *ListRecognitionSamplesResponseBody {
	return s.Body
}

func (s *ListRecognitionSamplesResponse) SetHeaders(v map[string]*string) *ListRecognitionSamplesResponse {
	s.Headers = v
	return s
}

func (s *ListRecognitionSamplesResponse) SetStatusCode(v int32) *ListRecognitionSamplesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecognitionSamplesResponse) SetBody(v *ListRecognitionSamplesResponseBody) *ListRecognitionSamplesResponse {
	s.Body = v
	return s
}

func (s *ListRecognitionSamplesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
