// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateKeyResponseBody) *CreateKeyResponse
	GetBody() *CreateKeyResponseBody
}

type CreateKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKeyResponse) GetBody() *CreateKeyResponseBody {
	return s.Body
}

func (s *CreateKeyResponse) SetHeaders(v map[string]*string) *CreateKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyResponse) SetStatusCode(v int32) *CreateKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyResponse) SetBody(v *CreateKeyResponseBody) *CreateKeyResponse {
	s.Body = v
	return s
}

func (s *CreateKeyResponse) Validate() error {
	return dara.Validate(s)
}
