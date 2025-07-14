// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebApplicationResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationBody) *CreateWebApplicationResponse
	GetBody() *WebApplicationBody
}

type CreateWebApplicationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateWebApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebApplicationResponse) GetBody() *WebApplicationBody {
	return s.Body
}

func (s *CreateWebApplicationResponse) SetHeaders(v map[string]*string) *CreateWebApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateWebApplicationResponse) SetStatusCode(v int32) *CreateWebApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebApplicationResponse) SetBody(v *WebApplicationBody) *CreateWebApplicationResponse {
	s.Body = v
	return s
}

func (s *CreateWebApplicationResponse) Validate() error {
	return dara.Validate(s)
}
