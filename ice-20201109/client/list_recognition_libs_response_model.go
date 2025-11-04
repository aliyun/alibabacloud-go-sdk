// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionLibsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecognitionLibsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecognitionLibsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecognitionLibsResponseBody) *ListRecognitionLibsResponse
	GetBody() *ListRecognitionLibsResponseBody
}

type ListRecognitionLibsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecognitionLibsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecognitionLibsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionLibsResponse) GoString() string {
	return s.String()
}

func (s *ListRecognitionLibsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecognitionLibsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecognitionLibsResponse) GetBody() *ListRecognitionLibsResponseBody {
	return s.Body
}

func (s *ListRecognitionLibsResponse) SetHeaders(v map[string]*string) *ListRecognitionLibsResponse {
	s.Headers = v
	return s
}

func (s *ListRecognitionLibsResponse) SetStatusCode(v int32) *ListRecognitionLibsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecognitionLibsResponse) SetBody(v *ListRecognitionLibsResponseBody) *ListRecognitionLibsResponse {
	s.Body = v
	return s
}

func (s *ListRecognitionLibsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
