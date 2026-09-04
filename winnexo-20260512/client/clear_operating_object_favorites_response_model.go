// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearOperatingObjectFavoritesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearOperatingObjectFavoritesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearOperatingObjectFavoritesResponse
	GetStatusCode() *int32
	SetBody(v *ClearOperatingObjectFavoritesResponseBody) *ClearOperatingObjectFavoritesResponse
	GetBody() *ClearOperatingObjectFavoritesResponseBody
}

type ClearOperatingObjectFavoritesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearOperatingObjectFavoritesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearOperatingObjectFavoritesResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearOperatingObjectFavoritesResponse) GoString() string {
	return s.String()
}

func (s *ClearOperatingObjectFavoritesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearOperatingObjectFavoritesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearOperatingObjectFavoritesResponse) GetBody() *ClearOperatingObjectFavoritesResponseBody {
	return s.Body
}

func (s *ClearOperatingObjectFavoritesResponse) SetHeaders(v map[string]*string) *ClearOperatingObjectFavoritesResponse {
	s.Headers = v
	return s
}

func (s *ClearOperatingObjectFavoritesResponse) SetStatusCode(v int32) *ClearOperatingObjectFavoritesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponse) SetBody(v *ClearOperatingObjectFavoritesResponseBody) *ClearOperatingObjectFavoritesResponse {
	s.Body = v
	return s
}

func (s *ClearOperatingObjectFavoritesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
