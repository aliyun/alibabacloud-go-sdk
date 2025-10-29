// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPlaylistItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPlaylistItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPlaylistItemsResponse
	GetStatusCode() *int32
	SetBody(v *AddPlaylistItemsResponseBody) *AddPlaylistItemsResponse
	GetBody() *AddPlaylistItemsResponseBody
}

type AddPlaylistItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPlaylistItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPlaylistItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPlaylistItemsResponse) GoString() string {
	return s.String()
}

func (s *AddPlaylistItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPlaylistItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPlaylistItemsResponse) GetBody() *AddPlaylistItemsResponseBody {
	return s.Body
}

func (s *AddPlaylistItemsResponse) SetHeaders(v map[string]*string) *AddPlaylistItemsResponse {
	s.Headers = v
	return s
}

func (s *AddPlaylistItemsResponse) SetStatusCode(v int32) *AddPlaylistItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPlaylistItemsResponse) SetBody(v *AddPlaylistItemsResponseBody) *AddPlaylistItemsResponse {
	s.Body = v
	return s
}

func (s *AddPlaylistItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
