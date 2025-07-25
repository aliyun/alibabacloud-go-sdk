// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdnsAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdnsAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdnsAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreatePdnsAppKeyResponseBody) *CreatePdnsAppKeyResponse
	GetBody() *CreatePdnsAppKeyResponseBody
}

type CreatePdnsAppKeyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePdnsAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdnsAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdnsAppKeyResponse) GoString() string {
	return s.String()
}

func (s *CreatePdnsAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdnsAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdnsAppKeyResponse) GetBody() *CreatePdnsAppKeyResponseBody {
	return s.Body
}

func (s *CreatePdnsAppKeyResponse) SetHeaders(v map[string]*string) *CreatePdnsAppKeyResponse {
	s.Headers = v
	return s
}

func (s *CreatePdnsAppKeyResponse) SetStatusCode(v int32) *CreatePdnsAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdnsAppKeyResponse) SetBody(v *CreatePdnsAppKeyResponseBody) *CreatePdnsAppKeyResponse {
	s.Body = v
	return s
}

func (s *CreatePdnsAppKeyResponse) Validate() error {
	return dara.Validate(s)
}
