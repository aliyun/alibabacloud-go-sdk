// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlbumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlbumResponse
	GetStatusCode() *int32
	SetBody(v *GetAlbumResponseBody) *GetAlbumResponse
	GetBody() *GetAlbumResponseBody
}

type GetAlbumResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlbumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlbumResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumResponse) GoString() string {
	return s.String()
}

func (s *GetAlbumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlbumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlbumResponse) GetBody() *GetAlbumResponseBody {
	return s.Body
}

func (s *GetAlbumResponse) SetHeaders(v map[string]*string) *GetAlbumResponse {
	s.Headers = v
	return s
}

func (s *GetAlbumResponse) SetStatusCode(v int32) *GetAlbumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlbumResponse) SetBody(v *GetAlbumResponseBody) *GetAlbumResponse {
	s.Body = v
	return s
}

func (s *GetAlbumResponse) Validate() error {
	return dara.Validate(s)
}
