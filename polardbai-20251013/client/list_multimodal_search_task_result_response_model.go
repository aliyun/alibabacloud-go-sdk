// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalSearchTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalSearchTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalSearchTaskResultResponseBody) *ListMultimodalSearchTaskResultResponse
	GetBody() *ListMultimodalSearchTaskResultResponseBody
}

type ListMultimodalSearchTaskResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalSearchTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalSearchTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskResultResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalSearchTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalSearchTaskResultResponse) GetBody() *ListMultimodalSearchTaskResultResponseBody {
	return s.Body
}

func (s *ListMultimodalSearchTaskResultResponse) SetHeaders(v map[string]*string) *ListMultimodalSearchTaskResultResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalSearchTaskResultResponse) SetStatusCode(v int32) *ListMultimodalSearchTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalSearchTaskResultResponse) SetBody(v *ListMultimodalSearchTaskResultResponseBody) *ListMultimodalSearchTaskResultResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalSearchTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
