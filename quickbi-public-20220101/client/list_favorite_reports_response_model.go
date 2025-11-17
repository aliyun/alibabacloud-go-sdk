// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFavoriteReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFavoriteReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFavoriteReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListFavoriteReportsResponseBody) *ListFavoriteReportsResponse
	GetBody() *ListFavoriteReportsResponseBody
}

type ListFavoriteReportsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFavoriteReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFavoriteReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFavoriteReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFavoriteReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFavoriteReportsResponse) GetBody() *ListFavoriteReportsResponseBody {
	return s.Body
}

func (s *ListFavoriteReportsResponse) SetHeaders(v map[string]*string) *ListFavoriteReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFavoriteReportsResponse) SetStatusCode(v int32) *ListFavoriteReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFavoriteReportsResponse) SetBody(v *ListFavoriteReportsResponseBody) *ListFavoriteReportsResponse {
	s.Body = v
	return s
}

func (s *ListFavoriteReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
