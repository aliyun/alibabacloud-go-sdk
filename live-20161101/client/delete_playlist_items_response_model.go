// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaylistItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePlaylistItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePlaylistItemsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePlaylistItemsResponseBody) *DeletePlaylistItemsResponse
	GetBody() *DeletePlaylistItemsResponseBody
}

type DeletePlaylistItemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePlaylistItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePlaylistItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaylistItemsResponse) GoString() string {
	return s.String()
}

func (s *DeletePlaylistItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePlaylistItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePlaylistItemsResponse) GetBody() *DeletePlaylistItemsResponseBody {
	return s.Body
}

func (s *DeletePlaylistItemsResponse) SetHeaders(v map[string]*string) *DeletePlaylistItemsResponse {
	s.Headers = v
	return s
}

func (s *DeletePlaylistItemsResponse) SetStatusCode(v int32) *DeletePlaylistItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlaylistItemsResponse) SetBody(v *DeletePlaylistItemsResponseBody) *DeletePlaylistItemsResponse {
	s.Body = v
	return s
}

func (s *DeletePlaylistItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
