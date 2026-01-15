// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalSearchModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalSearchModelResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalSearchModelResponseBody) *ListMultimodalSearchModelResponse
	GetBody() *ListMultimodalSearchModelResponseBody
}

type ListMultimodalSearchModelResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalSearchModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalSearchModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchModelResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalSearchModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalSearchModelResponse) GetBody() *ListMultimodalSearchModelResponseBody {
	return s.Body
}

func (s *ListMultimodalSearchModelResponse) SetHeaders(v map[string]*string) *ListMultimodalSearchModelResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalSearchModelResponse) SetStatusCode(v int32) *ListMultimodalSearchModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalSearchModelResponse) SetBody(v *ListMultimodalSearchModelResponseBody) *ListMultimodalSearchModelResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalSearchModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
