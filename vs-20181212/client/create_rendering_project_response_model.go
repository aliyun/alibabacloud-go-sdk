// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRenderingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRenderingProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateRenderingProjectResponseBody) *CreateRenderingProjectResponse
	GetBody() *CreateRenderingProjectResponseBody
}

type CreateRenderingProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRenderingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRenderingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateRenderingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRenderingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRenderingProjectResponse) GetBody() *CreateRenderingProjectResponseBody {
	return s.Body
}

func (s *CreateRenderingProjectResponse) SetHeaders(v map[string]*string) *CreateRenderingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateRenderingProjectResponse) SetStatusCode(v int32) *CreateRenderingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRenderingProjectResponse) SetBody(v *CreateRenderingProjectResponseBody) *CreateRenderingProjectResponse {
	s.Body = v
	return s
}

func (s *CreateRenderingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
