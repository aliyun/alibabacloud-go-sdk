// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalSearchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalSearchTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalSearchTaskResponseBody) *ListMultimodalSearchTaskResponse
	GetBody() *ListMultimodalSearchTaskResponseBody
}

type ListMultimodalSearchTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalSearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalSearchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalSearchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalSearchTaskResponse) GetBody() *ListMultimodalSearchTaskResponseBody {
	return s.Body
}

func (s *ListMultimodalSearchTaskResponse) SetHeaders(v map[string]*string) *ListMultimodalSearchTaskResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalSearchTaskResponse) SetStatusCode(v int32) *ListMultimodalSearchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalSearchTaskResponse) SetBody(v *ListMultimodalSearchTaskResponseBody) *ListMultimodalSearchTaskResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalSearchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
