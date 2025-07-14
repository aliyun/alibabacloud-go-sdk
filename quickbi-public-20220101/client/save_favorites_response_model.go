// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFavoritesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveFavoritesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveFavoritesResponse
	GetStatusCode() *int32
	SetBody(v *SaveFavoritesResponseBody) *SaveFavoritesResponse
	GetBody() *SaveFavoritesResponseBody
}

type SaveFavoritesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFavoritesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveFavoritesResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveFavoritesResponse) GoString() string {
	return s.String()
}

func (s *SaveFavoritesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveFavoritesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveFavoritesResponse) GetBody() *SaveFavoritesResponseBody {
	return s.Body
}

func (s *SaveFavoritesResponse) SetHeaders(v map[string]*string) *SaveFavoritesResponse {
	s.Headers = v
	return s
}

func (s *SaveFavoritesResponse) SetStatusCode(v int32) *SaveFavoritesResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFavoritesResponse) SetBody(v *SaveFavoritesResponseBody) *SaveFavoritesResponse {
	s.Body = v
	return s
}

func (s *SaveFavoritesResponse) Validate() error {
	return dara.Validate(s)
}
