// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultimodalSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultimodalSearchResponse
	GetStatusCode() *int32
	SetBody(v *MultimodalSearchOutput) *MultimodalSearchResponse
	GetBody() *MultimodalSearchOutput
}

type MultimodalSearchResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultimodalSearchOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultimodalSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s MultimodalSearchResponse) GoString() string {
	return s.String()
}

func (s *MultimodalSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultimodalSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultimodalSearchResponse) GetBody() *MultimodalSearchOutput {
	return s.Body
}

func (s *MultimodalSearchResponse) SetHeaders(v map[string]*string) *MultimodalSearchResponse {
	s.Headers = v
	return s
}

func (s *MultimodalSearchResponse) SetStatusCode(v int32) *MultimodalSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *MultimodalSearchResponse) SetBody(v *MultimodalSearchOutput) *MultimodalSearchResponse {
	s.Body = v
	return s
}

func (s *MultimodalSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
