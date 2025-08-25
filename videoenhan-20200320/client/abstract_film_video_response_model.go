// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbstractFilmVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbstractFilmVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbstractFilmVideoResponse
	GetStatusCode() *int32
	SetBody(v *AbstractFilmVideoResponseBody) *AbstractFilmVideoResponse
	GetBody() *AbstractFilmVideoResponseBody
}

type AbstractFilmVideoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbstractFilmVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbstractFilmVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s AbstractFilmVideoResponse) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbstractFilmVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbstractFilmVideoResponse) GetBody() *AbstractFilmVideoResponseBody {
	return s.Body
}

func (s *AbstractFilmVideoResponse) SetHeaders(v map[string]*string) *AbstractFilmVideoResponse {
	s.Headers = v
	return s
}

func (s *AbstractFilmVideoResponse) SetStatusCode(v int32) *AbstractFilmVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *AbstractFilmVideoResponse) SetBody(v *AbstractFilmVideoResponseBody) *AbstractFilmVideoResponse {
	s.Body = v
	return s
}

func (s *AbstractFilmVideoResponse) Validate() error {
	return dara.Validate(s)
}
