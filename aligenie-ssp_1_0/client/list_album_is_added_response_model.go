// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumIsAddedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlbumIsAddedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlbumIsAddedResponse
	GetStatusCode() *int32
	SetBody(v *ListAlbumIsAddedResponseBody) *ListAlbumIsAddedResponse
	GetBody() *ListAlbumIsAddedResponseBody
}

type ListAlbumIsAddedResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlbumIsAddedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlbumIsAddedResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedResponse) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlbumIsAddedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlbumIsAddedResponse) GetBody() *ListAlbumIsAddedResponseBody {
	return s.Body
}

func (s *ListAlbumIsAddedResponse) SetHeaders(v map[string]*string) *ListAlbumIsAddedResponse {
	s.Headers = v
	return s
}

func (s *ListAlbumIsAddedResponse) SetStatusCode(v int32) *ListAlbumIsAddedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlbumIsAddedResponse) SetBody(v *ListAlbumIsAddedResponseBody) *ListAlbumIsAddedResponse {
	s.Body = v
	return s
}

func (s *ListAlbumIsAddedResponse) Validate() error {
	return dara.Validate(s)
}
