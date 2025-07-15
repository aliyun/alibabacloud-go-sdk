// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDirectoriesResponseBody) *DeleteDirectoriesResponse
	GetBody() *DeleteDirectoriesResponseBody
}

type DeleteDirectoriesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDirectoriesResponse) GetBody() *DeleteDirectoriesResponseBody {
	return s.Body
}

func (s *DeleteDirectoriesResponse) SetHeaders(v map[string]*string) *DeleteDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoriesResponse) SetStatusCode(v int32) *DeleteDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirectoriesResponse) SetBody(v *DeleteDirectoriesResponseBody) *DeleteDirectoriesResponse {
	s.Body = v
	return s
}

func (s *DeleteDirectoriesResponse) Validate() error {
	return dara.Validate(s)
}
