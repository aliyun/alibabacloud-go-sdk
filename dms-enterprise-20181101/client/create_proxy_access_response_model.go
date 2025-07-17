// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProxyAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProxyAccessResponse
	GetStatusCode() *int32
	SetBody(v *CreateProxyAccessResponseBody) *CreateProxyAccessResponse
	GetBody() *CreateProxyAccessResponseBody
}

type CreateProxyAccessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProxyAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProxyAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyAccessResponse) GoString() string {
	return s.String()
}

func (s *CreateProxyAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProxyAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProxyAccessResponse) GetBody() *CreateProxyAccessResponseBody {
	return s.Body
}

func (s *CreateProxyAccessResponse) SetHeaders(v map[string]*string) *CreateProxyAccessResponse {
	s.Headers = v
	return s
}

func (s *CreateProxyAccessResponse) SetStatusCode(v int32) *CreateProxyAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProxyAccessResponse) SetBody(v *CreateProxyAccessResponseBody) *CreateProxyAccessResponse {
	s.Body = v
	return s
}

func (s *CreateProxyAccessResponse) Validate() error {
	return dara.Validate(s)
}
