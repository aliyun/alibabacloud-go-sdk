// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumDetailByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlbumDetailByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlbumDetailByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetAlbumDetailByIdResponseBody) *GetAlbumDetailByIdResponse
	GetBody() *GetAlbumDetailByIdResponseBody
}

type GetAlbumDetailByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlbumDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlbumDetailByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlbumDetailByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlbumDetailByIdResponse) GetBody() *GetAlbumDetailByIdResponseBody {
	return s.Body
}

func (s *GetAlbumDetailByIdResponse) SetHeaders(v map[string]*string) *GetAlbumDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAlbumDetailByIdResponse) SetStatusCode(v int32) *GetAlbumDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlbumDetailByIdResponse) SetBody(v *GetAlbumDetailByIdResponseBody) *GetAlbumDetailByIdResponse {
	s.Body = v
	return s
}

func (s *GetAlbumDetailByIdResponse) Validate() error {
	return dara.Validate(s)
}
