// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperatingObjectFavoritesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperatingObjectFavoritesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperatingObjectFavoritesResponse
	GetStatusCode() *int32
	SetBody(v *ListOperatingObjectFavoritesResponseBody) *ListOperatingObjectFavoritesResponse
	GetBody() *ListOperatingObjectFavoritesResponseBody
}

type ListOperatingObjectFavoritesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperatingObjectFavoritesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperatingObjectFavoritesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperatingObjectFavoritesResponse) GoString() string {
	return s.String()
}

func (s *ListOperatingObjectFavoritesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperatingObjectFavoritesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperatingObjectFavoritesResponse) GetBody() *ListOperatingObjectFavoritesResponseBody {
	return s.Body
}

func (s *ListOperatingObjectFavoritesResponse) SetHeaders(v map[string]*string) *ListOperatingObjectFavoritesResponse {
	s.Headers = v
	return s
}

func (s *ListOperatingObjectFavoritesResponse) SetStatusCode(v int32) *ListOperatingObjectFavoritesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponse) SetBody(v *ListOperatingObjectFavoritesResponseBody) *ListOperatingObjectFavoritesResponse {
	s.Body = v
	return s
}

func (s *ListOperatingObjectFavoritesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
