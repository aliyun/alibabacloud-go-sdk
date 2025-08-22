// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse
	GetBody() *DeleteProjectResponseBody
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectResponse) GetBody() *DeleteProjectResponseBody {
	return s.Body
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectResponse) Validate() error {
	return dara.Validate(s)
}
