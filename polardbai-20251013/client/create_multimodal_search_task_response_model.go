// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultimodalSearchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultimodalSearchTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultimodalSearchTaskResponseBody) *CreateMultimodalSearchTaskResponse
	GetBody() *CreateMultimodalSearchTaskResponseBody
}

type CreateMultimodalSearchTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultimodalSearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultimodalSearchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultimodalSearchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultimodalSearchTaskResponse) GetBody() *CreateMultimodalSearchTaskResponseBody {
	return s.Body
}

func (s *CreateMultimodalSearchTaskResponse) SetHeaders(v map[string]*string) *CreateMultimodalSearchTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMultimodalSearchTaskResponse) SetStatusCode(v int32) *CreateMultimodalSearchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultimodalSearchTaskResponse) SetBody(v *CreateMultimodalSearchTaskResponseBody) *CreateMultimodalSearchTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMultimodalSearchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
