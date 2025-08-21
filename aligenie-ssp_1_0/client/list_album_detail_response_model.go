// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlbumDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlbumDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListAlbumDetailResponseBody) *ListAlbumDetailResponse
	GetBody() *ListAlbumDetailResponseBody
}

type ListAlbumDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlbumDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlbumDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailResponse) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlbumDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlbumDetailResponse) GetBody() *ListAlbumDetailResponseBody {
	return s.Body
}

func (s *ListAlbumDetailResponse) SetHeaders(v map[string]*string) *ListAlbumDetailResponse {
	s.Headers = v
	return s
}

func (s *ListAlbumDetailResponse) SetStatusCode(v int32) *ListAlbumDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlbumDetailResponse) SetBody(v *ListAlbumDetailResponseBody) *ListAlbumDetailResponse {
	s.Body = v
	return s
}

func (s *ListAlbumDetailResponse) Validate() error {
	return dara.Validate(s)
}
