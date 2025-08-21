// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMusicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMusicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMusicResponse
	GetStatusCode() *int32
	SetBody(v *ListMusicResponseBody) *ListMusicResponse
	GetBody() *ListMusicResponseBody
}

type ListMusicResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMusicResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMusicResponse) GoString() string {
	return s.String()
}

func (s *ListMusicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMusicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMusicResponse) GetBody() *ListMusicResponseBody {
	return s.Body
}

func (s *ListMusicResponse) SetHeaders(v map[string]*string) *ListMusicResponse {
	s.Headers = v
	return s
}

func (s *ListMusicResponse) SetStatusCode(v int32) *ListMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMusicResponse) SetBody(v *ListMusicResponseBody) *ListMusicResponse {
	s.Body = v
	return s
}

func (s *ListMusicResponse) Validate() error {
	return dara.Validate(s)
}
