// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpenSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpenSearchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpenSearchResponseBody) *DeleteOpenSearchResponse
	GetBody() *DeleteOpenSearchResponseBody
}

type DeleteOpenSearchResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpenSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpenSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpenSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpenSearchResponse) GetBody() *DeleteOpenSearchResponseBody {
	return s.Body
}

func (s *DeleteOpenSearchResponse) SetHeaders(v map[string]*string) *DeleteOpenSearchResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpenSearchResponse) SetStatusCode(v int32) *DeleteOpenSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpenSearchResponse) SetBody(v *DeleteOpenSearchResponseBody) *DeleteOpenSearchResponse {
	s.Body = v
	return s
}

func (s *DeleteOpenSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
