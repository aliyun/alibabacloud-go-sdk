// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateHostShareKeyResponseBody) *CreateHostShareKeyResponse
	GetBody() *CreateHostShareKeyResponseBody
}

type CreateHostShareKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHostShareKeyResponse) GetBody() *CreateHostShareKeyResponseBody {
	return s.Body
}

func (s *CreateHostShareKeyResponse) SetHeaders(v map[string]*string) *CreateHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateHostShareKeyResponse) SetStatusCode(v int32) *CreateHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostShareKeyResponse) SetBody(v *CreateHostShareKeyResponseBody) *CreateHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *CreateHostShareKeyResponse) Validate() error {
	return dara.Validate(s)
}
