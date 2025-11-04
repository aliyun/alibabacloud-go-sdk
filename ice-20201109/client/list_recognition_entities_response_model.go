// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecognitionEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecognitionEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecognitionEntitiesResponseBody) *ListRecognitionEntitiesResponse
	GetBody() *ListRecognitionEntitiesResponseBody
}

type ListRecognitionEntitiesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecognitionEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecognitionEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListRecognitionEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecognitionEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecognitionEntitiesResponse) GetBody() *ListRecognitionEntitiesResponseBody {
	return s.Body
}

func (s *ListRecognitionEntitiesResponse) SetHeaders(v map[string]*string) *ListRecognitionEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListRecognitionEntitiesResponse) SetStatusCode(v int32) *ListRecognitionEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecognitionEntitiesResponse) SetBody(v *ListRecognitionEntitiesResponseBody) *ListRecognitionEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListRecognitionEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
