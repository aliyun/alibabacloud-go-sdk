// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaylistItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPlaylistItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPlaylistItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListPlaylistItemsResponseBody) *ListPlaylistItemsResponse
	GetBody() *ListPlaylistItemsResponseBody
}

type ListPlaylistItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlaylistItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlaylistItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistItemsResponse) GoString() string {
	return s.String()
}

func (s *ListPlaylistItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPlaylistItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPlaylistItemsResponse) GetBody() *ListPlaylistItemsResponseBody {
	return s.Body
}

func (s *ListPlaylistItemsResponse) SetHeaders(v map[string]*string) *ListPlaylistItemsResponse {
	s.Headers = v
	return s
}

func (s *ListPlaylistItemsResponse) SetStatusCode(v int32) *ListPlaylistItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlaylistItemsResponse) SetBody(v *ListPlaylistItemsResponseBody) *ListPlaylistItemsResponse {
	s.Body = v
	return s
}

func (s *ListPlaylistItemsResponse) Validate() error {
	return dara.Validate(s)
}
