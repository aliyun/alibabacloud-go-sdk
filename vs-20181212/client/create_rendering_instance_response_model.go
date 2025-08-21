// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateRenderingInstanceResponseBody) *CreateRenderingInstanceResponse
	GetBody() *CreateRenderingInstanceResponseBody
}

type CreateRenderingInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRenderingInstanceResponse) GetBody() *CreateRenderingInstanceResponseBody {
	return s.Body
}

func (s *CreateRenderingInstanceResponse) SetHeaders(v map[string]*string) *CreateRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateRenderingInstanceResponse) SetStatusCode(v int32) *CreateRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRenderingInstanceResponse) SetBody(v *CreateRenderingInstanceResponseBody) *CreateRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateRenderingInstanceResponse) Validate() error {
	return dara.Validate(s)
}
