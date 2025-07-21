// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientKeyResponseBody) *CreateClientKeyResponse
	GetBody() *CreateClientKeyResponseBody
}

type CreateClientKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateClientKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientKeyResponse) GetBody() *CreateClientKeyResponseBody {
	return s.Body
}

func (s *CreateClientKeyResponse) SetHeaders(v map[string]*string) *CreateClientKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateClientKeyResponse) SetStatusCode(v int32) *CreateClientKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientKeyResponse) SetBody(v *CreateClientKeyResponseBody) *CreateClientKeyResponse {
	s.Body = v
	return s
}

func (s *CreateClientKeyResponse) Validate() error {
	return dara.Validate(s)
}
