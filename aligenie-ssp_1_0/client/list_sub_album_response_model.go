// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAlbumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubAlbumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubAlbumResponse
	GetStatusCode() *int32
	SetBody(v *ListSubAlbumResponseBody) *ListSubAlbumResponse
	GetBody() *ListSubAlbumResponseBody
}

type ListSubAlbumResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubAlbumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubAlbumResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumResponse) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubAlbumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubAlbumResponse) GetBody() *ListSubAlbumResponseBody {
	return s.Body
}

func (s *ListSubAlbumResponse) SetHeaders(v map[string]*string) *ListSubAlbumResponse {
	s.Headers = v
	return s
}

func (s *ListSubAlbumResponse) SetStatusCode(v int32) *ListSubAlbumResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubAlbumResponse) SetBody(v *ListSubAlbumResponseBody) *ListSubAlbumResponse {
	s.Body = v
	return s
}

func (s *ListSubAlbumResponse) Validate() error {
	return dara.Validate(s)
}
