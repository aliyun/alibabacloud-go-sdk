// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRenderingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRenderingProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRenderingProjectResponseBody) *DeleteRenderingProjectResponse
	GetBody() *DeleteRenderingProjectResponseBody
}

type DeleteRenderingProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRenderingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRenderingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteRenderingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRenderingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRenderingProjectResponse) GetBody() *DeleteRenderingProjectResponseBody {
	return s.Body
}

func (s *DeleteRenderingProjectResponse) SetHeaders(v map[string]*string) *DeleteRenderingProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteRenderingProjectResponse) SetStatusCode(v int32) *DeleteRenderingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRenderingProjectResponse) SetBody(v *DeleteRenderingProjectResponseBody) *DeleteRenderingProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteRenderingProjectResponse) Validate() error {
	return dara.Validate(s)
}
