// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRenderingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRenderingProjectResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRenderingProjectResponseBody) *UpdateRenderingProjectResponse
	GetBody() *UpdateRenderingProjectResponseBody
}

type UpdateRenderingProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRenderingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRenderingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateRenderingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRenderingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRenderingProjectResponse) GetBody() *UpdateRenderingProjectResponseBody {
	return s.Body
}

func (s *UpdateRenderingProjectResponse) SetHeaders(v map[string]*string) *UpdateRenderingProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateRenderingProjectResponse) SetStatusCode(v int32) *UpdateRenderingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRenderingProjectResponse) SetBody(v *UpdateRenderingProjectResponseBody) *UpdateRenderingProjectResponse {
	s.Body = v
	return s
}

func (s *UpdateRenderingProjectResponse) Validate() error {
	return dara.Validate(s)
}
